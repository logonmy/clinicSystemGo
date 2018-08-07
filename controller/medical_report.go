package controller

import (
	"clinicSystemGo/model"
	"time"

	"github.com/kataras/iris"
)

// ReceiveTreatment 医生接诊统计
func ReceiveTreatment(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	departmentID := ctx.PostValue("department_id")
	startDateStr := ctx.PostValue("start_date")
	endDateStr := ctx.PostValue("end_date")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
		return
	}

	startDate, errs := time.Parse("2006-01-02", startDateStr)
	if errs != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "start_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}

	endDate, erre := time.Parse("2006-01-02", endDateStr)
	if erre != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "end_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}

	startDateStr = startDate.AddDate(0, 0, -1).Format("2006-01-02")
	endDateStr = endDate.AddDate(0, 0, 1).Format("2006-01-02")

	var queryOptions = map[string]interface{}{
		"clinic_id":     ToNullInt64(clinicID),
		"department_id": ToNullInt64(departmentID),
		"start_date":    ToNullString(startDateStr),
		"end_date":      ToNullString(endDateStr),
	}

	querySQL := `select 
	p.name as personnel_name,
	d.name as department_name,
	mpo.charge_project_type_id,
	sum(fee) as fee
	 from clinic_triage_patient ctp 
	left join department d on ctp.department_id = d.id
	left join personnel p on ctp.doctor_id = p.id 
	left join mz_paid_orders mpo on mpo.clinic_triage_patient_id = ctp.id 
	where ctp.doctor_id is not null and ctp.status >= 40 and mpo.charge_project_type_id is not null 
	and d.clinic_id = :clinic_id 
	and ctp.visit_date between :start_date and :end_date `

	if departmentID != "" {
		querySQL += ` and ctp.department_id = :department_id`
	}

	rows, err := model.DB.NamedQuery(querySQL+" group by (p.name, d.name, mpo.charge_project_type_id);", queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results})
}

// ExaminationStatistics 检查统计
func ExaminationStatistics(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	startDateStr := ctx.PostValue("start_date")
	endDateStr := ctx.PostValue("end_date")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
		return
	}

	startDate, errs := time.Parse("2006-01-02", startDateStr)
	if errs != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "start_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}

	endDate, erre := time.Parse("2006-01-02", endDateStr)
	if erre != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "end_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}

	startDateStr = startDate.AddDate(0, 0, -1).Format("2006-01-02")
	endDateStr = endDate.AddDate(0, 0, 1).Format("2006-01-02")

	var queryOptions = map[string]interface{}{
		"clinic_id":  ToNullInt64(clinicID),
		"start_date": ToNullString(startDateStr),
		"end_date":   ToNullString(endDateStr),
	}

	querySQL := `select 
	ep.created_time,
	ep.updated_time,
	p.id,
	p.name,
	p.sex,
	p.birthday,
	p.phone,
	ps.name as personnel_name,
	ce.name as clinic_examination_name,
	ep.times,
	ce.price
	from examination_patient ep 
 left join clinic_triage_patient ctp on ep.clinic_triage_patient_id = ctp.id
 left join clinic_patient cp on ctp.clinic_patient_id = cp.id 
 left join patient p on cp.patient_id = p.id 
 left join personnel ps on ps.id = ep.operation_id 
 left join clinic_examination ce on ce.id = ep.clinic_examination_id
 where ep.order_status = '30'
	and ps.clinic_id = :clinic_id 
	and ep.updated_time between :start_date and :end_date `

	rows, err := model.DB.NamedQuery(querySQL, queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results})
}

// LaboratoryStatistics 检验统计
func LaboratoryStatistics(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	startDateStr := ctx.PostValue("start_date")
	endDateStr := ctx.PostValue("end_date")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
		return
	}

	startDate, errs := time.Parse("2006-01-02", startDateStr)
	if errs != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "start_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}

	endDate, erre := time.Parse("2006-01-02", endDateStr)
	if erre != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "end_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}

	startDateStr = startDate.AddDate(0, 0, -1).Format("2006-01-02")
	endDateStr = endDate.AddDate(0, 0, 1).Format("2006-01-02")

	var queryOptions = map[string]interface{}{
		"clinic_id":  ToNullInt64(clinicID),
		"start_date": ToNullString(startDateStr),
		"end_date":   ToNullString(endDateStr),
	}

	querySQL := `
	select 
	 el.created_time,
	 el.updated_time,
	 p.id,
	 p.name,
	 p.sex,
	 p.birthday,
	 p.phone,
	 ps.name as personnel_name,
	 cl.name as clinic_laboratory_name,
	 el.times,
	 cl.price
	 from laboratory_patient el 
	left join clinic_triage_patient ctp on el.clinic_triage_patient_id = ctp.id
	left join clinic_patient cp on ctp.clinic_patient_id = cp.id 
	left join patient p on cp.patient_id = p.id 
	left join personnel ps on ps.id = el.operation_id 
	left join clinic_laboratory cl on cl.id = el.clinic_laboratory_id
	where el.order_status = '30' 
	and ps.clinic_id = :clinic_id 
	and el.updated_time between :start_date and :end_date `

	rows, err := model.DB.NamedQuery(querySQL, queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results})
}

// TreatmentStatistics 治疗统计
func TreatmentStatistics(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	startDateStr := ctx.PostValue("start_date")
	endDateStr := ctx.PostValue("end_date")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
		return
	}

	startDate, errs := time.Parse("2006-01-02", startDateStr)
	if errs != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "start_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}

	endDate, erre := time.Parse("2006-01-02", endDateStr)
	if erre != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "end_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}

	startDateStr = startDate.AddDate(0, 0, -1).Format("2006-01-02")
	endDateStr = endDate.AddDate(0, 0, 1).Format("2006-01-02")

	var queryOptions = map[string]interface{}{
		"clinic_id":  ToNullInt64(clinicID),
		"start_date": ToNullString(startDateStr),
		"end_date":   ToNullString(endDateStr),
	}

	querySQL := `
	select 
	tp.created_time,
	tp.updated_time,
	p.id,
	p.name,
	p.sex,
	p.birthday,
	p.phone,
	ps.name as personnel_name,
	ct.name as clinic_treatment_name,
	tp.times,
	ct.price
	from treatment_patient tp 
	left join clinic_triage_patient ctp on tp.clinic_triage_patient_id = ctp.id
	left join clinic_patient cp on ctp.clinic_patient_id = cp.id 
	left join patient p on cp.patient_id = p.id 
	left join personnel ps on ps.id = tp.operation_id 
	left join clinic_treatment ct on ct.id = tp.clinic_treatment_id
	where tp.order_status = '30' 
	and ps.clinic_id = :clinic_id 
	and tp.updated_time between :start_date and :end_date `

	rows, err := model.DB.NamedQuery(querySQL, queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results})
}

// RegisterStatistics 登记统计
func RegisterStatistics(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	startDateStr := ctx.PostValue("start_date")
	endDateStr := ctx.PostValue("end_date")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
		return
	}

	startDate, errs := time.Parse("2006-01-02", startDateStr)
	if errs != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "start_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}

	endDate, erre := time.Parse("2006-01-02", endDateStr)
	if erre != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "end_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}

	startDateStr = startDate.AddDate(0, 0, -1).Format("2006-01-02")
	endDateStr = endDate.AddDate(0, 0, 1).Format("2006-01-02")

	var queryOptions = map[string]interface{}{
		"clinic_id":  ToNullInt64(clinicID),
		"start_date": ToNullString(startDateStr),
		"end_date":   ToNullString(endDateStr),
	}

	querySQL := `select to_char(ctpo.created_time, 'YYYY-MM-DD') as visit_date, ctp.register_type, count(ctp.id) from clinic_triage_patient ctp
	left join clinic_triage_patient_operation ctpo on ctpo.clinic_triage_patient_id = ctp.id and times = 1 and type = 30 
	left join personnel p on p.id = ctp.doctor_id 
	where ctp.status >= 40 and p.clinic_id = :clinic_id 
	and ctpo.created_time between :start_date and :end_date `

	rows, err := model.DB.NamedQuery(querySQL+`
	group by (to_char(ctpo.created_time, 'YYYY-MM-DD'), ctp.register_type)
	order by visit_date DESC`, queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results})
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
