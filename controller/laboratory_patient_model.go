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

//LaboratoryModel 检验模板
type LaboratoryModel struct {
	ModelName                string                `json:"model_name"`
	LaboratoryPatientModelID int                   `json:"laboratory_patient_model_id"`
	OperationName            string                `json:"operation_name"`
	IsCommon                 bool                  `json:"is_common"`
	CreatedTime              time.Time             `json:"created_time"`
	Items                    []LaboratoryModelItem `json:"items"`
}

//LaboratoryModelItem 检验模板item
type LaboratoryModelItem struct {
	LaboratoryName     string      `json:"laboratory_name"`
	Times              int         `json:"times"`
	ClinicLaboratoryID int         `json:"clinic_laboratory_id"`
	Illustration       interface{} `json:"illustration"`
}

// LaboratoryPatientModelCreate 创建检验医嘱模板
func LaboratoryPatientModelCreate(ctx iris.Context) {
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
	row := model.DB.QueryRowx("select id from laboratory_patient_model where model_name=$1 limit 1", modelName)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
		return
	}
	laboratoryModel := FormatSQLRowToMap(row)
	_, ok := laboratoryModel["id"]
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
		"laboratory_patient_model_id",
		"clinic_laboratory_id",
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
	var laboratoryModelID int
	err := tx.QueryRow("insert into laboratory_patient_model (model_name,is_common,operation_id) values ($1,$2,$3) RETURNING id", modelName, isCommon, personnelID).Scan(&laboratoryModelID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	tSetStr := strings.Join(itemSets, ",")
	inserttSQL := "insert into laboratory_patient_model_item (" + tSetStr + ") values ($1,$2,$3,$4)"
	clinicLaboratorySQL := `select id from clinic_laboratory where id=$1`

	for _, v := range results {
		clinicLaboratoryID := v["clinic_laboratory_id"]
		times := v["times"]
		illustration := v["illustration"]

		trow := model.DB.QueryRowx(clinicLaboratorySQL, clinicLaboratoryID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "保存模板错误"})
			return
		}
		clinicLaboratory := FormatSQLRowToMap(trow)
		_, ok := clinicLaboratory["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "选择的检验医嘱错误"})
			return
		}

		_, errt := tx.Exec(inserttSQL,
			laboratoryModelID,
			ToNullInt64(clinicLaboratoryID),
			ToNullInt64(times),
			ToNullString(illustration),
		)
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

// LaboratoryPatientModelList 查询检验医嘱模板
func LaboratoryPatientModelList(ctx iris.Context) {
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

	countSQL := `select count(id) as total from laboratory_patient_model where id>0`
	selectSQL := `select lpm.id as laboratory_patient_model_id,cl.name as laboratory_name,p.name as operation_name,
	lpm.is_common,lpm.created_time,lpmi.clinic_laboratory_id, lpmi.illustration, lpm.model_name,lpmi.times from laboratory_patient_model lpm
	left join laboratory_patient_model_item lpmi on lpmi.laboratory_patient_model_id = lpm.id
	left join clinic_laboratory cl on lpmi.clinic_laboratory_id = cl.id
	left join personnel p on lpm.operation_id = p.id
	where lpm.id>0`

	if keyword != "" {
		countSQL += ` and model_name ~:keyword`
		selectSQL += ` and lpm.model_name ~:keyword`
	}

	if isCommon != "" {
		countSQL += ` and is_common = :is_common`
		selectSQL += ` and lpm.is_common= :is_common`
	}

	if operationID != "" {
		countSQL += ` and operation_id = :operation_id`
		selectSQL += ` and lpm.operation_id=:operation_id`
	}

	var queryOption = map[string]interface{}{
		"keyword":      ToNullInt64(keyword),
		"is_common":    ToNullString(isCommon),
		"operation_id": ToNullBool(operationID),
		"offset":       ToNullInt64(offset),
		"limit":        ToNullInt64(limit),
	}
	fmt.Println("countSQL===", countSQL)
	fmt.Println("selectSQL===", selectSQL)
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
	var models []LaboratoryModel
	for _, v := range result {
		modelName := v["model_name"]
		laboratoryPatientModelID := v["laboratory_patient_model_id"]
		laboratoryName := v["laboratory_name"]
		operationName := v["operation_name"]
		isCommon := v["is_common"]
		createdTime := v["created_time"]
		times := v["times"]
		clinicLaboratoryID := v["clinic_laboratory_id"]
		illustration := v["illustration"]
		has := false
		item := LaboratoryModelItem{
			LaboratoryName:     laboratoryName.(string),
			Times:              int(times.(int64)),
			ClinicLaboratoryID: int(clinicLaboratoryID.(int64)),
			Illustration:       illustration,
		}
		for k, pModel := range models {
			plaboratoryPatientModelID := pModel.LaboratoryPatientModelID
			items := pModel.Items
			if int(laboratoryPatientModelID.(int64)) == plaboratoryPatientModelID {
				models[k].Items = append(items, item)
				has = true
			}
		}
		if !has {
			var items []LaboratoryModelItem
			items = append(items, item)

			pmodel := LaboratoryModel{
				ModelName:                modelName.(string),
				LaboratoryPatientModelID: int(laboratoryPatientModelID.(int64)),
				OperationName:            operationName.(string),
				IsCommon:                 isCommon.(bool),
				CreatedTime:              createdTime.(time.Time),
				Items:                    items,
			}
			models = append(models, pmodel)
		}
	}

	ctx.JSON(iris.Map{"code": "200", "data": models, "page_info": pageInfo})
}

// LaboratoryPersonalPatientModelList 查询个人和通用检验医嘱模板
func LaboratoryPersonalPatientModelList(ctx iris.Context) {
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

	countSQL := `select count(id) as total from laboratory_patient_model where model_name ~$1 and (operation_id=$2 or is_common=true)`
	selectSQL := `select lpm.id as laboratory_patient_model_id,cl.name as laboratory_name,p.name as operation_name,
	lpm.is_common,lpm.created_time,lpmi.clinic_laboratory_id, lpmi.illustration, lpm.model_name,lpmi.times from laboratory_patient_model lpm
	left join laboratory_patient_model_item lpmi on lpmi.laboratory_patient_model_id = lpm.id
	left join clinic_laboratory cl on lpmi.clinic_laboratory_id = cl.id
	left join personnel p on lpm.operation_id = p.id
	where lpm.model_name ~$1 and (lpm.operation_id=$2 or lpm.is_common=true)`

	if isCommon != "" {
		countSQL += ` and is_common =` + isCommon
		selectSQL += ` and lpm.is_common=` + isCommon
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
	var models []LaboratoryModel
	for _, v := range result {
		modelName := v["model_name"]
		laboratoryPatientModelID := v["laboratory_patient_model_id"]
		laboratoryName := v["laboratory_name"]
		operationName := v["operation_name"]
		isCommon := v["is_common"]
		createdTime := v["created_time"]
		times := v["times"]
		clinicLaboratoryID := v["clinic_laboratory_id"]
		illustration := v["illustration"]
		has := false
		item := LaboratoryModelItem{
			LaboratoryName:     laboratoryName.(string),
			Times:              int(times.(int64)),
			ClinicLaboratoryID: int(clinicLaboratoryID.(int64)),
			Illustration:       illustration,
		}
		for k, pModel := range models {
			plaboratoryPatientModelID := pModel.LaboratoryPatientModelID
			items := pModel.Items
			if int(laboratoryPatientModelID.(int64)) == plaboratoryPatientModelID {
				models[k].Items = append(items, item)
				has = true
			}
		}
		if !has {
			var items []LaboratoryModelItem
			items = append(items, item)

			pmodel := LaboratoryModel{
				ModelName:                modelName.(string),
				LaboratoryPatientModelID: int(laboratoryPatientModelID.(int64)),
				OperationName:            operationName.(string),
				IsCommon:                 isCommon.(bool),
				CreatedTime:              createdTime.(time.Time),
				Items:                    items,
			}
			models = append(models, pmodel)
		}
	}

	ctx.JSON(iris.Map{"code": "200", "data": models, "page_info": pageInfo})
}

// LaboratoryPatientModelDetail 查询检验医嘱模板详情
func LaboratoryPatientModelDetail(ctx iris.Context) {
	laboratoryModelID := ctx.PostValue("laboratory_patient_model_id")

	selectmSQL := `select id as laboratory_patient_model_id,model_name,is_common,status from laboratory_patient_model where id=$1`
	mrows := model.DB.QueryRowx(selectmSQL, laboratoryModelID)
	if mrows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询失败"})
		return
	}
	laboratoryModel := FormatSQLRowToMap(mrows)

	selectiSQL := `select lpmi.clinic_laboratory_id,cl.name,lpmi.times,lpmi.illustration 
		from laboratory_patient_model_item lpmi
		left join laboratory_patient_model lpm on lpmi.laboratory_patient_model_id = lpm.id
		left join clinic_laboratory cl on lpmi.clinic_laboratory_id = cl.id
		where lpmi.laboratory_patient_model_id=$1`

	rows, err := model.DB.Queryx(selectiSQL, laboratoryModelID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	laboratoryModel["items"] = result
	ctx.JSON(iris.Map{"code": "200", "data": laboratoryModel})
}

// LaboratoryPatientModelUpdate 修改检验医嘱模板
func LaboratoryPatientModelUpdate(ctx iris.Context) {
	laboratoryModelID := ctx.PostValue("laboratory_patient_model_id")
	modelName := ctx.PostValue("model_name")
	isCommon := ctx.PostValue("is_common")
	items := ctx.PostValue("items")
	personnelID := ctx.PostValue("operation_id")

	if laboratoryModelID == "" || modelName == "" || items == "" {
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
	mrow := model.DB.QueryRowx("select id from laboratory_patient_model where id=$1 limit 1", laboratoryModelID)
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

	row := model.DB.QueryRowx("select id from laboratory_patient_model where model_name=$1 and id!=$2 limit 1", modelName, laboratoryModelID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	laboratoryModel := FormatSQLRowToMap(row)
	_, ok := laboratoryModel["id"]
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
		"laboratory_patient_model_id",
		"clinic_laboratory_id",
		"times",
		"illustration",
	}
	tSetStr := strings.Join(itemSets, ",")

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errb})
		return
	}
	updateSQL := `update laboratory_patient_model set model_name=$1,is_common=$2,
	operation_id=$3,updated_time=LOCALTIMESTAMP where id=$4`
	_, err := tx.Exec(updateSQL, modelName, isCommon, personnelID, laboratoryModelID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	deleteSQL := "delete from laboratory_patient_model_item where laboratory_patient_model_id=$1"
	fmt.Println("deleteSQL===", deleteSQL)
	_, errd := tx.Exec(deleteSQL, laboratoryModelID)
	if errd != nil {
		fmt.Println("errd ===", errd)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errd.Error()})
		return
	}
	inserttSQL := "insert into laboratory_patient_model_item (" + tSetStr + ") values ($1,$2,$3,$4)"

	for _, v := range results {
		clinicLaboratoryID := v["clinic_laboratory_id"]
		times := v["times"]
		illustration := v["illustration"]

		clinicLaboratorySQL := `select id from clinic_laboratory where id=$1`
		trow := model.DB.QueryRowx(clinicLaboratorySQL, clinicLaboratoryID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "保存模板错误"})
			return
		}
		clinicLaboratory := FormatSQLRowToMap(trow)
		_, ok := clinicLaboratory["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "选择的检验医嘱错误"})
			return
		}

		_, errt := tx.Exec(inserttSQL,
			ToNullInt64(laboratoryModelID),
			ToNullInt64(clinicLaboratoryID),
			ToNullInt64(times),
			ToNullString(illustration),
		)
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
