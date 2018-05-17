package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris"
)

// ChargeUnPayCreate 创建收费列表
func ChargeUnPayCreate(ctx iris.Context) {

	items := ctx.PostValue("items")
	if items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	var results []map[string]interface{}
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	timestamp := time.Now().Unix()
	random := strconv.FormatFloat(rand.Float64(), 'E', 10, 64)[2:5]
	orderSn := strconv.FormatInt(timestamp, 10) + random

	var sets []string

	for index, v := range results {

		var set []string

		if v["registration_id"] == nil || v["charge_project_type_id"] == nil || v["charge_project_id"] == nil || v["amount"] == nil || v["discount"] == nil || v["operation_id"] == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
			return
		}

		registrationID := strconv.FormatFloat(v["registration_id"].(float64), 'f', 0, 64)
		chargeProjecttypeid := strconv.FormatFloat(v["charge_project_type_id"].(float64), 'f', 0, 64)
		chargeprojectid := strconv.FormatFloat(v["charge_project_id"].(float64), 'f', 0, 64)
		amount := int(v["amount"].(float64))
		discount := int(v["discount"].(float64))
		operation := strconv.FormatFloat(v["operation_id"].(float64), 'f', 0, 64)

		var (
			name  string
			unit  string
			price int
			total int
			fee   int
		)

		switch chargeProjecttypeid {
		case "8":
			rows := model.DB.QueryRowx("select * from charge_project_treatment where project_type_id=" + chargeProjecttypeid + " AND id=" + chargeprojectid)
			result := FormatSQLRowToMap(rows)
			_, exsist := result["id"]
			if !exsist {
				ctx.JSON(iris.Map{"code": "-1", "msg": "未找到指定收费项 " + chargeprojectid})
				return
			}
			name = result["name"].(string)
			unit = "次"
			price = int(result["fee"].(int64))
			total = price * amount
			fee = total - discount
			orderSn = "F8-" + orderSn
			break
		default:
			ctx.JSON(iris.Map{"code": "-1", "msg": "charge_project_type_id 无效"})
			return
		}

		set = append(set, registrationID)
		set = append(set, chargeProjecttypeid)
		set = append(set, chargeprojectid)
		set = append(set, operation)
		set = append(set, "'"+orderSn+"'")
		set = append(set, strconv.Itoa(index+1))
		set = append(set, "'"+name+"'")
		set = append(set, "'"+unit+"'")
		set = append(set, strconv.Itoa(price))
		set = append(set, strconv.Itoa(amount))
		set = append(set, strconv.Itoa(total))
		set = append(set, strconv.Itoa(discount))
		set = append(set, strconv.Itoa(fee))
		setn := "(" + strings.Join(set, ",") + ")"

		sets = append(sets, setn)
	}

	setst := strings.Join(sets, ",")

	sql := "INSERT INTO mz_unpaid_orders (registration_id, charge_project_type_id, charge_project_id, operation_id, order_sn, soft_sn, name, unit, price, amount, total, discount, fee ) VALUES " + setst

	_, err1 := model.DB.Query(sql)

	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})

}

// ChargeUnPayDelete 删除缴费项目
func ChargeUnPayDelete(ctx iris.Context) {
	id := ctx.PostValue("id")
	if id == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	_, err := model.DB.Query("DELETE FROM mz_unpaid_orders id=" + id)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})

}

// ChargeUnPayList 根据预约编码查询待缴费列表
func ChargeUnPayList(ctx iris.Context) {
	clinicTriagePatientid := ctx.PostValue("clinic_triage_patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if clinicTriagePatientid == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if offset == "" {
		offset = "0"
	}

	if limit == "" {
		limit = "10"
	}

	_, err := strconv.Atoi(offset)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "offset 必须为数字"})
		return
	}
	_, err = strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "limit 必须为数字"})
		return
	}

	total := model.DB.QueryRowx(`select count(id) as total,sum(total) as charge_total,sum(discount) as discount_total,sum(fee) as charge_total_fee from mz_unpaid_orders where clinic_triage_patient_id=$1`, clinicTriagePatientid)

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	typesql := `select sum(total) as type_charge_total, charge_project_type_id from mz_unpaid_orders where clinic_triage_patient_id = $1 group by charge_project_type_id`

	typetotal, _ := model.DB.Queryx(typesql, clinicTriagePatientid)

	typetotalfomat := FormatSQLRowsToMapArray(typetotal)

	rowSQL := `select m.name,m.price,m.amount,m.total,m.discount,m.fee,p.name as doctor_name,d.name as department_name from mz_unpaid_orders m 
	left join personnel p on p.id = m.operation_id 
	left join department_personnel dp on dp.personnel_id = p.id 
	left join department d on d.id = dp.department_id 
	where m.clinic_triage_patient_id=$1 order by m.created_time DESC offset $2 limit $3`

	rows, err1 := model.DB.Queryx(rowSQL, clinicTriagePatientid, offset, limit)

	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo, "type_total": typetotalfomat})

}

// ChargePay 缴费
func ChargePay(ctx iris.Context) {

	registrationID := ctx.PostValue("registration_id")
	orderSn := ctx.PostValue("order_sn")
	softSn := ctx.PostValue("soft_sn")
	confrimID := ctx.PostValue("confrim_id")
	payTypeCode := ctx.PostValue("pay_type_code")
	payMethodCode := ctx.PostValue("pay_method_code")
	balanceMoney := ctx.PostValue("balance_money")
	totalMoney := ctx.PostValue("total_money")
	outTradeNo := ctx.PostValue("out_trade_no")

	discountRate := ctx.PostValue("discount_rate")

	derateMoney, _ := strconv.Atoi(ctx.PostValue("derate_money"))
	medicalMoney, _ := strconv.Atoi(ctx.PostValue("medical_money"))
	onCreditMoney, _ := strconv.Atoi(ctx.PostValue("on_credit_money"))
	voucherMoney, _ := strconv.Atoi(ctx.PostValue("voucher_money"))
	bonusPointsMoney, _ := strconv.Atoi(ctx.PostValue("bonus_points_money"))

	if registrationID == "" || orderSn == "" || softSn == "" || balanceMoney == "" || totalMoney == "" || confrimID == "" || payTypeCode == "" || payMethodCode == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	//查询分诊就录是否存在
	registrationRows := model.DB.QueryRowx("select r.*,cp.patient_id from registration r left join clinic_patient cp on r.clinic_patient_id = cp.id where r.id =" + registrationID)
	registration := FormatSQLRowToMap(registrationRows)
	if registration["id"] == nil {
		ctx.JSON(iris.Map{"code": "-1", "data": nil, "msg": "未找到指定问诊记录"})
		return
	}

	balanceMoneyInt, _ := strconv.Atoi(balanceMoney)
	totalMoneyInt, _ := strconv.Atoi(totalMoney)
	if discountRate == "" {
		discountRate = "100"
	}

	discountRateInt, _ := strconv.Atoi(discountRate)

	//判断实收金额
	sBalance := totalMoneyInt*(discountRateInt/100) - (derateMoney + medicalMoney + onCreditMoney + voucherMoney + bonusPointsMoney)

	discountMoney := totalMoneyInt * ((100 - discountRateInt) / 100)

	if sBalance != balanceMoneyInt {
		cmap := map[string]interface{}{
			"balance":      sBalance,
			"balanceMoney": balanceMoneyInt,
		}
		ctx.JSON(iris.Map{"code": "-1", "data": cmap, "msg": "实收款结算错误"})
		return
	}

	//判断应收金额是否正确
	rows := model.DB.QueryRowx("select SUM(fee) from mz_unpaid_orders where order_sn='" + orderSn + "' AND soft_sn in (" + softSn + ")")
	result := FormatSQLRowToMap(rows)
	var total int
	sum := result["sum"]

	if sum == nil {
		ctx.JSON(iris.Map{"code": "-1", "data": nil, "msg": "未找到指定收费项"})
		return
	}

	total = int(sum.(int64))

	if total != totalMoneyInt {
		tmap := map[string]interface{}{
			"total":      total,
			"totalMoney": totalMoneyInt,
		}
		ctx.JSON(iris.Map{"code": "-1", "data": tmap, "msg": "应收收款结算错误"})
		return
	}

	tx, err := model.DB.Beginx()

	if err != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
		return
	}

	//插入门诊缴费记录
	var recordID int
	insertPaid := "insert into mz_paid_record (registration_id,out_trade_no,soft_sns,order_sn,confrim_id,pay_type_code,pay_method_code,status,discount_money,derate_money,medical_money,on_credit_money,voucher_money,bonus_points_money,total_money,balance_money) " +
		"values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16) RETURNING id"
	insertPaidErr := tx.QueryRow(insertPaid, registrationID, outTradeNo, softSn, orderSn, confrimID, payTypeCode, payMethodCode, "SUCCESS", discountMoney, derateMoney, medicalMoney, onCreditMoney, voucherMoney, bonusPointsMoney, totalMoneyInt, balanceMoneyInt).Scan(&recordID)
	if insertPaidErr != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "2", "msg": insertPaidErr.Error()})
		return
	}

	//如果有挂账金额，添加挂账记录
	if onCreditMoney > 0 {

	}

	//插入门诊缴费流水表
	var ID int
	insertPaidDetail := "insert into mz_paid_record_detail (mz_paid_record_id,out_trade_no,soft_sns,order_sn,confrim_id,discount_money,derate_money,medical_money,on_credit_money,voucher_money,bonus_points_money,total_money,balance_money) " +
		"values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) RETURNING id"
	insertPaidDetailErr := tx.QueryRow(insertPaidDetail, recordID, outTradeNo, softSn, orderSn, confrimID, discountMoney, derateMoney, medicalMoney, onCreditMoney, voucherMoney, bonusPointsMoney, totalMoneyInt, balanceMoneyInt).Scan(&ID)
	if insertPaidDetailErr != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "3", "msg": insertPaidDetailErr.Error()})
		return
	}

	//插入交易流水表
	var sID int
	insertPaidCharge := "insert into charge_detail (pay_record_id,out_trade_no,in_out,patient_id,department_id,doctor_id,pay_type_code,pay_type_code_name,pay_method_code,pay_method_code_name,discount_money,derate_money,medical_money,on_credit_money,voucher_money,bonus_points_money,total_money,balance_money) " +
		"values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18) RETURNING id"
	insertPaidChargeErr := tx.QueryRow(insertPaidCharge, recordID, outTradeNo, "in", registration["patient_id"], registration["department_id"], registration["personnel_id"], "01", "门诊缴费", payMethodCode, "", discountMoney, derateMoney, medicalMoney, onCreditMoney, voucherMoney, bonusPointsMoney, totalMoneyInt, balanceMoneyInt).Scan(&sID)
	if insertPaidChargeErr != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "4", "msg": insertPaidChargeErr.Error()})
		return
	}

	//插入已缴费
	insertPaidOrders := "insert into mz_paid_orders (id,mz_paid_record_id,registration_id,charge_project_type_id,charge_project_id,order_sn,soft_sn,name,price,amount,unit,total,discount,fee,operation_id,confrim_id)" +
		" select id," + strconv.Itoa(recordID) + ",registration_id,charge_project_type_id,charge_project_id,order_sn,soft_sn,name,price,amount,unit,total,discount,fee,operation_id," + confrimID + " from mz_unpaid_orders where order_sn='" + orderSn + "' AND soft_sn in (" + softSn + ")"
	_, insertPaidOrdersErr := tx.Query(insertPaidOrders)
	if insertPaidOrdersErr != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "5", "msg": insertPaidOrdersErr.Error()})
		return
	}

	//删除未交费
	deleteUnPaid := "DELETE from mz_unpaid_orders where order_sn='" + orderSn + "' AND soft_sn in (" + softSn + ")"
	_, deleteUnPaidErr := tx.Query(deleteUnPaid)
	if deleteUnPaidErr != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "6", "msg": deleteUnPaidErr.Error()})
		return
	}

	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "7", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": recordID})

}

// ChargePaidList 根据预约编码查询已缴费缴费列表
func ChargePaidList(ctx iris.Context) {
	registrationid := ctx.PostValue("registration_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if registrationid == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if offset == "" {
		offset = "0"
	}

	if limit == "" {
		limit = "10"
	}

	_, err := strconv.Atoi(offset)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "offset 必须为数字"})
		return
	}
	_, err = strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "limit 必须为数字"})
		return
	}

	total := model.DB.QueryRowx(`select count(id) as total from mz_paid_orders where registration_id=$1`, registrationid)

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rowSQL := `select * from mz_paid_orders where registration_id=$1 offset $2 limit $3`

	rows, err1 := model.DB.Queryx(rowSQL, registrationid, offset, limit)

	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}

	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})

}
