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

	insertKeys := `(clinic_patient_id, register_personnel_id,register_type, visit_type)`
	insertValues := `($1, $2, 2, $3, )`
	if departmentID != "" {
		insertKeys = `(department_id, clinic_patient_id, register_personnel_id,register_type, visit_type)`
		insertValues = `(` + departmentID + `, $1, $2, 2, $3)`
	}

	insertSQL := "INSERT INTO clinic_triage_patient " + insertKeys + " VALUES " + insertValues + " RETURNING id"

	fmt.Println("insertSQL ======", insertSQL)

	var resultID int
	err = tx.QueryRow("INSERT INTO clinic_triage_patient "+insertKeys+" VALUES "+insertValues+" RETURNING id", clinicPatientID, personnelID, visitType).Scan(&resultID)
	if err != nil {
		fmt.Println("clinic_triage_patient ======", err)
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
	treatStatus := ctx.PostValue("treat_status")
	receptionStatus := ctx.PostValue("reception_status")
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
	left join clinic c on c.id = cp.clinic_id
	left join department d on ctp.department_id = d.id
	left join personnel rp on ctp.register_personnel_id = rp.id
	left join personnel tp on ctp.triage_personnel_id = tp.id
	left join patient p on cp.patient_id = p.id 
	where cp.clinic_id = $1 and (p.cert_no ~ $2 or p.name ~ $2 or p.phone ~ $2)`

	rowSQL := `select 
	ctp.id as clinic_triage_patient_id,ctp.clinic_patient_id,ctp.visit_date, ctp.treat_status, cp.clinic_id, c.name as clinic_name,
	p.id as patient_id, p.name as patient_name, p.birthday, p.sex, p.cert_no, p.phone,
	doc.id as doctor_id, doc.name as doctor_name,
	ctp.reception_time as reception_time,
	ctp.created_time as register_time,
	ctp.triage_time,
	ctp.complete_time,
	ctp.department_id, d.name as department_name,
	ctp.register_personnel_id, rp.name as register_personnel_name,
	ctp.triage_personnel_id, tp.name as triage_personnel_name
	from clinic_triage_patient ctp left join clinic_patient cp  on ctp.clinic_patient_id = cp.id
	left join clinic c on c.id = cp.clinic_id
	left join department d on ctp.department_id = d.id
	left join personnel rp on ctp.register_personnel_id = rp.id
	left join personnel tp on ctp.triage_personnel_id = tp.id
	left join personnel doc on ctp.doctor_id = doc.id
	left join patient p on cp.patient_id = p.id 
	where cp.clinic_id = $1 and (p.cert_no ~ $2 or p.name ~ $2 or p.phone ~ $2)`

	if treatStatus != "" {
		countSQL += " and ctp.treat_status=" + treatStatus
		rowSQL += " and ctp.treat_status=" + treatStatus
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

	if receptionStatus != "" {
		fmt.Println("====", receptionStatus)
		if receptionStatus == "true" {
			countSQL += " and ctp.reception_time is not null"
			rowSQL += " and ctp.reception_time is not null"
		} else {
			countSQL += " and ctp.reception_time is null"
			rowSQL += " and ctp.reception_time is null"
		}
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

//PersonnelChoose 线下分诊选择医生 换诊
func PersonnelChoose(ctx iris.Context) {
	doctorVisitScheduleID := ctx.PostValue("doctor_visit_schedule_id")
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	triagePersonnelID := ctx.PostValue("triage_personnel_id")

	if doctorVisitScheduleID == "" || clinicTriagePatientID == "" || triagePersonnelID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	ctprow := model.DB.QueryRowx("select id,reception_time,clinic_patient_id from clinic_triage_patient where id=$1", clinicTriagePatientID)
	clinicTriagePatient := FormatSQLRowToMap(ctprow)
	_, ctpok := clinicTriagePatient["id"]
	if !ctpok {
		ctx.JSON(iris.Map{"code": "1", "msg": "就诊人不存在"})
		return
	}

	receptionTime, ok := clinicTriagePatient["reception_time"]
	if ok && receptionTime != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "该就诊人已接诊"})
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
	err = tx.QueryRow("update clinic_triage_patient set doctor_id=$1,triage_personnel_id=$2,department_id=$3,treat_status=true,triage_time=LOCALTIMESTAMP,updated_time=LOCALTIMESTAMP where id=$4 RETURNING id", doctorID, triagePersonnelID, deparmentID, clinicTriagePatientID).Scan(&resultID)
	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
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
		where treat_status=true and reception_time is null and visit_date=current_date and doctor_id=dvs.personnel_id) as wait_total,
	(select count(ctped.id)from clinic_triage_patient ctped where 
		treat_status=true and reception_time is not null and visit_date=current_date and doctor_id=dvs.personnel_id) as triaged_total
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

	_, err1 := model.DB.Exec("delete from body_sign where clinicTriagePatientID=" + clinicTriagePatientID)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err1})
		return
	}

	insertSQL := "insert into body_sign (" + insertKeyStr + ") values (" + insertValueStr + ") RETURNING id;"
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
	_, err1 := model.DB.Exec("delete from pre_medical_record where clinicTriagePatientID=" + clinicTriagePatientID)
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

	_, err1 := model.DB.Exec("delete from pre_diagnosis where clinicTriagePatientID=" + clinicTriagePatientID)
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
	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	row := model.DB.QueryRowx("select id from clinic_triage_patient where id=$1", clinicTriagePatientID)
	clinicTriagePatient := FormatSQLRowToMap(row)
	_, ok := clinicTriagePatient["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "就诊人不存在"})
		return
	}
	_, err := model.DB.Exec("update clinic_triage_patient set reception_time=LOCALTIMESTAMP where id=$1", clinicTriagePatientID)
	if err != nil {
		fmt.Println("接诊错误", err)
		ctx.JSON(iris.Map{"code": "1", "msg": "就诊失败"})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "msg": "ok"})
}
