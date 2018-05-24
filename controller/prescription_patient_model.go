package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

// PrescriptionWesternPatientModelCreate 创建西药处方模板
func PrescriptionWesternPatientModelCreate(ctx iris.Context) {
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
	row := model.DB.QueryRowx("select id from prescription_western_patient_model where model_name=$1 limit 1", modelName)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
		return
	}
	prescriptionModel := FormatSQLRowToMap(row)
	_, ok := prescriptionModel["id"]
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

	var itemValues []string
	itemSets := []string{
		"prescription_western_patient_model_id",
		"clinic_drug_id",
		"once_dose",
		"once_dose_unit_name",
		"route_administration_name",
		"frequency_name",
		"amount",
		"fetch_address",
		"eff_day",
		"illustration",
	}

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errb})
		return
	}
	var prescriptionWesternPatientModelID string
	err := tx.QueryRow("insert into prescription_western_patient_model (model_name,is_common,operation_id) values ($1,$2,$3) RETURNING id", modelName, isCommon, personnelID).Scan(&prescriptionWesternPatientModelID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	for _, v := range results {
		clinicDrugID := v["clinic_drug_id"]
		onceDose := v["once_dose"]
		onceDoseUnitName := v["once_dose_unit_name"]
		routeAdministrationName := v["route_administration_name"]
		frequencyName := v["frequency_name"]
		times := v["amount"]
		illustration := v["illustration"]
		fetchAddress := v["fetch_address"]
		effDay := v["eff_day"]

		var s []string
		clinicDrugSQL := `select id from clinic_drug where id=$1`
		trow := model.DB.QueryRowx(clinicDrugSQL, clinicDrugID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "保存模板错误"})
			return
		}
		clinicDrug := FormatSQLRowToMap(trow)
		_, ok := clinicDrug["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "选择的药品错误"})
			return
		}
		s = append(s, prescriptionWesternPatientModelID, clinicDrugID, onceDose, onceDoseUnitName, routeAdministrationName, frequencyName, times, fetchAddress)
		if effDay == "" {
			s = append(s, `null`)
		} else {
			s = append(s, effDay)
		}

		if illustration == "" {
			s = append(s, `null`)
		} else {
			s = append(s, "'"+illustration+"'")
		}
		tstr := "(" + strings.Join(s, ",") + ")"
		itemValues = append(itemValues, tstr)
	}
	tSetStr := strings.Join(itemSets, ",")
	tValueStr := strings.Join(itemValues, ",")

	inserttSQL := "insert into prescription_western_patient_model_item (" + tSetStr + ") values " + tValueStr
	fmt.Println("inserttSQL===", inserttSQL)

	_, errt := tx.Exec(inserttSQL)
	if errt != nil {
		fmt.Println("errt ===", errt)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errt.Error()})
		return
	}
	errc := tx.Commit()
	if errc != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errc.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// PrescriptionWesternPatientModelList 查询西药处方模板
func PrescriptionWesternPatientModelList(ctx iris.Context) {
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

	countSQL := `select count(id) as total from prescription_western_patient_model where model_name ~$1`
	selectSQL := `select pwpm.id as prescription_patient_model_id,
	pwpm.is_common,
	pwpm.created_time,
	pwpm.updated_time,
	p.name as operation_name,
	pwpm.model_name,
	pwpmi.clinic_drug_id,
	d.name as drug_name,
	d.specification,
	pwpmi.once_dose,
	pwpmi.once_dose_unit_name,
	pwpmi.route_administration_name,
	pwpmi.frequency_name, 
	pwpmi.eff_day,
	pwpmi.amount,
	d.packing_unit_name, 
	pwpmi.fetch_address,
	pwpmi.illustration,
	d.type,sum(ds.stock_amount) as stock_amount
	from prescription_western_patient_model pwpm
	left join prescription_western_patient_model_item pwpmi on pwpmi.prescription_western_patient_model_id = pwpm.id
	left join clinic_drug cd on pwpmi.clinic_drug_id = cd.id 
	left join drug d on cd.drug_id = d.id		
	left join drug_stock ds on ds.clinic_drug_id = cd.id
  left join personnel p on pwpm.operation_id = p.id
	where pwpm.model_name ~$1
	group by pwpm.id,pwpm.is_common,pwpm.created_time,pwpm.updated_time,p.name,pwpm.model_name,pwpmi.clinic_drug_id,
		d.name,d.specification,pwpmi.once_dose,pwpmi.once_dose_unit_name,pwpmi.route_administration_name,
		pwpmi.frequency_name, pwpmi.eff_day,pwpmi.amount,d.packing_unit_name, pwpmi.fetch_address,pwpmi.illustration,d.type`

	if isCommon != "" {
		countSQL += ` and is_common =` + isCommon
		selectSQL += ` and pwpm.is_common=` + isCommon
	}

	if operationID != "" {
		countSQL += ` and operation_id =` + operationID
		selectSQL += ` and pwpm.operation_id=` + operationID
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

	resData := FormatPrescriptionModel(result)

	ctx.JSON(iris.Map{"code": "200", "data": resData, "page_info": pageInfo})
}

// PrescriptionWesternPersonalPatientModelList 查询个人和通用西药处方模板
func PrescriptionWesternPersonalPatientModelList(ctx iris.Context) {
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

	countSQL := `select count(id) as total from prescription_western_patient_model where model_name ~$1 and (operation_id=$2 or is_common=true)`
	selectSQL := `select pwpm.id as prescription_patient_model_id,
	pwpm.is_common,
	pwpm.created_time,
	pwpm.updated_time,
	p.name as operation_name,
	pwpm.model_name,
	pwpmi.clinic_drug_id,
	d.name as drug_name,
	d.specification,
	cd.stock_amount,
	pwpmi.once_dose,
	pwpmi.once_dose_unit_name,
	pwpmi.route_administration_name,
	pwpmi.frequency_name, 
	pwpmi.eff_day,
	pwpmi.amount,
	d.packing_unit_name, 
	pwpmi.fetch_address,
	pwpmi.illustration,
	d.type
	from prescription_western_patient_model pwpm
	left join prescription_western_patient_model_item pwpmi on pwpmi.prescription_western_patient_model_id = pwpm.id
	left join clinic_drug cd on pwpmi.clinic_drug_id = cd.id 
		left join drug d on cd.drug_id = d.id
    left join personnel p on pwpm.operation_id = p.id
	where pwpm.model_name ~$1 and (pwpm.operation_id=$2 or pwpm.is_common=true)`

	if isCommon != "" {
		countSQL = countSQL + ` and is_common =` + isCommon
		selectSQL += ` and pwpm.is_common=` + isCommon
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

	resData := FormatPrescriptionModel(result)

	ctx.JSON(iris.Map{"code": "200", "data": resData, "page_info": pageInfo})
}

// PrescriptionWesternPatientModelDetail 查询西药处方模板详情
func PrescriptionWesternPatientModelDetail(ctx iris.Context) {
	prescriptionWesternPatientModelID := ctx.PostValue("prescription_patient_model_id")

	selectmSQL := `select id as prescription_patient_model_id,model_name,is_common,status from prescription_western_patient_model where id=$1`
	mrows := model.DB.QueryRowx(selectmSQL, prescriptionWesternPatientModelID)
	if mrows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询失败"})
		return
	}
	prescriptionModel := FormatSQLRowToMap(mrows)

	selectiSQL := `select pwpmi.*,d.name as drug_name from prescription_western_patient_model_item pwpmi
		left join prescription_western_patient_model pwpm on pwpmi.prescription_western_patient_model_id = pwpm.id
		left join clinic_drug cd on pwpmi.clinic_drug_id = cd.id
		left join drug d on cd.drug_id = d.id
		where pwpmi.prescription_western_patient_model_id=$1`

	rows, err := model.DB.Queryx(selectiSQL, prescriptionWesternPatientModelID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	prescriptionModel["items"] = result
	ctx.JSON(iris.Map{"code": "200", "data": prescriptionModel})
}

// PrescriptionWesternPatientModelUpdate 修改西药处方模板
func PrescriptionWesternPatientModelUpdate(ctx iris.Context) {
	prescriptionWesternPatientModelID := ctx.PostValue("prescription_patient_model_id")
	modelName := ctx.PostValue("model_name")
	isCommon := ctx.PostValue("is_common")
	items := ctx.PostValue("items")
	personnelID := ctx.PostValue("operation_id")

	if prescriptionWesternPatientModelID == "" || modelName == "" || items == "" {
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
	mrow := model.DB.QueryRowx("select id from prescription_western_patient_model where id=$1 limit 1", prescriptionWesternPatientModelID)
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

	row := model.DB.QueryRowx("select id from prescription_western_patient_model where model_name=$1 and id!=$2 limit 1", modelName, prescriptionWesternPatientModelID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	prescriptionModel := FormatSQLRowToMap(row)
	_, ok := prescriptionModel["id"]
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

	var itemValues []string
	itemSets := []string{
		"prescription_western_patient_model_id",
		"clinic_drug_id",
		"once_dose",
		"once_dose_unit_name",
		"route_administration_name",
		"frequency_name",
		"amount",
		"fetch_address",
		"eff_day",
		"illustration",
	}

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errb})
		return
	}
	updateSQL := `update prescription_western_patient_model set model_name=$1,is_common=$2,
	operation_id=$3,updated_time=LOCALTIMESTAMP where id=$4`
	_, err := tx.Exec(updateSQL, modelName, isCommon, personnelID, prescriptionWesternPatientModelID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	for _, v := range results {
		clinicDrugID := v["clinic_drug_id"]
		onceDose := v["once_dose"]
		onceDoseUnitName := v["once_dose_unit_name"]
		routeAdministrationName := v["route_administration_name"]
		frequencyName := v["frequency_name"]
		times := v["amount"]
		illustration := v["illustration"]
		fetchAddress := v["fetch_address"]
		effDay := v["eff_day"]

		var s []string
		clinicDrugSQL := `select id from clinic_drug where id=$1`
		trow := model.DB.QueryRowx(clinicDrugSQL, clinicDrugID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "保存模板错误"})
			return
		}
		clinicDrug := FormatSQLRowToMap(trow)
		_, ok := clinicDrug["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "选择的药品错误"})
			return
		}
		s = append(s, prescriptionWesternPatientModelID, clinicDrugID, onceDose, onceDoseUnitName, routeAdministrationName, frequencyName, times, fetchAddress)
		if effDay == "" {
			s = append(s, `null`)
		} else {
			s = append(s, effDay)
		}

		if illustration == "" {
			s = append(s, `null`)
		} else {
			s = append(s, "'"+illustration+"'")
		}
		tstr := "(" + strings.Join(s, ",") + ")"
		itemValues = append(itemValues, tstr)
	}
	tSetStr := strings.Join(itemSets, ",")
	tValueStr := strings.Join(itemValues, ",")

	deleteSQL := "delete from prescription_western_patient_model_item where prescription_western_patient_model_id=$1"
	fmt.Println("deleteSQL===", deleteSQL)
	_, errd := tx.Exec(deleteSQL, prescriptionWesternPatientModelID)
	if errd != nil {
		fmt.Println("errd ===", errd)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errd.Error()})
		return
	}

	inserttSQL := "insert into prescription_western_patient_model_item (" + tSetStr + ") values " + tValueStr
	fmt.Println("inserttSQL===", inserttSQL)
	_, errt := tx.Exec(inserttSQL)
	if errt != nil {
		fmt.Println("errt ===", errt)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errt.Error()})
		return
	}
	errc := tx.Commit()
	if errc != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errc.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// PrescriptionChinesePatientModelCreate 创建中药处方模板
func PrescriptionChinesePatientModelCreate(ctx iris.Context) {
	modelName := ctx.PostValue("model_name")
	isCommon := ctx.PostValue("is_common")

	routeAdministrationName := ctx.PostValue("route_administration_name")
	frequencyName := ctx.PostValue("frequency_name")
	amount := ctx.PostValue("amount")
	fetchAddress := ctx.PostValue("fetch_address")
	effDay := ctx.PostValue("eff_day")
	medicineIllustration := ctx.PostValue("medicine_illustration")

	items := ctx.PostValue("items")
	personnelID := ctx.PostValue("operation_id")
	fmt.Println("amount===", amount)

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

	row := model.DB.QueryRowx("select id from prescription_chinese_patient_model where model_name=$1 limit 1", modelName)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
		return
	}
	prescriptionModel := FormatSQLRowToMap(row)
	_, ok := prescriptionModel["id"]
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

	var itemValues []string
	itemSets := []string{
		"prescription_chinese_patient_model_id",
		"clinic_drug_id",
		"once_dose",
		"once_dose_unit_name",
		"amount",
		"special_illustration",
	}

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errb})
		return
	}
	var prescriptionChinesePatientModelID string
	err := tx.QueryRow(`insert into prescription_chinese_patient_model 
		(model_name,is_common,operation_id,route_administration_name,frequency_name,amount,fetch_address,eff_day,medicine_illustration) 
		values ($1,$2,$3,$4,$5,$6,$7,$8,$9) 
		RETURNING id`, modelName, isCommon, personnelID, routeAdministrationName, frequencyName, amount, fetchAddress, effDay, medicineIllustration).Scan(&prescriptionChinesePatientModelID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	for _, v := range results {
		clinicDrugID := v["clinic_drug_id"]
		onceDose := v["once_dose"]
		onceDoseUnitName := v["once_dose_unit_name"]
		times := v["amount"]
		illustration := v["special_illustration"]
		var s []string
		clinicDrugSQL := `select id from clinic_drug where id=$1`
		trow := model.DB.QueryRowx(clinicDrugSQL, clinicDrugID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "保存模板错误"})
			return
		}
		clinicDrug := FormatSQLRowToMap(trow)
		_, ok := clinicDrug["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "选择的药品错误"})
			return
		}
		s = append(s, prescriptionChinesePatientModelID, clinicDrugID, onceDose, onceDoseUnitName, times)
		if illustration == "" {
			s = append(s, `null`)
		} else {
			s = append(s, "'"+illustration+"'")
		}
		tstr := "(" + strings.Join(s, ",") + ")"
		itemValues = append(itemValues, tstr)
	}
	tSetStr := strings.Join(itemSets, ",")
	tValueStr := strings.Join(itemValues, ",")

	inserttSQL := "insert into prescription_chinese_patient_model_item (" + tSetStr + ") values " + tValueStr
	fmt.Println("inserttSQL===", inserttSQL)

	_, errt := tx.Exec(inserttSQL)
	if errt != nil {
		fmt.Println("errt ===", errt)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errt.Error()})
		return
	}
	errc := tx.Commit()
	if errc != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errc.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// PrescriptionChinesePatientModelList 查询中药处方模板
func PrescriptionChinesePatientModelList(ctx iris.Context) {
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

	countSQL := `select count(id) as total from prescription_chinese_patient_model where model_name ~$1`
	selectSQL := `select 
	pcpm.id as prescription_patient_model_id,
	pcpm.is_common,
	pcpm.created_time,
	pcpm.updated_time,
	p.name as operation_name,
	pcpm.model_name,
	pcpm.route_administration_name as info_route_administration_name,
	pcpm.eff_day as info_eff_day,
	pcpm.amount as info_amount,
	pcpm.frequency_name as info_frequencyName,
	pcpm.fetch_address as info_fetch_address,
	pcpm.medicine_illustration,
	pcpmi.clinic_drug_id,
	d.name as drug_name,
	d.type,
	pcpmi.once_dose,
	pcpmi.once_dose_unit_name,
	pcpmi.special_illustration,
	pcpmi.amount,sum(ds.stock_amount) as stock_amount
	from prescription_chinese_patient_model pcpm
	left join prescription_chinese_patient_model_item pcpmi on pcpmi.prescription_chinese_patient_model_id = pcpm.id
	left join clinic_drug cd on pcpmi.clinic_drug_id = cd.id 
	left join drug d on cd.drug_id = d.id
	left join drug_stock ds on ds.clinic_drug_id = cd.id
	left join personnel p on pcpm.operation_id = p.id
	where pcpm.model_name ~$1
	group by pcpm.id,pcpm.is_common,pcpm.created_time,pcpm.updated_time,p.name,pcpm.model_name,pcpm.route_administration_name,
		pcpm.eff_day,pcpm.amount,pcpm.frequency_name,pcpm.fetch_address,pcpm.medicine_illustration,pcpmi.clinic_drug_id,
		d.name,d.type,pcpmi.once_dose,pcpmi.once_dose_unit_name,pcpmi.special_illustration,pcpmi.amount`
	fmt.Println("countSQL===", countSQL)
	fmt.Println("selectSQL===", selectSQL)
	if isCommon != "" {
		countSQL += ` and is_common =` + isCommon
		selectSQL += ` and pcpm.is_common=` + isCommon
	}

	if operationID != "" {
		countSQL += ` and operation_id =` + operationID
		selectSQL += ` and pcpm.operation_id=` + operationID
	}

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

	resData := FormatPrescriptionModel(result)

	ctx.JSON(iris.Map{"code": "200", "data": resData, "page_info": pageInfo})
}

// PrescriptionChinesePersonalPatientModelList 查询个人和通用中药处方模板
func PrescriptionChinesePersonalPatientModelList(ctx iris.Context) {
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

	countSQL := `select count(id) as total from prescription_chinese_patient_model where model_name ~$1 and (operation_id=$2 or is_common=true)`
	selectSQL := `select 
	pcpm.id as prescription_patient_model_id,
	pcpm.is_common,
	pcpm.created_time,
	pcpm.updated_time,
	p.name as operation_name,
	pcpm.model_name,
	pcpm.route_administration_name as info_route_administration_name,
	pcpm.eff_day as info_eff_day,
	pcpm.amount as info_amount,
	pcpm.frequency_name as info_frequencyName,
	pcpm.fetch_address as info_fetch_address,
	pcpm.medicine_illustration,
	pcpmi.clinic_drug_id,
	d.name as drug_name,
	d.type,
	cd.stock_amount,
	pcpmi.once_dose,
	pcpmi.once_dose_unit_name,
	pcpmi.special_illustration,
	pcpmi.amount
	from prescription_chinese_patient_model pcpm
	left join prescription_chinese_patient_model_item pcpmi on pcpmi.prescription_chinese_patient_model_id = pcpm.id
	left join clinic_drug cd on pcpmi.clinic_drug_id = cd.id 
	left join drug d on cd.drug_id = d.id
	left join personnel p on pcpm.operation_id = p.id
	where pcpm.model_name ~$1 and (pcpm.operation_id=$2 or pcpm.is_common=true)`

	if isCommon != "" {
		countSQL += ` and is_common =` + isCommon
		selectSQL += ` and pcpm.is_common=` + isCommon
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

	resData := FormatPrescriptionModel(result)

	ctx.JSON(iris.Map{"code": "200", "data": resData, "page_info": pageInfo})
}

// PrescriptionChinesePatientModelDetail 查询中药处方模板详情
func PrescriptionChinesePatientModelDetail(ctx iris.Context) {
	prescriptionChinesePatientModelID := ctx.PostValue("prescription_patient_model_id")

	selectmSQL := `select pcpm.id as prescription_patient_model_id,pcpm.model_name,pcpm.is_common,pcpm.status,pcpm.route_administration_name,
		pcpm.frequency_name,pcpm.amount,pcpm.eff_day,pcpm.fetch_address,pcpm.medicine_illustration
		from prescription_chinese_patient_model pcpm
		where pcpm.id=$1`
	mrows := model.DB.QueryRowx(selectmSQL, prescriptionChinesePatientModelID)
	if mrows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询失败"})
		return
	}
	prescriptionModel := FormatSQLRowToMap(mrows)

	selectiSQL := `select pcpmi.*,d.name as drug_name
		from prescription_chinese_patient_model_item pcpmi
		left join prescription_chinese_patient_model pwpm on pcpmi.prescription_chinese_patient_model_id = pwpm.id
		left join clinic_drug cd on pcpmi.clinic_drug_id = cd.id
		left join drug d on cd.drug_id = d.id
		where pcpmi.prescription_chinese_patient_model_id=$1`

	rows, err := model.DB.Queryx(selectiSQL, prescriptionChinesePatientModelID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	prescriptionModel["items"] = result
	ctx.JSON(iris.Map{"code": "200", "data": prescriptionModel})
}

// PrescriptionChinesePatientModelUpdate 修改中药药处方模板
func PrescriptionChinesePatientModelUpdate(ctx iris.Context) {
	prescriptionChinesePatientModelID := ctx.PostValue("prescription_patient_model_id")
	modelName := ctx.PostValue("model_name")
	isCommon := ctx.PostValue("is_common")

	routeAdministrationName := ctx.PostValue("route_administration_name")
	frequencyName := ctx.PostValue("frequency_name")
	amount := ctx.PostValue("amount")
	fetchAddress := ctx.PostValue("fetch_address")
	effDay := ctx.PostValue("eff_day")
	medicineIllustration := ctx.PostValue("medicine_illustration")

	items := ctx.PostValue("items")
	personnelID := ctx.PostValue("operation_id")
	fmt.Println("amount===", amount)

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

	mrow := model.DB.QueryRowx("select id from prescription_chinese_patient_model where id=$1 limit 1", prescriptionChinesePatientModelID)
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

	row := model.DB.QueryRowx("select id from prescription_chinese_patient_model where model_name=$1 and id!=$2 limit 1", modelName, prescriptionChinesePatientModelID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	prescriptionModel := FormatSQLRowToMap(row)
	_, ok := prescriptionModel["id"]
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

	var itemValues []string
	itemSets := []string{
		"prescription_chinese_patient_model_id",
		"clinic_drug_id",
		"once_dose",
		"once_dose_unit_name",
		"amount",
		"special_illustration",
	}

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errb})
		return
	}

	updateSQL := `update prescription_chinese_patient_model set model_name=$1,is_common=$2,
		operation_id=$3,route_administration_name=$4,frequency_name=$5,amount=$6,fetch_address=$7,
		eff_day=$8,medicine_illustration=$9,updated_time=LOCALTIMESTAMP where id=$10`
	_, err := tx.Exec(updateSQL, modelName, isCommon, personnelID, routeAdministrationName, frequencyName, amount, fetchAddress, effDay, medicineIllustration, prescriptionChinesePatientModelID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	for _, v := range results {
		clinicDrugID := v["clinic_drug_id"]
		onceDose := v["once_dose"]
		onceDoseUnitName := v["once_dose_unit_name"]
		times := v["amount"]
		illustration := v["special_illustration"]
		var s []string
		clinicDrugSQL := `select id from clinic_drug where id=$1`
		trow := model.DB.QueryRowx(clinicDrugSQL, clinicDrugID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "保存模板错误"})
			return
		}
		clinicDrug := FormatSQLRowToMap(trow)
		_, ok := clinicDrug["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "选择的药品错误"})
			return
		}
		s = append(s, prescriptionChinesePatientModelID, clinicDrugID, onceDose, onceDoseUnitName, times)
		if illustration == "" {
			s = append(s, `null`)
		} else {
			s = append(s, "'"+illustration+"'")
		}
		tstr := "(" + strings.Join(s, ",") + ")"
		itemValues = append(itemValues, tstr)
	}
	tSetStr := strings.Join(itemSets, ",")
	tValueStr := strings.Join(itemValues, ",")

	deleteSQL := "delete from prescription_chinese_patient_model_item where prescription_chinese_patient_model_id=$1"
	fmt.Println("deleteSQL===", deleteSQL)
	_, errd := tx.Exec(deleteSQL, prescriptionChinesePatientModelID)
	if errd != nil {
		fmt.Println("errd ===", errd)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errd.Error()})
		return
	}

	inserttSQL := "insert into prescription_chinese_patient_model_item (" + tSetStr + ") values " + tValueStr
	fmt.Println("inserttSQL===", inserttSQL)

	_, errt := tx.Exec(inserttSQL)
	if errt != nil {
		fmt.Println("errt ===", errt)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errt.Error()})
		return
	}
	errc := tx.Commit()
	if errc != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errc.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}
