package controller

import (
	"clinicSystemGo/model"

	"github.com/kataras/iris"
)

// MedicalRecordCreate 创建病历
func MedicalRecordCreate(ctx iris.Context) {

	patientID := ctx.PostValue("patient_id")
	chiefComplaint := ctx.PostValue("chief_complaint")

	registrationID := ctx.PostValue("registration_id")
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

	if patientID == "" || chiefComplaint == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	sql := `INSERT INTO  medical_record ( patient_id, registration_id, morbidity_date, chief_complaint, history_of_present_illness, history_of_past_illness, family_medical_history, allergic_history, allergic_reaction, immunizations, body_examination, diagnosis, cure_suggestion, remark, operation_id, files ) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16) RETURNING id`

	var id int
	err := model.DB.QueryRow(sql, patientID, registrationID, morbidityDate, chiefComplaint, historyOfPresentIllness, historyOfPastIllness, familyMedicalHistory, allergicHistory, allergicReaction, immunizations, bodyExamination, diagnosis, cureSuggestion, remark, operationID, files).Scan(&id)
	if err != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": id})
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
