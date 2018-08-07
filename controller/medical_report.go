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
	and personnel.clinic_id = :clinic_id 
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
	and personnel.clinic_id = :clinic_id 
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
	and personnel.clinic_id = :clinic_id 
	and tp.updated_time between :start_date and :end_date `

	rows, err := model.DB.NamedQuery(querySQL, queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results})
}

// RegisterStatistics 治疗统计
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
