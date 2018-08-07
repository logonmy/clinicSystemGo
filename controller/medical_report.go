package controller

import (
	"clinicSystemGo/model"

	"github.com/kataras/iris"
)

// ReceiveTreatment 医生接诊统计
func ReceiveTreatment(ctx iris.Context) {

}

// OutPatietnRecords 门诊记录
func OutPatietnRecords(ctx iris.Context) {
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")
	patientName := ctx.PostValue("patient_name")
	phone := ctx.PostValue("phone")
	doctorID := ctx.PostValue("doctor_id")
	operationID := ctx.PostValue("operation_id")
	clinicID := ctx.PostValue("clinic_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if startDate == "" || endDate == "" || clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if offset == "" {
		offset = "0"
	}

	if limit == "" {
		limit = "10"
	}

	var queryOptions = map[string]interface{}{
		"phone":       ToNullString(phone),
		"patientName": ToNullString(patientName),
		"operationID": ToNullInt64(operationID),
		"clinicID":    ToNullInt64(clinicID),
		"doctorID":    ToNullInt64(doctorID),
		"offset":      ToNullInt64(offset),
		"limit":       ToNullInt64(limit),
		"startDate":   ToNullString(startDate),
		"endDate":     ToNullString(endDate),
	}

	SQL := ` from clinic_triage_patient ctp 
	left join clinic_patient cp on cp.id = ctp.clinic_patient_id 
	left join patient p on p.id = cp.patient_id 
	left join medical_record mr on mr.clinic_triage_patient_id = ctp.id and mr.is_default = true 
	left join clinic_triage_patient_operation ctpo on ctpo.clinic_triage_patient_id = ctp.id and type = 10 
	left join personnel ctpop on ctpo.personnel_id = ctpop.id
	left join department dept on dept.id = ctp.department_id 
	left join personnel doc on doc.id = ctp.doctor_id where cp.clinic_id = :clinicID and ctp.visit_date between :startDate and :endDate and ctp.status = 40 
	`

	if patientName != "" {
		SQL += ` and p.name ~:patientName `
	}

	if phone != "" {
		SQL += ` and p.phone ~:phone `
	}

	if operationID != "" {
		SQL += ` and ctpop.id = :operationID `
	}

	if doctorID != "" {
		SQL += ` and doc.id = :doctorID`
	}

	querySQL := `select ctp.visit_date,p.name as patient_name,cp.patient_id,
	p.sex,p.birthday,p.phone,p.profession,p.province,p.city,p.district,p.address,
	mr.morbidity_date,mr.diagnosis,ctp.visit_type,ctpop.name as opreation_name,
	doc.name as doctor_name,dept.name as dept_name
	` + SQL + ` order by ctp.visit_date ASC offset :offset limit :limit`

	countSQL := `select count(*) as total, count(DISTINCT cp.patient_id) as person_amount ` + SQL

	items, _ := model.DB.NamedQuery(querySQL, queryOptions)
	itemMap := FormatSQLRowsToMapArray(items)

	count, _ := model.DB.NamedQuery(countSQL, queryOptions)

	pageInfo := FormatSQLRowsToMapArray(count)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	ctx.JSON(iris.Map{"code": "200", "data": itemMap, "page_info": pageInfo})

}

// OutPatietnType 门诊日志
func OutPatietnType(ctx iris.Context) {
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	clinicID := ctx.PostValue("clinic_id")

	if startDate == "" || endDate == "" || clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if offset == "" {
		offset = "0"
	}

	if limit == "" {
		limit = "10"
	}

	var queryOptions = map[string]interface{}{
		"startDate": ToNullString(startDate),
		"endDate":   ToNullString(endDate),
		"offset":    ToNullInt64(offset),
		"limit":     ToNullInt64(limit),
	}

	querySQL := `select ctpo.created_time FROM clinic_triage_patient ctp 
	left join clinic_triage_patient_operation ctpo on ctpo.clinic_triage_patient_id = ctp.id and type = 30 
	where ctp.status = 40 and ctpo.created_time between :startDate and :endDate group by ctpo.created_time 
	order by ctpo.created_time ASC offset :offset limit : limit
	`

	items, _ := model.DB.NamedQuery(querySQL, queryOptions)
	itemMap := FormatSQLRowsToMapArray(items)

	ctx.JSON(iris.Map{"code": "200", "data": itemMap})

}
