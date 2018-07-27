package controller

import (
	"clinicSystemGo/model"

	"strconv"
	"time"

	"github.com/kataras/iris"
)

// OnCreditRepay 挂账还款
func OnCreditRepay(ctx iris.Context) {

	recordID := ctx.PostValue("on_credit_record_id")
	money, e := ctx.PostValueInt64("money")
	operation := ctx.PostValue("operation_id")
	outTradeNo := ctx.PostValue("out_trade_no")
	payCode := ctx.PostValue("pay_method_code")

	if e != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "输入金额无效，请修改后重新提交"})
		return
	}

	if recordID == "" || money == 0 || operation == "" || payCode == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
		return
	}

	row := model.DB.QueryRowx(`select * from on_credit_record where id = $1`, recordID)
	record := FormatSQLRowToMap(row)

	if record["id"] == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "未找到指定的挂费记录"})
		return
	}

	registrationRow := model.DB.QueryRowx(`select r.*,cp.patient_id from registration r left join clinic_patient cp on r.clinic_patient_id = cp.id where r.id = $1`, record["registration_id"])
	registration := FormatSQLRowToMap(registrationRow)

	if record["remain_pay_money"].(int64) == 0 {
		ctx.JSON(iris.Map{"code": "2", "msg": "该挂账记录已还清"})
		return
	}

	if record["remain_pay_money"].(int64) < money {
		ctx.JSON(iris.Map{"code": "3", "msg": "缴费金额超过待缴费金额, 请重新确认"})
		return
	}

	tx, err := model.DB.Beginx()

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	detailSQL := "INSERT INTO on_credit_record_detail (on_credit_record_id, type, should_repay_moeny,repay_moeny,remain_repay_moeny,operation_id,pay_method_code) VALUES ($1,$2,$3,$4,$5,$6,$7) returning id"

	var recordDetailID int
	err2 := tx.QueryRow(detailSQL, record["id"], 1, record["remain_pay_money"], money, record["remain_pay_money"].(int64)-money, operation, payCode).Scan(&recordDetailID)
	if err2 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "2", "msg": err2.Error()})
		return
	}

	chargeSQL := `INSERT INTO charge_detail(
		pay_record_id, out_trade_no, in_out, patient_id, department_id, doctor_id, pay_type_code, pay_type_code_name, pay_method_code, pay_method_code_name,on_credit_money, total_money, balance_money)
	 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	_, chargeDetialErr := tx.Exec(chargeSQL, recordDetailID, outTradeNo, "in", registration["patient_id"], registration["department_id"], registration["personnel_id"], "03", "挂账还款", payCode, "", money*-1, 0, money)
	if chargeDetialErr != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "2", "msg": chargeDetialErr.Error()})
		return
	}

	_, err3 := tx.Exec(`UPDATE on_credit_record SET remain_pay_money=$1,already_pay_money=$2,updated_time=$3 where id=$4`, record["remain_pay_money"].(int64)-money, record["already_pay_money"].(int64)+money, time.Now(), recordID)
	if err3 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "3", "msg": err3.Error()})
		return
	}

	cmerr := tx.Commit()
	if cmerr != nil {
		ctx.JSON(iris.Map{"code": "3", "msg": cmerr.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})

}

// OnCreditTraigePatient  查询有挂账的就诊记录
func OnCreditTraigePatient(ctx iris.Context) {
	keyword := ctx.PostValue("keyword")
	clinicID := ctx.PostValue("clinic_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")
	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
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

	sql := ` from clinic_triage_patient ctp 
	left join clinic_patient cp on cp.id = ctp.clinic_patient_id 
	left join personnel doc on doc.id = ctp.doctor_id 
	left join department d on d.id = ctp.department_id  
	left join patient p on p.id = cp.patient_id 
	left join clinic_triage_patient_operation register on ctp.id = register.clinic_triage_patient_id and register.type = 10
	left join personnel triage_personnel on triage_personnel.id = register.personnel_id 
	left join (select clinic_triage_patient_id,sum(on_credit_money) - sum(already_pay_money) as charge_total_fee from on_credit_record group by(clinic_triage_patient_id)) up on up.clinic_triage_patient_id = ctp.id 
	where up.charge_total_fee > 0 AND cp.clinic_id=$1 AND ctp.updated_time BETWEEN $2 and $3 AND (p.name ~$4 OR p.cert_no ~$4 OR p.phone ~$4) `

	countsql := `select count(*) as total` + sql
	querysql := `select 
	up.charge_total_fee,
	ctp.id as clinic_triage_patient_id,
	ctp.clinic_patient_id as clinic_patient_id,
	ctp.updated_time,
	ctp.created_time as register_time,
	triage_personnel.name as register_personnel_name,
	ctp.status,
	ctp.visit_date,
	ctp.register_type,
	p.name as patient_name,
	p.birthday,
	p.sex,
	p.phone,
	cp.patient_id,
	doc.name as doctor_name,
	d.name as department_name ` + sql + ` offset $5 limit $6`

	total := model.DB.QueryRowx(countsql, clinicID, startDate, endDate, keyword)

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.Queryx(querysql, clinicID, startDate, endDate, keyword, offset, limit)
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

// OnCreditList 获取挂账记录
func OnCreditList(ctx iris.Context) {
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

	total := model.DB.QueryRowx(`select count(id) as total,sum(on_credit_money) as on_credit_money_total,sum(already_pay_money) as already_pay_money_total from on_credit_record where clinic_triage_patient_id=$1`, clinicTriagePatientid)
	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rowSQL := `select 
	ocr.id as on_credit_record_id,
	ocr.on_credit_money,ocr.already_pay_money,ocr.created_time,
	ctp.visit_date,
	p.name as doctor_name,
	c.name as clinic_name,
	d.name as department_name 
	from on_credit_record ocr
	left join clinic_triage_patient ctp on ctp.id = ocr.clinic_triage_patient_id 
	left join personnel p on p.id = ctp.doctor_id 
	left join department d on d.id = ctp.department_id 
	left join clinic c on c.id = d.clinic_id 
	where ocr.clinic_triage_patient_id=$1 order by ocr.already_pay_money ASC,ocr.created_time DESC offset $2 limit $3`

	rows, err1 := model.DB.Queryx(rowSQL, clinicTriagePatientid, offset, limit)

	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})

}
