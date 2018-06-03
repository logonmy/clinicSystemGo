package controller

import (
	"clinicSystemGo/model"
	"strconv"

	"github.com/kataras/iris"
)

// LaboratoryTriageList 获取检验记录（包括 待检验，已检验，检验中）
func LaboratoryTriageList(ctx iris.Context) {
	status := ctx.PostValue("order_status")
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if clinicTriagePatientID == "" || status == "" {
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

	SQL := `FROM laboratory_patient lp 
	left join clinic_laboratory cl on cl.id = lp.clinic_laboratory_id
	left join 
	where lp.clinic_triage_patient_id = $1 and lp.order_status = $2 and lp.charge_project_type_id = 3`
	countsql := "select count(lp.*) as total " + SQL

	total := model.DB.QueryRowx(countsql, clinicTriagePatientID, status)
	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit
	querysql := `select lp.id,lp.name,lp.amount,lp.charge_project_type_id,
	cl.specification,cl.manu_factory_name,cl.dose_form_name,ds.stock_amount ` + SQL + `offset $3 limit $4`

	rows, _ := model.DB.Queryx(querysql, clinicTriagePatientID, status, offset, limit)
	results := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

// LaboratoryTriageWaiting 获取待检验的分诊记录
func LaboratoryTriageWaiting(ctx iris.Context) {
	laboratoryTriageList(ctx, "10")
}

// LaboratoryTriageChecked 获取已检验的分诊记录
func LaboratoryTriageChecked(ctx iris.Context) {
	laboratoryTriageList(ctx, "30")
}

// LaboratoryTriageChecking 获取检验中的分诊记录
func LaboratoryTriageChecking(ctx iris.Context) {
	laboratoryTriageList(ctx, "20")
}

// 获取各状态的分诊记录
func laboratoryTriageList(ctx iris.Context, status string) {
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
	left join medical_record mr on mr.clinic_triage_patient_id = ctp.id
	left join clinic_triage_patient_operation register on ctp.id = register.clinic_triage_patient_id and register.type = 10
	left join personnel triage_personnel on triage_personnel.id = register.personnel_id 
	left join (select clinic_triage_patient_id,count(*) as total_count from laboratory_patient where order_status = $1 group by(clinic_triage_patient_id)) up on up.clinic_triage_patient_id = ctp.id 
	where up.total_count > 0 AND cp.clinic_id=$2 AND ctp.updated_time BETWEEN $3 and $4 AND (p.name ~$5 OR p.cert_no ~$5 OR p.phone ~$5) `

	countsql := `select count(*) as total` + sql
	querysql := `select 
	up.total_count,
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
	doc.name as doctor_name,
	d.name as department_name,
	mr.diagnosis,
	mr.allergic_history ` + sql + `offset $6 limit $7`

	total := model.DB.QueryRowx(countsql, status, clinicID, startDate, endDate, keyword)

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.Queryx(querysql, status, clinicID, startDate, endDate, keyword, offset, limit)
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

//LaboratoryTriageReportSave 保存检验报告
func LaboratoryTriageReportSave(ctx iris.Context) {
	// clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
}
