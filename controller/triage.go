package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"
	"strings"
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
	if name == "" || birthday == "" || sex == "" || phone == "" || patientChannelID == "" || clinicID == "" || personnelID == "" || visitType == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
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
		ctx.JSON(iris.Map{"code": "1", "msg": "登记失败"})
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
		VALUES (`+insertValues+`) RETURNING id`, name, birthday, sex, phone, address, profession, remark, patientChannelID, province, city, district).Scan(&patientID)
		if err != nil {
			tx.Rollback()
			fmt.Println("err2 ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
	} else {
		updateSQL := `update patient set name= $1,birthday=$2,sex=$3, phone=$4, address=$5,profession = $6,remark= $7 ,patient_channel_id = $8 , province = $9, city = $10, district = $11 where id = $12`
		if certNo != "" {
			updateSQL = `update patient set cert_no = ` + certNo + `, name= $1,birthday=$2,sex=$3, phone=$4, address=$5,profession = $6,remark= $7 ,patient_channel_id = $8, province = $9, city = $10, district = $11  where id = $12`
		}
		_, err = tx.Exec(updateSQL, name, birthday, sex, phone, address, profession, remark, patientChannelID, province, city, district, patientID)
		if err != nil {
			tx.Rollback()
			fmt.Println("err3 ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
	}

	fmt.Println("' ======= '", patientID)

	row = model.DB.QueryRowx("select * from clinic_patient where patient_id= $1 and clinic_id = $2", patientID, clinicID)
	if row == nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": "登记失败"})
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
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
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
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": nil})
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
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
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
	ctp.updated_time, 
	ctp.created_time as register_time, 
	triage_personnel.name as register_personnel_name, 
	ctp.status, 
	ctp.visit_date,
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
			ctx.JSON(iris.Map{"code": "-1", "msg": "开始日期必须大于结束日期"})
			return
		}
		countSQL += " and ctp.created_time between date'" + startDate + "' - integer '1' and date '" + endDate + "' + integer '1'"
		rowSQL += " and ctp.created_time between date'" + startDate + "' - integer '1' and date '" + endDate + "' + integer '1'"
	}

	total := model.DB.QueryRowx(countSQL, clinicID, keyword)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	fmt.Println("pageInfo ====", pageInfo)
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

//PersonnelChoose 分诊、换诊
func PersonnelChoose(ctx iris.Context) {
	doctorVisitScheduleID := ctx.PostValue("doctor_visit_schedule_id")
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	triagePersonnelID := ctx.PostValue("triage_personnel_id")

	if doctorVisitScheduleID == "" || clinicTriagePatientID == "" || triagePersonnelID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	fmt.Println("clinicTriagePatientID =========", clinicTriagePatientID)
	ctprow := model.DB.QueryRowx("select id, status, clinic_patient_id from clinic_triage_patient where id=$1", clinicTriagePatientID)
	clinicTriagePatient := FormatSQLRowToMap(ctprow)
	_, ctpok := clinicTriagePatient["id"]
	if !ctpok {
		ctx.JSON(iris.Map{"code": "1", "msg": "就诊人不存在"})
		return
	}

	_, ok := clinicTriagePatient["status"]
	if ok {
		fmt.Println("status.(string) ======", int(clinicTriagePatient["status"].(int64)))
		status := int(clinicTriagePatient["status"].(int64))
		if status >= 30 {
			ctx.JSON(iris.Map{"code": "1", "msg": "该就诊人已接诊"})
			return
		}
	} else {
		ctx.JSON(iris.Map{"code": "1", "msg": "状态错误，请重试"})
		return
	}

	dvsrow := model.DB.QueryRowx("select id,department_id,personnel_id,am_pm,visit_date from doctor_visit_schedule where id=$1", doctorVisitScheduleID)
	doctorVisitSchedule := FormatSQLRowToMap(dvsrow)
	_, dvsok := doctorVisitSchedule["id"]
	if !dvsok {
		ctx.JSON(iris.Map{"code": "1", "msg": "分诊医生号源不存在"})
		return
	}

	clinicPatientID := clinicTriagePatient["clinic_patient_id"]
	deparmentID := doctorVisitSchedule["department_id"]
	doctorID := doctorVisitSchedule["personnel_id"]
	amPm := doctorVisitSchedule["am_pm"]
	visitDate := doctorVisitSchedule["visit_date"]

	tx, err := model.DB.Begin()
	var resultID int
	err = tx.QueryRow("update clinic_triage_patient set doctor_id=$1, department_id=$2,status=20,updated_time=LOCALTIMESTAMP where id=$3 RETURNING id", doctorID, deparmentID, clinicTriagePatientID).Scan(&resultID)
	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
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

	var registrationID interface{}
	rrow := model.DB.QueryRowx("select id from registration where clinic_triage_patient_id=$1", clinicTriagePatientID)
	registration := FormatSQLRowToMap(rrow)
	_, rok := registration["id"]
	if !rok {
		err = tx.QueryRow(`INSERT INTO registration (
			clinic_patient_id, department_id, clinic_triage_patient_id,personnel_id,visit_date,am_pm,operation_id) 
			VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`, clinicPatientID, deparmentID, clinicTriagePatientID, doctorID, visitDate, amPm, triagePersonnelID).Scan(&registrationID)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
	} else {
		registrationID = registration["id"]
		_, err = tx.Exec(`update registration set
			department_id=$1,personnel_id=$2,visit_date=$3,am_pm=$4,operation_id=$5,updated_time=LOCALTIMESTAMP where id=$6`, deparmentID, doctorID, visitDate, amPm, triagePersonnelID, registrationID)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": registrationID})
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
	where p.clinic_id = $1 and (p.name like '%' || $2 || '%') and dvs.am_pm=$3 and dvs.visit_date=current_date`

	if deparmentID != "" {
		countSQL += " and dvs.department_id=" + deparmentID
		selectSQL += " and dvs.department_id=" + deparmentID
	}

	total := model.DB.QueryRowx(`select count(dvs.id) as total `+countSQL, clinicID, keyword, ampm)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
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

// TriageCompleteBodySign 完善体征信息
func TriageCompleteBodySign(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	weight := ctx.PostValue("weight")
	height := ctx.PostValue("height")
	bmi := ctx.PostValue("bmi")
	bloodType := ctx.PostValue("blood_type")
	rhBloodType := ctx.PostValue("rh_blood_type")
	temperatureType := ctx.PostValue("temperature_type")
	temperature := ctx.PostValue("temperature")
	breathe := ctx.PostValue("breathe")
	pulse := ctx.PostValue("pulse")
	systolicBloodPressure := ctx.PostValue("systolic_blood_pressure")
	diastolicBloodPressure := ctx.PostValue("diastolic_blood_pressure")
	bloodSugarTime := ctx.PostValue("blood_sugar_time")
	bloodSugarType := ctx.PostValue("blood_sugar_type")
	leftVision := ctx.PostValue("left_vision")
	rightVision := ctx.PostValue("right_vision")
	oxygenSaturation := ctx.PostValue("oxygen_saturation")
	painScore := ctx.PostValue("pain_score")

	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	var insertKeys []string
	var insertValues []string

	if clinicTriagePatientID != "" {
		insertKeys = append(insertKeys, "clinic_triage_patient_id")
		insertValues = append(insertValues, clinicTriagePatientID)
	}
	if weight != "" {
		insertKeys = append(insertKeys, "weight")
		insertValues = append(insertValues, weight)
	}
	if height != "" {
		insertKeys = append(insertKeys, "height")
		insertValues = append(insertValues, height)
	}
	if bmi != "" {
		insertKeys = append(insertKeys, "bmi")
		insertValues = append(insertValues, bmi)
	}
	if bloodType != "" {
		insertKeys = append(insertKeys, "blood_type")
		insertValues = append(insertValues, "'"+bloodType+"'")
	}
	if rhBloodType != "" {
		insertKeys = append(insertKeys, "rh_blood_type")
		insertValues = append(insertValues, "'"+rhBloodType+"'")
	}
	if temperatureType != "" {
		insertKeys = append(insertKeys, "temperature_type")
		insertValues = append(insertValues, temperatureType)
	}
	if temperature != "" {
		insertKeys = append(insertKeys, "temperature")
		insertValues = append(insertValues, temperature)
	}
	if breathe != "" {
		insertKeys = append(insertKeys, "breathe")
		insertValues = append(insertValues, breathe)
	}
	if pulse != "" {
		insertKeys = append(insertKeys, "pulse")
		insertValues = append(insertValues, pulse)
	}
	if systolicBloodPressure != "" {
		insertKeys = append(insertKeys, "systolic_blood_pressure")
		insertValues = append(insertValues, systolicBloodPressure)
	}
	if diastolicBloodPressure != "" {
		insertKeys = append(insertKeys, "diastolic_blood_pressure")
		insertValues = append(insertValues, diastolicBloodPressure)
	}
	if bloodSugarTime != "" {
		insertKeys = append(insertKeys, "blood_sugar_time")
		insertValues = append(insertValues, "'"+bloodSugarTime+"'")
	}
	if bloodSugarType != "" {
		insertKeys = append(insertKeys, "blood_sugar_type")
		insertValues = append(insertValues, bloodSugarType)
	}
	if leftVision != "" {
		insertKeys = append(insertKeys, "left_vision")
		insertValues = append(insertValues, "'"+leftVision+"'")
	}
	if rightVision != "" {
		insertKeys = append(insertKeys, "right_vision")
		insertValues = append(insertValues, "'"+rightVision+"'")
	}
	if oxygenSaturation != "" {
		insertKeys = append(insertKeys, "oxygen_saturation")
		insertValues = append(insertValues, oxygenSaturation)
	}
	if painScore != "" {
		insertKeys = append(insertKeys, "pain_score")
		insertValues = append(insertValues, painScore)
	}

	insertKeyStr := strings.Join(insertKeys, ",")
	insertValueStr := strings.Join(insertValues, ",")

	_, err1 := model.DB.Exec("delete from body_sign where clinic_triage_patient_id=" + clinicTriagePatientID)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err1})
		return
	}

	insertSQL := "insert into body_sign (" + insertKeyStr + ") values (" + insertValueStr + ") RETURNING id;"
	fmt.Println("insertSQL ========", insertSQL)
	_, err := model.DB.Exec(insertSQL)

	if err != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "保存成功"})
}

// TriageCompletePreMedicalRecord 完善诊前病历
func TriageCompletePreMedicalRecord(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	hasAllergicHistory := ctx.PostValue("has_allergic_history")
	allergicHistory := ctx.PostValue("allergic_history")
	personalMedicalHistory := ctx.PostValue("personal_medical_history")
	familyMedicalHistory := ctx.PostValue("family_medical_history")
	vaccination := ctx.PostValue("vaccination")
	menarcheAge := ctx.PostValue("menarche_age")
	menstrualPeriodStartDay := ctx.PostValue("menstrual_period_start_day")
	menstrualPeriodEndDay := ctx.PostValue("menstrual_period_end_day")
	menstrualCycleStartDay := ctx.PostValue("menstrual_cycle_start_day")
	menstrualCycleEndDay := ctx.PostValue("menstrual_cycle_end_day")
	menstrualLastDay := ctx.PostValue("menstrual_last_day")
	gestationalWeeks := ctx.PostValue("gestational_weeks")
	childbearingHistory := ctx.PostValue("childbearing_history")

	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	var insertKeys []string
	var insertValues []string

	if clinicTriagePatientID != "" {
		insertKeys = append(insertKeys, "clinic_triage_patient_id")
		insertValues = append(insertValues, clinicTriagePatientID)
	}
	if hasAllergicHistory != "" {
		insertKeys = append(insertKeys, "has_allergic_history")
		insertValues = append(insertValues, hasAllergicHistory)
	}
	if allergicHistory != "" {
		insertKeys = append(insertKeys, "allergic_history")
		insertValues = append(insertValues, "'"+allergicHistory+"'")
	}
	if personalMedicalHistory != "" {
		insertKeys = append(insertKeys, "personal_medical_history")
		insertValues = append(insertValues, "'"+personalMedicalHistory+"'")
	}
	if familyMedicalHistory != "" {
		insertKeys = append(insertKeys, "family_medical_history")
		insertValues = append(insertValues, "'"+familyMedicalHistory+"'")
	}
	if vaccination != "" {
		insertKeys = append(insertKeys, "vaccination")
		insertValues = append(insertValues, "'"+vaccination+"'")
	}
	if menarcheAge != "" {
		insertKeys = append(insertKeys, "menarche_age")
		insertValues = append(insertValues, menarcheAge)
	}
	if menstrualPeriodStartDay != "" {
		insertKeys = append(insertKeys, "menstrual_period_start_day")
		insertValues = append(insertValues, "'"+menstrualPeriodStartDay+"'")
	}
	if menstrualPeriodEndDay != "" {
		insertKeys = append(insertKeys, "menstrual_period_end_day")
		insertValues = append(insertValues, "'"+menstrualPeriodEndDay+"'")
	}
	if menstrualCycleStartDay != "" {
		insertKeys = append(insertKeys, "menstrual_cycle_start_day")
		insertValues = append(insertValues, "'"+menstrualCycleStartDay+"'")
	}
	if menstrualCycleEndDay != "" {
		insertKeys = append(insertKeys, "menstrual_cycle_end_day")
		insertValues = append(insertValues, "'"+menstrualCycleEndDay+"'")
	}
	if menstrualLastDay != "" {
		insertKeys = append(insertKeys, "menstrual_last_day")
		insertValues = append(insertValues, "'"+menstrualLastDay+"'")
	}
	if gestationalWeeks != "" {
		insertKeys = append(insertKeys, "gestational_weeks")
		insertValues = append(insertValues, gestationalWeeks)
	}
	if childbearingHistory != "" {
		insertKeys = append(insertKeys, "childbearing_history")
		insertValues = append(insertValues, "'"+childbearingHistory+"'")
	}

	insertKeyStr := strings.Join(insertKeys, ",")
	insertValueStr := strings.Join(insertValues, ",")

	insertSQL := "insert into pre_medical_record (" + insertKeyStr + ") values (" + insertValueStr + ") RETURNING id;"
	_, err1 := model.DB.Exec("delete from pre_medical_record where clinic_triage_patient_id=" + clinicTriagePatientID)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err1})
		return
	}
	_, err := model.DB.Exec(insertSQL)

	if err != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "保存成功"})
}

// TriageCompletePreDiagnosis 完善诊前欲诊
func TriageCompletePreDiagnosis(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	chiefComplaint := ctx.PostValue("chief_complaint")
	historyOfRresentIllness := ctx.PostValue("history_of_rresent_illness")
	historyOfPastIllness := ctx.PostValue("history_of_past_illness")
	physicalExamination := ctx.PostValue("physical_examination")
	remark := ctx.PostValue("remark")

	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	var insertKeys []string
	var insertValues []string

	if clinicTriagePatientID != "" {
		insertKeys = append(insertKeys, "clinic_triage_patient_id")
		insertValues = append(insertValues, clinicTriagePatientID)
	}

	if chiefComplaint != "" {
		insertKeys = append(insertKeys, "chief_complaint")
		insertValues = append(insertValues, "'"+chiefComplaint+"'")
	}

	if historyOfRresentIllness != "" {
		insertKeys = append(insertKeys, "history_of_rresent_illness")
		insertValues = append(insertValues, "'"+historyOfRresentIllness+"'")
	}

	if historyOfPastIllness != "" {
		insertKeys = append(insertKeys, "history_of_past_illness")
		insertValues = append(insertValues, "'"+historyOfPastIllness+"'")
	}

	if physicalExamination != "" {
		insertKeys = append(insertKeys, "physical_examination")
		insertValues = append(insertValues, "'"+physicalExamination+"'")
	}

	if remark != "" {
		insertKeys = append(insertKeys, "remark")
		insertValues = append(insertValues, "'"+remark+"'")
	}

	insertKeyStr := strings.Join(insertKeys, ",")
	insertValueStr := strings.Join(insertValues, ",")

	insertSQL := "insert into pre_diagnosis (" + insertKeyStr + ") values (" + insertValueStr + ") RETURNING id;"

	_, err1 := model.DB.Exec("delete from pre_diagnosis where clinic_triage_patient_id=" + clinicTriagePatientID)
	_, err := model.DB.Exec(insertSQL)

	if err1 != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err1})
		return
	}

	if err != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "保存成功"})
}

//TriageReception 医生接诊病人
func TriageReception(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	triagePersonnelID := ctx.PostValue("recept_personnel_id")
	if clinicTriagePatientID == "" || triagePersonnelID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	row := model.DB.QueryRowx("select id,status from clinic_triage_patient where id=$1", clinicTriagePatientID)
	clinicTriagePatient := FormatSQLRowToMap(row)
	_, ok := clinicTriagePatient["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "就诊人不存在"})
		return
	}
	fmt.Println("clinicTriagePatient", clinicTriagePatient)
	status := clinicTriagePatient["status"]
	fmt.Println("ssss", status)
	if status.(int64) != 20 {
		ctx.JSON(iris.Map{"code": "1", "msg": "当前状态不能接诊"})
		return
	}
	tx, err := model.DB.Begin()
	_, err = tx.Exec("update clinic_triage_patient set status=30,updated_time=LOCALTIMESTAMP where id=$1", clinicTriagePatientID)
	if err != nil {
		fmt.Println("接诊错误", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": "接诊失败"})
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
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
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
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
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
