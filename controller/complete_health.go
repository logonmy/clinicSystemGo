package controller

import (
	"clinicSystemGo/model"

	"github.com/kataras/iris"
)

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
	bloodSugarConcentration := ctx.PostValue("blood_sugar_concentration")
	leftVision := ctx.PostValue("left_vision")
	rightVision := ctx.PostValue("right_vision")
	oxygenSaturation := ctx.PostValue("oxygen_saturation")
	painScore := ctx.PostValue("pain_score")
	remark := ctx.PostValue("remark")

	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	tx, err := model.DB.Begin()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	_, err = tx.Exec("delete from body_sign where clinic_triage_patient_id=$1", clinicTriagePatientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	insertSQL := `insert into body_sign (
		clinic_triage_patient_id,
		weight,
		height,
		bmi,
		blood_type,
		rh_blood_type,
		temperature_type,
		temperature,
		breathe,
		pulse,
		systolic_blood_pressure,
		diastolic_blood_pressure,
		blood_sugar_time,
		blood_sugar_type,
		blood_sugar_concentration,
		left_vision,
		right_vision,
		oxygen_saturation,
		pain_score,
		remark
	) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20) RETURNING id;`
	_, err = tx.Exec(insertSQL,
		ToNullInt64(clinicTriagePatientID),
		ToNullFloat64(weight),
		ToNullFloat64(height),
		ToNullFloat64(bmi),
		ToNullString(bloodType),
		ToNullInt64(rhBloodType),
		ToNullInt64(temperatureType),
		ToNullFloat64(temperature),
		ToNullInt64(breathe),
		ToNullInt64(pulse),
		ToNullInt64(systolicBloodPressure),
		ToNullInt64(diastolicBloodPressure),
		ToNullString(bloodSugarTime),
		ToNullInt64(bloodSugarType),
		ToNullFloat64(bloodSugarConcentration),
		ToNullString(leftVision),
		ToNullString(rightVision),
		ToNullFloat64(oxygenSaturation),
		ToNullInt64(painScore),
		ToNullString(remark))

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
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
	remark := ctx.PostValue("remark")
	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	tx, err := model.DB.Begin()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	_, err = tx.Exec("delete from pre_medical_record where clinic_triage_patient_id=$1", clinicTriagePatientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	insertSQL := `insert into pre_medical_record (
		clinic_triage_patient_id,
		has_allergic_history,
		allergic_history,
		personal_medical_history,
		family_medical_history,
		vaccination,
		menarche_age,
		menstrual_period_start_day,
		menstrual_period_end_day,
		menstrual_cycle_start_day,
		menstrual_cycle_end_day,
		menstrual_last_day,
		gestational_weeks,
		childbearing_history,
		remark
	) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id;`
	_, err = tx.Exec(insertSQL,
		ToNullInt64(clinicTriagePatientID),
		ToNullBool(hasAllergicHistory),
		ToNullString(allergicHistory),
		ToNullString(personalMedicalHistory),
		ToNullString(familyMedicalHistory),
		ToNullString(vaccination),
		ToNullInt64(menarcheAge),
		ToNullString(menstrualPeriodStartDay),
		ToNullString(menstrualPeriodEndDay),
		ToNullString(menstrualCycleStartDay),
		ToNullString(menstrualCycleEndDay),
		ToNullString(menstrualLastDay),
		ToNullInt64(gestationalWeeks),
		ToNullString(childbearingHistory),
		ToNullString(remark))
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
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
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	tx, err := model.DB.Begin()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	_, err = tx.Exec("delete from pre_diagnosis where clinic_triage_patient_id=$1", clinicTriagePatientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	insertSQL := `insert into pre_diagnosis (
		clinic_triage_patient_id,
		chief_complaint,
		history_of_rresent_illness,
		history_of_past_illness,
		physical_examination,
		remark
	) values ($1, $2, $3, $4, $5, $6) RETURNING id;`

	_, err = tx.Exec(insertSQL,
		ToNullInt64(clinicTriagePatientID),
		ToNullString(chiefComplaint),
		ToNullString(historyOfRresentIllness),
		ToNullString(historyOfPastIllness),
		ToNullString(physicalExamination),
		ToNullString(remark))

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "保存成功"})
}

// GetHealthRecord 获取健康档案
func GetHealthRecord(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")

	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	bodySignRow := model.DB.QueryRowx(`select * from body_sign where clinic_triage_patient_id = $1`, clinicTriagePatientID)
	preMedicalRecordRow := model.DB.QueryRowx(`select * from pre_medical_record where clinic_triage_patient_id = $1`, clinicTriagePatientID)
	preDiagnosisRow := model.DB.QueryRowx(`select * from pre_diagnosis where clinic_triage_patient_id = $1`, clinicTriagePatientID)
	bodySignMap := FormatSQLRowToMap(bodySignRow)
	preMedicalRecordMap := FormatSQLRowToMap(preMedicalRecordRow)
	preDiagnosisMap := FormatSQLRowToMap(preDiagnosisRow)

	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "body_sign": bodySignMap, "pre_medical_record": preMedicalRecordMap, "pre_diagnosis": preDiagnosisMap})
}
