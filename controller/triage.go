package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris"
)

// TriageRegister 就诊患者登记
func TriageRegister(ctx iris.Context) {
	certNo := ctx.PostValue("cert_no")
	name := ctx.PostValue("name")
	birthday := ctx.PostValue("birthday")
	sex := ctx.PostValue("sex")
	phone := ctx.PostValue("phone")
	address := ctx.PostValue("address")
	profession := ctx.PostValue("profession")
	remark := ctx.PostValue("remark")
	patientChannelID := ctx.PostValue("patient_channel_id")
	clinicID := ctx.PostValue("clinic_id")
	personnelID := ctx.PostValue("personnel_id")
	departmentID := ctx.PostValue("department_id")
	if certNo == "" || name == "" || birthday == "" || sex == "" || phone == "" || patientChannelID == "" || clinicID == "" || personnelID == "" || departmentID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	todayHour := time.Now().Hour()
	amPm := "p"
	if todayHour < 12 {
		amPm = "a"
	}
	row := model.DB.QueryRowx("select * from doctor_visit_schedule where visit_date = CURRENT_DATE and am_pm = $1 and department_id = $2 and personnel_id = $3", amPm, departmentID, personnelID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "登记失败"})
		return
	}
	schedule := FormatSQLRowToMap(row)
	_, ok := schedule["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "号源不存在"})
		return
	}
	row = model.DB.QueryRowx("select * from patient where cert_no = $1", certNo)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "登记失败"})
		return
	}
	tx, err := model.DB.Begin()
	patient := FormatSQLRowToMap(row)
	_, ok = patient["id"]
	patientID := patient["id"]
	if !ok {
		err = tx.QueryRow(`INSERT INTO patient (
		cert_no, name, birthday, sex, phone, address, profession, remark, patient_channel_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`, certNo, name, birthday, sex, phone, address, profession, remark, patientChannelID).Scan(&patientID)
		if err != nil {
			tx.Rollback()
			fmt.Println("err2 ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
	} else {
		_, err = tx.Exec(`update patient set cert_no = $1, name= $2,birthday=$3,sex=$4, phone=$5, address=$6,profession = $7,remark= $8 ,patient_channel_id = $9  where id = $10`, certNo, name, birthday, sex, phone, address, profession, remark, patientChannelID, patientID)
		if err != nil {
			tx.Rollback()
			fmt.Println("err2 ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
	}

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

	var resultID int
	err = tx.QueryRow("INSERT INTO clinic_triage_patient (department_id, clinic_patient_id, register_personnel_id,register_type) VALUES ($1, $2, $3, 2) RETURNING id", departmentID, clinicPatientID, personnelID).Scan(&resultID)
	if err != nil {
		fmt.Println("clinic_triage_patient ======", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": nil})
}

// TriagePatientList 当日登记就诊人列表
func TriagePatientList(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	keyword := ctx.PostValue("keyword")
	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	rowSQL := `select 
	ctp.id,ctp.clinic_patient_id,ctp.visit_date, ctp.treat_status, cp.clinic_id, c.name as clinic_name,
	p.id as patient_id, p.name as patient_name, p.birthday, p.sex, p.cert_no, p.phone,
	ctp.created_time as register_time,
	ctp.department_id, d.name as department_name,
	ctp.register_personnel_id, rp.name as register_personnel_name,
	ctp.triage_personnel_id, tp.name as triage_personnel_name
	from clinic_triage_patient ctp left join clinic_patient cp  on ctp.clinic_patient_id = cp.id
	left join clinic c on c.id = cp.clinic_id
	left join department d on ctp.department_id = d.id
	left join personnel rp on ctp.register_personnel_id = rp.id
	left join personnel tp on ctp.triage_personnel_id = tp.id
	left join patient p on cp.patient_id = p.id 
	where cp.clinic_id = $1 and ctp.visit_date = CURRENT_DATE 
	and (p.cert_no like '%' || $2 || '%' or p.name like '%' || $2 || '%' or p.phone like '%' || $2 || '%') `
	rows, err1 := model.DB.Queryx(rowSQL, clinicID, keyword)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result})
}

//PersonnelChoose 线下分诊选择医生
func PersonnelChoose(ctx iris.Context) {
	clinicPatientID := ctx.PostValue("clinic_patient_id")
	deparmentID := ctx.PostValue("department_id")
	clinicTriagePatientID := ctx.PostValue("id")
	doctorID := ctx.PostValue("doctor_id")
	triagePersonnelID := ctx.PostValue("personnel_id")
	amPm := ctx.PostValue("am_pm")
	visitTypeCode := ctx.PostValue("visit_type_code")

	if clinicPatientID == "" || deparmentID == "" || clinicTriagePatientID == "" || doctorID == "" || triagePersonnelID == "" || amPm == "" || visitTypeCode == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	tx, err := model.DB.Begin()
	var resultID int
	err = tx.QueryRow("update clinic_triage_patient set doctor_id=$1,triage_personnel_id=$2,treat_status=true,triage_time=LOCALTIMESTAMP where id=$3 RETURNING id", doctorID, triagePersonnelID, clinicTriagePatientID).Scan(&resultID)
	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	var registrationID int
	err = tx.QueryRow(`INSERT INTO registration (
		clinic_patient_id, department_id, clinic_triage_patient_id,personnel_id,visit_date,am_pm,visit_type_code,operation_id) 
		VALUES ($1, $2, $3, $4, CURRENT_DATE, $5, $6, $7) RETURNING id`, clinicPatientID, deparmentID, clinicTriagePatientID, doctorID, amPm, visitTypeCode, triagePersonnelID).Scan(&registrationID)
	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
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
		where treat_status=false and visit_date=current_date and doctor_id=dvs.personnel_id) as waitTotal,
	(select count(ctped.id)from clinic_triage_patient ctped where 
		treat_status=true and visit_date=current_date and doctor_id=dvs.personnel_id) as triagedTotal
	from doctor_visit_schedule dvs 
	left join department d on dvs.department_id = d.id 
	left join personnel p on dvs.personnel_id = p.id
	where p.clinic_id = $1 and (p.name like '%' || $2 || '%') and dvs.am_pm=$3 and dvs.visit_date=current_date`

	if deparmentID != "" {
		countSQL += " and dvs.department_id=" + deparmentID
		selectSQL += " and dvs.department_id=" + deparmentID
	}

	total := model.DB.QueryRowx(`select count(distinct(dvs.personnel_id,dvs.department_id)) as total `+countSQL, clinicID, keyword, ampm)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rowSQL := `select distinct(p.name, d.name), p.name as doctor_name, d.name as department_name, dvs.id, dvs.am_pm, dvs.visit_date, ` + selectSQL + " offset $4 limit $5"

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
