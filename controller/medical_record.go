package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strings"

	"github.com/kataras/iris"
)

// MedicalRecordCreate 创建主病历
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

	if clinicTriagePatientID == "" || chiefComplaint == "" || operationID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from medical_record where clinic_triage_patient_id=$1", clinicTriagePatientID)
	clinicTriagePatient := FormatSQLRowToMap(row)
	_, ok := clinicTriagePatient["id"]
	if !ok {
		sql := `INSERT INTO  medical_record ( clinic_triage_patient_id, morbidity_date, chief_complaint, history_of_present_illness, history_of_past_illness, family_medical_history, allergic_history, allergic_reaction, immunizations, body_examination, diagnosis, cure_suggestion, remark, operation_id, files, is_default ) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16) RETURNING id`

		var id int
		err := model.DB.QueryRow(sql, clinicTriagePatientID, morbidityDate, chiefComplaint, historyOfPresentIllness, historyOfPastIllness, familyMedicalHistory, allergicHistory, allergicReaction, immunizations, bodyExamination, diagnosis, cureSuggestion, remark, operationID, files, true).Scan(&id)
		if err != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	} else {
		sql := `UPDATE medical_record SET morbidity_date=$1, chief_complaint=$2, history_of_present_illness=$3, history_of_past_illness=$4, family_medical_history=$5, allergic_history=$6, allergic_reaction=$7, immunizations=$8, body_examination=$9, diagnosis=$10, cure_suggestion=$11, remark=$12, operation_id=$13, files=$14, updated_time=LOCALTIMESTAMP WHERE clinic_triage_patient_id=$15`

		_, err := model.DB.Exec(sql, morbidityDate, chiefComplaint, historyOfPresentIllness, historyOfPastIllness, familyMedicalHistory, allergicHistory, allergicReaction, immunizations, bodyExamination, diagnosis, cureSuggestion, remark, operationID, files, clinicTriagePatientID)
		if err != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}

	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// MedicalRecordRenew 续写病历
func MedicalRecordRenew(ctx iris.Context) {

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

	if clinicTriagePatientID == "" || chiefComplaint == "" || operationID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from medical_record where clinic_triage_patient_id=$1 and is_default=true", clinicTriagePatientID)
	medicalRecord := FormatSQLRowToMap(row)
	_, ok := medicalRecord["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "没有主病历,不能续写"})
		return
	}
	sql := `INSERT INTO  medical_record ( clinic_triage_patient_id, morbidity_date, chief_complaint, history_of_present_illness, history_of_past_illness, family_medical_history, allergic_history, allergic_reaction, immunizations, body_examination, diagnosis, cure_suggestion, remark, operation_id, files ) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)`

	_, err := model.DB.Exec(sql, clinicTriagePatientID, morbidityDate, chiefComplaint, historyOfPresentIllness, historyOfPastIllness, familyMedicalHistory, allergicHistory, allergicReaction, immunizations, bodyExamination, diagnosis, cureSuggestion, remark, operationID, files)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// MedicalRecordRenewUpdate 续写病历修改
func MedicalRecordRenewUpdate(ctx iris.Context) {

	medicalRecordID := ctx.PostValue("medical_record_id")
	chiefComplaint := ctx.PostValue("chief_complaint")
	files := ctx.PostValue("files")
	operationID := ctx.PostValue("operation_id")

	if medicalRecordID == "" || chiefComplaint == "" || operationID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id,is_default from medical_record where id=$1", medicalRecordID)
	medicalRecord := FormatSQLRowToMap(row)

	_, ok := medicalRecord["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改的病历不存在"})
		return
	}
	isDefault := medicalRecord["is_default"]

	if isDefault != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "主病历不能修改"})
		return
	}

	sql := `UPDATE medical_record SET chief_complaint=$2, operation_id=$3, files=$4, updated_time=LOCALTIMESTAMP WHERE id=$1`

	_, err := model.DB.Exec(sql, medicalRecordID, chiefComplaint, operationID, files)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// MedicalRecordRenewDelete 续写病历删除
func MedicalRecordRenewDelete(ctx iris.Context) {

	medicalRecordID := ctx.PostValue("medical_record_id")

	if medicalRecordID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id,is_default from medical_record where id=$1", medicalRecordID)
	medicalRecord := FormatSQLRowToMap(row)

	_, ok := medicalRecord["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "删除的病历不存在"})
		return
	}

	isDefault := medicalRecord["is_default"]

	if isDefault != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "主病历不能删除"})
		return
	}

	sql := `delete from medical_record WHERE id=$1`

	_, err := model.DB.Exec(sql, medicalRecordID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// MedicalRecordFindByTriageID 通过id查找
func MedicalRecordFindByTriageID(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")

	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	row := model.DB.QueryRowx("select id as clinic_triage_patient_id,*  from medical_record where clinic_triage_patient_id=$1", clinicTriagePatientID)
	medicalRecord := FormatSQLRowToMap(row)
	ctx.JSON(iris.Map{"code": "200", "data": medicalRecord})
	return
}

// MedicalRecordModelCreate 创建病历模板
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
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	sql := `INSERT INTO  medical_record_model ( model_name, is_common, chief_complaint, history_of_present_illness, history_of_past_illness, family_medical_history, allergic_history, allergic_reaction, immunizations, body_examination, diagnosis, cure_suggestion, remark, operation_id ) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) RETURNING id`

	var id int
	err := model.DB.QueryRow(sql, modelName, isCommon, chiefComplaint, historyOfPresentIllness, historyOfPastIllness, familyMedicalHistory, allergicHistory, allergicReaction, immunizations, bodyExamination, diagnosis, cureSuggestion, remark, operationID).Scan(&id)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": id})
}

// MedicalRecordModelUpdate 修改病历模板
func MedicalRecordModelUpdate(ctx iris.Context) {
	medicalRecordModelID := ctx.PostValue("medical_record_model_id")
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

	if modelName == "" || medicalRecordModelID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	mrow := model.DB.QueryRowx("select id from medical_record_model where id=$1 limit 1", medicalRecordModelID)
	if mrow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	models := FormatSQLRowToMap(mrow)
	_, mok := models["id"]
	if !mok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改的模板不存在"})
		return
	}

	row := model.DB.QueryRowx("select id from medical_record_model where model_name=$1 and id!=$2 limit 1", modelName, medicalRecordModelID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	examinationModel := FormatSQLRowToMap(row)
	_, ok := examinationModel["id"]
	if ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "模板名称已存在"})
		return
	}

	prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", operationID)
	if prow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改模板失败,操作员错误"})
		return
	}
	personnel := FormatSQLRowToMap(prow)
	_, pok := personnel["id"]
	if !pok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "操作员错误"})
		return
	}

	medicalRecordModelMap := map[string]interface{}{
		"id":                         medicalRecordModelID,
		"model_name":                 modelName,
		"is_common":                  isCommon,
		"chief_complaint":            chiefComplaint,
		"history_of_present_illness": historyOfPresentIllness,
		"history_of_past_illness":    historyOfPastIllness,
		"family_medical_history":     familyMedicalHistory,
		"allergic_history":           allergicHistory,
		"allergic_reaction":          allergicReaction,
		"immunizations":              immunizations,
		"body_examination":           bodyExamination,
		"diagnosis":                  diagnosis,
		"cure_suggestion":            cureSuggestion,
		"remark":                     remark,
		"operation_id":               ToNullInt64(operationID),
	}

	var s []string
	s = append(s, "id=:id", "model_name=:model_name", "chief_complaint=:chief_complaint",
		"history_of_present_illness=:history_of_present_illness", "history_of_past_illness=:history_of_past_illness",
		"family_medical_history=:family_medical_history", "allergic_history=:allergic_history", "allergic_reaction=:allergic_reaction",
		"immunizations=:immunizations", "body_examination=:body_examination", "diagnosis=:diagnosis", "cure_suggestion=:cure_suggestion",
		"remark=:remark", "operation_id=:operation_id")

	if isCommon != "" {
		s = append(s, "is_common=:is_common")
	}
	joinSQL := strings.Join(s, ",")
	medicalRecordUpdateSQL := `update medical_record_model set ` + joinSQL + ` where id=:id`

	_, err2 := model.DB.NamedExec(medicalRecordUpdateSQL, medicalRecordModelMap)
	if err2 != nil {
		fmt.Println(" err2====", err2)
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// MedicalRecordListByPID 获取病历列表
func MedicalRecordListByPID(ctx iris.Context) {
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

	countSQL := `select count(mr.id) as total from medical_record mr 
	left join clinic_triage_patient ctp on ctp.id = mr.clinic_triage_patient_id 
	left join clinic_patient cp on ctp.clinic_patient_id = cp.id 
	where cp.patient_id = $1`
	selectSQL := `	select mr.*,
	c.name as clinic_name,
	cp.id as clinic_patient_id,
	ctp.created_time as registion_time,
	ctp.visit_type,
	p.name as doctor_name, 
	d.name as department_name
	from medical_record mr 
	left join clinic_triage_patient ctp on ctp.id = mr.clinic_triage_patient_id 
	left join personnel p on p.id = ctp.doctor_id 
	left join department d on d.id = ctp.department_id 
	left join clinic_patient cp on ctp.clinic_patient_id = cp.id 
	left join clinic c on cp.clinic_id = c.id 
	where cp.patient_id = $1 ORDER BY mr.created_time DESC offset $2 limit $3`

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
	operationID := ctx.PostValue("operation_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}

	countSQL := `select count(id) as total from medical_record_model where model_name ~$1 and deleted_time is null`
	selectSQL := `select mrm.*, p.name as operation_name from medical_record_model mrm
	left join personnel p on mrm.operation_id = p.id
	where mrm.model_name ~$1 and mrm.deleted_time is null`

	if isCommon != "" {
		countSQL += ` and is_common =` + isCommon
		selectSQL += ` and mrm.is_common=` + isCommon
	}

	if operationID != "" {
		countSQL += ` and operation_id =` + operationID
		selectSQL += ` and mrm.operation_id=` + operationID
	}
	fmt.Println("countSQL===", countSQL)
	fmt.Println("selectSQL===", selectSQL)
	total := model.DB.QueryRowx(countSQL, keyword)

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err1 := model.DB.Queryx(selectSQL+" ORDER BY created_time DESC offset $2 limit $3", keyword, offset, limit)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})

}

// MedicalRecordModelListByOperation 医生查询模板
func MedicalRecordModelListByOperation(ctx iris.Context) {
	keyword := ctx.PostValue("keyword")
	isCommon := ctx.PostValue("is_common")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	operationID := ctx.PostValue("operation_id")

	if operationID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}

	countSQL := `select count(id) as total from medical_record_model where model_name ~$1 and deleted_time is null AND (operation_id = $2 or is_common = true)`
	if isCommon != "" {
		countSQL = countSQL + ` and is_common =` + isCommon
	}
	selectSQL := `select * from medical_record_model where model_name ~$1 and deleted_time is null AND (operation_id = $2 or is_common = true)`
	if isCommon != "" {
		selectSQL = selectSQL + ` and is_common =` + isCommon
	}
	selectSQL = selectSQL + " ORDER BY created_time DESC offset $3 limit $4"

	total := model.DB.QueryRowx(countSQL, keyword, operationID)

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err1 := model.DB.Queryx(selectSQL, keyword, operationID, offset, limit)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})
}

// MedicalRecordModelDelete 删除病历模板
func MedicalRecordModelDelete(ctx iris.Context) {
	medicalRecordModelID := ctx.PostValue("medical_record_model_id")

	if medicalRecordModelID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	mrow := model.DB.QueryRowx("select id from medical_record_model where id=$1 limit 1", medicalRecordModelID)
	if mrow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	models := FormatSQLRowToMap(mrow)
	_, mok := models["id"]
	if !mok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改的模板不存在"})
		return
	}

	medicalRecordUpdateSQL := `update medical_record_model set deleted_time=LOCALTIMESTAMP where id=$1`

	_, err2 := model.DB.Exec(medicalRecordUpdateSQL, medicalRecordModelID)
	if err2 != nil {
		fmt.Println(" err2====", err2)
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}
