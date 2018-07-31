package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris"
)

// TriageRegister 就诊患者登记
func TriageRegister(ctx iris.Context) {
	qPatientID := ctx.PostValue("patient_id")
	certNo := ctx.PostValue("cert_no")
	name := ctx.PostValue("name")
	birthday := ctx.PostValue("birthday")
	sex := ctx.PostValue("sex")
	phone := ctx.PostValue("phone")
	province := ctx.PostValue("province")
	city := ctx.PostValue("city")
	district := ctx.PostValue("district")
	address := ctx.PostValue("address")
	profession := ctx.PostValue("profession")
	remark := ctx.PostValue("remark")
	patientChannelID := ctx.PostValue("patient_channel_id")
	clinicID := ctx.PostValue("clinic_id")
	visitType := ctx.PostValue("visit_type")
	personnelID := ctx.PostValue("personnel_id")
	departmentID := ctx.PostValue("department_id")
	if name == "" || birthday == "" || sex == "" || phone == "" || clinicID == "" || personnelID == "" || visitType == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	var row *sqlx.Row

	if qPatientID != "" {
		row = model.DB.QueryRowx("select * from patient where id = $1", qPatientID)
	} else if certNo != "" {
		row = model.DB.QueryRowx("select * from patient where cert_no = $1", certNo)
	} else {
		row = model.DB.QueryRowx("select * from patient where name = $1 and birthday = $2 and phone = $3", name, birthday, phone)
	}

	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "登记失败"})
		return
	}
	tx, err := model.DB.Begin()
	patient := FormatSQLRowToMap(row)
	_, ok := patient["id"]
	patientID := patient["id"]
	if !ok {
		insertKeys := `name, birthday, sex, phone, address, profession, remark, patient_channel_id, province, city, district`
		insertValues := `$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11`
		if certNo != "" {
			insertKeys += ", cert_no"
			insertValues += ", " + certNo
		}
		err = tx.QueryRow(`INSERT INTO patient (`+insertKeys+`) 
		VALUES (`+insertValues+`) RETURNING id`, name, birthday, sex, phone, address, profession, remark, ToNullInt64(patientChannelID), province, city, district).Scan(&patientID)
		if err != nil {
			tx.Rollback()
			fmt.Println("err2 ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	} else {
		updateSQL := `update patient set name=$1,birthday=$2,sex=$3, phone=$4, address=$5,profession = $6,remark= $7 ,patient_channel_id = $8 , province = $9, city = $10, district = $11 where id = $12`
		if certNo != "" {
			updateSQL = `update patient set cert_no = ` + certNo + `, name= $1,birthday=$2,sex=$3, phone=$4, address=$5,profession = $6,remark= $7 ,patient_channel_id = $8, province = $9, city = $10, district = $11  where id = $12`
		}
		_, err = tx.Exec(updateSQL, ToNullString(name), ToNullString(birthday), ToNullInt64(sex), ToNullString(phone), ToNullString(address), ToNullString(profession), ToNullString(remark), ToNullInt64(patientChannelID), ToNullString(province), ToNullString(city), ToNullString(district), patientID)
		if err != nil {
			tx.Rollback()
			fmt.Println("err3 ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	}

	fmt.Println("' ======= '", patientID)

	row = model.DB.QueryRowx("select * from clinic_patient where patient_id= $1 and clinic_id = $2", patientID, clinicID)
	if row == nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": "登记失败"})
		return
	}
	clinicPatient := FormatSQLRowToMap(row)
	_, ok = clinicPatient["id"]
	var clinicPatientID interface{}
	if !ok {
		err = tx.QueryRow("INSERT INTO clinic_patient (patient_id, clinic_id, personnel_id) VALUES ($1, $2, $3) RETURNING id", patientID, clinicID, personnelID).Scan(&clinicPatientID)
		if err != nil {
			fmt.Println("clinic_patient ======", err)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	} else {
		clinicPatientID = clinicPatient["id"]
	}

	insertKeys := `(clinic_patient_id, register_type, visit_type, status)`
	insertValues := `($1, 2, $2, 10)`
	if departmentID != "" {
		insertKeys = `(department_id, clinic_patient_id, register_type, visit_type, status)`
		insertValues = `(` + departmentID + `, $1, 2, $2, 10)`
	}

	insertSQL := "INSERT INTO clinic_triage_patient " + insertKeys + " VALUES " + insertValues + " RETURNING id"

	fmt.Println("insertSQL ======", insertSQL)

	var clinicTriagePatientID int
	err = tx.QueryRow(insertSQL, clinicPatientID, visitType).Scan(&clinicTriagePatientID)
	if err != nil {
		fmt.Println("clinic_triage_patient ======", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	_, err = tx.Exec("INSERT INTO clinic_triage_patient_operation(clinic_triage_patient_id, type, times, personnel_id) VALUES ($1, $2, $3, $4)", clinicTriagePatientID, 10, 1, personnelID)

	if err != nil {
		fmt.Println("clinic_triage_patient_operation ======", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": nil})
}

// TriagePatientDetail 登记患者详情
func TriagePatientDetail(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	querySQL := `select p.*, ctp.visit_type, ctp.clinic_patient_id from clinic_triage_patient ctp
	left join clinic_patient cp on ctp.clinic_patient_id = cp.id
	left join patient p on cp.patient_id = p.id
	where ctp.id = $1`
	row := model.DB.QueryRowx(querySQL, clinicTriagePatientID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询患者"})
		return
	}
	result := FormatSQLRowToMap(row)
	ctx.JSON(iris.Map{"code": "200", "data": result})
}

// TriagePatientList 当日登记就诊人列表 分诊记录
func TriagePatientList(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	keyword := ctx.PostValue("keyword")
	// status := ctx.PostValue("status")
	statusStart := ctx.PostValue("status_start")
	statusEnd := ctx.PostValue("status_end")
	registerType := ctx.PostValue("register_type")
	personnelID := ctx.PostValue("personnel_id")
	deparmentID := ctx.PostValue("department_id")
	isToday := ctx.PostValue("is_today")
	startDate := ctx.PostValue("startDate")
	endDate := ctx.PostValue("endDate")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	if startDate != "" && endDate == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请选择结束日期"})
		return
	}
	if startDate == "" && endDate != "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请选择开始日期"})
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

	countSQL := `select count(ctp.id) as total 
	from clinic_triage_patient ctp left join clinic_patient cp  on ctp.clinic_patient_id = cp.id
	left join department d on ctp.department_id = d.id
	left join personnel doc on ctp.doctor_id = doc.id
	left join patient p on cp.patient_id = p.id
	left join clinic_triage_patient_operation register on ctp.id = register.clinic_triage_patient_id and register.type = 10
	left join personnel triage_personnel on triage_personnel.id = register.personnel_id
	where cp.clinic_id = $1 and (p.cert_no ~ $2 or p.name ~ $2 or p.phone ~ $2)`

	rowSQL := `select 
	ctp.id as clinic_triage_patient_id, 
	ctp.clinic_patient_id as clinic_patient_id,
	ctp.updated_time, 
	ctp.created_time as register_time, 
	triage_personnel.name as register_personnel_name, 
	ctp.status, 
	ctp.visit_date,
	ctp.register_type,
	cp.patient_id,
	p.name as patient_name, 
	p.birthday,
	p.sex, 
	p.phone,
	doc.name as doctor_name,
	d.name as department_name
	from clinic_triage_patient ctp left join clinic_patient cp  on ctp.clinic_patient_id = cp.id
	left join department d on ctp.department_id = d.id
	left join personnel doc on ctp.doctor_id = doc.id
	left join patient p on cp.patient_id = p.id
	left join clinic_triage_patient_operation register on ctp.id = register.clinic_triage_patient_id and register.type = 10
	left join personnel triage_personnel on triage_personnel.id = register.personnel_id
	where cp.clinic_id = $1 and (p.cert_no ~ $2 or p.name ~ $2 or p.phone ~ $2)`

	if statusStart != "" && statusEnd != "" {
		countSQL += " and ctp.status BETWEEN " + statusStart + " AND " + statusEnd
		rowSQL += " and ctp.status BETWEEN " + statusStart + " AND " + statusEnd
	}

	if isToday == "true" {
		countSQL += " and ctp.visit_date= current_date "
		rowSQL += " and ctp.visit_date= current_date "
	}

	if registerType != "" {
		countSQL += " and ctp.register_type=" + registerType
		rowSQL += " and ctp.register_type=" + registerType
	}

	if personnelID != "" {
		countSQL += " and ctp.doctor_id=" + personnelID
		rowSQL += " and ctp.doctor_id=" + personnelID
	}

	if deparmentID != "" {
		countSQL += " and ctp.department_id=" + deparmentID
		rowSQL += " and ctp.department_id=" + deparmentID
	}

	if startDate != "" && endDate != "" {
		if startDate > endDate {
			ctx.JSON(iris.Map{"code": "-1", "msg": "开始日期必须小于结束日期"})
			return
		}
		countSQL += " and ctp.created_time between date'" + startDate + "' - integer '1' and date '" + endDate + "' + integer '1'"
		rowSQL += " and ctp.created_time between date'" + startDate + "' - integer '1' and date '" + endDate + "' + integer '1'"
	}

	total := model.DB.QueryRowx(countSQL, clinicID, keyword)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err1 := model.DB.Queryx(rowSQL+" order by ctp.id DESC offset $3 limit $4", clinicID, keyword, offset, limit)
	if err1 != nil {
		fmt.Println("err1 =======", err1)
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})
}

// RecptionPatientList 接诊就诊人列表
func RecptionPatientList(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	personnelID := ctx.PostValue("personnel_id")
	queryType := ctx.PostValue("query_type") // 待接诊 0, 已接诊 1
	startDate := ctx.PostValue("startDate")
	endDate := ctx.PostValue("endDate")
	keyword := ctx.PostValue("keyword")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if clinicID == "" || personnelID == "" || queryType == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	if startDate != "" && endDate == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请选择结束日期"})
		return
	}
	if startDate == "" && endDate != "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请选择开始日期"})
		return
	}

	if offset == "" {
		offset = "0"
	}

	if limit == "" {
		limit = "10"
	}
	statusStart := 40
	statusEnd := 90
	if queryType == "0" {
		statusStart = 20
		statusEnd = 30
	}

	queryMap := map[string]interface{}{
		"clinic_id":    ToNullInt64(clinicID),
		"personnel_id": ToNullInt64(personnelID),
		"keyword":      ToNullString(keyword),
		"status_start": statusStart,
		"status_end":   statusEnd,
		"startDate":    startDate,
		"endDate":      endDate,
		"offset":       ToNullInt64(offset),
		"limit":        ToNullInt64(limit),
	}

	row := model.DB.QueryRowx(`select * from personnel where id = $1`, personnelID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "医生不存在"})
		return
	}
	personnel := FormatSQLRowToMap(row)
	isClinicAdmin, ok := personnel["is_clinic_admin"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "医生不存在"})
		return
	}

	countSQL := `select count(*) as total from clinic_triage_patient ctp left join clinic_patient cp  on ctp.clinic_patient_id = cp.id
	left join department d on ctp.department_id = d.id
	left join personnel doc on ctp.doctor_id = doc.id
	left join patient p on cp.patient_id = p.id
	left join clinic_triage_patient_operation register on ctp.id = register.clinic_triage_patient_id and register.type = 10
	left join personnel triage_personnel on triage_personnel.id = register.personnel_id
	where cp.clinic_id = :clinic_id and ctp.status BETWEEN :status_start and :status_end `

	querySQL := `select 
	ctp.id as clinic_triage_patient_id, 
	ctp.clinic_patient_id as clinic_patient_id,
	ctp.updated_time, 
	ctp.created_time as register_time, 
	triage_personnel.name as register_personnel_name, 
	ctp.status, 
	ctp.visit_date,
	ctp.register_type,
	cp.patient_id,
	p.name as patient_name, 
	p.birthday,
	p.sex, 
	p.phone,
	doc.name as doctor_name,
	d.name as department_name
	from clinic_triage_patient ctp left join clinic_patient cp  on ctp.clinic_patient_id = cp.id
	left join department d on ctp.department_id = d.id
	left join personnel doc on ctp.doctor_id = doc.id
	left join patient p on cp.patient_id = p.id
	left join clinic_triage_patient_operation register on ctp.id = register.clinic_triage_patient_id and register.type = 10
	left join personnel triage_personnel on triage_personnel.id = register.personnel_id
	where cp.clinic_id = :clinic_id and ctp.status BETWEEN :status_start and :status_end `

	if queryType == "1" {
		if startDate != "" && endDate != "" {
			countSQL += " and ctp.created_time between date(:startDate)-integer '1' and date(:endDate) + integer '1'"
			querySQL += " and ctp.created_time between date(:startDate)-integer '1' and date(:endDate) + integer '1'"
		}
	}
	if keyword != "" {
		countSQL += ` and (p.cert_no ~:keyword or p.name ~:keyword or p.phone ~:keyword) `
		querySQL += ` and (p.cert_no ~:keyword or p.name ~:keyword or p.phone ~:keyword) `
	}

	if !isClinicAdmin.(bool) {
		countSQL += " and ctp.doctor_id=:personnel_id "
		querySQL += " and ctp.doctor_id=:personnel_id "
	}

	total, err2 := model.DB.NamedQuery(countSQL, queryMap)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-2", "msg": err2.Error()})
		return
	}
	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err1 := model.DB.NamedQuery(querySQL+" order by ctp.id DESC offset :offset limit :limit", queryMap)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})

}

//PersonnelChoose 分诊、换诊
func PersonnelChoose(ctx iris.Context) {
	doctorVisitScheduleID := ctx.PostValue("doctor_visit_schedule_id")
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	triagePersonnelID := ctx.PostValue("triage_personnel_id")

	if doctorVisitScheduleID == "" || clinicTriagePatientID == "" || triagePersonnelID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	fmt.Println("clinicTriagePatientID =========", clinicTriagePatientID)
	ctprow := model.DB.QueryRowx("select id, status, clinic_patient_id from clinic_triage_patient where id=$1", clinicTriagePatientID)
	clinicTriagePatient := FormatSQLRowToMap(ctprow)
	_, ctpok := clinicTriagePatient["id"]
	if !ctpok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "就诊人不存在"})
		return
	}

	_, ok := clinicTriagePatient["status"]
	if ok {
		fmt.Println("status.(string) ======", int(clinicTriagePatient["status"].(int64)))
		status := int(clinicTriagePatient["status"].(int64))
		if status >= 30 {
			ctx.JSON(iris.Map{"code": "-1", "msg": "该就诊人已接诊"})
			return
		}
	} else {
		ctx.JSON(iris.Map{"code": "-1", "msg": "状态错误，请重试"})
		return
	}

	dvsrow := model.DB.QueryRowx("select id,department_id,personnel_id,am_pm,visit_date from doctor_visit_schedule where id=$1", doctorVisitScheduleID)
	doctorVisitSchedule := FormatSQLRowToMap(dvsrow)
	_, dvsok := doctorVisitSchedule["id"]
	if !dvsok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "分诊医生号源不存在"})
		return
	}

	deparmentID := doctorVisitSchedule["department_id"]
	doctorID := doctorVisitSchedule["personnel_id"]

	tx, err := model.DB.Begin()
	var resultID int
	err = tx.QueryRow("update clinic_triage_patient set doctor_id=$1, department_id=$2,status=20,updated_time=LOCALTIMESTAMP where id=$3 RETURNING id", doctorID, deparmentID, clinicTriagePatientID).Scan(&resultID)
	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	opRow := model.DB.QueryRowx("select count(*) + 1 as times from clinic_triage_patient_operation where type = 20 and clinic_triage_patient_id = $1", clinicTriagePatientID)
	operation := FormatSQLRowToMap(opRow)
	times := operation["times"]
	_, err = tx.Exec("INSERT INTO clinic_triage_patient_operation(clinic_triage_patient_id, type, times, personnel_id) VALUES ($1, $2, $3, $4)", clinicTriagePatientID, 20, times, triagePersonnelID)

	if err != nil {
		fmt.Println("clinic_triage_patient_operation ======", err, times, clinicTriagePatientID)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": nil})
}

//TriagePersonnelList 分诊医生列表
func TriagePersonnelList(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	deparmentID := ctx.PostValue("department_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	keyword := ctx.PostValue("keyword")
	todayHour := time.Now().Hour()
	ampm := "p"
	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
		return
	}
	if todayHour < 12 {
		ampm = "a"
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

	countSQL := `from doctor_visit_schedule dvs 
	left join department d on dvs.department_id = d.id 
	left join personnel p on dvs.personnel_id = p.id 
	where p.clinic_id = $1 and (p.name like '%' || $2 || '%') and dvs.am_pm=$3 and dvs.visit_date=current_date`

	selectSQL := `(select count(ctp.id) from clinic_triage_patient ctp 
		where status=20 and visit_date=current_date and doctor_id=dvs.personnel_id) as wait_total,
	(select count(ctped.id)from clinic_triage_patient ctped where 
	status >=30 and visit_date=current_date and doctor_id=dvs.personnel_id) as triaged_total
	from doctor_visit_schedule dvs 
	left join department d on dvs.department_id = d.id 
	left join personnel p on dvs.personnel_id = p.id
	where p.clinic_id = $1 and p.name ~$2 and dvs.am_pm=$3 and dvs.visit_date=current_date`

	if deparmentID != "" {
		countSQL += " and dvs.department_id=" + deparmentID
		selectSQL += " and dvs.department_id=" + deparmentID
	}

	total := model.DB.QueryRowx(`select count(dvs.id) as total `+countSQL, clinicID, keyword, ampm)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rowSQL := `select p.name as doctor_name, d.name as department_name, dvs.id as doctor_visit_schedule_id, ` + selectSQL + " offset $4 limit $5"

	rows, err1 := model.DB.Queryx(rowSQL, clinicID, keyword, ampm, offset, limit)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})
}

//TriageReception 医生接诊病人
func TriageReception(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	triagePersonnelID := ctx.PostValue("recept_personnel_id")
	if clinicTriagePatientID == "" || triagePersonnelID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	row := model.DB.QueryRowx("select id,doctor_id, status from clinic_triage_patient where id=$1", clinicTriagePatientID)
	clinicTriagePatient := FormatSQLRowToMap(row)
	_, ok := clinicTriagePatient["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "就诊人不存在"})
		return
	}
	doctorID := clinicTriagePatient["doctor_id"]
	doctorIDStr := strconv.Itoa(int(doctorID.(int64)))
	if triagePersonnelID != doctorIDStr {
		ctx.JSON(iris.Map{"code": "-1", "msg": "该患者您不能接诊"})
		return
	}
	fmt.Println("clinicTriagePatient", clinicTriagePatient)
	status := clinicTriagePatient["status"]
	fmt.Println("ssss", status)
	if status.(int64) != 20 {
		ctx.JSON(iris.Map{"code": "-1", "msg": "当前状态不能接诊"})
		return
	}
	tx, err := model.DB.Begin()
	_, err = tx.Exec("update clinic_triage_patient set status=30,doctor_id=$1, updated_time=LOCALTIMESTAMP where id=$2", triagePersonnelID, clinicTriagePatientID)
	if err != nil {
		fmt.Println("接诊错误", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": "接诊失败"})
		return
	}

	_, err = tx.Exec("INSERT INTO clinic_triage_patient_operation(clinic_triage_patient_id, type, times, personnel_id) VALUES ($1, $2, $3, $4)", clinicTriagePatientID, 30, 1, triagePersonnelID)

	if err != nil {
		fmt.Println("clinic_triage_patient_operation ======", err, clinicTriagePatientID)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok"})
}

// TriageComplete 医生完成接诊
func TriageComplete(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	triagePersonnelID := ctx.PostValue("recept_personnel_id")
	if clinicTriagePatientID == "" || triagePersonnelID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	row := model.DB.QueryRowx("select id,status from clinic_triage_patient where id=$1", clinicTriagePatientID)
	clinicTriagePatient := FormatSQLRowToMap(row)
	_, ok := clinicTriagePatient["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "就诊人不存在"})
		return
	}
	status := clinicTriagePatient["status"]
	if status.(int64) != 30 {
		ctx.JSON(iris.Map{"code": "-1", "msg": "当前状态不能完成接诊"})
		return
	}
	tx, err := model.DB.Begin()
	_, err = tx.Exec("update clinic_triage_patient set status=40,updated_time=LOCALTIMESTAMP where id=$1", clinicTriagePatientID)
	if err != nil {
		fmt.Println("完成接诊错误", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": "完成接诊失败"})
		return
	}

	_, err = tx.Exec("INSERT INTO clinic_triage_patient_operation(clinic_triage_patient_id, type, times, personnel_id) VALUES ($1, $2, $3, $4)", clinicTriagePatientID, 40, 1, triagePersonnelID)

	if err != nil {
		fmt.Println("clinic_triage_patient_operation ======", err, clinicTriagePatientID)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok"})

}

//AppointmentsByDate 按日期统计挂号记录
func AppointmentsByDate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	departmentID := ctx.PostValue("department_id")
	personnelID := ctx.PostValue("personnel_id")
	startDateStr := ctx.PostValue("start_date")
	offsetStr := ctx.PostValue("offset")
	limitStr := ctx.PostValue("limit")
	dayLong := ctx.PostValue("day_long")
	if clinicID == "" || startDateStr == "" || dayLong == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if offsetStr == "" {
		offsetStr = "0"
	}

	if limitStr == "" {
		limitStr = "10"
	}

	long, err1 := strconv.Atoi(dayLong)
	offset, err01 := strconv.Atoi(offsetStr)
	limit, err02 := strconv.Atoi(limitStr)

	if err1 != nil || long < 1 {
		ctx.JSON(iris.Map{"code": "-1", "msg": "day_long 必须为大于0 的数字"})
		return
	}

	if err01 != nil || offset < 0 {
		ctx.JSON(iris.Map{"code": "-1", "msg": "offset 必须为大于-1 的数字"})
		return
	}

	if err02 != nil || limit < 0 {
		ctx.JSON(iris.Map{"code": "-1", "msg": "limit 必须为大于-1 的数字"})
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "start_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}

	endDate := startDate.AddDate(0, 0, long-1)

	countClinicSQL := `select count(ctp.clinic_patient_id) as count, ctp.visit_date, ctp.am_pm from clinic_triage_patient ctp 
	left join department d on ctp.department_id = d.id
	where ctp.register_type = 1 and d.clinic_id = $1 and ctp.visit_date between $2 and $3 `

	if departmentID != "" {
		countClinicSQL += ` and ctp.department_id = ` + departmentID
	}

	if personnelID != "" {
		countClinicSQL += ` and ctp.doctor_id = ` + personnelID
	}

	countClinicRows, err1 := model.DB.Queryx(countClinicSQL+` group by (ctp.visit_date, ctp.am_pm)`, clinicID, startDate, endDate)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	clinicArray := FormatSQLRowsToMapArray(countClinicRows)

	countSQL := `select count(dp.id) as total from department_personnel dp 
	left join department d on dp.department_id = d.id
	left join personnel p on dp.personnel_id = p.id
	where dp.type = 2 and d.clinic_id = $1 `

	doctorListSQL := `select dp.department_id, dp.personnel_id, d.name as department_name, p.name as personnel_name from department_personnel dp 
	left join department d on dp.department_id = d.id
	left join personnel p on dp.personnel_id = p.id
	where dp.type = 2 and d.clinic_id = $1 `

	doctorCountSQL := `select count(clinic_patient_id) as count, ctp.visit_date, ctp.am_pm, ctp.department_id, ctp.doctor_id from clinic_triage_patient ctp 
	left join department d on ctp.department_id = d.id 
	where d.clinic_id = $1 and ctp.register_type = 1 and ctp.visit_date between $2 and $3 `

	if departmentID != "" {
		countSQL += ` and dp.department_id = ` + departmentID
		doctorListSQL += ` and dp.department_id = ` + departmentID
		doctorCountSQL += ` and ctp.department_id = ` + departmentID
	}

	if personnelID != "" {
		countSQL += ` and dp.personnel_id = ` + personnelID
		doctorListSQL += ` and dp.personnel_id = ` + personnelID
		doctorCountSQL += ` and ctp.doctor_id = ` + personnelID
	}

	doctorDataSQL := `select dp.department_id, dp.department_name,
	dp.personnel_id, dp.personnel_name,
	 dpc.count, dpc.visit_date, dpc.am_pm
	 from 
	 (` + doctorListSQL + ` offset $4 limit $5) dp left join (` + doctorCountSQL + ` group by (ctp.visit_date, ctp.am_pm, ctp.department_id, ctp.doctor_id)) dpc 
	 on dp.department_id = dpc.department_id and dpc.doctor_id = dp.personnel_id;`

	doctorDataRows, err2 := model.DB.Queryx(doctorDataSQL, clinicID, startDate, endDate, offset, limit)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}

	doctorArray := FormatSQLRowsToMapArray(doctorDataRows)

	pageInfoRow := model.DB.QueryRowx(countSQL, clinicID)

	pageInfo := FormatSQLRowToMap(pageInfoRow)

	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "clinic_array": clinicArray, "doctor_array": doctorArray, "page_info": pageInfo})
}

//ReceiveRecord 获取病人历史已接诊记录
func ReceiveRecord(ctx iris.Context) {
	clinicPatientID := ctx.PostValue("clinic_patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	keyword := ctx.PostValue("keyword")
	if clinicPatientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if offset == "" {
		offset = "0"
	}

	if limit == "" {
		limit = "10"
	}

	countSQL := `select count (*) as total from clinic_triage_patient ctp
	left join medical_record mr on ctp.id = mr.clinic_triage_patient_id and mr.is_default = true
	where ctp.status > 30 and ctp.clinic_patient_id = :clinic_patient_id`

	querySQL := `select 
	ctp.id as clinic_triage_patient_id,
	ctpo.created_time,
	ctp.visit_type,
	d.name as department_name,
	p.name as doctor_name,
	mr.diagnosis,
	(select count(*) from prescription_western_patient where clinic_triage_patient_id = ctp.id) as pwp_count,
	(select count(*) from prescription_chinese_patient where clinic_triage_patient_id = ctp.id) as pcp_count,
	(select count(*) from treatment_patient where clinic_triage_patient_id = ctp.id) as tp_count,
	(select count(*) from laboratory_patient where clinic_triage_patient_id = ctp.id) as lp_count,
	(select count(*) from examination_patient where clinic_triage_patient_id = ctp.id) as ep_count,
	(select count(*) from material_patient where clinic_triage_patient_id = ctp.id) as mp_count,
	(select count(*) from other_cost_patient where clinic_triage_patient_id = ctp.id) as ocp_count
	from clinic_triage_patient ctp
	left join clinic_triage_patient_operation ctpo on ctp.id = ctpo.clinic_triage_patient_id and type = 30 and times = 1
	left join department d on ctp.department_id = d.id
	left join personnel p on ctp.doctor_id = p.id
	left join medical_record mr on ctp.id = mr.clinic_triage_patient_id and mr.is_default = true
	where ctp.status > 30 and ctp.clinic_patient_id = :clinic_patient_id`

	if keyword != "" {
		countSQL += ` and mr.diagnosis ~*:keyword`
		querySQL += ` and mr.diagnosis ~*:keyword`
	}

	var queryOptions = map[string]interface{}{
		"clinic_patient_id": ToNullInt64(clinicPatientID),
		"keyword":           ToNullString(keyword),
		"offset":            ToNullInt64(offset),
		"limit":             ToNullInt64(limit),
	}

	pageInfoRows, err := model.DB.NamedQuery(countSQL, queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	pageInfoArray := FormatSQLRowsToMapArray(pageInfoRows)
	pageInfo := pageInfoArray[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err := model.DB.NamedQuery(querySQL+` order by ctpo.created_time DESC offset :offset limit :limit`, queryOptions)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
	}

	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})
}

//TriagePatientVisitDetail 获取病人就诊信息详情
func TriagePatientVisitDetail(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")

	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	querySQL := `select mr.*,p.name as operation_name,
	(select string_agg (cl.name, '、') as clinic_laboratory_name from laboratory_patient lp
			left join clinic_laboratory cl on cl.id = lp.clinic_laboratory_id where clinic_triage_patient_id = ctp.id group by clinic_triage_patient_id),
	(select string_agg (ce.name, '、') as clinic_examination_name from examination_patient ep 
			left join clinic_examination ce on ce.id = ep.clinic_examination_id where clinic_triage_patient_id = ctp.id group by clinic_triage_patient_id),
	(select string_agg (ct.name, '，') as clinic_treatment_name from treatment_patient tp 
			left join clinic_treatment ct on ct.id = tp.clinic_treatment_id where clinic_triage_patient_id = ctp.id group by clinic_triage_patient_id)
	from clinic_triage_patient ctp
	left join medical_record mr on mr.clinic_triage_patient_id = ctp.id and mr.is_default = true
	left join personnel p on p.id = mr.operation_id
	where ctp.id=$1`

	queryMedicalRecordSQL := `select mr.*,p.name as operation_name
	from medical_record mr
	left join personnel p on p.id = mr.operation_id
	where mr.clinic_triage_patient_id=$1 and mr.is_default is null`

	queryWesternSQL := `select 
		cd.name,
		cd.packing_unit_name,
		cd.specification,
		pwp.order_sn,
		pwp.eff_day,
		pwp.illustration,
		pwp.amount,
		pwp.route_administration_name,
		pwp.once_dose,
		pwp.once_dose_unit_name 
		from prescription_western_patient pwp
		left join clinic_drug cd on cd.id = pwp.clinic_drug_id 
		where pwp.clinic_triage_patient_id=$1`

	queryChineseSQL := `select 
		cd.name as drug_name,
		cd.packing_unit_name,
		pci.once_dose,
		pci.once_dose_unit_name,
		pci.amount,
		pcp.order_sn,
		pcp.amount as info_amount,
		pcp.id as prescription_patient_id,
		pcp.medicine_illustration,
		pcp.route_administration_name as info_route_administration_name,
		pcp.frequency_name as info_frequency_name,
		pcp.eff_day as info_eff_day
		from prescription_chinese_patient pcp 
		left join prescription_chinese_item pci on pci.prescription_chinese_patient_id = pcp.id
		left join clinic_drug cd on cd.id = pci.clinic_drug_id
		where pcp.clinic_triage_patient_id=$1`

	row := model.DB.QueryRowx(querySQL, clinicTriagePatientID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询详情错误"})
	}
	result := FormatSQLRowToMap(row)

	rowsMedicalRecord, err := model.DB.Queryx(queryMedicalRecordSQL, clinicTriagePatientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	resultsMedicalRecord := FormatSQLRowsToMapArray(rowsMedicalRecord)

	rowsWestern, err := model.DB.Queryx(queryWesternSQL, clinicTriagePatientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	resultsWestern := FormatSQLRowsToMapArray(rowsWestern)

	rowsChinese, err := model.DB.Queryx(queryChineseSQL, clinicTriagePatientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	resultsChinese := FormatSQLRowsToMapArray(rowsChinese)

	result["medical_records"] = resultsMedicalRecord
	result["prescription_western_patient"] = resultsWestern
	result["prescription_chinese_patient"] = FormatPrescription(resultsChinese)

	ctx.JSON(iris.Map{"code": "200", "data": result})
}

// TriagePatientRecord 分诊记录报告
func TriagePatientRecord(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")

	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	resData := make(map[string]interface{})

	selectLaboratorySQL := `select 
	lp.id as laboratory_patient_id,
	lp.clinic_triage_patient_id,
	cl.name as clinic_laboratory_name,
	cl.laboratory_sample,
	lp.clinic_laboratory_id,
	lpr.id as laboratory_patient_record_id,
	lpr.remark,
	lp.created_time as order_time,
	lpr.created_time as report_time,
	doc.name as report_doctor_name,
	p.name as order_doctor_name
	FROM laboratory_patient lp 
	left join clinic_laboratory cl on cl.id = lp.clinic_laboratory_id
	left join mz_paid_orders mo on mo.clinic_triage_patient_id = lp.clinic_triage_patient_id and mo.charge_project_type_id=3 and lp.clinic_laboratory_id=mo.charge_project_id
	left join laboratory_patient_record lpr on lpr.laboratory_patient_id = lp.id
	left join personnel doc on doc.id = lpr.operation_id
	left join personnel p on p.id = lp.operation_id
	where lp.clinic_triage_patient_id = $1 and lp.order_status=$2`

	laboratoryRows, _ := model.DB.Queryx(selectLaboratorySQL, clinicTriagePatientID, "30")
	laboratoryResults := FormatSQLRowsToMapArray(laboratoryRows)

	selectExaminationSQL := `select 
	ep.id as examination_patient_id,
	ep.clinic_triage_patient_id,
	ce.name as clinic_examination_name,
	ce.organ,
	ep.clinic_examination_id,
	mpr.id as examination_patient_record_id,
	mpr.picture_examination,
	mpr.result_examination,
	mpr.conclusion_examination,
	mpr.created_time as report_time,
	ep.created_time as order_time,
	doc.name as report_doctor_name,
	p.name as order_doctor_name
	FROM examination_patient ep 
	left join clinic_examination ce on ce.id = ep.clinic_examination_id
	left join mz_paid_orders mo on mo.clinic_triage_patient_id = ep.clinic_triage_patient_id and mo.charge_project_type_id=4 and ep.clinic_examination_id=mo.charge_project_id
	left join examination_patient_record mpr on mpr.examination_patient_id = ep.id
	left join personnel doc on doc.id = mpr.operation_id
	left join personnel p on p.id = ep.operation_id
	where mo.id is not NULL and ep.clinic_triage_patient_id = $1 and ep.order_status=$2`

	examinationRows, _ := model.DB.Queryx(selectExaminationSQL, clinicTriagePatientID, "30")
	examinationresults := FormatSQLRowsToMapArray(examinationRows)

	resData["laboratory_results"] = laboratoryResults
	resData["examination_results"] = examinationresults

	ctx.JSON(iris.Map{"code": "200", "data": resData})
}

// QuickReception 快速接诊
func QuickReception(ctx iris.Context) {
	qPatientID := ctx.PostValue("patient_id")
	certNo := ctx.PostValue("cert_no")
	name := ctx.PostValue("name")
	birthday := ctx.PostValue("birthday")
	sex := ctx.PostValue("sex")
	phone := ctx.PostValue("phone")
	province := ctx.PostValue("province")
	city := ctx.PostValue("city")
	district := ctx.PostValue("district")
	address := ctx.PostValue("address")
	clinicID := ctx.PostValue("clinic_id")
	visitType := ctx.PostValue("visit_type")
	personnelID := ctx.PostValue("personnel_id")
	departmentID := ctx.PostValue("department_id")
	if departmentID == "" || name == "" || birthday == "" || sex == "" || phone == "" || clinicID == "" || personnelID == "" || visitType == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	var row *sqlx.Row

	if qPatientID != "" {
		row = model.DB.QueryRowx("select * from patient where id = $1", qPatientID)
	} else if certNo != "" {
		row = model.DB.QueryRowx("select * from patient where cert_no = $1", certNo)
	} else {
		row = model.DB.QueryRowx("select * from patient where name = $1 and birthday = $2 and phone = $3", name, birthday, phone)
	}

	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "登记失败"})
		return
	}
	tx, err := model.DB.Begin()
	patient := FormatSQLRowToMap(row)
	_, ok := patient["id"]
	patientID := patient["id"]
	if !ok {
		insertKeys := `name, birthday, sex, phone, address, province, city, district`
		insertValues := `$1, $2, $3, $4, $5, $6, $7, $8`
		if certNo != "" {
			insertKeys += ", cert_no"
			insertValues += ", " + certNo
		}
		err = tx.QueryRow(`INSERT INTO patient (`+insertKeys+`) 
		VALUES (`+insertValues+`) RETURNING id`, name, birthday, sex, phone, address, province, city, district).Scan(&patientID)
		if err != nil {
			tx.Rollback()
			fmt.Println("err2 ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	} else {
		updateSQL := `update patient set name=$1,birthday=$2,sex=$3, phone=$4, address=$5, province = $6, city = $7, district = $8 where id = $9`
		if certNo != "" {
			updateSQL = `update patient set cert_no = ` + certNo + `, name= $1,birthday=$2,sex=$3, phone=$4, address=$5, province = $6, city = $7, district = $8  where id = $9`
		}
		_, err = tx.Exec(updateSQL, ToNullString(name), ToNullString(birthday), ToNullInt64(sex), ToNullString(phone), ToNullString(address), ToNullString(province), ToNullString(city), ToNullString(district), patientID)
		if err != nil {
			tx.Rollback()
			fmt.Println("err3 ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	}

	fmt.Println("' ======= '", patientID)

	row = model.DB.QueryRowx("select * from clinic_patient where patient_id= $1 and clinic_id = $2", patientID, clinicID)
	if row == nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": "登记失败"})
		return
	}
	clinicPatient := FormatSQLRowToMap(row)
	_, ok = clinicPatient["id"]
	var clinicPatientID interface{}
	if !ok {
		err = tx.QueryRow("INSERT INTO clinic_patient (patient_id, clinic_id, personnel_id) VALUES ($1, $2, $3) RETURNING id", patientID, clinicID, personnelID).Scan(&clinicPatientID)
		if err != nil {
			fmt.Println("clinic_patient ======", err)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	} else {
		clinicPatientID = clinicPatient["id"]
	}

	insertKeys := `(department_id, doctor_id, clinic_patient_id, visit_type, register_type, status)`
	insertValues := `($1, $2, $3, $4, 3, 30)`

	insertSQL := "INSERT INTO clinic_triage_patient " + insertKeys + " VALUES " + insertValues + " RETURNING id"

	fmt.Println("insertSQL ======", insertSQL)

	var clinicTriagePatientID interface{}
	err = tx.QueryRow(insertSQL, departmentID, personnelID, clinicPatientID, visitType).Scan(&clinicTriagePatientID)
	if err != nil {
		fmt.Println("clinic_triage_patient ======", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	_, err = tx.Exec("INSERT INTO clinic_triage_patient_operation(clinic_triage_patient_id, type, times, personnel_id) VALUES ($1, $2, $3, $4)", clinicTriagePatientID, 30, 1, personnelID)

	if err != nil {
		fmt.Println("clinic_triage_patient_operation ======", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	resData := make(map[string]interface{})
	resData["patient_id"] = patientID
	resData["clinic_triage_patient_id"] = clinicTriagePatientID

	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": resData})
}
