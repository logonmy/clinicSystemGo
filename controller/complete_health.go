package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/kataras/iris"
)

// GetLastBodySign 获取最后一次体征信息
func GetLastBodySign(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")

	if patientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	querySQL := `select p.id as patient_id, bs.* from body_sign bs left join clinic_triage_patient ctp on bs.clinic_triage_patient_id = ctp.id
	left join clinic_patient cp on ctp.clinic_patient_id = cp.id
	left join patient p on cp.patient_id = p.id where p.id = $1 order by bs.id DESC LIMIT 1`

	row := model.DB.QueryRowx(querySQL, patientID)
	result := FormatSQLRowToMap(row)
	ctx.JSON(iris.Map{"code": "200", "data": result})
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

	selectSQL := `select ctp.id, cp.patient_id
	from clinic_triage_patient ctp
	left join clinic_patient cp on cp.id = ctp.clinic_patient_id
	where ctp.id=$1`

	row := model.DB.QueryRowx(selectSQL, clinicTriagePatientID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询分诊记录失败"})
		return
	}

	clinicTriagePatient := FormatSQLRowToMap(row)
	_, ok := clinicTriagePatient["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "分诊记录不存在"})
		return
	}
	patientID := clinicTriagePatient["patient_id"].(int64)
	patientIDStr := strconv.FormatInt(patientID, 10)
	recordTime := time.Now().Format("2006-01-02")

	tx, err := model.DB.Begin()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	_, err = tx.Exec("delete from body_sign where clinic_triage_patient_id=$1", clinicTriagePatientID)
	if err != nil {
		tx.Rollback()
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
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	if height != "" {
		var results []map[string]string
		results = append(results, map[string]string{"record_time": recordTime, "height": height, "upsert_type": "insert"})
		err := upsertPatientHeight(patientIDStr, results)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	}

	if weight != "" {
		var results []map[string]string
		results = append(results, map[string]string{"record_time": recordTime, "weight": weight, "upsert_type": "insert"})
		err := upsertPatientWeight(patientIDStr, results)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	}

	if bmi != "" {
		var results []map[string]string
		results = append(results, map[string]string{"record_time": recordTime, "bmi": bmi, "upsert_type": "insert"})
		err := upsertPatientBmi(patientIDStr, results)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	}

	if bloodType != "" {
		var results []map[string]string
		results = append(results, map[string]string{"record_time": recordTime, "blood_type": bloodType, "upsert_type": "insert"})
		err := upsertPatientBloodType(patientIDStr, results)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	}

	if rhBloodType != "" {
		var results []map[string]string
		results = append(results, map[string]string{"record_time": recordTime, "rh_blood_type": rhBloodType, "upsert_type": "insert"})
		err := upsertPatientRhBloodType(patientIDStr, results)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	}

	if temperature != "" && temperatureType != "" {
		var results []map[string]string
		results = append(results, map[string]string{"record_time": recordTime, "temperature_type": temperatureType, "temperature": temperature, "upsert_type": "insert"})
		err := upsertPatientTemperature(patientIDStr, results)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	}

	if breathe != "" {
		var results []map[string]string
		results = append(results, map[string]string{"record_time": recordTime, "breathe": breathe, "upsert_type": "insert"})
		err := upsertPatientBreathe(patientIDStr, results)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	}

	if pulse != "" {
		var results []map[string]string
		results = append(results, map[string]string{"record_time": recordTime, "pulse": pulse, "upsert_type": "insert"})
		err := upsertPatientPulse(patientIDStr, results)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	}

	if systolicBloodPressure != "" && diastolicBloodPressure != "" {
		var results []map[string]string
		results = append(results, map[string]string{"record_time": recordTime, "systolic_blood_pressure": systolicBloodPressure, "diastolic_blood_pressure": diastolicBloodPressure, "upsert_type": "insert"})
		err := upsertPatientBloodPressure(patientIDStr, results)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	}

	if leftVision != "" && rightVision != "" {
		var results []map[string]string
		results = append(results, map[string]string{"record_time": recordTime, "left_vision": leftVision, "right_vision": rightVision, "upsert_type": "insert"})
		err := upsertPatientVision(patientIDStr, results)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	}

	if oxygenSaturation != "" {
		var results []map[string]string
		results = append(results, map[string]string{"record_time": recordTime, "oxygen_saturation": oxygenSaturation, "upsert_type": "insert"})
		err := upsertPatientOxygenSaturation(patientIDStr, results)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
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

	selectSQL := `select cp.patient_id,ctp.id
	from clinic_triage_patient ctp
	left join clinic_patient cp on cp.id = ctp.clinic_patient_id
	where ctp.id=$1`
	ctrow := model.DB.QueryRowx(selectSQL, clinicTriagePatientID)
	if ctrow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询分诊记录错误"})
		return
	}

	clinicTriagePatient := FormatSQLRowToMap(ctrow)
	_, ok := clinicTriagePatient["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "分诊记录不存在"})
		return
	}
	patientID := clinicTriagePatient["patient_id"]

	tx, err := model.DB.Begin()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	_, err = tx.Exec("delete from pre_medical_record where clinic_triage_patient_id=$1", clinicTriagePatientID)
	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	_, perr := tx.Exec("delete from personal_medical_record where patient_id=$1", patientID)
	if perr != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": perr.Error()})
		return
	}

	insertpSQL := `insert into personal_medical_record (
		patient_id,
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

	_, err = tx.Exec(insertpSQL,
		ToNullInt64(patientID.(string)),
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
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	_, perr = tx.Exec(insertSQL,
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
	if perr != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": perr.Error()})
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

// UpsertPatientHeight 修改身高
func UpsertPatientHeight(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	items := ctx.PostValue("items")
	if patientID == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	var results []map[string]string
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = upsertPatientHeight(patientID, results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "保存成功"})
}

// upsertPatientHeight 修改身高
func upsertPatientHeight(patientID string, results []map[string]string) error {
	if patientID == "" || results == nil {
		return errors.New("缺少参数")
	}

	tx, err := model.DB.Beginx()

	if err != nil {
		return err
	}

	for _, v := range results {
		recordTime, rok := v["record_time"]
		height, hok := v["height"]
		upsertType, tok := v["upsert_type"]
		if !rok || recordTime == "" || !hok || height == "" || !tok || upsertType == "" {
			return errors.New("缺少参数")
		}

		var err error
		if upsertType == "update" {
			exceSQL := `update patient_height set height = $1 where patient_id = $2 and record_time = $3`
			_, err = tx.Exec(exceSQL, ToNullFloat64(height), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "insert" {
			exceSQL := `insert into patient_height (height, patient_id, record_time) values ($1, $2, $3)`
			_, err = tx.Exec(exceSQL, ToNullFloat64(height), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "delete" {
			exceSQL := `delete from patient_height where patient_id = $1 and record_time = $2`
			_, err = tx.Exec(exceSQL, ToNullInt64(patientID), ToNullString(recordTime))
		} else {
			return errors.New("upsert_type 值为 update，insert，delete ")
		}
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// UpsertPatientWeight 修改体重
func UpsertPatientWeight(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	items := ctx.PostValue("items")
	if patientID == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	var results []map[string]string
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = upsertPatientWeight(patientID, results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "保存成功"})
}

// upsertPatientWeight 修改体重
func upsertPatientWeight(patientID string, results []map[string]string) error {
	if patientID == "" || results == nil {
		return errors.New("缺少参数")
	}

	tx, err := model.DB.Beginx()

	if err != nil {
		return err
	}

	for _, v := range results {
		recordTime, rok := v["record_time"]
		weight, hok := v["weight"]
		upsertType, tok := v["upsert_type"]
		if !rok || recordTime == "" || !hok || weight == "" || !tok || upsertType == "" {
			return errors.New("缺少参数")
		}

		var err error
		if upsertType == "update" {
			exceSQL := `update patient_weight set weight = $1 where patient_id = $2 and record_time = $3`
			_, err = tx.Exec(exceSQL, ToNullFloat64(weight), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "insert" {
			exceSQL := `insert into patient_weight (weight, patient_id, record_time) values ($1, $2, $3)`
			_, err = tx.Exec(exceSQL, ToNullFloat64(weight), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "delete" {
			exceSQL := `delete from patient_weight where patient_id = $1 and record_time = $2`
			_, err = tx.Exec(exceSQL, ToNullInt64(patientID), ToNullString(recordTime))
		} else {
			return errors.New("upsert_type 值为 update，insert，delete ")
		}
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// UpsertPatientBmi 修改BMI
func UpsertPatientBmi(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	items := ctx.PostValue("items")
	if patientID == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	var results []map[string]string
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = upsertPatientBmi(patientID, results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "保存成功"})
}

// upsertPatientBmi 修改BMI
func upsertPatientBmi(patientID string, results []map[string]string) error {
	if patientID == "" || results == nil {
		return errors.New("缺少参数")
	}

	tx, err := model.DB.Beginx()

	if err != nil {
		return err
	}

	for _, v := range results {
		recordTime, rok := v["record_time"]
		bmi, hok := v["bmi"]
		upsertType, tok := v["upsert_type"]
		if !rok || recordTime == "" || !hok || bmi == "" || !tok || upsertType == "" {
			return errors.New("缺少参数")
		}

		var err error
		if upsertType == "update" {
			exceSQL := `update patient_bmi set bmi = $1 where patient_id = $2 and record_time = $3`
			_, err = tx.Exec(exceSQL, ToNullFloat64(bmi), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "insert" {
			exceSQL := `insert into patient_bmi (bmi, patient_id, record_time) values ($1, $2, $3)`
			_, err = tx.Exec(exceSQL, ToNullFloat64(bmi), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "delete" {
			exceSQL := `delete from patient_bmi where patient_id = $1 and record_time = $2`
			_, err = tx.Exec(exceSQL, ToNullInt64(patientID), ToNullString(recordTime))
		} else {
			return errors.New("upsert_type 值为 update，insert，delete ")
		}
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// UpsertPatientBloodType 修改血型
func UpsertPatientBloodType(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	items := ctx.PostValue("items")
	if patientID == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	var results []map[string]string
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = upsertPatientBloodType(patientID, results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "保存成功"})
}

// upsertPatientBloodType 修改血型
func upsertPatientBloodType(patientID string, results []map[string]string) error {
	if patientID == "" || results == nil {
		return errors.New("缺少参数")
	}

	tx, err := model.DB.Beginx()

	if err != nil {
		return err
	}

	for _, v := range results {
		recordTime, rok := v["record_time"]
		bloodType, hok := v["blood_type"]
		upsertType, tok := v["upsert_type"]
		if !rok || recordTime == "" || !hok || bloodType == "" || !tok || upsertType == "" {
			return errors.New("缺少参数")
		}

		var err error
		if upsertType == "update" {
			exceSQL := `update patient_blood_type set blood_type = $1 where patient_id = $2 and record_time = $3`
			_, err = tx.Exec(exceSQL, ToNullString(bloodType), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "insert" {
			exceSQL := `insert into patient_blood_type (blood_type, patient_id, record_time) values ($1, $2, $3)`
			_, err = tx.Exec(exceSQL, ToNullString(bloodType), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "delete" {
			exceSQL := `delete from patient_blood_type where patient_id = $1 and record_time = $2`
			_, err = tx.Exec(exceSQL, ToNullInt64(patientID), ToNullString(recordTime))
		} else {
			return errors.New("upsert_type 值为 update，insert，delete ")
		}
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// UpsertPatientRhBloodType 修改RH血型
func UpsertPatientRhBloodType(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	items := ctx.PostValue("items")
	if patientID == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	var results []map[string]string
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = upsertPatientRhBloodType(patientID, results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "保存成功"})
}

// upsertPatientRhBloodType 修改RH血型
func upsertPatientRhBloodType(patientID string, results []map[string]string) error {
	if patientID == "" || results == nil {
		return errors.New("缺少参数")
	}

	tx, err := model.DB.Beginx()

	if err != nil {
		return err
	}

	for _, v := range results {
		recordTime, rok := v["record_time"]
		RhBloodType, hok := v["rh_blood_type"]
		upsertType, tok := v["upsert_type"]
		if !rok || recordTime == "" || !hok || RhBloodType == "" || !tok || upsertType == "" {
			return errors.New("缺少参数")
		}

		var err error
		if upsertType == "update" {
			exceSQL := `update patient_rh_blood_type set rh_blood_type = $1 where patient_id = $2 and record_time = $3`
			_, err = tx.Exec(exceSQL, ToNullInt64(RhBloodType), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "insert" {
			exceSQL := `insert into patient_rh_blood_type (rh_blood_type, patient_id, record_time) values ($1, $2, $3)`
			_, err = tx.Exec(exceSQL, ToNullInt64(RhBloodType), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "delete" {
			exceSQL := `delete from patient_rh_blood_type where patient_id = $1 and record_time = $2`
			_, err = tx.Exec(exceSQL, ToNullInt64(patientID), ToNullString(recordTime))
		} else {
			return errors.New("upsert_type 值为 update，insert，delete ")
		}
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// UpsertPatientTemperature 修改体温
func UpsertPatientTemperature(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	items := ctx.PostValue("items")
	if patientID == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	var results []map[string]string
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = upsertPatientTemperature(patientID, results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "保存成功"})
}

// upsertPatientTemperature 修改体温
func upsertPatientTemperature(patientID string, results []map[string]string) error {
	if patientID == "" || results == nil {
		return errors.New("缺少参数")
	}

	tx, err := model.DB.Beginx()

	if err != nil {
		return err
	}

	for _, v := range results {
		recordTime, rok := v["record_time"]
		temperatureType, ttok := v["temperature_type"]
		temperature, tok := v["temperature"]
		upsertType, uok := v["upsert_type"]
		if !rok || recordTime == "" || !ttok || temperatureType == "" || !tok || temperature == "" || !uok || upsertType == "" {
			return errors.New("缺少参数")
		}

		var err error
		if upsertType == "update" {
			exceSQL := `update patient_temperature set temperature_type = $1,temperature = $2 where patient_id = $3 and record_time = $4`
			_, err = tx.Exec(exceSQL, ToNullInt64(temperatureType), ToNullFloat64(temperature), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "insert" {
			exceSQL := `insert into patient_temperature (temperature_type, temperature, patient_id, record_time) values ($1, $2, $3, $4)`
			_, err = tx.Exec(exceSQL, ToNullInt64(temperatureType), ToNullFloat64(temperature), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "delete" {
			exceSQL := `delete from patient_temperature where patient_id = $1 and record_time = $2`
			_, err = tx.Exec(exceSQL, ToNullInt64(patientID), ToNullString(recordTime))
		} else {
			return errors.New("upsert_type 值为 update，insert，delete ")
		}
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// UpsertPatientBreathe 修改呼吸
func UpsertPatientBreathe(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	items := ctx.PostValue("items")
	if patientID == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	var results []map[string]string
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = upsertPatientBreathe(patientID, results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "保存成功"})
}

// upsertPatientBreathe 修改呼吸
func upsertPatientBreathe(patientID string, results []map[string]string) error {
	if patientID == "" || results == nil {
		return errors.New("缺少参数")
	}

	tx, err := model.DB.Beginx()

	if err != nil {
		return err
	}

	for _, v := range results {
		recordTime, rok := v["record_time"]
		breathe, hok := v["breathe"]
		upsertType, tok := v["upsert_type"]
		if !rok || recordTime == "" || !hok || breathe == "" || !tok || upsertType == "" {
			return errors.New("缺少参数")
		}

		var err error
		if upsertType == "update" {
			exceSQL := `update patient_breathe set breathe = $1 where patient_id = $2 and record_time = $3`
			_, err = tx.Exec(exceSQL, ToNullInt64(breathe), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "insert" {
			exceSQL := `insert into patient_breathe (breathe, patient_id, record_time) values ($1, $2, $3)`
			_, err = tx.Exec(exceSQL, ToNullInt64(breathe), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "delete" {
			exceSQL := `delete from patient_breathe where patient_id = $1 and record_time = $2`
			_, err = tx.Exec(exceSQL, ToNullInt64(patientID), ToNullString(recordTime))
		} else {
			return errors.New("upsert_type 值为 update，insert，delete ")
		}
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// UpsertPatientPulse 修改脉搏
func UpsertPatientPulse(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	items := ctx.PostValue("items")
	if patientID == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	var results []map[string]string
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = upsertPatientPulse(patientID, results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "保存成功"})
}

// upsertPatientPulse 修改脉搏
func upsertPatientPulse(patientID string, results []map[string]string) error {
	if patientID == "" || results == nil {
		return errors.New("缺少参数")
	}

	tx, err := model.DB.Beginx()

	if err != nil {
		return err
	}

	for _, v := range results {
		recordTime, rok := v["record_time"]
		pulse, hok := v["pulse"]
		upsertType, tok := v["upsert_type"]
		if !rok || recordTime == "" || !hok || pulse == "" || !tok || upsertType == "" {
			return errors.New("缺少参数")
		}

		var err error
		if upsertType == "update" {
			exceSQL := `update patient_pulse set pulse = $1 where patient_id = $2 and record_time = $3`
			_, err = tx.Exec(exceSQL, ToNullInt64(pulse), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "insert" {
			exceSQL := `insert into patient_pulse (pulse, patient_id, record_time) values ($1, $2, $3)`
			_, err = tx.Exec(exceSQL, ToNullInt64(pulse), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "delete" {
			exceSQL := `delete from patient_pulse where patient_id = $1 and record_time = $2`
			_, err = tx.Exec(exceSQL, ToNullInt64(patientID), ToNullString(recordTime))
		} else {
			return errors.New("upsert_type 值为 update，insert，delete ")
		}
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// UpsertPatientBloodPressure 修改血压
func UpsertPatientBloodPressure(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	items := ctx.PostValue("items")
	if patientID == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	var results []map[string]string
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = upsertPatientBloodPressure(patientID, results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "保存成功"})
}

// upsertPatientBloodPressure 修改血压
func upsertPatientBloodPressure(patientID string, results []map[string]string) error {
	if patientID == "" || results == nil {
		return errors.New("缺少参数")
	}

	tx, err := model.DB.Beginx()

	if err != nil {
		return err
	}

	for _, v := range results {
		recordTime, rok := v["record_time"]
		systolicBloodPressure, ttok := v["systolic_blood_pressure"]
		diastolicBloodPressure, tok := v["diastolic_blood_pressure"]
		upsertType, uok := v["upsert_type"]
		if !rok || recordTime == "" || !ttok || systolicBloodPressure == "" || !tok || diastolicBloodPressure == "" || !uok || upsertType == "" {
			return errors.New("缺少参数")
		}

		var err error
		if upsertType == "update" {
			exceSQL := `update patient_blood_pressure set systolic_blood_pressure = $1,diastolic_blood_pressure = $2 where patient_id = $3 and record_time = $4`
			_, err = tx.Exec(exceSQL, ToNullInt64(systolicBloodPressure), ToNullInt64(diastolicBloodPressure), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "insert" {
			exceSQL := `insert into patient_blood_pressure (systolic_blood_pressure, diastolic_blood_pressure, patient_id, record_time) values ($1, $2, $3, $4)`
			_, err = tx.Exec(exceSQL, ToNullInt64(systolicBloodPressure), ToNullInt64(diastolicBloodPressure), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "delete" {
			exceSQL := `delete from patient_blood_pressure where patient_id = $1 and record_time = $2`
			_, err = tx.Exec(exceSQL, ToNullInt64(patientID), ToNullString(recordTime))
		} else {
			return errors.New("upsert_type 值为 update，insert，delete ")
		}
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// UpsertPatientVision 修改视力
func UpsertPatientVision(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	items := ctx.PostValue("items")
	if patientID == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	var results []map[string]string
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = upsertPatientVision(patientID, results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "保存成功"})
}

// upsertPatientVision 修改视力
func upsertPatientVision(patientID string, results []map[string]string) error {
	if patientID == "" || results == nil {
		return errors.New("缺少参数")
	}

	tx, err := model.DB.Beginx()

	if err != nil {
		return err
	}

	for _, v := range results {
		recordTime, rok := v["record_time"]
		leftVision, ttok := v["left_vision"]
		rightVision, tok := v["right_vision"]
		upsertType, uok := v["upsert_type"]
		if !rok || recordTime == "" || !ttok || leftVision == "" || !tok || rightVision == "" || !uok || upsertType == "" {
			return errors.New("缺少参数")
		}

		var err error
		if upsertType == "update" {
			exceSQL := `update patient_vision set left_vision = $1,right_vision = $2 where patient_id = $3 and record_time = $4`
			_, err = tx.Exec(exceSQL, ToNullString(leftVision), ToNullString(rightVision), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "insert" {
			exceSQL := `insert into patient_vision (left_vision, right_vision, patient_id, record_time) values ($1, $2, $3, $4)`
			_, err = tx.Exec(exceSQL, ToNullString(leftVision), ToNullString(rightVision), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "delete" {
			exceSQL := `delete from patient_vision where patient_id = $1 and record_time = $2`
			_, err = tx.Exec(exceSQL, ToNullInt64(patientID), ToNullString(recordTime))
		} else {
			return errors.New("upsert_type 值为 update，insert，delete ")
		}
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// UpsertPatientOxygenSaturation 修改氧饱和度
func UpsertPatientOxygenSaturation(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	items := ctx.PostValue("items")
	if patientID == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	var results []map[string]string
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = upsertPatientOxygenSaturation(patientID, results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "保存成功"})
}

// upsertPatientOxygenSaturation 修改氧饱和度
func upsertPatientOxygenSaturation(patientID string, results []map[string]string) error {
	if patientID == "" || results == nil {
		return errors.New("缺少参数")
	}

	tx, err := model.DB.Beginx()

	if err != nil {
		return err
	}

	for _, v := range results {
		recordTime, rok := v["record_time"]
		oxygenSaturation, hok := v["oxygen_saturation"]
		upsertType, tok := v["upsert_type"]
		if !rok || recordTime == "" || !hok || oxygenSaturation == "" || !tok || upsertType == "" {
			return errors.New("缺少参数")
		}

		var err error
		if upsertType == "update" {
			exceSQL := `update patient_oxygen_saturation set oxygen_saturation = $1 where patient_id = $2 and record_time = $3`
			_, err = tx.Exec(exceSQL, ToNullFloat64(oxygenSaturation), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "insert" {
			exceSQL := `insert into patient_oxygen_saturation (oxygen_saturation, patient_id, record_time) values ($1, $2, $3)`
			_, err = tx.Exec(exceSQL, ToNullFloat64(oxygenSaturation), ToNullInt64(patientID), ToNullString(recordTime))
		} else if upsertType == "delete" {
			exceSQL := `delete from patient_oxygen_saturation where patient_id = $1 and record_time = $2`
			_, err = tx.Exec(exceSQL, ToNullInt64(patientID), ToNullString(recordTime))
		} else {
			return errors.New("upsert_type 值为 update，insert，delete ")
		}
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

//PatientHeightList 患者身高记录
func PatientHeightList(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if patientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
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

	countSQL := `select count(id) as total from patient_height where patient_id=$1`
	rowSQL := `SELECT * FROM patient_height WHERE patient_id = $1`

	total := model.DB.QueryRowx(countSQL, patientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.Queryx(rowSQL, patientID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	patientHeightLists := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": patientHeightLists, "page_info": pageInfo})
}

//PatientWeightList 患者体重记录
func PatientWeightList(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if patientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
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

	countSQL := `select count(id) as total from patient_weight where patient_id=$1`
	rowSQL := `SELECT * FROM patient_weight WHERE patient_id = $1`

	total := model.DB.QueryRowx(countSQL, patientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.Queryx(rowSQL, patientID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	patientHeightLists := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": patientHeightLists, "page_info": pageInfo})
}

//PatientBmiList 患者BMI记录
func PatientBmiList(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if patientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
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

	countSQL := `select count(id) as total from patient_bmi where patient_id=$1`
	rowSQL := `SELECT * FROM patient_bmi WHERE patient_id = $1`

	total := model.DB.QueryRowx(countSQL, patientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.Queryx(rowSQL, patientID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	patientHeightLists := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": patientHeightLists, "page_info": pageInfo})
}

//PatientBloodTypeList 患者血型记录
func PatientBloodTypeList(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if patientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
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

	countSQL := `select count(id) as total from patient_blood_type where patient_id=$1`
	rowSQL := `SELECT * FROM patient_blood_type WHERE patient_id = $1`

	total := model.DB.QueryRowx(countSQL, patientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.Queryx(rowSQL, patientID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	patientHeightLists := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": patientHeightLists, "page_info": pageInfo})
}

//PatientRhBloodTypeList 患者RH血型记录
func PatientRhBloodTypeList(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if patientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
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

	countSQL := `select count(id) as total from patient_rh_blood_type where patient_id=$1`
	rowSQL := `SELECT * FROM patient_rh_blood_type WHERE patient_id = $1`

	total := model.DB.QueryRowx(countSQL, patientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.Queryx(rowSQL, patientID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	patientHeightLists := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": patientHeightLists, "page_info": pageInfo})
}

//PatientTemperatureList 患者体温记录
func PatientTemperatureList(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if patientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
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

	countSQL := `select count(id) as total from patient_temperature where patient_id=$1`
	rowSQL := `SELECT * FROM patient_temperature WHERE patient_id = $1`

	total := model.DB.QueryRowx(countSQL, patientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.Queryx(rowSQL, patientID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	patientHeightLists := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": patientHeightLists, "page_info": pageInfo})
}

//PatientBreatheList 患者呼吸记录
func PatientBreatheList(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if patientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
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

	countSQL := `select count(id) as total from patient_breathe where patient_id=$1`
	rowSQL := `SELECT * FROM patient_breathe WHERE patient_id = $1`

	total := model.DB.QueryRowx(countSQL, patientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.Queryx(rowSQL, patientID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	patientHeightLists := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": patientHeightLists, "page_info": pageInfo})
}

//PatientPulseList 患者脉搏记录
func PatientPulseList(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if patientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
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

	countSQL := `select count(id) as total from patient_pulse where patient_id=$1`
	rowSQL := `SELECT * FROM patient_pulse WHERE patient_id = $1`

	total := model.DB.QueryRowx(countSQL, patientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.Queryx(rowSQL, patientID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	patientHeightLists := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": patientHeightLists, "page_info": pageInfo})
}

//PatientBloodPressureList 患者血压记录
func PatientBloodPressureList(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if patientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
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

	countSQL := `select count(id) as total from patient_blood_pressure where patient_id=$1`
	rowSQL := `SELECT * FROM patient_blood_pressure WHERE patient_id = $1`

	total := model.DB.QueryRowx(countSQL, patientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.Queryx(rowSQL, patientID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	patientHeightLists := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": patientHeightLists, "page_info": pageInfo})
}

//PatientVisionList 患者视力记录
func PatientVisionList(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if patientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
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

	countSQL := `select count(id) as total from patient_vision where patient_id=$1`
	rowSQL := `SELECT * FROM patient_vision WHERE patient_id = $1`

	total := model.DB.QueryRowx(countSQL, patientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.Queryx(rowSQL, patientID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	patientHeightLists := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": patientHeightLists, "page_info": pageInfo})
}

//PatientOxygenSaturationList 患者氧饱和度记录
func PatientOxygenSaturationList(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if patientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
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

	countSQL := `select count(id) as total from patient_oxygen_saturation where patient_id=$1`
	rowSQL := `SELECT * FROM patient_oxygen_saturation WHERE patient_id = $1`

	total := model.DB.QueryRowx(countSQL, patientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.Queryx(rowSQL, patientID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	patientHeightLists := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": patientHeightLists, "page_info": pageInfo})
}
