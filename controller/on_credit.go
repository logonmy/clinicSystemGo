package controller

import (
	"clinicSystemGo/model"

	"time"

	"github.com/kataras/iris"
)

// OnCreditCreate 创建挂账记录
func OnCreditCreate(ctx iris.Context) {
	registration := ctx.PostValue("registration_id")
	money := ctx.PostValue("on_credit_money")
	operation := ctx.PostValue("operation_id")

	if registration == "" || money == "" || operation == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	tx, err := model.DB.Beginx()

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	var ID int
	recordSQL := "INSERT INTO on_credit_record (registration_id,on_credit_money,remain_pay_money,operation_id) VALUES(" + registration + "," + money + "," + money + "," + operation + ") RETURNING id"
	err = tx.QueryRow(recordSQL).Scan(&ID)
	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	detailSQL := "INSERT INTO on_credit_record_detail (on_credit_record_id, type, should_repay_moeny,repay_moeny,remain_repay_moeny,operation_id) VALUES ($1,$2,$3,$4,$5,$6)"

	_, err2 := tx.Exec(detailSQL, ID, 0, money, 0, money, operation)
	if err2 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "2", "msg": err2.Error()})
		return
	}

	cmerr := tx.Commit()
	if cmerr != nil {
		ctx.JSON(iris.Map{"code": "3", "msg": cmerr.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})

}

// OnCreditRegisttionList 获取有挂账的挂号记录
func OnCreditRegisttionList(ctx iris.Context) {
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")
	keyword := ctx.PostValue("keyword")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if offset == "" {
		offset = "0"
	}

	currentTime := time.Now().Local()

	if limit == "" {
		limit = "10"
	}

	if startDate == "" {
		startDate = currentTime.Format("2006-01-02")
	}

	if endDate == "" {
		endDate = currentTime.Add(24 * time.Hour).Format("2006-01-02")
	}
	selectSQL := `select p.name,p.sex,p.birthday,cr.updated_time,cr.remain_pay_money,cr.id as on_credit_record_id `
	countSQL := `select count(cr.id) as total `

	sql := `from on_credit_record cr 
	left join registration r on cr.registration_id = r.id 
	left join clinic_patient cp on r.clinic_patient_id = cp.id
	left join patient p on cp.patient_id = p.id
	where cr.created_time between $1 and $2 and (p.name ~$3 or p.phone ~$3 or p.cert_no ~$3)`

	rows, err := model.DB.Queryx(selectSQL+sql+" order by cr.updated_time DESC offset $4 limit $5", startDate, endDate, keyword, offset, limit)

	total := model.DB.QueryRowx(countSQL+sql, startDate, endDate, keyword)
	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	if err != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
		return
	}

	results := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "pageInfo": pageInfo})

}

// OnCreditRegisttionDetail 挂账详情
func OnCreditRegisttionDetail(ctx iris.Context) {
	recordID := ctx.PostValue("on_credit_record_id")

	if recordID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
		return
	}

	sql := `select crd.created_time,c.name as clinic_name,r.visit_date,d.name as departmentName,p.name as doctorName,crd.should_repay_moeny,crd.repay_moeny,crd.remain_repay_moeny from on_credit_record_detail crd 
	left join on_credit_record cr on cr.id = crd.on_credit_record_id 
	left join registration r on cr.registration_id = r.id 
	left join department d on d.id = r.department_id 
	left join personnel p on p.id = r.personnel_id 
	left join clinic c on c.id = p.clinic_id where on_credit_record_id = $1`

	rows, err := model.DB.Queryx(sql, recordID)

	if err != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
		return
	}

	results := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})

}

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
		ctx.JSON(iris.Map{"code": "1", "msg": "未找到指定的挂费记录"})
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
