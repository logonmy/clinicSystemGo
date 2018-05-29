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

//TreatmentModel 治疗模板
type TreatmentModel struct {
	ModelName               string               `json:"model_name"`
	TreatmentPatientModelID int                  `json:"treatment_patient_model_id"`
	OperationName           string               `json:"operation_name"`
	IsCommon                bool                 `json:"is_common"`
	CreatedTime             time.Time            `json:"created_time"`
	Items                   []TreatmentModelItem `json:"items"`
}

//TreatmentModelItem 治疗模板item
type TreatmentModelItem struct {
	TreatmentName     string      `json:"treatment_name"`
	ClinicTreatmentID int         `json:"clinic_treatment_id"`
	Illustration      interface{} `json:"illustration"`
	Times             int         `json:"times"`
	UnitName          string      `json:"unit_name"`
}

// TreatmentPatientModelCreate 创建治疗医嘱模板
func TreatmentPatientModelCreate(ctx iris.Context) {
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
	row := model.DB.QueryRowx("select id from treatment_patient_model where model_name=$1 limit 1", modelName)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
		return
	}
	treatmentModel := FormatSQLRowToMap(row)
	_, ok := treatmentModel["id"]
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
		"treatment_patient_model_id",
		"clinic_treatment_id",
		"times",
		"illustration",
	}
	tSetStr := strings.Join(itemSets, ",")
	inserttSQL := "insert into treatment_patient_model_item (" + tSetStr + ") values ($1,$2,$3,$4)"

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errb})
		return
	}
	var treatmentModelID int
	err := tx.QueryRow("insert into treatment_patient_model (model_name,is_common,operation_id) values ($1,$2,$3) RETURNING id", modelName, isCommon, personnelID).Scan(&treatmentModelID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	clinicTreatmentSQL := `select id from clinic_treatment where id=$1`

	for _, v := range results {
		clinicTreatmentID := v["clinic_treatment_id"]
		times := v["times"]
		illustration := v["illustration"]

		trow := model.DB.QueryRowx(clinicTreatmentSQL, clinicTreatmentID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "保存模板错误"})
			return
		}
		clinicTreatment := FormatSQLRowToMap(trow)
		_, ok := clinicTreatment["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "选择的治疗医嘱错误"})
			return
		}

		_, errt := tx.Exec(inserttSQL, treatmentModelID, ToNullInt64(clinicTreatmentID), ToNullInt64(times), ToNullString(illustration))
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

// TreatmentPatientModelList 查询治疗医嘱模板
func TreatmentPatientModelList(ctx iris.Context) {
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

	countSQL := `select count(id) as total from treatment_patient_model where id >0`
	selectSQL := `select 
	tpm.id as treatment_patient_model_id,
	tpmi.clinic_treatment_id,
	tpmi.illustration,
	ct.name as treatment_name,
	p.name as operation_name,
	tpm.is_common,
	tpm.created_time,
	tpm.model_name,
	tpmi.illustration,
	ct.unit_name,
	tpmi.times from treatment_patient_model tpm
	left join treatment_patient_model_item tpmi on tpmi.treatment_patient_model_id = tpm.id
	left join clinic_treatment ct on tpmi.clinic_treatment_id = ct.id
	left join personnel p on tpm.operation_id = p.id
	where tpm.id >0`

	if keyword != "" {
		countSQL += ` and model_name ~:keyword`
		selectSQL += ` and tpm.model_name ~:keyword`
	}
	if isCommon != "" {
		countSQL += ` and is_common =:is_common`
		selectSQL += ` and tpm.is_common=:is_common`
	}
	if operationID != "" {
		countSQL += ` and operation_id =:operation_id`
		selectSQL += ` and tpm.operation_id=:operation_id`
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
	var models []TreatmentModel
	for _, v := range result {
		modelName := v["model_name"]
		treatmentPatientModelID := v["treatment_patient_model_id"]
		treatmentName := v["treatment_name"]
		clinicTreatmentID := v["clinic_treatment_id"]
		illustration := v["illustration"]
		operationName := v["operation_name"]
		isCommon := v["is_common"]
		createdTime := v["created_time"]
		times := v["times"]
		unitName := v["unit_name"]

		item := TreatmentModelItem{
			TreatmentName:     treatmentName.(string),
			Times:             int(times.(int64)),
			ClinicTreatmentID: int(clinicTreatmentID.(int64)),
			Illustration:      illustration,
			UnitName:          unitName.(string),
		}

		has := false
		for k, pModel := range models {
			ptreatmentPatientModelID := pModel.TreatmentPatientModelID
			items := pModel.Items
			if int(treatmentPatientModelID.(int64)) == ptreatmentPatientModelID {
				models[k].Items = append(items, item)
				has = true
			}
		}
		if !has {
			var items []TreatmentModelItem

			items = append(items, item)

			pmodel := TreatmentModel{
				ModelName:               modelName.(string),
				TreatmentPatientModelID: int(treatmentPatientModelID.(int64)),
				OperationName:           operationName.(string),
				IsCommon:                isCommon.(bool),
				CreatedTime:             createdTime.(time.Time),
				Items:                   items,
			}
			models = append(models, pmodel)
		}
	}

	ctx.JSON(iris.Map{"code": "200", "data": models, "page_info": pageInfo})
}

// TreatmentPersonalPatientModelList 查询个人和通用治疗医嘱模板
func TreatmentPersonalPatientModelList(ctx iris.Context) {
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

	countSQL := `select count(id) as total from treatment_patient_model where model_name ~$1 and (operation_id=$2 or is_common=true)`
	selectSQL := `select tpm.id as treatment_patient_model_id,tpmi.clinic_treatment_id,tpmi.illustration,ct.name as treatment_name,p.name as operation_name,
	tpm.is_common,tpm.created_time,tpm.model_name,tpmi.times from treatment_patient_model tpm
	left join treatment_patient_model_item tpmi on tpmi.treatment_patient_model_id = tpm.id
	left join clinic_treatment ct on tpmi.clinic_treatment_id = ct.id
	left join personnel p on tpm.operation_id = p.id
	where tpm.model_name ~$1 and (tpm.operation_id=$2 or tpm.is_common=true)`

	if isCommon != "" {
		countSQL += ` and is_common =` + isCommon
		selectSQL += ` and tpm.is_common=` + isCommon
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
	var models []TreatmentModel
	for _, v := range result {
		modelName := v["model_name"]
		treatmentPatientModelID := v["treatment_patient_model_id"]
		treatmentName := v["treatment_name"]
		clinicTreatmentID := v["clinic_treatment_id"]
		illustration := v["illustration"]
		operationName := v["operation_name"]
		isCommon := v["is_common"]
		createdTime := v["created_time"]
		times := v["times"]
		has := false
		for k, pModel := range models {
			ptreatmentPatientModelID := pModel.TreatmentPatientModelID
			items := pModel.Items
			if int(treatmentPatientModelID.(int64)) == ptreatmentPatientModelID {
				item := TreatmentModelItem{
					TreatmentName:     treatmentName.(string),
					Times:             int(times.(int64)),
					ClinicTreatmentID: int(clinicTreatmentID.(int64)),
					Illustration:      illustration,
				}
				models[k].Items = append(items, item)
				has = true
			}
		}
		if !has {
			var items []TreatmentModelItem
			item := TreatmentModelItem{
				TreatmentName:     treatmentName.(string),
				Times:             int(times.(int64)),
				ClinicTreatmentID: int(clinicTreatmentID.(int64)),
				Illustration:      illustration,
			}
			items = append(items, item)

			pmodel := TreatmentModel{
				ModelName:               modelName.(string),
				TreatmentPatientModelID: int(treatmentPatientModelID.(int64)),
				OperationName:           operationName.(string),
				IsCommon:                isCommon.(bool),
				CreatedTime:             createdTime.(time.Time),
				Items:                   items,
			}
			models = append(models, pmodel)
		}
	}

	ctx.JSON(iris.Map{"code": "200", "data": models, "page_info": pageInfo})
}

// TreatmentPatientModelDetail 查询治疗医嘱模板详情
func TreatmentPatientModelDetail(ctx iris.Context) {
	treatmentModelID := ctx.PostValue("treatment_patient_model_id")

	selectmSQL := `select id as treatment_patient_model_id,model_name,is_common,status from treatment_patient_model where id=$1`
	mrows := model.DB.QueryRowx(selectmSQL, treatmentModelID)
	if mrows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询失败"})
		return
	}
	treatmentModel := FormatSQLRowToMap(mrows)

	selectiSQL := `select tpmi.clinic_treatment_id,ct.name,tpmi.times,tpmi.illustration 
		from treatment_patient_model_item tpmi
		left join treatment_patient_model tpm on tpmi.treatment_patient_model_id = tpm.id
		left join clinic_treatment ct on tpmi.clinic_treatment_id = ct.id
		where tpmi.treatment_patient_model_id=$1`

	rows, err := model.DB.Queryx(selectiSQL, treatmentModelID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	treatmentModel["items"] = result
	ctx.JSON(iris.Map{"code": "200", "data": treatmentModel})
}

// TreatmentPatientModelUpdate 修改治疗医嘱模板
func TreatmentPatientModelUpdate(ctx iris.Context) {
	treatmentModelID := ctx.PostValue("treatment_patient_model_id")
	modelName := ctx.PostValue("model_name")
	isCommon := ctx.PostValue("is_common")
	items := ctx.PostValue("items")
	personnelID := ctx.PostValue("operation_id")

	if treatmentModelID == "" || modelName == "" || items == "" {
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
	mrow := model.DB.QueryRowx("select id from treatment_patient_model where id=$1 limit 1", treatmentModelID)
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

	row := model.DB.QueryRowx("select id from treatment_patient_model where model_name=$1 and id!=$2 limit 1", modelName, treatmentModelID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	treatmentModel := FormatSQLRowToMap(row)
	_, ok := treatmentModel["id"]
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
		"treatment_patient_model_id",
		"clinic_treatment_id",
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
	updateSQL := `update treatment_patient_model set model_name=$1,is_common=$2,
	operation_id=$3,updated_time=LOCALTIMESTAMP where id=$4`
	_, err := tx.Exec(updateSQL, modelName, isCommon, personnelID, treatmentModelID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	deleteSQL := "delete from treatment_patient_model_item where treatment_patient_model_id=$1"
	_, errd := tx.Exec(deleteSQL, treatmentModelID)
	if errd != nil {
		fmt.Println("errd ===", errd)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errd.Error()})
		return
	}
	tSetStr := strings.Join(itemSets, ",")
	inserttSQL := "insert into treatment_patient_model_item (" + tSetStr + ") values ()"

	for _, v := range results {
		clinicTreatmentID := v["clinic_treatment_id"]
		times := v["times"]
		illustration := v["illustration"]

		clinicTreatmentSQL := `select id from clinic_treatment where id=$1`
		trow := model.DB.QueryRowx(clinicTreatmentSQL, clinicTreatmentID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "保存模板错误"})
			return
		}
		clinicTreatment := FormatSQLRowToMap(trow)
		_, ok := clinicTreatment["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "选择的治疗医嘱错误"})
			return
		}

		_, errt := tx.Exec(inserttSQL, ToNullInt64(treatmentModelID), ToNullInt64(clinicTreatmentID), ToNullInt64(times), ToNullString(illustration))
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
