package controller

import (
	"clinicSystemGo/model"

	"github.com/kataras/iris"
)

// MedicalRecordCreate 创建病历
func MedicalRecordCreate(ctx iris.Context) {

	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	chiefComplaint := ctx.PostValue("chief_complaint")

	morbidityDate := ctx.PostValue("morbidity_date")
	historyOfPresentIllness := ctx.PostValue("history_of_present_illness")
	historyOfPastIllness := ctx.PostValue("history_of_past_illness")
	familyMedicalHistory := ctx.PostValue("family_medical_history")
	allergicHistory := ctx.PostValue("allergic_history")
	allergicReaction := ctx.PostValue("allergic_reaction")
	immunizations := ctx.PostValue("immunizations")
	bodyExamination := ctx.PostValue("body_examination")
	diagnosis := ctx.PostValue("diagnosis")
	cureSuggestion := ctx.PostValue("cure_suggestion")
	remark := ctx.PostValue("remark")
	files := ctx.PostValue("files")
	operationID := ctx.PostValue("operation_id")

	if clinicTriagePatientID == "" || chiefComplaint == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from medical_record where clinic_triage_patient_id=$1", clinicTriagePatientID)
	clinicTriagePatient := FormatSQLRowToMap(row)
	_, ok := clinicTriagePatient["id"]
	if !ok {
		sql := `INSERT INTO  medical_record ( clinic_triage_patient_id, morbidity_date, chief_complaint, history_of_present_illness, history_of_past_illness, family_medical_history, allergic_history, allergic_reaction, immunizations, body_examination, diagnosis, cure_suggestion, remark, operation_id, files ) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15) RETURNING id`

		var id int
		err := model.DB.QueryRow(sql, clinicTriagePatientID, morbidityDate, chiefComplaint, historyOfPresentIllness, historyOfPastIllness, familyMedicalHistory, allergicHistory, allergicReaction, immunizations, bodyExamination, diagnosis, cureSuggestion, remark, operationID, files).Scan(&id)
		if err != nil {
			ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
			return
		}
	} else {
		sql := `UPDATE medical_record SET morbidity_date=$1, chief_complaint=$2, history_of_present_illness=$3, history_of_past_illness=$4, family_medical_history=$5, allergic_history=$6, allergic_reaction=$7, immunizations=$8, body_examination=$9, diagnosis=$10, cure_suggestion=$11, remark=$12, operation_id=$13, files=$14, updated_time=LOCALTIMESTAMP WHERE clinic_triage_patient_id=$15`

		_, err := model.DB.Exec(sql, morbidityDate, chiefComplaint, historyOfPresentIllness, historyOfPastIllness, familyMedicalHistory, allergicHistory, allergicReaction, immunizations, bodyExamination, diagnosis, cureSuggestion, remark, operationID, files, clinicTriagePatientID)
		if err != nil {
			ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
			return
		}

	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// MedicalRecordFindByTriageID 通过id查找
func MedicalRecordFindByTriageID(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")

	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	row := model.DB.QueryRowx("select id as clinic_triage_patient_id,*  from medical_record where clinic_triage_patient_id=$1", clinicTriagePatientID)
	medicalRecord := FormatSQLRowToMap(row)
	ctx.JSON(iris.Map{"code": "200", "data": medicalRecord})
	return
}

// MedicalRecordModelCreate 创建病历
func MedicalRecordModelCreate(ctx iris.Context) {

	modelName := ctx.PostValue("model_name")
	isCommon := ctx.PostValue("is_common")

	chiefComplaint := ctx.PostValue("chief_complaint")
	historyOfPresentIllness := ctx.PostValue("history_of_present_illness")
	historyOfPastIllness := ctx.PostValue("history_of_past_illness")
	familyMedicalHistory := ctx.PostValue("family_medical_history")
	allergicHistory := ctx.PostValue("allergic_history")
	allergicReaction := ctx.PostValue("allergic_reaction")
	immunizations := ctx.PostValue("immunizations")
	bodyExamination := ctx.PostValue("body_examination")
	diagnosis := ctx.PostValue("diagnosis")
	cureSuggestion := ctx.PostValue("cure_suggestion")
	remark := ctx.PostValue("remark")
	operationID := ctx.PostValue("operation_id")

	if modelName == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	sql := `INSERT INTO  medical_record_model ( model_name, is_common, chief_complaint, history_of_present_illness, history_of_past_illness, family_medical_history, allergic_history, allergic_reaction, immunizations, body_examination, diagnosis, cure_suggestion, remark, operation_id ) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) RETURNING id`

	var id int
	err := model.DB.QueryRow(sql, modelName, isCommon, chiefComplaint, historyOfPresentIllness, historyOfPastIllness, familyMedicalHistory, allergicHistory, allergicReaction, immunizations, bodyExamination, diagnosis, cureSuggestion, remark, operationID).Scan(&id)
	if err != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": id})
}

// MedicalRecordListByPID 获取病历列表
func MedicalRecordListByPID(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if patientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}

	countSQL := `select count(id) as total from medocal_record where patient_id = $1`
	selectSQL := `select * from medocal_record where patient_id = $1 ORDER BY created_time DESC offset $2 limit $3`

	total := model.DB.QueryRowx(countSQL, patientID)

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err1 := model.DB.Queryx(selectSQL, patientID, offset, limit)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})

}

// MedicalRecordModelList 查询模板
func MedicalRecordModelList(ctx iris.Context) {
	keyword := ctx.PostValue("keyword")
	isCommon := ctx.PostValue("is_common")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}

	countSQL := `select count(id) as total from medocal_record_model where model_name ~$1`
	if isCommon != "" {
		countSQL = countSQL + ` and is_common =` + isCommon
	}
	selectSQL := `select model_name,is_common,created_time from medocal_record_model where model_name ~$1 ORDER BY created_time DESC offset $2 limit $3`
	if isCommon != "" {
		selectSQL = `select model_name,is_common,created_time from medocal_record_model where model_name ~$1 AND is_common=` + isCommon + ` ORDER BY created_time DESC offset $2 limit $3`
	}

	total := model.DB.QueryRowx(countSQL, keyword)

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err1 := model.DB.Queryx(selectSQL, keyword, offset, limit)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})

}
