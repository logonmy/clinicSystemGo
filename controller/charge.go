package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
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
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
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
	insertPaidCharge := "insert into charge_detail (pay_record_id,record_type,out_trade_no,in_out,patient_id,department_id,doctor_id,pay_type_code,pay_type_code_name,pay_method_code,pay_method_code_name,discount_money,derate_money,medical_money,on_credit_money,voucher_money,bonus_points_money,total_money,balance_money) " +
		"values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19) RETURNING id"
	insertPaidChargeErr := tx.QueryRow(insertPaidCharge, recordID, 1, outTradeNo, "in", registration["patient_id"], registration["department_id"], registration["personnel_id"], "01", "门诊缴费", payMethodCode, "", discountMoney, derateMoney, medicalMoney, onCreditMoney, voucherMoney, bonusPointsMoney, totalMoneyInt, balanceMoneyInt).Scan(&sID)
	if insertPaidChargeErr != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "4", "msg": insertPaidChargeErr.Error()})
		return
	}

	//插入已缴费
	insertPaidOrders := "insert into mz_paid_orders (id,mz_paid_record_id,registration_id,charge_project_type_id,charge_project_id,order_sn,soft_sn,name,cost,price,amount,unit,total,discount,fee,operation_id,confrim_id)" +
		" select id," + strconv.Itoa(recordID) + ",registration_id,charge_project_type_id,charge_project_id,order_sn,soft_sn,name,cost,price,amount,unit,total,discount,fee,operation_id," + confrimID + " from mz_unpaid_orders where order_sn='" + orderSn + "' AND soft_sn in (" + softSn + ")"
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
	authCode := ctx.PostValue("auth_code")

	if baerr != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "无效的金额"})
		return
	}

	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	ordersIds := ctx.PostValue("orders_ids")
	operationID := ctx.PostValue("operation_id")
	payMethodCode := ctx.PostValue("pay_method_code")

	outTradeNo := GetTradeNo("T2")

	if clinicTriagePatientID == "" || ordersIds == "" || operationID == "" || payMethodCode == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if (payMethodCode == "1" || payMethodCode == "2") && authCode == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少认证码"})
		return
	}

	row := model.DB.QueryRowx(`select count(*) as count, sum(fee) as charge_total from mz_unpaid_orders where id in (` + ordersIds + `) AND clinic_triage_patient_id = ` + clinicTriagePatientID)
	orderArray := strings.Split(ordersIds, ",")
	rowMap := FormatSQLRowToMap(row)

	if int(rowMap["count"].(int64)) != len(orderArray) {
		ctx.JSON(iris.Map{"code": "2", "msg": "存在未知缴费项"})
		return
	}

	if rowMap["charge_total"] == nil {
		ctx.JSON(iris.Map{"code": "2", "msg": "收费金额异常"})
		return
	}

	totalMoney := rowMap["charge_total"]

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
		chaerr := charge(outTradeNo, outTradeNo)
		if chaerr != nil {
			fmt.Println(chaerr)
			ctx.JSON(iris.Map{"code": "-1", "msg": "缴费通知失败"})
			return
		}
		ctx.JSON(iris.Map{"code": "200", "msg": "直接缴费成功"})
		return
	} else if payMethodCode == "1" || payMethodCode == "2" {

		requestIP := ctx.Host()
		requestIP = requestIP[0:strings.LastIndex(requestIP, ":")]
		if strings.ToLower(requestIP) == "localhost" {
			requestIP = "127.0.0.1"
		}

		merID := "ali"
		payModel := "alipay_f2f"
		if payMethodCode == "2" {
			merID = "wx"
			payModel = "weixin_f2f"
		}

		result := FaceToFace(outTradeNo, authCode, merID, payModel, "mz", strconv.Itoa(balanceMoney), "门诊缴费", requestIP, "", "")
		if result["code"].(string) != "200" {
			if result["code"].(string) == "2" {
				_, err := model.DB.Exec("update mz_paid_record set status = 'WATTING_PASSWORD' where out_trade_no = $1", outTradeNo)
				if err != nil {
					ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
					return
				}
				ctx.JSON(iris.Map{"code": "300", "msg": result["msg"], "data": outTradeNo})
			} else {
				model.DB.Exec("update mz_paid_record set status = 'TRADE_FAILED' where out_trade_no = $1", outTradeNo)
				ctx.JSON(iris.Map{"code": result["code"], "msg": result["msg"]})
				return
			}
		} else {
			data := result["data"].(map[string]interface{})
			err := charge(outTradeNo, data["transaction_id"].(string))
			if err != nil {
				ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
				return
			}
			ctx.JSON(iris.Map{"code": "200", "data": 1})
			return
		}
	} else {
		ctx.JSON(iris.Map{"code": "200", "msg": "不支持的缴费方式"})
		return
	}
}

// ChargePaymentQuery 获取支付状态
func ChargePaymentQuery(ctx iris.Context) {
	outTradeNo := ctx.PostValue("out_trade_no")
	if outTradeNo == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	res := QueryHcOrder(outTradeNo, "")

	if res["code"] == "200" {
		data := res["data"].(map[string]interface{})
		tradeStatus := data["trade_status"].(string)
		if tradeStatus == "SUCCESS" {
			err := charge(outTradeNo, data["transaction_id"].(string))
			if err != nil {
				fmt.Println("缴费通知失败", err.Error())
			}
		}
		ctx.JSON(iris.Map{"code": "200", "data": res["data"]})
	} else {
		ctx.JSON(iris.Map{"code": "-1", "msg": res["msg"]})
	}

}

// ChargePaymentRefund 门诊退费
func ChargePaymentRefund(ctx iris.Context) {
	outTradeNo := ctx.PostValue("out_trade_no")
	refundIDs := ctx.PostValue("refundIds")
	operarton := ctx.PostValue("operation_id")
	if outTradeNo == "" || refundIDs == "" || operarton == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	var results []string
	err := json.Unmarshal([]byte(refundIDs), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	paidRecordRow := model.DB.QueryRowx("select * from mz_paid_record where out_trade_no = $1", outTradeNo)
	paidRecord := FormatSQLRowToMap(paidRecordRow)
	if paidRecord["id"] == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "未找到指定的缴费记录"})
		return
	}

	triagePatient := model.DB.QueryRowx("select * from clinic_triage_patient where id = $1", paidRecord["clinic_triage_patient_id"])
	triage := FormatSQLRowToMap(triagePatient)

	balanceMoney := paidRecord["balance_money"].(int64)
	totalMoney := paidRecord["total_money"].(int64)

	tx, txErr := model.DB.Beginx()
	if txErr != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": txErr.Error()})
		return
	}

	refundFee := int64(0)

	traditionalMedicalFee := int64(0)
	westernMedicineFee := int64(0)
	examinationFee := int64(0)
	labortoryFee := int64(0)
	treatmentFee := int64(0)
	diagnosisTreatmentFee := int64(0)
	materialFee := int64(0)
	otherFee := int64(0)

	for _, id := range results {
		row := model.DB.QueryRowx(`select * from mz_paid_orders where id = $1 and mz_paid_record_id = $2`, id, paidRecord["id"])
		rowMap := FormatSQLRowToMap(row)

		if rowMap["id"] == nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": "存在未知收费项"})
			return
		}

		orderStatus := rowMap["order_status"].(string)
		refundStatus := rowMap["refund_status"].(bool)
		projectType := rowMap["charge_project_type_id"].(int64)
		fee := rowMap["fee"].(int64)
		if refundStatus == true {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": "存在已退费的项目：" + rowMap["name"].(string)})
			return
		}
		if orderStatus != "10" && orderStatus != "40" {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": "存在无法退费的项目：" + rowMap["name"].(string)})
			return
		}
		tx.Exec("update mz_paid_orders set refund_status=true,refund_id=$1 where id=$2", operarton, rowMap["id"])
		refundFee += fee

		switch projectType {
		case 1:
			westernMedicineFee += fee * -1
		case 2:
			traditionalMedicalFee += fee * -1
		case 3:
			labortoryFee += fee * -1
		case 4:
			examinationFee += fee * -1
		case 5:
			materialFee += fee * -1
		case 6:
			otherFee += fee * -1
		case 7:
			treatmentFee += fee * -1
		case 8:
			diagnosisTreatmentFee += fee * -1
		}

	}

	totalFefundFee := refundFee * -1

	// 部分退款
	if refundFee != totalMoney {
		// 存在优惠
		if balanceMoney != totalMoney {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": "存在优惠的情况下 无法部分退费"})
			return
		}
	} else {
		if balanceMoney != totalMoney {
			refundFee = balanceMoney
		}
	}

	outRefundNo := GetTradeNo("R2")

	refundNo, refundErr := refundNotice(outTradeNo, outRefundNo, refundFee)
	if refundErr != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": refundErr.Error()})
		return
	}

	var refundID int
	er := tx.QueryRow(`insert into mz_refund_record (mz_paid_record_id,out_refund_no,refund_no,status,orders_ids,refund_money,operation_id) values 
	($1,$2,$3,$4,$5,$6,$7) RETURNING id`, paidRecord["id"], outRefundNo, refundNo, 2, strings.Join(results, ","), refundFee*-1, operarton).Scan(&refundID)
	if er != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": er.Error()})
		return
	}

	cash := int64(0)
	wechat := int64(0)
	alipay := int64(0)
	bank := int64(0)

	switch paidRecord["pay_method_code"].(string) {
	case "3":
		bank = refundFee * -1
	case "4":
		cash = refundFee * -1
	case "2":
		wechat = refundFee * -1
	case "1":
		alipay = refundFee * -1
	}

	_, inserErr := tx.Exec(`insert into charge_detail 
	(
		pay_record_id,record_type,out_trade_no,out_refund_no,in_out,clinic_patient_id,department_id,doctor_id,traditional_medical_fee,western_medicine_fee,
		examination_fee,labortory_fee,treatment_fee,diagnosis_treatment_fee,material_fee,other_fee,discount_money,derate_money,medical_money,bonus_points_money,
		on_credit_money,total_money,balance_money,operation_id,cash,wechat,alipay,bank) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28)`,
		refundID, 1, paidRecord["out_trade_no"], outRefundNo, "out", triage["clinic_patient_id"], triage["department_id"], triage["doctor_id"], traditionalMedicalFee, westernMedicineFee,
		examinationFee, labortoryFee, treatmentFee, diagnosisTreatmentFee, materialFee, otherFee, paidRecord["discount_money"], paidRecord["derate_money"], paidRecord["medical_money"], paidRecord["bonus_points_money"],
		paidRecord["on_credit_money"], totalFefundFee, refundFee*-1, operarton, cash, wechat, alipay, bank)

	if inserErr != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": inserErr.Error()})
		return
	}

	cerr := tx.Commit()
	if cerr != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": cerr.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "msg": ""})

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

	if outTradeNo == "" || tradeNo == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	err := charge(outTradeNo, tradeNo)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "缴费成功"})
	return
}

// 处理缴费通知
func charge(outTradeNo string, tradeNo string) error {
	payment := model.DB.QueryRowx("select * from mz_paid_record where out_trade_no = $1", outTradeNo)
	pay := FormatSQLRowToMap(payment)
	_, ok := pay["id"]
	if !ok {
		return errors.New("未找到指定的待缴费单")
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
		  clinic_triage_patient_id, on_credit_money, trade_no, operation_id, type, remark) 
			VALUES ($1, $2, $3, $4, $5, $6)`
		_, creditErr := tx.Exec(creditSQL, triageID, onCreditMoney, tradeNo, operationID, 0, "门诊缴费挂账")
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

	//更新材料费库存
	materialsql := `select charge_project_id,amount from mz_unpaid_orders where id in (` + pay["orders_ids"].(string) + `) and charge_project_type_id=5`
	materialrows, _ := model.DB.Queryx(materialsql)
	materials := FormatSQLRowsToMapArray(materialrows)

	for _, material := range materials {
		updateMaterialStock(tx, material["charge_project_id"].(int64), material["amount"].(int64))
	}

	//插入交易明细表
	triagePatient := model.DB.QueryRowx("select * from clinic_triage_patient where id = $1", pay["clinic_triage_patient_id"])
	triage := FormatSQLRowToMap(triagePatient)

	typesql := `select sum(fee) as type_charge_total, charge_project_type_id from mz_unpaid_orders where id in (` + pay["orders_ids"].(string) + `) group by charge_project_type_id`
	typetotal, _ := model.DB.Queryx(typesql)
	typearray := FormatSQLRowsToMapArray(typetotal)

	typeMoney := map[string]interface{}{}

	for _, item := range typearray {
		typeMoney[strconv.FormatInt(item["charge_project_type_id"].(int64), 10)] = item["type_charge_total"]
	}

	for index := 0; index < 9; index++ {
		i := strconv.Itoa(index)
		if typeMoney[i] == nil {
			typeMoney[i] = 0
		}
	}

	cash := int64(0)
	wechat := int64(0)
	alipay := int64(0)
	bank := int64(0)
	payCode := pay["pay_method_code"]
	switch payCode.(string) {
	case "1":
		alipay = pay["balance_money"].(int64)
	case "2":
		wechat = pay["balance_money"].(int64)
	case "3":
		bank = pay["balance_money"].(int64)
	case "4":
		cash = pay["balance_money"].(int64)
	}

	insertDetails := `insert into charge_detail (pay_record_id,record_type,out_trade_no,in_out,clinic_patient_id,department_id,doctor_id,
		traditional_medical_fee,western_medicine_fee,examination_fee,labortory_fee,treatment_fee,diagnosis_treatment_fee,
		material_fee,retail_fee,other_fee,discount_money,derate_money,medical_money,voucher_money,bonus_points_money,
		on_credit_money,total_money,balance_money,operation_id,cash,wechat,alipay,bank) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29)`
	_, insertDetailErr := tx.Exec(insertDetails, pay["id"], 1, outTradeNo, "in", triage["clinic_patient_id"], triage["department_id"], triage["doctor_id"], typeMoney["2"], typeMoney["1"], typeMoney["4"], typeMoney["3"], typeMoney["7"], typeMoney["8"], typeMoney["5"], 0, typeMoney["6"],
		pay["discount_money"], pay["derate_money"], pay["medical_money"], pay["voucher_money"], pay["bonus_points_money"], pay["on_credit_money"], pay["total_money"], pay["balance_money"], confrimID, cash, wechat, alipay, bank)
	if insertDetailErr != nil {
		tx.Rollback()
		return insertDetailErr
	}

	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func refundNotice(outTradeNo string, outRefundNo string, refundFee int64) (string, error) {
	payment := model.DB.QueryRowx("select * from mz_paid_record where out_trade_no = $1", outTradeNo)
	pay := FormatSQLRowToMap(payment)
	_, ok := pay["id"]
	if !ok {
		return "", errors.New("未找到指定的待缴费单")
	}

	refundNo := outRefundNo

	// --支付方式编码，1-支付宝，2-微信, 3-银行卡, 4-现金

	methodCode := pay["pay_method_code"].(string)
	if methodCode == "1" || methodCode == "2" {

		res := HcRefund(outTradeNo, strconv.FormatInt(refundFee, 10), outRefundNo, "", "门诊退费")
		if res["code"] != "200" {
			return "", errors.New(res["msg"].(string))
		}

		data := res["data"].(map[string]interface{})
		refundNo = data["refund_id"].(string)
	}

	if methodCode == "3" {
		return "", errors.New("不支持的退费方式")
	}

	_, err := model.DB.Exec("update mz_paid_record set refund_money = $1+refund_money where id = $2", refundFee*-1, pay["id"])
	return refundNo, err
}

//updateMaterialStock
func updateMaterialStock(tx *sqlx.Tx, clinicMaterialID int64, amount int64) error {
	if amount < 0 {
		return errors.New("库存数量有误")
	}
	if amount == 0 {
		return nil
	}

	timeNow := time.Now().Format("2006-01-02")
	row := model.DB.QueryRowx("select * from material_stock where clinic_material_id = $1 and stock_amount > 0 and eff_date > $2 ORDER by created_time asc limit 1", clinicMaterialID, timeNow)
	rowMap := FormatSQLRowToMap(row)
	if rowMap["stock_amount"] == nil {
		return errors.New("库存不足")
	}

	stockAmount := rowMap["stock_amount"].(int64)

	if stockAmount >= amount {
		_, err := tx.Exec("update material_stock set stock_amount = $1 where id = $2", stockAmount-amount, rowMap["id"])
		if err != nil {
			tx.Rollback()
			return err
		}

		return nil
	}
	_, err := tx.Exec("update material_stock set 0 where id = $1", rowMap["id"])
	if err != nil {
		tx.Rollback()
		return err
	}

	return updateMaterialStock(tx, clinicMaterialID, amount-stockAmount)
}

// BusinessTransaction 获取交易流水
func BusinessTransaction(ctx iris.Context) {
	oprationName := ctx.PostValue("oprationName")
	patientName := ctx.PostValue("patientName")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")

	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}

	if startDate == "" || endDate == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请输入正确的时间范围"})
		return
	}

	queryMap := map[string]interface{}{
		"oprationName": ToNullString(oprationName),
		"patientName":  ToNullString(patientName),
		"offset":       ToNullInt64(offset),
		"limit":        ToNullInt64(limit),
		"startDate":    ToNullString(startDate),
		"endDate":      ToNullString(endDate),
	}

	sql := `FROM charge_detail cd
	left join personnel p on p.id = cd.operation_id  
	left join department d on d.id = cd.department_id 
	left join clinic_patient cp on cp.id = cd.clinic_patient_id 
	left join patient pa on pa.id = cp.patient_id 
	left join personnel doc on doc.id = cd.doctor_id 
	where cd.created_time BETWEEN :startDate and :endDate `

	if patientName != "" {
		sql += ` and pa.name ~:patientName `
	}

	if oprationName != "" {
		sql += ` and p.name ~:oprationName `
	}

	rows, err1 := model.DB.NamedQuery("SELECT cp.patient_id,cd.*,p.name as operation,d.name as departmentName, pa.name as patientName, cp.id as pid, doc.name as doctorName "+sql+" offset :offset limit :limit", queryMap)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	total, err2 := model.DB.NamedQuery(`SELECT COUNT (*) as total,
	sum(wechat) as wechat,sum(cash) as cash,
	sum(total_money) as total_money,sum(balance_money) as balance_money,
	sum(bank) as bank,sum(alipay) as alipay,sum(discount_money) as discount_money,
	sum(derate_money) as derate_money, sum(medical_money) as medical_money,
	sum(on_credit_money) as on_credit_money,sum(voucher_money) as voucher_money,
	sum(bonus_points_money) as bonus_points_money,sum(traditional_medical_fee) as traditional_medical_fee,
	sum(western_medicine_fee) as western_medicine_fee,sum(examination_fee) as examination_fee,
	sum(labortory_fee) as labortory_fee, sum(treatment_fee) as treatment_fee,sum(diagnosis_treatment_fee) as diagnosis_treatment_fee,
	sum(material_fee) as material_fee, sum(retail_fee) as retail_fee,sum(other_fee) as other_fee
	`+sql, queryMap)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit
	results := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})

}

// BusinessTransactionAnalysis 获取分析类交易
func BusinessTransactionAnalysis(ctx iris.Context) {
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")

	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}

	if startDate == "" || endDate == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请输入正确的时间范围"})
		return
	}

	queryMap := map[string]interface{}{
		"offset":    ToNullInt64(offset),
		"limit":     ToNullInt64(limit),
		"startDate": ToNullString(startDate),
		"endDate":   ToNullString(endDate),
	}

	sql := `FROM charge_detail where created_time BETWEEN :startDate and :endDate `

	rows, err1 := model.DB.NamedQuery(`SELECT to_char(created_time, 'YYYY-MM-DD') as created_time,
	sum(wechat) as wechat,sum(cash) as cash,
	sum(total_money) as total_money,sum(balance_money) as balance_money,
	sum(bank) as bank,sum(alipay) as alipay,sum(discount_money) as discount_money,
	sum(derate_money) as derate_money, sum(medical_money) as medical_money,
	sum(on_credit_money) as on_credit_money,sum(voucher_money) as voucher_money,
	sum(bonus_points_money) as bonus_points_money,sum(traditional_medical_fee) as traditional_medical_fee,
	sum(western_medicine_fee) as western_medicine_fee,sum(examination_fee) as examination_fee,
	sum(labortory_fee) as labortory_fee, sum(treatment_fee) as treatment_fee,sum(diagnosis_treatment_fee) as diagnosis_treatment_fee,
	sum(material_fee) as material_fee, sum(retail_fee) as retail_fee,sum(other_fee) as other_fee `+sql+" GROUP by to_char(created_time, 'YYYY-MM-DD') order by created_time ASC offset :offset limit :limit", queryMap)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	total, err2 := model.DB.NamedQuery(`SELECT COUNT (*) as total,
	sum(wechat) as wechat,sum(cash) as cash,
	sum(total_money) as total_money,sum(balance_money) as balance_money,
	sum(bank) as bank,sum(alipay) as alipay,sum(discount_money) as discount_money,
	sum(derate_money) as derate_money, sum(medical_money) as medical_money,
	sum(on_credit_money) as on_credit_money,sum(voucher_money) as voucher_money,
	sum(bonus_points_money) as bonus_points_money,sum(traditional_medical_fee) as traditional_medical_fee,
	sum(western_medicine_fee) as western_medicine_fee,sum(examination_fee) as examination_fee,
	sum(labortory_fee) as labortory_fee, sum(treatment_fee) as treatment_fee,sum(diagnosis_treatment_fee) as diagnosis_treatment_fee,
	sum(material_fee) as material_fee, sum(retail_fee) as retail_fee,sum(other_fee) as other_fee
	`+sql, queryMap)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit
	results := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})

}

// BusinessTransactionMonth 获取交易流水月报表
func BusinessTransactionMonth(ctx iris.Context) {
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")

	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}

	if startDate == "" || endDate == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请输入正确的时间范围"})
		return
	}

	queryMap := map[string]interface{}{
		"offset":    ToNullInt64(offset),
		"limit":     ToNullInt64(limit),
		"startDate": ToNullString(startDate),
		"endDate":   ToNullString(endDate),
	}

	sql := `from charge_detail where created_time BETWEEN :startDate and :endDate `

	rows, err1 := model.DB.NamedQuery(`select count(DISTINCT clinic_patient_id) as total,
	to_char(created_time, 'YYYY-MM-DD') as date,
	sum(total_money) as total_money,sum(balance_money) as balance_money,
	sum(wechat) as wechat,sum(cash) as cash,
	sum(bank) as bank,sum(alipay) as alipay,
	sum(discount_money) as discount_money,
	sum(derate_money) as derate_money, sum(medical_money) as medical_money,
	sum(on_credit_money) as on_credit_money,sum(voucher_money) as voucher_money,
	sum(bonus_points_money) as bonus_points_money 
	`+sql+" GROUP by to_char(created_time, 'YYYY-MM-DD') order by date ASC offset :offset limit :limit", queryMap)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	total, err2 := model.DB.NamedQuery(` SELECT count(distinct to_char(created_time, 'YYYY-MM-DD')) as total,
	count(DISTINCT clinic_patient_id) as people_count,
	sum(wechat) as wechat,sum(cash) as cash, 
	sum(total_money) as total_money,sum(balance_money) as balance_money,
	sum(bank) as bank,sum(alipay) as alipay,sum(discount_money) as discount_money,
	sum(derate_money) as derate_money, sum(medical_money) as medical_money,
	sum(on_credit_money) as on_credit_money,sum(voucher_money) as voucher_money,
	sum(bonus_points_money) as bonus_points_money 
	`+sql, queryMap)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-2", "msg": err2.Error()})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit
	results := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

// BusinessTransactionDetail 获取交易详情
func BusinessTransactionDetail(ctx iris.Context) {
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")
	patientName := ctx.PostValue("patientName")
	phone := ctx.PostValue("phone")
	porjectName := ctx.PostValue("porjectName")
	io := ctx.PostValue("in_out") // 退费或收费

	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}

	if startDate == "" || endDate == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请输入正确的时间范围"})
		return
	}

	queryMap := map[string]interface{}{
		"offset":      ToNullInt64(offset),
		"limit":       ToNullInt64(limit),
		"startDate":   ToNullString(startDate),
		"endDate":     ToNullString(endDate),
		"patientName": ToNullString(patientName),
		"phone":       ToNullString(phone),
		"porjectName": ToNullString(porjectName),
		"io":          ToNullString(io),
	}

	sql := `from charge_detail cd left join mz_paid_orders mpo on mpo.mz_paid_record_id = cd.pay_record_id and cd.record_type = 1 
	left join personnel ope on cd.operation_id = ope.id 
	left join charge_project_type cpt on cpt.id = mpo.charge_project_type_id 
	left join clinic_triage_patient ctp on mpo.clinic_triage_patient_id = ctp.id 
	left join personnel doc on doc.id = cd.doctor_id 
	left join department d on d.id = cd.department_id 
	left join clinic_patient cp on cp.id = cd.clinic_patient_id 
	left join patient pa on pa.id = cp.patient_id 
	left join drug_retail dr on (dr.out_trade_no = cd.out_trade_no and cd.in_out = 'in' and dr.amount > 0) or (cd.out_refund_no = dr.out_refund_no and cd.in_out = 'out') 
	left join clinic_drug drug on drug.id = dr.clinic_drug_id 
	where cd.created_time BETWEEN :startDate and :endDate `

	if patientName != "" {
		sql += " and pa.name ~:patientName "
	}

	if phone != "" {
		sql += " and pa.phone ~:phone "
	}

	if porjectName != "" {
		sql += " and (mpo.name ~:porjectName or drug.name ~:porjectName)"
	}

	if io != "" {
		sql += " and cd.in_out ~:io "
	}

	rows, err1 := model.DB.NamedQuery(`select 
		drug.packing_unit_name as drug_unit,drug.name as drug_name,drug.ret_price as drug_price,
		dr.amount as drug_mount, dr.total_fee as drug_total,
	  mpo.name,mpo.price,mpo.amount,
		mpo.total,mpo.fee,mpo.unit,ope.name as operarion,cpt.name as charge_project_type,
		cd.created_time,cd.out_trade_no,cd.record_type,ctp.visit_date,doc.name as doctorName,d.name as deptName,
		pa.name as patientName,pa.sex,pa.birthday,pa.phone 
		 `+sql+" order by cd.created_time ASC offset :offset limit :limit", queryMap)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	total, err2 := model.DB.NamedQuery(`SELECT COUNT (*) as total, sum(dr.total_fee)+sum(mpo.total) as total_fee, sum(dr.total_fee)+sum(mpo.fee) as banance_fee `+sql, queryMap)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-2", "msg": err2.Error()})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit
	results := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

// BusinessTransactionCredit 获取挂账详情
func BusinessTransactionCredit(ctx iris.Context) {
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")
	keyword := ctx.PostValue("keyword")

	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}

	if startDate == "" || endDate == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请输入正确的时间范围"})
		return
	}

	queryMap := map[string]interface{}{
		"offset":    ToNullInt64(offset),
		"limit":     ToNullInt64(limit),
		"startDate": ToNullString(startDate),
		"endDate":   ToNullString(endDate),
		"keyword":   ToNullString(keyword),
	}

	sql := `from on_credit_record ocr 
	left join personnel ope on ocr.operation_id = ope.id 
	left join clinic_triage_patient ctp on ocr.clinic_triage_patient_id = ctp.id 
	left join clinic_patient cp on cp.id = ctp.clinic_patient_id 
	left join patient pa on pa.id = cp.patient_id 
	where ocr.created_time BETWEEN :startDate and :endDate `

	if keyword != "" {
		sql += " and (pa.name ~:keyword or pa.phone ~:keyword or cp.id ~:keyword) "
	}

	rows, err1 := model.DB.NamedQuery(`select ocr.on_credit_money,ocr.created_time,
		ope.name as operation,pa.name as patientname,pa.sex,pa.phone,cp.id as pid 
		 `+sql+"order by ocr.created_time ASC offset :offset limit :limit", queryMap)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	total, err2 := model.DB.NamedQuery(`SELECT COUNT (*) as total `+sql, queryMap)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-2", "msg": err2.Error()})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit
	results := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})

}

// PatientChargeList 患者支付列表
func PatientChargeList(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if patientID == "" {
		ctx.JSON(iris.Map{"code": "-2", "msg": "参数错误"})
		return
	}

	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}

	countSQL := `select count(*) as total from charge_detail cd 
	left join clinic_patient cp on cd.clinic_patient_id = cp.id
	left join clinic c on cp.clinic_id = c.id
	left join mz_paid_record mpr on cd.out_trade_no = mpr.out_trade_no
	left join clinic_triage_patient_operation ctpo on mpr.clinic_triage_patient_id = ctpo.clinic_triage_patient_id and type = '30' and times = 1
	where cp.patient_id = $1`

	total := model.DB.QueryRowx(countSQL, patientID)
	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	querySQL := `select cd.*, c.name as clinic_name, ctpo.created_time as visit_time, mpr.id as mz_paid_record_id from charge_detail cd 
	left join clinic_patient cp on cd.clinic_patient_id = cp.id
	left join clinic c on cp.clinic_id = c.id
	left join mz_paid_record mpr on cd.out_trade_no = mpr.out_trade_no
	left join clinic_triage_patient_operation ctpo on mpr.clinic_triage_patient_id = ctpo.clinic_triage_patient_id and type = '30' and times = 1
	where cp.patient_id = $1 order by id desc offset $2 limit $3`
	rows, err := model.DB.Queryx(querySQL, patientID, offset, limit)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-2", "msg": err.Error()})
		return
	}
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}
