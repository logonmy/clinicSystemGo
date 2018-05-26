package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"errors"
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
	totalID, _ := model.DB.Queryx(`select id from mz_unpaid_orders where clinic_triage_patient_id=$1`, clinicTriagePatientid)
	pageInfo := FormatSQLRowToMap(total)
	totalIds := FormatSQLRowsToMapArray(totalID)
	arr := []string{}
	for _, item := range totalIds {
		id := item["id"]
		arr = append(arr, strconv.FormatInt(id.(int64), 10))
	}
	arrStr := strings.Join(arr, ",")
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit
	pageInfo["totalIds"] = arrStr

	typesql := `select sum(total) as type_charge_total, charge_project_type_id from mz_unpaid_orders where clinic_triage_patient_id = $1 group by charge_project_type_id`

	typetotal, _ := model.DB.Queryx(typesql, clinicTriagePatientid)

	typetotalfomat := FormatSQLRowsToMapArray(typetotal)

	rowSQL := `select m.id as mz_unpaid_orders_id,m.name,m.price,m.amount,m.total,m.discount,m.fee,p.name as doctor_name,d.name as department_name from mz_unpaid_orders m 
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

// ChargePaymentCreate 创建支付订单
func ChargePaymentCreate(ctx iris.Context) {
	discountMoney, _ := strconv.Atoi(ctx.PostValue("discount_money"))
	derateMoney, _ := strconv.Atoi(ctx.PostValue("derate_money"))
	medicalMoney, _ := strconv.Atoi(ctx.PostValue("medical_money"))
	onCreditMoney, _ := strconv.Atoi(ctx.PostValue("on_credit_money"))
	voucherMoney, _ := strconv.Atoi(ctx.PostValue("voucher_money"))
	bonusPointsMoney, _ := strconv.Atoi(ctx.PostValue("bonus_points_money"))
	money, baerr := strconv.Atoi(ctx.PostValue("balance_money"))

	if baerr != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "无效的金额"})
		return
	}

	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	ordersIds := ctx.PostValue("orders_ids")
	operationID := ctx.PostValue("operation_id")
	payMethodCode := ctx.PostValue("pay_method_code")

	outTradeNo := time.Now().Format("20060102150405")

	if clinicTriagePatientID == "" || ordersIds == "" || operationID == "" || payMethodCode == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx(`select count(*) as count, sum(fee) as charge_toral from mz_unpaid_orders where id in (` + ordersIds + `) AND clinic_triage_patient_id = ` + clinicTriagePatientID)
	orderArray := strings.Split(ordersIds, ",")
	rowMap := FormatSQLRowToMap(row)

	if int(rowMap["count"].(int64)) != len(orderArray) {
		ctx.JSON(iris.Map{"code": "2", "msg": "存在未知缴费项"})
		return
	}

	if rowMap["charge_toral"] == nil {
		ctx.JSON(iris.Map{"code": "2", "msg": "收费金额异常"})
		return
	}

	totalMoney := rowMap["charge_toral"]

	balanceMoney := int(totalMoney.(int64)) - (derateMoney + voucherMoney + discountMoney + bonusPointsMoney + onCreditMoney + medicalMoney)

	if balanceMoney != money {
		balanceMoneyStr := strconv.Itoa(balanceMoney)
		moneyStr := strconv.Itoa(money)
		ctx.JSON(iris.Map{"code": "2", "msg": "收费金额与应收金额不匹配，应收: " + balanceMoneyStr + "分钱，实收：" + moneyStr + "分钱！"})
		return
	}

	if balanceMoney < 0 {
		ctx.JSON(iris.Map{"code": "3", "msg": "收费金额不能为负数"})
		return
	}

	tx, txerr := model.DB.Beginx()
	if txerr != nil {
		ctx.JSON(iris.Map{"code": "2", "msg": txerr.Error()})
		return
	}

	_, err := tx.Exec(`INSERT INTO mz_paid_record ( 
		clinic_triage_patient_id, out_trade_no, orders_ids, operation_id, pay_method_code, status, derate_money, voucher_money, discount_money, bonus_points_money, on_credit_money, medical_money, total_money, balance_money)
		VALUES ($1, $2, $3, $4, $5, 'WATTING_FOR_PAY', $6, $7, $8, $9, $10, $11, $12, $13)`, clinicTriagePatientID, outTradeNo, ordersIds, operationID, payMethodCode, derateMoney, voucherMoney, discountMoney, bonusPointsMoney, onCreditMoney, medicalMoney, totalMoney, balanceMoney)
	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "3", "msg": err.Error()})
		return
	}

	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "7", "msg": err.Error()})
		return
	}

	// 如果金额为0 直接通知
	if balanceMoney == 0 || payMethodCode == "4" {
		chaerr := charge(outTradeNo, outTradeNo, int64(balanceMoney))
		if chaerr != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "缴费通知失败"})
			return
		}
		ctx.JSON(iris.Map{"code": "300", "msg": "直接缴费成功"})
		return
	}

	data := map[string]interface{}{}
	data["total_money"] = totalMoney
	data["balance_money"] = balanceMoney
	data["out_trade_no"] = outTradeNo

	ctx.JSON(iris.Map{"code": "200", "data": data, "msg": "成功"})
}

// ChargePaidList 根据预约编码查询已缴费缴费列表
func ChargePaidList(ctx iris.Context) {
	mzPaidRecordID := ctx.PostValue("mz_paid_record_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if mzPaidRecordID == "" {
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

	payment := model.DB.QueryRowx(`select * from mz_paid_record where id=$1`, mzPaidRecordID)
	total := model.DB.QueryRowx(`select count(id) as total,sum(total) as charge_total,sum(discount) as discount_total,sum(fee) as charge_total_fee from mz_paid_orders where mz_paid_record_id=$1`, mzPaidRecordID)
	pageInfo := FormatSQLRowToMap(payment)
	totalMap := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit
	pageInfo["total"] = totalMap["total"]
	pageInfo["charge_total"] = totalMap["charge_total"]
	pageInfo["discount_total"] = totalMap["discount_total"]
	pageInfo["charge_total_fee"] = totalMap["charge_total_fee"]

	typesql := `select sum(total) as type_charge_total, charge_project_type_id from mz_paid_orders where mz_paid_record_id = $1 group by charge_project_type_id`

	typetotal, _ := model.DB.Queryx(typesql, mzPaidRecordID)

	typetotalfomat := FormatSQLRowsToMapArray(typetotal)

	rowSQL := `select m.id as mz_paid_orders_id,m.name,m.price,m.amount,m.total,m.discount,m.fee,p.name as doctor_name,d.name as department_name from mz_paid_orders m 
	left join personnel p on p.id = m.operation_id 
	left join department_personnel dp on dp.personnel_id = p.id 
	left join department d on d.id = dp.department_id
	where m.mz_paid_record_id=$1 order by m.created_time DESC offset $2 limit $3`

	rows, err1 := model.DB.Queryx(rowSQL, mzPaidRecordID, offset, limit)

	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo, "type_total": typetotalfomat})

}

// ChargeNotice 缴费通知
func ChargeNotice(ctx iris.Context) {
	outTradeNo := ctx.PostValue("out_trade_no")
	tradeNo := ctx.PostValue("trade_no")
	money, merr := ctx.PostValueInt64("money") //钱以分为单位
	if merr != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "无效的输入金额"})
		return
	}

	if outTradeNo == "" || tradeNo == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	err := charge(outTradeNo, tradeNo, money)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "缴费成功"})
	return
}

// 处理缴费通知
func charge(outTradeNo string, tradeNo string, money int64) error {
	payment := model.DB.QueryRowx("select * from mz_paid_record where out_trade_no = $1", outTradeNo)
	pay := FormatSQLRowToMap(payment)
	_, ok := pay["id"]
	if !ok {
		return errors.New("未找到指定的待缴费单")
	}
	balanceMoney := pay["balance_money"]
	if balanceMoney.(int64) != money {
		return errors.New("缴费金额不匹配")
	}
	tx, txErr := model.DB.Beginx()
	if txErr != nil {
		return txErr
	}
	_, updateErr := tx.Exec(`update mz_paid_record set trade_no = $1,updated_time = LOCALTIMESTAMP,status='TRADE_SUCCESS' WHERE out_trade_no = $2`, tradeNo, outTradeNo)
	if updateErr != nil {
		tx.Rollback()
		return updateErr
	}

	operationID := pay["operation_id"]

	onCreditMoney := pay["on_credit_money"]
	if onCreditMoney.(int64) != 0 {
		triageID := pay["clinic_triage_patient_id"]
		creditSQL := `INSERT INTO on_credit_record( 
		  clinic_triage_patient_id, on_credit_money, trade_no, operation_id) 
			VALUES ($1, $2, $3, $4)`
		_, creditErr := tx.Exec(creditSQL, triageID, onCreditMoney, tradeNo, operationID)
		if creditErr != nil {
			tx.Rollback()
			return updateErr
		}
	}

	orderIDs := pay["orders_ids"]
	recordID := pay["id"]
	confrimID := pay["operation_id"]

	//插入已缴费
	insertPaidOrders := "insert into mz_paid_orders (id,mz_paid_record_id,clinic_triage_patient_id,charge_project_type_id,charge_project_id,order_sn,soft_sn,name,price,amount,unit,total,discount,fee,operation_id,confrim_id)" +
		" select id," + strconv.Itoa(int(recordID.(int64))) + ",clinic_triage_patient_id,charge_project_type_id,charge_project_id,order_sn,soft_sn,name,price,amount,unit,total,discount,fee,operation_id," + strconv.Itoa(int(confrimID.(int64))) + " from mz_unpaid_orders where id in (" + orderIDs.(string) + ")"
	_, insertPaidOrdersErr := tx.Query(insertPaidOrders)
	if insertPaidOrdersErr != nil {
		tx.Rollback()
		return insertPaidOrdersErr
	}

	//删除未交费
	deleteUnPaid := "DELETE from mz_unpaid_orders where id in (" + orderIDs.(string) + ")"
	_, deleteUnPaidErr := tx.Query(deleteUnPaid)
	if deleteUnPaidErr != nil {
		tx.Rollback()
		return deleteUnPaidErr
	}

	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
