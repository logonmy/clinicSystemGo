package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris"
)

//ExaminationModel 检查模板
type ExaminationModel struct {
	ModelName                 string                 `json:"model_name"`
	ExaminationPatientModelID int                    `json:"examination_patient_model_id"`
	OperationName             string                 `json:"operation_name"`
	IsCommon                  bool                   `json:"is_common"`
	CreatedTime               time.Time              `json:"created_time"`
	Items                     []ExaminationModelItem `json:"items"`
}

//ExaminationModelItem 检查模板item
type ExaminationModelItem struct {
	ExaminationName     string      `json:"examination_name"`
	Times               int         `json:"times"`
	ClinicExaminationID int         `json:"clinic_examination_id"`
	Illustration        interface{} `json:"illustration"`
}

// ExaminationPatientModelCreate 创建检查医嘱模板
func ExaminationPatientModelCreate(ctx iris.Context) {
	modelName := ctx.PostValue("model_name")
	isCommon := ctx.PostValue("is_common")
	items := ctx.PostValue("items")
	personnelID := ctx.PostValue("operation_id")

	if modelName == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	var results []map[string]string
	errj := json.Unmarshal([]byte(items), &results)
	fmt.Println("results===", results)
	if errj != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": errj.Error()})
		return
	}
	row := model.DB.QueryRowx("select id from examination_patient_model where model_name=$1 limit 1", modelName)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
		return
	}
	examinationModel := FormatSQLRowToMap(row)
	_, ok := examinationModel["id"]
	if ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "模板名称已存在"})
		return
	}

	prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", personnelID)
	if prow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "保存模板失败,操作员错误"})
		return
	}
	personnel := FormatSQLRowToMap(prow)
	_, pok := personnel["id"]
	if !pok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "操作员错误"})
		return
	}

	itemSets := []string{
		"examination_patient_model_id",
		"clinic_examination_id",
		"organ",
		"times",
		"illustration",
	}

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errb})
		return
	}
	var examinationModelID int
	err := tx.QueryRow("insert into examination_patient_model (model_name,is_common,operation_id) values ($1,$2,$3) RETURNING id", modelName, isCommon, personnelID).Scan(&examinationModelID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	clinicExaminationSQL := `select id from clinic_examination where id=$1`
	tSetStr := strings.Join(itemSets, ",")
	inserttSQL := "insert into examination_patient_model_item (" + tSetStr + ") values ($1,$2,$3,$4)"

	for _, v := range results {
		clinicExaminationID := v["clinic_examination_id"]
		times := v["times"]
		organ := v["organ"]
		illustration := v["illustration"]

		trow := model.DB.QueryRowx(clinicExaminationSQL, clinicExaminationID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "保存模板错误"})
			return
		}
		clinicExamination := FormatSQLRowToMap(trow)
		_, ok := clinicExamination["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "选择的检查医嘱错误"})
			return
		}

		_, errt := tx.Exec(inserttSQL, examinationModelID, ToNullInt64(clinicExaminationID), ToNullInt64(times), ToNullString(illustration), ToNullString(organ))
		if errt != nil {
			fmt.Println("errt ===", errt)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": errt.Error()})
			return
		}
	}

	errc := tx.Commit()
	if errc != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errc.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// ExaminationPatientModelList 查询检查医嘱模板
func ExaminationPatientModelList(ctx iris.Context) {
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

	countSQL := `select count(id) as total from examination_patient_model where id>0`
	selectSQL := `select epm.id as examination_patient_model_id,ce.name as examination_name,p.name as operation_name,
	epm.is_common,epm.created_time,epmi.clinic_examination_id,epmi.illustration,epm.model_name,epmi.times from examination_patient_model epm
	left join examination_patient_model_item epmi on epmi.examination_patient_model_id = epm.id
	left join clinic_examination ce on epmi.clinic_examination_id = ce.id
	left join personnel p on epm.operation_id = p.id
	where epm.id >0`

	if keyword != "" {
		countSQL += ` and model_name ~:keyword`
		selectSQL += ` and epm.model_name ~:keyword`
	}
	if isCommon != "" {
		countSQL += ` and is_common =:is_common`
		selectSQL += ` and epm.is_common=:is_common`
	}
	if operationID != "" {
		countSQL += ` and operation_id =:operation_id`
		selectSQL += ` and epm.operation_id=:operation_id`
	}

	var queryOption = map[string]interface{}{
		"operation_id": ToNullInt64(operationID),
		"keyword":      keyword,
		"is_common":    ToNullBool(isCommon),
		"offset":       ToNullInt64(offset),
		"limit":        ToNullInt64(limit),
	}

	fmt.Println("countSQL===", countSQL)
	fmt.Println("selectSQL===", selectSQL, queryOption)
	total, err := model.DB.NamedQuery(countSQL, queryOption)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err1 := model.DB.NamedQuery(selectSQL+" ORDER BY created_time DESC offset :offset limit :limit", queryOption)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	fmt.Println("result ====", result)
	var models []ExaminationModel
	for _, v := range result {
		modelName := v["model_name"]
		examinationPatientModelID := v["examination_patient_model_id"]
		examinationName := v["examination_name"]
		operationName := v["operation_name"]
		isCommon := v["is_common"]
		createdTime := v["created_time"]
		times := v["times"]
		clinicExaminationID := v["clinic_examination_id"]
		illustration := v["illustration"]
		has := false
		for k, pModel := range models {
			pexaminationPatientModelID := pModel.ExaminationPatientModelID
			items := pModel.Items
			if int(examinationPatientModelID.(int64)) == pexaminationPatientModelID {
				item := ExaminationModelItem{
					ExaminationName:     examinationName.(string),
					Times:               int(times.(int64)),
					ClinicExaminationID: int(clinicExaminationID.(int64)),
					Illustration:        illustration,
				}
				models[k].Items = append(items, item)
				has = true
			}
		}
		if !has {
			var items []ExaminationModelItem
			item := ExaminationModelItem{
				ExaminationName:     examinationName.(string),
				Times:               int(times.(int64)),
				ClinicExaminationID: int(clinicExaminationID.(int64)),
				Illustration:        illustration,
			}
			items = append(items, item)

			pmodel := ExaminationModel{
				ModelName:                 modelName.(string),
				ExaminationPatientModelID: int(examinationPatientModelID.(int64)),
				OperationName:             operationName.(string),
				IsCommon:                  isCommon.(bool),
				CreatedTime:               createdTime.(time.Time),
				Items:                     items,
			}
			models = append(models, pmodel)
		}
	}

	ctx.JSON(iris.Map{"code": "200", "data": models, "page_info": pageInfo})
}

// ExaminationPersonalPatientModelList 查询个人和通用检查医嘱模板
func ExaminationPersonalPatientModelList(ctx iris.Context) {
	keyword := ctx.PostValue("keyword")
	isCommon := ctx.PostValue("is_common")
	operationID := ctx.PostValue("operation_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

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

	countSQL := `select count(id) as total from examination_patient_model where model_name ~$1 and (operation_id=$2 or is_common=true)`
	selectSQL := `select epm.id as examination_patient_model_id,ce.name as examination_name,p.name as operation_name,
	epm.is_common,epm.created_time,epmi.clinic_examination_id,epmi.illustration,epm.model_name,epmi.times from examination_patient_model epm
	left join examination_patient_model_item epmi on epmi.examination_patient_model_id = epm.id
	left join clinic_examination ce on epmi.clinic_examination_id = ce.id
	left join personnel p on epm.operation_id = p.id
	where epm.model_name ~$1 and (epm.operation_id=$2 or epm.is_common=true)`

	if isCommon != "" {
		countSQL += ` and is_common =` + isCommon
		selectSQL += ` and epm.is_common=` + isCommon
	}

	fmt.Println("countSQL===", countSQL)
	fmt.Println("selectSQL===", selectSQL)
	total := model.DB.QueryRowx(countSQL, keyword, operationID)

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err1 := model.DB.Queryx(selectSQL+" ORDER BY created_time DESC offset $3 limit $4", keyword, operationID, offset, limit)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	var models []ExaminationModel
	for _, v := range result {
		modelName := v["model_name"]
		examinationPatientModelID := v["examination_patient_model_id"]
		examinationName := v["examination_name"]
		operationName := v["operation_name"]
		isCommon := v["is_common"]
		createdTime := v["created_time"]
		times := v["times"]
		clinicExaminationID := v["clinic_examination_id"]
		illustration := v["illustration"]
		has := false
		item := ExaminationModelItem{
			ExaminationName:     examinationName.(string),
			Times:               int(times.(int64)),
			ClinicExaminationID: int(clinicExaminationID.(int64)),
			Illustration:        illustration,
		}
		for k, pModel := range models {
			pexaminationPatientModelID := pModel.ExaminationPatientModelID
			items := pModel.Items
			if int(examinationPatientModelID.(int64)) == pexaminationPatientModelID {
				models[k].Items = append(items, item)
				has = true
			}
		}
		if !has {
			var items []ExaminationModelItem
			items = append(items, item)

			pmodel := ExaminationModel{
				ModelName:                 modelName.(string),
				ExaminationPatientModelID: int(examinationPatientModelID.(int64)),
				OperationName:             operationName.(string),
				IsCommon:                  isCommon.(bool),
				CreatedTime:               createdTime.(time.Time),
				Items:                     items,
			}
			models = append(models, pmodel)
		}
	}

	ctx.JSON(iris.Map{"code": "200", "data": models, "page_info": pageInfo})
}

// ExaminationPatientModelDetail 查询检查医嘱模板详情
func ExaminationPatientModelDetail(ctx iris.Context) {
	examinationModelID := ctx.PostValue("examination_patient_model_id")

	selectmSQL := `select id as examination_patient_model_id,model_name,is_common,status from examination_patient_model where id=$1`
	mrows := model.DB.QueryRowx(selectmSQL, examinationModelID)
	if mrows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询失败"})
		return
	}
	examinationModel := FormatSQLRowToMap(mrows)

	selectiSQL := `select epmi.clinic_examination_id,ce.name,epmi.times,epmi.illustration 
		from examination_patient_model_item epmi
		left join examination_patient_model epm on epmi.examination_patient_model_id = epm.id
		left join clinic_examination ce on epmi.clinic_examination_id = ce.id
		where epmi.examination_patient_model_id=$1`

	rows, err := model.DB.Queryx(selectiSQL, examinationModelID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	examinationModel["items"] = result
	ctx.JSON(iris.Map{"code": "200", "data": examinationModel})
}

// ExaminationPatientModelUpdate 修改检查医嘱模板
func ExaminationPatientModelUpdate(ctx iris.Context) {
	examinationModelID := ctx.PostValue("examination_patient_model_id")
	modelName := ctx.PostValue("model_name")
	isCommon := ctx.PostValue("is_common")
	items := ctx.PostValue("items")
	personnelID := ctx.PostValue("operation_id")

	if examinationModelID == "" || modelName == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	var results []map[string]string
	errj := json.Unmarshal([]byte(items), &results)
	fmt.Println("results===", results)
	if errj != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": errj.Error()})
		return
	}
	mrow := model.DB.QueryRowx("select id from examination_patient_model where id=$1 limit 1", examinationModelID)
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

	row := model.DB.QueryRowx("select id from examination_patient_model where model_name=$1 and id!=$2 limit 1", modelName, examinationModelID)
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

	prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", personnelID)
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

	itemSets := []string{
		"examination_patient_model_id",
		"clinic_examination_id",
		"times",
		"illustration",
	}
	tSetStr := strings.Join(itemSets, ",")

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errb})
		return
	}
	updateSQL := `update examination_patient_model set model_name=$1,is_common=$2,
	operation_id=$3,updated_time=LOCALTIMESTAMP where id=$4`
	_, err := tx.Exec(updateSQL, modelName, isCommon, personnelID, examinationModelID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	deleteSQL := "delete from examination_patient_model_item where examination_patient_model_id=$1"
	_, errd := tx.Exec(deleteSQL, examinationModelID)
	if errd != nil {
		fmt.Println("errd ===", errd)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errd.Error()})
		return
	}
	clinicExaminationSQL := `select id from clinic_examination where id=$1`
	inserttSQL := "insert into examination_patient_model_item (" + tSetStr + ") values ($1,$2,$3,$4)"

	for _, v := range results {
		clinicExaminationID := v["clinic_examination_id"]
		times := v["times"]
		illustration := v["illustration"]

		trow := model.DB.QueryRowx(clinicExaminationSQL, clinicExaminationID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "保存模板错误"})
			return
		}
		clinicExamination := FormatSQLRowToMap(trow)
		_, ok := clinicExamination["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "选择的检查医嘱错误"})
			return
		}

		_, errt := tx.Exec(inserttSQL, ToNullInt64(examinationModelID), ToNullInt64(clinicExaminationID), ToNullInt64(times), ToNullString(illustration))
		if errt != nil {
			fmt.Println("errt ===", errt)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "1", "msg": errt.Error()})
			return
		}
	}
	errc := tx.Commit()
	if errc != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errc.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}
