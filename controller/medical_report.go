package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"
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

	var queryOptions = map[string]interface{}{
		"clinic_id":     ToNullInt64(clinicID),
		"department_id": ToNullInt64(departmentID),
		"start_date":    ToNullString(startDateStr),
		"end_date":      ToNullString(endDateStr),
	}

	querySQL := `select 
	p.name as personnel_name,
	d.name as department_name,
	sum(fee) as total_fee,
	count(ctp.id) tatol_count,
	sum(case when mpo.charge_project_type_id = 1 then fee ELSE 0 end) as west_pre_fee,
	sum(case when mpo.charge_project_type_id = 2 then fee ELSE 0 end) as east_pre_fee,
	sum(case when mpo.charge_project_type_id = 3 then fee ELSE 0 end) as labora_pre_fee,
	sum(case when mpo.charge_project_type_id = 4 then fee ELSE 0 end) as exam_pre_fee,
	sum(case when mpo.charge_project_type_id = 5 then fee ELSE 0 end) as material_fee,
	sum(case when mpo.charge_project_type_id = 6 then fee ELSE 0 end) as other_fee,
	sum(case when mpo.charge_project_type_id = 7 then fee ELSE 0 end) as treatement_fee,
	sum(case when mpo.charge_project_type_id = 8 then fee ELSE 0 end) as diagnosis_treatment_fee
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

	rows, err := model.DB.NamedQuery(querySQL+" group by (p.name, d.name);", queryOptions)
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
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
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

	endDate, erre := time.Parse("2006-01-02", endDateStr)
	if erre != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "end_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}

	endDateStr = endDate.AddDate(0, 0, 1).Format("2006-01-02")

	var queryOptions = map[string]interface{}{
		"clinic_id":  ToNullInt64(clinicID),
		"start_date": ToNullString(startDateStr),
		"end_date":   ToNullString(endDateStr),
		"offset":     ToNullInt64(offset),
		"limit":      ToNullInt64(limit),
	}

	countSQL := `select count(*) as total from examination_patient ep 
 left join clinic_triage_patient ctp on ep.clinic_triage_patient_id = ctp.id
 left join clinic_patient cp on ctp.clinic_patient_id = cp.id 
 left join patient p on cp.patient_id = p.id 
 left join personnel ps on ps.id = ep.operation_id 
 left join clinic_examination ce on ce.id = ep.clinic_examination_id
 where ep.order_status = '30'
	and ps.clinic_id = :clinic_id 
	and ep.updated_time between :start_date and :end_date`
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

	patientCountSQL := `select 
	count(distinct p.name) as patient_count
	from examination_patient ep 
 left join clinic_triage_patient ctp on ep.clinic_triage_patient_id = ctp.id
 left join clinic_patient cp on ctp.clinic_patient_id = cp.id 
 left join patient p on cp.patient_id = p.id 
 left join personnel ps on ps.id = ep.operation_id 
 left join clinic_examination ce on ce.id = ep.clinic_examination_id
 where ep.order_status = '30'
	and ps.clinic_id = :clinic_id 
	and ep.updated_time between :start_date and :end_date`

	timesSQL := `select 
	sum(ep.times) as total_times
	from examination_patient ep 
 left join clinic_triage_patient ctp on ep.clinic_triage_patient_id = ctp.id
 left join clinic_patient cp on ctp.clinic_patient_id = cp.id 
 left join patient p on cp.patient_id = p.id 
 left join personnel ps on ps.id = ep.operation_id 
 left join clinic_examination ce on ce.id = ep.clinic_examination_id
 where ep.order_status = '30'
	and ps.clinic_id = :clinic_id 
	and ep.updated_time between :start_date and :end_date`

	patientCountRows, err := model.DB.NamedQuery(patientCountSQL, queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	patientCount := FormatSQLRowsToMapArray(patientCountRows)[0]["patient_count"]

	timesRows, err := model.DB.NamedQuery(timesSQL, queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	tiemsTotal := FormatSQLRowsToMapArray(timesRows)[0]["total_times"]

	total, err := model.DB.NamedQuery(countSQL, queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	pageInfo := FormatSQLRowsToMapArray(total)[0]

	pageInfo["patient_count"] = patientCount
	pageInfo["tiems_total"] = tiemsTotal
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err := model.DB.NamedQuery(querySQL+"offset :offset limit :limit", queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

// LaboratoryStatistics 检验统计
func LaboratoryStatistics(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	startDateStr := ctx.PostValue("start_date")
	endDateStr := ctx.PostValue("end_date")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
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

	endDate, erre := time.Parse("2006-01-02", endDateStr)
	if erre != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "end_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}

	endDateStr = endDate.AddDate(0, 0, 1).Format("2006-01-02")

	var queryOptions = map[string]interface{}{
		"clinic_id":  ToNullInt64(clinicID),
		"start_date": ToNullString(startDateStr),
		"end_date":   ToNullString(endDateStr),
		"offset":     ToNullInt64(offset),
		"limit":      ToNullInt64(limit),
	}

	countSQL := `select count(*) as total
	from laboratory_patient el 
 left join clinic_triage_patient ctp on el.clinic_triage_patient_id = ctp.id
 left join clinic_patient cp on ctp.clinic_patient_id = cp.id 
 left join patient p on cp.patient_id = p.id 
 left join personnel ps on ps.id = el.operation_id 
 left join clinic_laboratory cl on cl.id = el.clinic_laboratory_id
 where el.order_status = '30' 
 and ps.clinic_id = :clinic_id 
 and el.updated_time between :start_date and :end_date`
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

	patientCountSQL := `select 
	count(distinct p.name) as patient_count
	from laboratory_patient el 
	left join clinic_triage_patient ctp on el.clinic_triage_patient_id = ctp.id
	left join clinic_patient cp on ctp.clinic_patient_id = cp.id 
	left join patient p on cp.patient_id = p.id 
	left join personnel ps on ps.id = el.operation_id 
	left join clinic_laboratory cl on cl.id = el.clinic_laboratory_id
	where el.order_status = '30' 
	and ps.clinic_id = 1
	and el.updated_time between :start_date and :end_date `

	timesSQL := `	select 
	sum(el.times) as total_times
 	from laboratory_patient el 
	left join clinic_triage_patient ctp on el.clinic_triage_patient_id = ctp.id
	left join clinic_patient cp on ctp.clinic_patient_id = cp.id 
	left join patient p on cp.patient_id = p.id 
	left join personnel ps on ps.id = el.operation_id 
	left join clinic_laboratory cl on cl.id = el.clinic_laboratory_id
	where el.order_status = '30' 
	and ps.clinic_id = 1
	and el.updated_time between :start_date and :end_date `

	patientCountRows, err := model.DB.NamedQuery(patientCountSQL, queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	patientCount := FormatSQLRowsToMapArray(patientCountRows)[0]["patient_count"]

	timesRows, err := model.DB.NamedQuery(timesSQL, queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	tiemsTotal := FormatSQLRowsToMapArray(timesRows)[0]["total_times"]

	total, err := model.DB.NamedQuery(countSQL, queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]

	pageInfo["patient_count"] = patientCount
	pageInfo["tiems_total"] = tiemsTotal
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err := model.DB.NamedQuery(querySQL+"offset :offset limit :limit", queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

// TreatmentStatistics 治疗统计
func TreatmentStatistics(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	startDateStr := ctx.PostValue("start_date")
	endDateStr := ctx.PostValue("end_date")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
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

	endDate, erre := time.Parse("2006-01-02", endDateStr)
	if erre != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "end_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}
	endDateStr = endDate.AddDate(0, 0, 1).Format("2006-01-02")

	var queryOptions = map[string]interface{}{
		"clinic_id":  ToNullInt64(clinicID),
		"start_date": ToNullString(startDateStr),
		"end_date":   ToNullString(endDateStr),
		"offset":     ToNullInt64(offset),
		"limit":      ToNullInt64(limit),
	}

	countSQL := `select count(*) as total
	from treatment_patient tp 
	left join clinic_triage_patient ctp on tp.clinic_triage_patient_id = ctp.id
	left join clinic_patient cp on ctp.clinic_patient_id = cp.id 
	left join patient p on cp.patient_id = p.id 
	left join personnel ps on ps.id = tp.operation_id 
	left join clinic_treatment ct on ct.id = tp.clinic_treatment_id
	where tp.order_status = '30' 
	and ps.clinic_id = :clinic_id 
	and tp.updated_time between :start_date and :end_date `
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

	patientCountSQL := `select 
	count(distinct p.name) as patient_count
	from treatment_patient tp 
	left join clinic_triage_patient ctp on tp.clinic_triage_patient_id = ctp.id
	left join clinic_patient cp on ctp.clinic_patient_id = cp.id 
	left join patient p on cp.patient_id = p.id 
	left join personnel ps on ps.id = tp.operation_id 
	left join clinic_treatment ct on ct.id = tp.clinic_treatment_id
	where tp.order_status = '30' 
	and ps.clinic_id = :clinic_id 
	and tp.updated_time between :start_date and :end_date`

	timesSQL := `	select 
	sum(tp.times) as total_times
	from treatment_patient tp 
	left join clinic_triage_patient ctp on tp.clinic_triage_patient_id = ctp.id
	left join clinic_patient cp on ctp.clinic_patient_id = cp.id 
	left join patient p on cp.patient_id = p.id 
	left join personnel ps on ps.id = tp.operation_id 
	left join clinic_treatment ct on ct.id = tp.clinic_treatment_id
	where tp.order_status = '30' 
	and ps.clinic_id = :clinic_id 
	and tp.updated_time between :start_date and :end_date`

	patientCountRows, err := model.DB.NamedQuery(patientCountSQL, queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	patientCount := FormatSQLRowsToMapArray(patientCountRows)[0]["patient_count"]

	timesRows, err := model.DB.NamedQuery(timesSQL, queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	tiemsTotal := FormatSQLRowsToMapArray(timesRows)[0]["total_times"]

	total, err := model.DB.NamedQuery(countSQL, queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["patient_count"] = patientCount
	pageInfo["tiems_total"] = tiemsTotal
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err := model.DB.NamedQuery(querySQL+"offset :offset limit :limit", queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
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

	endDate, erre := time.Parse("2006-01-02", endDateStr)
	if erre != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "end_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}
	endDateStr = endDate.AddDate(0, 0, 1).Format("2006-01-02")

	var queryOptions = map[string]interface{}{
		"clinic_id":  ToNullInt64(clinicID),
		"start_date": ToNullString(startDateStr),
		"end_date":   ToNullString(endDateStr),
	}

	querySQL := `select to_char(ctpo.created_time, 'YYYY-MM-DD') as visit_date,  
	count(ctp.id) as total_count,
	count(case when ctp.register_type = 1 then null else 1 end) as register_count,
	count(case when ctp.register_type = 1 then 1 else null end) as appointment_count
	from clinic_triage_patient ctp
		left join clinic_triage_patient_operation ctpo on ctpo.clinic_triage_patient_id = ctp.id and times = 1 and type = 30 
		left join personnel p on p.id = ctp.doctor_id 
	where ctp.status >= 40 and p.clinic_id = :clinic_id 
	and ctpo.created_time between :start_date and :end_date `

	rows, err := model.DB.NamedQuery(querySQL+`
	group by (to_char(ctpo.created_time, 'YYYY-MM-DD'), ctp.register_type)
	order by visit_date ASC`, queryOptions)
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

// OutPatietnType 接诊类型
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
		"clinicID":  ToNullInt64(clinicID),
	}

	querySQL := `select to_char(ctpo.created_time, 'YYYY-MM-DD') as created_time, 
	sum(case when ctp.visit_type = 1 then 1 else 0 end) as type1,
	sum(case when ctp.visit_type = 2 then 1 else 0 end) as type2, 
	sum(case when ctp.visit_type = 3 then 1 else 0 end) as type3 
	FROM clinic_triage_patient ctp 
	left join clinic_triage_patient_operation ctpo on ctpo.clinic_triage_patient_id = ctp.id and type = 30 
	left join personnel p on p.id = ctpo.personnel_id 
	where ctp.status = 40 and ctpo.created_time between :startDate and :endDate and p.clinic_id = :clinicID
	group by to_char(ctpo.created_time, 'YYYY-MM-DD') 
	order by created_time ASC offset :offset limit :limit
	`

	countSQL := `select count(1) as total from (
	select 1
	FROM clinic_triage_patient ctp 
	left join clinic_triage_patient_operation ctpo on ctpo.clinic_triage_patient_id = ctp.id and type = 30 
	left join personnel p on p.id = ctpo.personnel_id 
	where ctp.status = 40 and ctpo.created_time between :startDate and :endDate and p.clinic_id = :clinicID
	group by to_char(ctpo.created_time, 'YYYY-MM-DD') ) as a`

	totalSQL := `select 
	sum(case when ctp.visit_type = 1 then 1 else 0 end) as type1,
	sum(case when ctp.visit_type = 2 then 1 else 0 end) as type2, 
	sum(case when ctp.visit_type = 3 then 1 else 0 end) as type3 
	FROM clinic_triage_patient ctp 
	left join clinic_triage_patient_operation ctpo on ctpo.clinic_triage_patient_id = ctp.id and type = 30 
	left join personnel p on p.id = ctpo.personnel_id 
	where ctp.status = 40 and ctpo.created_time between :startDate and :endDate and p.clinic_id = :clinicID`

	items, _ := model.DB.NamedQuery(querySQL, queryOptions)
	itemMap := FormatSQLRowsToMapArray(items)

	count, _ := model.DB.NamedQuery(countSQL, queryOptions)
	pageInfo := FormatSQLRowsToMapArray(count)[0]

	total, _ := model.DB.NamedQuery(totalSQL, queryOptions)
	totalMap := FormatSQLRowsToMapArray(total)[0]

	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	ctx.JSON(iris.Map{"code": "200", "data": itemMap, "page_info": pageInfo, "total": totalMap})

}

// OutPatietnDepartment 科室统计
func OutPatietnDepartment(ctx iris.Context) {

}

// OutPatietnEfficiencyStatistics 门诊效率统计
func OutPatietnEfficiencyStatistics(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	departmentID := ctx.PostValue("department_id")
	startDateStr := ctx.PostValue("start_date")
	endDateStr := ctx.PostValue("end_date")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
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

	countSQL := `	select 
	count(*) as total
	FROM clinic_triage_patient ctp
	left join clinic_patient cp on cp.id = ctp.clinic_patient_id
	left join patient p on p.id = cp.patient_id
	left join department d on d.id = ctp.department_id
	where ctp.status = '40' and cp.clinic_id =:clinic_id and ctp.visit_date between :start_date and :end_date`
	querySQL := `
	select 
	ctp.visit_date,
	p.name as patient_name,
	ctp.department_id,
	ctp.status,
	ctp.id as clinic_triage_patient_id,
	cp.patient_id,
	d.name as department_name,
	(select max(created_time) from clinic_triage_patient_operation where clinic_triage_patient_id=ctp.id and type='10') as register_time,
	(select max(created_time) from clinic_triage_patient_operation where clinic_triage_patient_id=ctp.id and type='20') as triage_time,
	(select max(created_time) from clinic_triage_patient_operation where clinic_triage_patient_id=ctp.id and type='30') as reception_time,
	(select max(created_time) from clinic_triage_patient_operation where clinic_triage_patient_id=ctp.id and type='40') as finish_time,
	(select max(created_time) from mz_paid_orders where clinic_triage_patient_id=ctp.id) as pay_time
	FROM clinic_triage_patient ctp
	left join clinic_patient cp on cp.id = ctp.clinic_patient_id
	left join patient p on p.id = cp.patient_id
	left join department d on d.id = ctp.department_id
	where ctp.status = '40' and cp.clinic_id =:clinic_id and ctp.visit_date between :start_date and :end_date`

	if departmentID != "" {
		countSQL += " and ctp.department_id =:department_id"
		querySQL += " and ctp.department_id =:department_id"
	}

	var queryOptions = map[string]interface{}{
		"clinic_id":     ToNullInt64(clinicID),
		"department_id": ToNullInt64(departmentID),
		"start_date":    ToNullString(startDateStr),
		"end_date":      ToNullString(endDateStr),
		"offset":        ToNullInt64(offset),
		"limit":         ToNullInt64(limit),
	}

	total, err := model.DB.NamedQuery(countSQL, queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	totalRows, err := model.DB.NamedQuery(querySQL, queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	totalResults := FormatSQLRowsToMapArray(totalRows)

	var triageFinishedTime float64
	var triageFinishedCount float64

	var receptionedTime float64
	var receptionedCount float64

	var receptionFinishedTime float64
	var receptionFinishedCount float64

	var payFinishedTime float64
	var payFinishedCount float64

	for _, v := range totalResults {
		registerTime := v["register_time"]
		triageTime := v["triage_time"]
		receptionTime := v["reception_time"]
		finishTime := v["finish_time"]
		payTime := v["pay_time"]

		if registerTime != nil && triageTime != nil {
			triageFinishTime := triageTime.(time.Time).Sub(registerTime.(time.Time)).Minutes()
			triageFinishedTime += triageFinishTime
			triageFinishedCount++
		}

		if triageTime != nil && receptionTime != nil {
			receptionFinishTime := receptionTime.(time.Time).Sub(triageTime.(time.Time)).Minutes()
			receptionedTime += receptionFinishTime
			receptionedCount++
		}

		if receptionTime != nil && finishTime != nil {
			receptionTotalTime := finishTime.(time.Time).Sub(receptionTime.(time.Time)).Minutes()
			receptionFinishedTime += receptionTotalTime
			receptionFinishedCount++
		}

		if finishTime != nil && payTime != nil {
			payFinishTime := payTime.(time.Time).Sub(finishTime.(time.Time)).Minutes()
			payFinishedTime += payFinishTime
			payFinishedCount++
		}
	}

	averageTriageFinishedTime, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", triageFinishedTime/triageFinishedCount), 64)
	averageReceptionedTime, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", receptionedTime/receptionedCount), 64)
	averageReceptionFinishedTime, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", receptionFinishedTime/receptionFinishedCount), 64)
	averagePayFinishedTime, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", payFinishedTime/payFinishedCount), 64)

	pageInfo["average_triage_finished_time"] = averageTriageFinishedTime
	pageInfo["average_receptioned_time"] = averageReceptionedTime
	pageInfo["average_reception_finished_time"] = averageReceptionFinishedTime
	pageInfo["average_pay_finished_time"] = averagePayFinishedTime

	rows, err := model.DB.NamedQuery(querySQL+" order by ctp.visit_date asc offset :offset limit :limit", queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}
