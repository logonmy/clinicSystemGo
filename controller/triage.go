package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
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
	ctp.register_type,
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

	deparmentID := doctorVisitSchedule["department_id"]
	doctorID := doctorVisitSchedule["personnel_id"]

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

	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
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

// TriageComplete 医生完成接诊
func TriageComplete(ctx iris.Context) {
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
	status := clinicTriagePatient["status"]
	if status.(int64) != 30 {
		ctx.JSON(iris.Map{"code": "1", "msg": "当前状态不能完成接诊"})
		return
	}
	tx, err := model.DB.Begin()
	_, err = tx.Exec("update clinic_triage_patient set status=40,updated_time=LOCALTIMESTAMP where id=$1", clinicTriagePatientID)
	if err != nil {
		fmt.Println("完成接诊错误", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": "完成接诊失败"})
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

//TreatmentPatientCreate 开治疗
func TreatmentPatientCreate(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	personnelID := ctx.PostValue("personnel_id")
	items := ctx.PostValue("items")

	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	if items == "" {
		ctx.JSON(iris.Map{"code": "200", "data": nil})
		return
	}

	var results []map[string]string
	errj := json.Unmarshal([]byte(items), &results)
	fmt.Println("results===", results)
	if errj != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": errj.Error()})
		return
	}
	row := model.DB.QueryRowx(`select id,status from clinic_triage_patient where id=$1 limit 1`, clinicTriagePatientID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "保存治疗失败,分诊记录错误"})
		return
	}

	prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", personnelID)
	if prow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "保存治疗失败,操作员错误"})
		return
	}
	clinicTriagePatient := FormatSQLRowToMap(row)
	personnel := FormatSQLRowToMap(prow)

	_, ok := clinicTriagePatient["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "分诊记录不存在"})
		return
	}
	status := clinicTriagePatient["status"]
	if status.(int64) != 30 {
		ctx.JSON(iris.Map{"code": "1", "msg": "分诊记录当前状态错误"})
		return
	}
	_, pok := personnel["id"]
	if !pok {
		ctx.JSON(iris.Map{"code": "1", "msg": "操作员错误"})
		return
	}

	var mzUnpaidOrdersValues []string
	mzUnpaidOrdersSets := []string{
		"clinic_triage_patient_id",
		"charge_project_type_id",
		"charge_project_id",
		"order_sn",
		"soft_sn",
		"name",
		"price",
		"amount",
		"unit",
		"total",
		"fee",
		"operation_id",
	}

	var treatmentPatientValues []string
	treatmentPatientSets := []string{
		"clinic_triage_patient_id",
		"clinic_treatment_id",
		"order_sn",
		"soft_sn",
		"times",
		"operation_id",
		"illustration",
	}

	for index, v := range results {
		clinicTreatmentID := v["clinic_treatment_id"]
		times := v["times"]
		illustration := v["illustration"]
		fmt.Println("clinicTreatmentID====", clinicTreatmentID)
		var st []string
		var sm []string
		treatmentSQL := `select ct.id as clinic_treatment_id,ct.price,ct.is_discount,t.name,du.name as dose_unit_name from clinic_treatment ct
			left join treatment t on t.id = ct.treatment_id
			left join dose_unit du on du.id = t.unit_id
			where ct.id=$1`
		trow := model.DB.QueryRowx(treatmentSQL, clinicTreatmentID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "治疗项错误"})
			return
		}
		treatment := FormatSQLRowToMap(trow)
		fmt.Println("====", treatment)
		_, ok := treatment["clinic_treatment_id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "选择的治疗项错误"})
			return
		}
		orderSn := FormatPayOrderSn(clinicTriagePatientID, "7")
		price := treatment["price"].(int64)
		name := treatment["name"].(string)
		unitName := treatment["dose_unit_name"].(string)
		amount, _ := strconv.Atoi(times)
		total := int(price) * amount

		st = append(st, clinicTriagePatientID, clinicTreatmentID, "'"+orderSn+"'", strconv.Itoa(index), times, personnelID)
		sm = append(sm, clinicTriagePatientID, "7", clinicTreatmentID, "'"+orderSn+"'", strconv.Itoa(index), "'"+name+"'", strconv.FormatInt(price, 10), strconv.Itoa(amount), "'"+unitName+"'", strconv.Itoa(total), strconv.Itoa(total), personnelID)

		if illustration == "" {
			st = append(st, `null`)
		} else {
			st = append(st, "'"+illustration+"'")
		}

		tstr := "(" + strings.Join(st, ",") + ")"
		treatmentPatientValues = append(treatmentPatientValues, tstr)
		mstr := "(" + strings.Join(sm, ",") + ")"
		mzUnpaidOrdersValues = append(mzUnpaidOrdersValues, mstr)
	}
	tSetStr := strings.Join(treatmentPatientSets, ",")
	tValueStr := strings.Join(treatmentPatientValues, ",")

	mSetStr := strings.Join(mzUnpaidOrdersSets, ",")
	mvValueStr := strings.Join(mzUnpaidOrdersValues, ",")

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errb})
		return
	}

	_, errdtp := tx.Exec("delete from treatment_patient where clinic_triage_patient_id=$1", clinicTriagePatientID)
	if errdtp != nil {
		fmt.Println("errdtp ===", errdtp)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errdtp.Error()})
		return
	}
	_, errdm := tx.Exec("delete from mz_unpaid_orders where clinic_triage_patient_id=$1 and charge_project_type_id=7", clinicTriagePatientID)
	if errdm != nil {
		fmt.Println("errdm ===", errdm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errdm.Error()})
		return
	}

	inserttSQL := "insert into treatment_patient (" + tSetStr + ") values " + tValueStr
	fmt.Println("inserttSQL===", inserttSQL)

	_, errt := tx.Exec(inserttSQL)
	if errt != nil {
		fmt.Println("errt ===", errt)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errt.Error()})
		return
	}

	insertmSQL := "insert into mz_unpaid_orders (" + mSetStr + ") values " + mvValueStr
	fmt.Println("insertmSQL===", insertmSQL)

	_, errm := tx.Exec(insertmSQL)
	if errm != nil {
		fmt.Println("errm ===", errm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": "请检查是否漏填"})
		return
	}
	errc := tx.Commit()
	if errc != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errc.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//TreatmentPatientGet 查询治疗
func TreatmentPatientGet(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	rows, err := model.DB.Queryx(`select tp.*, t.name, du.name as unit_name from treatment_patient tp left join clinic_treatment ct on tp.clinic_treatment_id = ct.id 
		left join treatment t on ct.treatment_id = t.id
		left join dose_unit du on t.unit_id = du.id
		where tp.clinic_triage_patient_id = $1`, clinicTriagePatientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
	}

	result := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": result})
}

//LaboratoryPatientCreate 开检验
func LaboratoryPatientCreate(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	personnelID := ctx.PostValue("personnel_id")
	items := ctx.PostValue("items")

	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	if items == "" {
		ctx.JSON(iris.Map{"code": "200", "data": nil})
		return
	}

	var results []map[string]string
	errj := json.Unmarshal([]byte(items), &results)
	fmt.Println("results===", results)
	if errj != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": errj.Error()})
		return
	}
	row := model.DB.QueryRowx(`select id,status from clinic_triage_patient where id=$1 limit 1`, clinicTriagePatientID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "保存检验失败,分诊记录错误"})
		return
	}

	prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", personnelID)
	if prow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "保存检验失败,操作员错误"})
		return
	}
	clinicTriagePatient := FormatSQLRowToMap(row)
	personnel := FormatSQLRowToMap(prow)

	_, ok := clinicTriagePatient["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "分诊记录不存在"})
		return
	}
	status := clinicTriagePatient["status"]
	if status.(int64) != 30 {
		ctx.JSON(iris.Map{"code": "1", "msg": "分诊记录当前状态错误"})
		return
	}
	_, pok := personnel["id"]
	if !pok {
		ctx.JSON(iris.Map{"code": "1", "msg": "操作员错误"})
		return
	}

	var mzUnpaidOrdersValues []string
	mzUnpaidOrdersSets := []string{
		"clinic_triage_patient_id",
		"charge_project_type_id",
		"charge_project_id",
		"order_sn",
		"soft_sn",
		"name",
		"price",
		"amount",
		"unit",
		"total",
		"fee",
		"operation_id",
	}

	var clinicLaboratoryValues []string
	clinicLaboratorySets := []string{
		"clinic_triage_patient_id",
		"clinic_laboratory_id",
		"order_sn",
		"soft_sn",
		"times",
		"operation_id",
		"illustration",
	}

	for index, v := range results {
		clinicLaboratoryID := v["clinic_laboratory_id"]
		times := v["times"]
		illustration := v["illustration"]
		fmt.Println("clinicLaboratoryID====", clinicLaboratoryID)
		var sl []string
		var sm []string
		laboratorySQL := `select cl.id as clinic_laboratory_id,cl.price,cl.is_discount,t.name,du.name as dose_unit_name from clinic_laboratory cl
			left join laboratory t on t.id = cl.laboratory_id
			left join dose_unit du on du.id = t.unit_id
			where cl.id=$1`
		trow := model.DB.QueryRowx(laboratorySQL, clinicLaboratoryID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "检验项错误"})
			return
		}
		laboratory := FormatSQLRowToMap(trow)
		fmt.Println("====", laboratory)
		_, ok := laboratory["clinic_laboratory_id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "选择的检验项错误"})
			return
		}
		orderSn := FormatPayOrderSn(clinicTriagePatientID, "3")
		price := laboratory["price"].(int64)
		name := laboratory["name"].(string)
		unitName := laboratory["dose_unit_name"].(string)
		amount, _ := strconv.Atoi(times)
		total := int(price) * amount

		sl = append(sl, clinicTriagePatientID, clinicLaboratoryID, "'"+orderSn+"'", strconv.Itoa(index), times, personnelID)
		sm = append(sm, clinicTriagePatientID, "3", clinicLaboratoryID, "'"+orderSn+"'", strconv.Itoa(index), "'"+name+"'", strconv.FormatInt(price, 10), strconv.Itoa(amount), "'"+unitName+"'", strconv.Itoa(total), strconv.Itoa(total), personnelID)

		if illustration == "" {
			sl = append(sl, `null`)
		} else {
			sl = append(sl, "'"+illustration+"'")
		}

		tstr := "(" + strings.Join(sl, ",") + ")"
		clinicLaboratoryValues = append(clinicLaboratoryValues, tstr)
		mstr := "(" + strings.Join(sm, ",") + ")"
		mzUnpaidOrdersValues = append(mzUnpaidOrdersValues, mstr)
	}
	tSetStr := strings.Join(clinicLaboratorySets, ",")
	tValueStr := strings.Join(clinicLaboratoryValues, ",")

	mSetStr := strings.Join(mzUnpaidOrdersSets, ",")
	mvValueStr := strings.Join(mzUnpaidOrdersValues, ",")

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errb})
		return
	}
	_, errdlp := tx.Exec("delete from laboratory_patient where clinic_triage_patient_id=$1", clinicTriagePatientID)
	if errdlp != nil {
		fmt.Println("errdlp ===", errdlp)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errdlp.Error()})
		return
	}
	_, errdm := tx.Exec("delete from mz_unpaid_orders where clinic_triage_patient_id=$1 and charge_project_type_id=3", clinicTriagePatientID)
	if errdm != nil {
		fmt.Println("errdm ===", errdm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errdm.Error()})
		return
	}

	inserttSQL := "insert into laboratory_patient (" + tSetStr + ") values " + tValueStr
	fmt.Println("inserttSQL===", inserttSQL)

	_, errt := tx.Exec(inserttSQL)
	if errt != nil {
		fmt.Println("errt ===", errt)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errt.Error()})
		return
	}

	insertmSQL := "insert into mz_unpaid_orders (" + mSetStr + ") values " + mvValueStr
	fmt.Println("insertmSQL===", insertmSQL)

	_, errm := tx.Exec(insertmSQL)
	if errm != nil {
		fmt.Println("errm ===", errm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": "请检查是否漏填"})
		return
	}
	errc := tx.Commit()
	if errc != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errc.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//PrescriptionWesternPatientCreate 开西/成药处方
func PrescriptionWesternPatientCreate(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	personnelID := ctx.PostValue("personnel_id")
	items := ctx.PostValue("items")

	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	if items == "" {
		ctx.JSON(iris.Map{"code": "200", "data": nil})
		return
	}

	var results []map[string]string
	errj := json.Unmarshal([]byte(items), &results)
	fmt.Println("results===", results)
	if errj != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": errj.Error()})
		return
	}
	row := model.DB.QueryRowx(`select id,status from clinic_triage_patient where id=$1 limit 1`, clinicTriagePatientID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "保存西/成药处方失败,分诊记录错误"})
		return
	}

	prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", personnelID)
	if prow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "保存西/成药处方失败,操作员错误"})
		return
	}
	clinicTriagePatient := FormatSQLRowToMap(row)
	personnel := FormatSQLRowToMap(prow)

	_, ok := clinicTriagePatient["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "分诊记录不存在"})
		return
	}
	status := clinicTriagePatient["status"]
	if status.(int64) != 30 {
		ctx.JSON(iris.Map{"code": "1", "msg": "分诊记录当前状态错误"})
		return
	}
	_, pok := personnel["id"]
	if !pok {
		ctx.JSON(iris.Map{"code": "1", "msg": "操作员错误"})
		return
	}

	var mzUnpaidOrdersValues []string
	mzUnpaidOrdersSets := []string{
		"clinic_triage_patient_id",
		"charge_project_type_id",
		"charge_project_id",
		"order_sn",
		"soft_sn",
		"name",
		"price",
		"amount",
		"unit",
		"total",
		"fee",
		"operation_id",
	}

	var prescriptionWesternPatientValues []string
	prescriptionWesternPatientSets := []string{
		"clinic_triage_patient_id",
		"drug_stock_id",
		"order_sn",
		"soft_sn",
		"once_dose",
		"once_dose_unit_id",
		"route_administration_id",
		"frequency_id",
		"amount",
		"fetch_address",
		"operation_id",
		"eff_day",
		"illustration",
	}

	for index, v := range results {
		drugStockID := v["drug_stock_id"]
		onceDose := v["once_dose"]
		onceDoseUnitID := v["once_dose_unit_id"]
		routeAdministrationID := v["route_administration_id"]
		frequencyID := v["frequency_id"]
		times := v["amount"]
		illustration := v["illustration"]
		fetchAddress := v["fetch_address"]
		effDay := v["eff_day"]
		fmt.Println("drugStockID====", drugStockID)
		var sl []string
		var sm []string
		laboratorySQL := `select ds.id,d.name,ds.ret_price,du.name as packing_unit_name from drug_stock ds
			left join drug d on d.id = ds.drug_id
			left join dose_unit du on du.id = ds.packing_unit_id
			where ds.id=$1`
		trow := model.DB.QueryRowx(laboratorySQL, drugStockID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "西/成药处方项错误"})
			return
		}
		drugStock := FormatSQLRowToMap(trow)
		fmt.Println("====", drugStock)
		_, ok := drugStock["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "选择的西/成药处方药错误"})
			return
		}

		orderSn := FormatPayOrderSn(clinicTriagePatientID, "1")
		price := drugStock["ret_price"].(int64)
		name := drugStock["name"].(string)
		unitName := drugStock["packing_unit_name"].(string)
		amount, _ := strconv.Atoi(times)
		total := int(price) * amount

		sl = append(sl, clinicTriagePatientID, drugStockID, "'"+orderSn+"'", strconv.Itoa(index), onceDose, onceDoseUnitID, routeAdministrationID, frequencyID, times, fetchAddress, personnelID)
		sm = append(sm, clinicTriagePatientID, "1", drugStockID, "'"+orderSn+"'", strconv.Itoa(index), "'"+name+"'", strconv.FormatInt(price, 10), strconv.Itoa(amount), "'"+unitName+"'", strconv.Itoa(total), strconv.Itoa(total), personnelID)

		if illustration == "" {
			sl = append(sl, `null`)
		} else {
			sl = append(sl, "'"+illustration+"'")
		}

		if effDay == "" {
			sl = append(sl, `null`)
		} else {
			sl = append(sl, effDay)
		}

		tstr := "(" + strings.Join(sl, ",") + ")"
		prescriptionWesternPatientValues = append(prescriptionWesternPatientValues, tstr)
		mstr := "(" + strings.Join(sm, ",") + ")"
		mzUnpaidOrdersValues = append(mzUnpaidOrdersValues, mstr)
	}
	tSetStr := strings.Join(prescriptionWesternPatientSets, ",")
	tValueStr := strings.Join(prescriptionWesternPatientValues, ",")

	mSetStr := strings.Join(mzUnpaidOrdersSets, ",")
	mvValueStr := strings.Join(mzUnpaidOrdersValues, ",")

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errb})
		return
	}
	_, errdlp := tx.Exec("delete from prescription_western_patient where clinic_triage_patient_id=$1", clinicTriagePatientID)
	if errdlp != nil {
		fmt.Println("errdlp ===", errdlp)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errdlp.Error()})
		return
	}
	_, errdm := tx.Exec("delete from mz_unpaid_orders where clinic_triage_patient_id=$1 and charge_project_type_id=1", clinicTriagePatientID)
	if errdm != nil {
		fmt.Println("errdm ===", errdm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errdm.Error()})
		return
	}

	inserttSQL := "insert into prescription_western_patient (" + tSetStr + ") values " + tValueStr
	fmt.Println("inserttSQL===", inserttSQL)

	_, errt := tx.Exec(inserttSQL)
	if errt != nil {
		fmt.Println("errt ===", errt)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errt.Error()})
		return
	}

	insertmSQL := "insert into mz_unpaid_orders (" + mSetStr + ") values " + mvValueStr
	fmt.Println("insertmSQL===", insertmSQL)

	_, errm := tx.Exec(insertmSQL)
	if errm != nil {
		fmt.Println("errm ===", errm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": "请检查是否漏填"})
		return
	}
	errc := tx.Commit()
	if errc != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errc.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//LaboratoryPatientGet 获取检验
func LaboratoryPatientGet(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	rows, err := model.DB.Queryx(`select lp.*, l.name from laboratory_patient lp left join clinic_laboratory cl on lp.clinic_laboratory_id = cl.id 
		left join laboratory l on cl.laboratory_id = l.id
		where lp.clinic_triage_patient_id = $1`, clinicTriagePatientID)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
	}

	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result})
}

//ExaminationPatientCreate 开检查
func ExaminationPatientCreate(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	personnelID := ctx.PostValue("personnel_id")
	items := ctx.PostValue("items")

	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	if items == "" {
		ctx.JSON(iris.Map{"code": "200", "data": nil})
		return
	}

	var results []map[string]string
	errj := json.Unmarshal([]byte(items), &results)
	fmt.Println("results===", results)
	if errj != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": errj.Error()})
		return
	}
	row := model.DB.QueryRowx(`select id,status from clinic_triage_patient where id=$1 limit 1`, clinicTriagePatientID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "保存检查失败,分诊记录错误"})
		return
	}

	prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", personnelID)
	if prow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "保存检查失败,操作员错误"})
		return
	}
	clinicTriagePatient := FormatSQLRowToMap(row)
	personnel := FormatSQLRowToMap(prow)

	_, ok := clinicTriagePatient["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "分诊记录不存在"})
		return
	}
	status := clinicTriagePatient["status"]
	if status.(int64) != 30 {
		ctx.JSON(iris.Map{"code": "1", "msg": "分诊记录当前状态错误"})
		return
	}
	_, pok := personnel["id"]
	if !pok {
		ctx.JSON(iris.Map{"code": "1", "msg": "操作员错误"})
		return
	}

	var mzUnpaidOrdersValues []string
	mzUnpaidOrdersSets := []string{
		"clinic_triage_patient_id",
		"charge_project_type_id",
		"charge_project_id",
		"order_sn",
		"soft_sn",
		"name",
		"price",
		"amount",
		"unit",
		"total",
		"fee",
		"operation_id",
	}

	var clinicExaminationValues []string
	clinicExaminationSets := []string{
		"clinic_triage_patient_id",
		"clinic_examination_id",
		"order_sn",
		"soft_sn",
		"times",
		"organ",
		"operation_id",
		"illustration",
	}

	orderSn := FormatPayOrderSn(clinicTriagePatientID, "4")

	for index, v := range results {
		clinicExaminationID := v["clinic_examination_id"]
		times := v["times"]
		illustration := v["illustration"]
		organ := v["organ"]
		fmt.Println("clinicExaminationID====", clinicExaminationID)
		var sl []string
		var sm []string
		examinationSQL := `select ce.id as clinic_examination_id,ce.price,ce.is_discount,e.name,du.name as dose_unit_name from clinic_examination ce
		left join examination e on e.id = ce.examination_id
		left join dose_unit du on du.id = e.unit_id
		where ce.id=$1`
		trow := model.DB.QueryRowx(examinationSQL, clinicExaminationID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "检查项错误"})
			return
		}
		examination := FormatSQLRowToMap(trow)
		fmt.Println("====", examination)
		_, ok := examination["clinic_examination_id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "选择的检查项错误"})
			return
		}
		price := examination["price"].(int64)
		name := examination["name"].(string)
		unitName := examination["dose_unit_name"].(string)
		amount, _ := strconv.Atoi(times)
		total := int(price) * amount

		sl = append(sl, clinicTriagePatientID, clinicExaminationID, "'"+orderSn+"'", strconv.Itoa(index), times, "'"+organ+"'", personnelID)
		sm = append(sm, clinicTriagePatientID, "4", clinicExaminationID, "'"+orderSn+"'", strconv.Itoa(index), "'"+name+"'", strconv.FormatInt(price, 10), strconv.Itoa(amount), "'"+unitName+"'", strconv.Itoa(total), strconv.Itoa(total), personnelID)

		if illustration == "" {
			sl = append(sl, `null`)
		} else {
			sl = append(sl, "'"+illustration+"'")
		}

		tstr := "(" + strings.Join(sl, ",") + ")"
		clinicExaminationValues = append(clinicExaminationValues, tstr)
		mstr := "(" + strings.Join(sm, ",") + ")"
		mzUnpaidOrdersValues = append(mzUnpaidOrdersValues, mstr)
	}
	tSetStr := strings.Join(clinicExaminationSets, ",")
	tValueStr := strings.Join(clinicExaminationValues, ",")

	mSetStr := strings.Join(mzUnpaidOrdersSets, ",")
	mvValueStr := strings.Join(mzUnpaidOrdersValues, ",")

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errb})
		return
	}
	_, errdlp := tx.Exec("delete from examination_patient where clinic_triage_patient_id=$1", clinicTriagePatientID)
	if errdlp != nil {
		fmt.Println("errdlp ===", errdlp)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errdlp.Error()})
		return
	}
	_, errdm := tx.Exec("delete from mz_unpaid_orders where clinic_triage_patient_id=$1 and charge_project_type_id=4", clinicTriagePatientID)
	if errdm != nil {
		fmt.Println("errdm ===", errdm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errdm.Error()})
		return
	}

	inserttSQL := "insert into examination_patient (" + tSetStr + ") values " + tValueStr
	fmt.Println("inserttSQL===", inserttSQL)

	_, errt := tx.Exec(inserttSQL)
	if errt != nil {
		fmt.Println("errt ===", errt)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errt.Error()})
		return
	}

	insertmSQL := "insert into mz_unpaid_orders (" + mSetStr + ") values " + mvValueStr
	fmt.Println("insertmSQL===", insertmSQL)

	_, errm := tx.Exec(insertmSQL)
	if errm != nil {
		fmt.Println("errm ===", errm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": "请检查是否漏填"})
		return
	}
	errc := tx.Commit()
	if errc != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errc.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//ExaminationPatientGet 获取检查
func ExaminationPatientGet(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	rows, err := model.DB.Queryx(`select ep.*, e.name from examination_patient ep left join clinic_examination ce on ep.clinic_examination_id = ce.id 
		left join examination e on ce.examination_id = e.id
		where ep.clinic_triage_patient_id = $1`, clinicTriagePatientID)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
	}

	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result})
}

//MaterialPatientCreate 开材料费
func MaterialPatientCreate(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	personnelID := ctx.PostValue("personnel_id")
	items := ctx.PostValue("items")

	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	if items == "" {
		ctx.JSON(iris.Map{"code": "200", "data": nil})
		return
	}

	var results []map[string]string
	errj := json.Unmarshal([]byte(items), &results)
	fmt.Println("results===", results)
	if errj != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": errj.Error()})
		return
	}
	row := model.DB.QueryRowx(`select id,status from clinic_triage_patient where id=$1 limit 1`, clinicTriagePatientID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "保存检查失败,分诊记录错误"})
		return
	}

	prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", personnelID)
	if prow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "保存检查失败,操作员错误"})
		return
	}
	clinicTriagePatient := FormatSQLRowToMap(row)
	personnel := FormatSQLRowToMap(prow)

	_, ok := clinicTriagePatient["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "分诊记录不存在"})
		return
	}
	status := clinicTriagePatient["status"]
	if status.(int64) != 30 {
		ctx.JSON(iris.Map{"code": "1", "msg": "分诊记录当前状态错误"})
		return
	}
	_, pok := personnel["id"]
	if !pok {
		ctx.JSON(iris.Map{"code": "1", "msg": "操作员错误"})
		return
	}

	var mzUnpaidOrdersValues []string
	mzUnpaidOrdersSets := []string{
		"clinic_triage_patient_id",
		"charge_project_type_id",
		"charge_project_id",
		"order_sn",
		"soft_sn",
		"name",
		"price",
		"amount",
		"unit",
		"total",
		"fee",
		"operation_id",
	}

	var materialStockValues []string
	materialStockSets := []string{
		"clinic_triage_patient_id",
		"clinic_examination_id",
		"order_sn",
		"soft_sn",
		"times",
		"operation_id",
		"illustration",
	}

	orderSn := FormatPayOrderSn(clinicTriagePatientID, "5")

	for index, v := range results {
		clinicExaminationID := v["clinic_examination_id"]
		times := v["times"]
		illustration := v["illustration"]
		fmt.Println("clinicExaminationID====", clinicExaminationID)
		var sl []string
		var sm []string
		materialStockSQL := `select ms.id as material_stock_id,ms.price,ms.is_discount,m.name,du.name as dose_unit_name from material_stock ms
		left join material m on m.id = ms.material_id
		left join dose_unit du on du.id = m.unit_id
		where ms.id=$1`
		trow := model.DB.QueryRowx(materialStockSQL, clinicExaminationID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "检查项错误"})
			return
		}
		materialStock := FormatSQLRowToMap(trow)
		fmt.Println("====", materialStock)
		_, ok := materialStock["clinic_examination_id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "选择的检查项错误"})
			return
		}
		price := materialStock["price"].(int64)
		name := materialStock["name"].(string)
		unitName := materialStock["dose_unit_name"].(string)
		amount, _ := strconv.Atoi(times)
		total := int(price) * amount

		sl = append(sl, clinicTriagePatientID, clinicExaminationID, "'"+orderSn+"'", strconv.Itoa(index), times, personnelID)
		sm = append(sm, clinicTriagePatientID, "4", clinicExaminationID, "'"+orderSn+"'", strconv.Itoa(index), "'"+name+"'", strconv.FormatInt(price, 10), strconv.Itoa(amount), "'"+unitName+"'", strconv.Itoa(total), strconv.Itoa(total), personnelID)

		if illustration == "" {
			sl = append(sl, `null`)
		} else {
			sl = append(sl, "'"+illustration+"'")
		}

		tstr := "(" + strings.Join(sl, ",") + ")"
		materialStockValues = append(materialStockValues, tstr)
		mstr := "(" + strings.Join(sm, ",") + ")"
		mzUnpaidOrdersValues = append(mzUnpaidOrdersValues, mstr)
	}
	tSetStr := strings.Join(materialStockSets, ",")
	tValueStr := strings.Join(materialStockValues, ",")

	mSetStr := strings.Join(mzUnpaidOrdersSets, ",")
	mvValueStr := strings.Join(mzUnpaidOrdersValues, ",")

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errb})
		return
	}
	_, errdlp := tx.Exec("delete from material_patient where clinic_triage_patient_id=$1", clinicTriagePatientID)
	if errdlp != nil {
		fmt.Println("errdlp ===", errdlp)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errdlp.Error()})
		return
	}
	_, errdm := tx.Exec("delete from mz_unpaid_orders where clinic_triage_patient_id=$1 and charge_project_type_id=5", clinicTriagePatientID)
	if errdm != nil {
		fmt.Println("errdm ===", errdm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errdm.Error()})
		return
	}

	inserttSQL := "insert into material_patient (" + tSetStr + ") values " + tValueStr
	fmt.Println("inserttSQL===", inserttSQL)

	_, errt := tx.Exec(inserttSQL)
	if errt != nil {
		fmt.Println("errt ===", errt)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errt.Error()})
		return
	}

	insertmSQL := "insert into mz_unpaid_orders (" + mSetStr + ") values " + mvValueStr
	fmt.Println("insertmSQL===", insertmSQL)

	_, errm := tx.Exec(insertmSQL)
	if errm != nil {
		fmt.Println("errm ===", errm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": "请检查是否漏填"})
		return
	}
	errc := tx.Commit()
	if errc != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errc.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//MaterialPatientGet 获取
func MaterialPatientGet(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	rows, err := model.DB.Queryx(`select mp.*, m.name from material_patient mp left join material_stock ms on material_stock_id = ms.id 
		left join material m on ms.material_id = m.id
		where mp.clinic_triage_patient_id = $1`, clinicTriagePatientID)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
	}

	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result})
}
