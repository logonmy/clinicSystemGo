package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

//LaboratoryCreate 检验医嘱创建
func LaboratoryCreate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	pyCode := ctx.PostValue("py_code")
	idcCode := ctx.PostValue("idc_code")
	unitName := ctx.PostValue("unit_name")
	timeReport := ctx.PostValue("time_report")
	clinicalSignificance := ctx.PostValue("clinical_significance")
	remark := ctx.PostValue("remark")
	laboratorySample := ctx.PostValue("laboratory_sample")
	cuvetteColorName := ctx.PostValue("cuvette_color_name")

	mergeFlag := ctx.PostValue("merge_flag")
	cost := ctx.PostValue("cost")
	price := ctx.PostValue("price")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")
	isDelivery := ctx.PostValue("is_delivery")

	if clinicID == "" || name == "" || price == "" || status == "" || isDiscount == "" || isDelivery == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from clinic where id=$1 limit 1", clinicID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)
	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据错误"})
		return
	}

	laboratorySets := []string{"name"}
	laboratoryValues := []string{"'" + name + "'"}

	clinicLaboratorySets := []string{"clinic_id"}
	clinicLaboratoryValues := []string{clinicID}

	if enName != "" {
		laboratorySets = append(laboratorySets, "en_name")
		laboratoryValues = append(laboratoryValues, "'"+enName+"'")
	}
	if pyCode != "" {
		laboratorySets = append(laboratorySets, "py_code")
		laboratoryValues = append(laboratoryValues, "'"+pyCode+"'")
	}
	if idcCode != "" {
		laboratorySets = append(laboratorySets, "idc_code")
		laboratoryValues = append(laboratoryValues, "'"+idcCode+"'")
	}
	if unitName != "" {
		laboratorySets = append(laboratorySets, "unit_name")
		laboratoryValues = append(laboratoryValues, "'"+unitName+"'")
	}
	if timeReport != "" {
		laboratorySets = append(laboratorySets, "time_report")
		laboratoryValues = append(laboratoryValues, "'"+timeReport+"'")
	}
	if clinicalSignificance != "" {
		laboratorySets = append(laboratorySets, "clinical_significance")
		laboratoryValues = append(laboratoryValues, "'"+clinicalSignificance+"'")
	}
	if laboratorySample != "" {
		laboratorySets = append(laboratorySets, "laboratory_sample")
		laboratoryValues = append(laboratoryValues, "'"+laboratorySample+"'")
	}
	if cuvetteColorName != "" {
		laboratorySets = append(laboratorySets, "cuvette_color_name")
		laboratoryValues = append(laboratoryValues, "'"+cuvetteColorName+"'")
	}

	if status != "" {
		clinicLaboratorySets = append(clinicLaboratorySets, "status")
		clinicLaboratoryValues = append(clinicLaboratoryValues, status)
	}
	if remark != "" {
		clinicLaboratorySets = append(laboratorySets, "remark")
		clinicLaboratoryValues = append(laboratoryValues, "'"+remark+"'")
	}
	if mergeFlag != "" {
		clinicLaboratorySets = append(clinicLaboratorySets, "merge_flag")
		clinicLaboratoryValues = append(clinicLaboratoryValues, mergeFlag)
	}
	if cost != "" {
		clinicLaboratorySets = append(clinicLaboratorySets, "cost")
		clinicLaboratoryValues = append(clinicLaboratoryValues, cost)
	}
	if price != "" {
		clinicLaboratorySets = append(clinicLaboratorySets, "price")
		clinicLaboratoryValues = append(clinicLaboratoryValues, price)
	}
	if isDelivery != "" {
		clinicLaboratorySets = append(clinicLaboratorySets, "is_delivery")
		clinicLaboratoryValues = append(clinicLaboratoryValues, isDelivery)
	}
	if isDiscount != "" {
		clinicLaboratorySets = append(clinicLaboratorySets, "is_discount")
		clinicLaboratoryValues = append(clinicLaboratoryValues, isDiscount)
	}

	laboratoryItemSetStr := strings.Join(laboratorySets, ",")
	laboratoryItemValueStr := strings.Join(laboratoryValues, ",")

	laboratoryInsertSQL := "insert into laboratory (" + laboratoryItemSetStr + ") values (" + laboratoryItemValueStr + ") RETURNING id;"
	fmt.Println("laboratoryInsertSQL==", laboratoryInsertSQL)

	tx, err := model.DB.Begin()
	var laboratoryID string

	lrow := model.DB.QueryRowx("select id from laboratory where name=$1 limit 1", name)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
		return
	}
	laboratory := FormatSQLRowToMap(lrow)
	_, lok := laboratory["id"]
	if !lok {
		err = tx.QueryRow(laboratoryInsertSQL).Scan(&laboratoryID)
		if err != nil {
			fmt.Println("err ===", err)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "1", "msg": err})
			return
		}
	} else {
		laboratoryID = strconv.Itoa(int(laboratory["id"].(int64)))
	}

	fmt.Println("laboratoryID====", laboratoryID)

	clinicLaboratorySets = append(clinicLaboratorySets, "laboratory_id")
	clinicLaboratoryValues = append(clinicLaboratoryValues, laboratoryID)

	clinicLaboratorySetStr := strings.Join(clinicLaboratorySets, ",")
	clinicLaboratoryValueStr := strings.Join(clinicLaboratoryValues, ",")

	clinicLaboratoryInsertSQL := "insert into clinic_laboratory (" + clinicLaboratorySetStr + ") values (" + clinicLaboratoryValueStr + ")"
	fmt.Println("clinicLaboratoryInsertSQL==", clinicLaboratoryInsertSQL)

	_, err2 := tx.Exec(clinicLaboratoryInsertSQL)
	if err2 != nil {
		fmt.Println(" err2====", err2)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}

	err3 := tx.Commit()
	if err3 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": laboratoryID})
}

//LaboratoryAssociation 关联检验项目
func LaboratoryAssociation(ctx iris.Context) {
	clinicLaboratoryID := ctx.PostValue("clinic_laboratory_id")
	items := ctx.PostValue("items")

	if clinicLaboratoryID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	if items == "" {
		ctx.JSON(iris.Map{"code": "200", "data": nil})
		return
	}

	var results []map[string]string
	errj := json.Unmarshal([]byte(items), &results)
	fmt.Println("===", results)
	if errj != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": errj.Error()})
		return
	}
	row := model.DB.QueryRowx("select id from clinic_laboratory where id=$1 limit 1", clinicLaboratoryID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "关联项目失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)

	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "所选诊所检验医嘱不存在"})
		return
	}
	var values []string
	sets := []string{
		"clinic_laboratory_id",
		"clinic_laboratory_item_id",
		"default_result",
	}
	for _, v := range results {
		clinicLaboratoryItemID := v["clinic_laboratory_item_id"]
		laboratoryItemName := v["name"]
		defaultResult := v["default_result"]
		var s []string
		row := model.DB.QueryRowx(`select id from clinic_laboratory_item where id=$1 limit 1`, clinicLaboratoryItemID)
		if row == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "关联项目失败"})
			return
		}
		clinicLaboratoryItem := FormatSQLRowToMap(row)
		_, ok := clinicLaboratoryItem["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "关联的" + laboratoryItemName + "检验项目不存在"})
			return
		}
		s = append(s, clinicLaboratoryID, clinicLaboratoryItemID)
		if defaultResult != "" {
			s = append(s, "'"+defaultResult+"'")
		} else {
			s = append(s, `null`)
		}
		str := strings.Join(s, ",")
		str = "(" + str + ")"
		values = append(values, str)
	}
	setStr := strings.Join(sets, ",")
	valueStr := strings.Join(values, ",")

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errb})
		return
	}
	_, errd := tx.Exec("delete from clinic_laboratory_association where clinic_laboratory_id=$1", clinicLaboratoryID)
	if errd != nil {
		fmt.Println("errd ===", errd)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errd.Error()})
		return
	}

	insertSQL := "insert into clinic_laboratory_association (" + setStr + ") values " + valueStr
	fmt.Println("insertSQL===", insertSQL)

	_, err := tx.Exec(insertSQL)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": "请检验是否漏填"})
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

//AssociationList 检验医嘱关联项目列表
func AssociationList(ctx iris.Context) {
	clinicLaboratoryID := ctx.PostValue("clinic_laboratory_id")
	if clinicLaboratoryID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select cls.clinic_laboratory_item_id,li.id as laboratory_item_id,li.name,li.en_name,li.unit_name,li.is_special,cls.default_result,
	li.data_type,lir.reference_sex,lir.stomach_status,lir.is_pregnancy,lir.reference_max,lir.reference_min,cli.status
	from clinic_laboratory_association cls
	left join clinic_laboratory_item cli on cls.clinic_laboratory_item_id = cli.id
	left join laboratory_item li on cli.laboratory_item_id = li.id
	left join laboratory_item_reference lir on lir.laboratory_item_id = li.id
	where cls.clinic_laboratory_id=$1`

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL, clinicLaboratoryID)
	results = FormatSQLRowsToMapArray(rows)
	laboratoryItems := FormatLaboratoryItem(results)

	ctx.JSON(iris.Map{"code": "200", "data": laboratoryItems})
}

//LaboratoryList 检验医嘱列表
func LaboratoryList(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	keyword := ctx.PostValue("keyword")
	status := ctx.PostValue("status")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if clinicID == "" {
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

	row := model.DB.QueryRowx("select id from clinic where id=$1 limit 1", clinicID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "查询失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)

	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "所在诊所不存在"})
		return
	}

	countSQL := `select count(cl.id) as total from clinic_laboratory cl
		left join laboratory l on cl.laboratory_id = l.id
		where cl.clinic_id=$1`
	selectSQL := `select l.id as laboratory_id,cl.id as clinic_laboratory_id,l.name as laboratory_name,l.unit_name,cl.price,l.py_code,cl.is_discount,
		l.remark,cl.status
		from clinic_laboratory cl
		left join laboratory l on cl.laboratory_id = l.id
		where cl.clinic_id=$1`

	if keyword != "" {
		countSQL += " and l.name ~'" + keyword + "'"
		selectSQL += " and l.name ~'" + keyword + "'"
	}
	if status != "" {
		countSQL += " and cl.status=" + status
		selectSQL += " and cl.status=" + status
	}

	fmt.Println("countSQL===", countSQL)
	fmt.Println("selectSQL===", selectSQL)
	total := model.DB.QueryRowx(countSQL, clinicID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL+" offset $2 limit $3", clinicID, offset, limit)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

//LaboratoryDetail 检验医嘱详情
func LaboratoryDetail(ctx iris.Context) {
	clinicLaboratoryID := ctx.PostValue("clinic_laboratory_id")

	if clinicLaboratoryID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	sql := `select l.id as laboratory_id,cl.id as clinic_laboratory_id,l.name,l.en_name,l.unit_name,l.py_code,l.idc_code,l.remark,
		l.time_report,l.clinical_significance,l.laboratory_sample,l.cuvette_color_name,
		cl.cost,cl.is_discount,cl.status,cl.merge_flag,cl.price,cl.is_delivery
		from clinic_laboratory cl
		left join laboratory l on cl.laboratory_id = l.id
		where cl.id=$1`
	fmt.Println("sql==", sql)
	arows := model.DB.QueryRowx(sql, clinicLaboratoryID)
	if arows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询结果不存在"})
		return
	}
	result := FormatSQLRowToMap(arows)
	ctx.JSON(iris.Map{"code": "200", "data": result})
}

//LaboratoryUpdate 检验医嘱修改
func LaboratoryUpdate(ctx iris.Context) {
	clinicLaboratoryID := ctx.PostValue("clinic_laboratory_id")

	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	pyCode := ctx.PostValue("py_code")
	idcCode := ctx.PostValue("idc_code")
	unitName := ctx.PostValue("unit_name")
	timeReport := ctx.PostValue("time_report")
	clinicalSignificance := ctx.PostValue("clinical_significance")
	remark := ctx.PostValue("remark")
	laboratorySample := ctx.PostValue("laboratory_sample")
	cuvetteColorName := ctx.PostValue("cuvette_color_name")

	mergeFlag := ctx.PostValue("merge_flag")
	cost := ctx.PostValue("cost")
	price := ctx.PostValue("price")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")
	isDelivery := ctx.PostValue("is_delivery")

	if clinicLaboratoryID == "" || name == "" || price == "" || status == "" || isDiscount == "" || isDelivery == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if mergeFlag == "" {
		mergeFlag = "0"
	}

	crow := model.DB.QueryRowx("select id,clinic_id,laboratory_id from clinic_laboratory where id=$1 limit 1", clinicLaboratoryID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicLaboratory := FormatSQLRowToMap(crow)
	_, rok := clinicLaboratory["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所检验医嘱数据错误"})
		return
	}
	laboratoryID := strconv.FormatInt(clinicLaboratory["laboratory_id"].(int64), 10)
	clinicID := strconv.FormatInt(clinicLaboratory["clinic_id"].(int64), 10)
	fmt.Println("laboratoryID====", laboratoryID)
	fmt.Println("clinicID====", clinicID)

	lrow := model.DB.QueryRowx("select id from laboratory where name=$1 and id!=$2 limit 1", name, laboratoryID)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	laboratoryItem := FormatSQLRowToMap(lrow)
	_, lok := laboratoryItem["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "1", "msg": "检验医嘱名称已存在"})
		return
	}

	laboratoryUpdateSQL := `update laboratory set name=$1,en_name=$2,py_code=$3,idc_code=$4,
		unit_name=$5,time_report=$6,clinical_significance=$7,remark=$8,laboratory_sample=$9,cuvette_color_name=$10 where id=$11`
	fmt.Println("laboratoryUpdateSQL==", laboratoryUpdateSQL)

	clinicLaboratoryUpdateSQL := `update clinic_laboratory set clinic_id=$1,laboratory_id=$2,cost=$3,price=$4,status=$5,is_discount=$6,is_delivery=$7,merge_flag=$8 where id=$9`
	fmt.Println("clinicLaboratoryUpdateSQL==", clinicLaboratoryUpdateSQL)

	tx, err := model.DB.Begin()
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}

	_, err1 := tx.Exec(laboratoryUpdateSQL, name, enName, pyCode, idcCode, unitName, timeReport, clinicalSignificance, remark, laboratorySample, cuvetteColorName, laboratoryID)
	if err1 != nil {
		fmt.Println(" err1====", err1)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	_, err2 := tx.Exec(clinicLaboratoryUpdateSQL, clinicID, laboratoryID, cost, price, status, isDiscount, isDelivery, mergeFlag, clinicLaboratoryID)
	if err2 != nil {
		fmt.Println(" err2====", err2)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}

	err3 := tx.Commit()
	if err3 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": clinicLaboratoryID})
}

//LaboratoryItemCreate 检验项目创建
func LaboratoryItemCreate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	instrumentCode := ctx.PostValue("instrument_code")
	unitName := ctx.PostValue("unit_name")
	clinicalSignificance := ctx.PostValue("clinical_significance")
	dataType := ctx.PostValue("data_type")

	isSpecial := ctx.PostValue("is_special")
	referenceMax := ctx.PostValue("reference_max")
	referenceMin := ctx.PostValue("reference_min")
	items := ctx.PostValue("items")

	status := ctx.PostValue("status")
	isDelivery := ctx.PostValue("is_delivery")

	if clinicID == "" || name == "" || dataType == "" || isSpecial == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from clinic where id=$1 limit 1", clinicID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)
	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据错误"})
		return
	}

	lrow := model.DB.QueryRowx("select id from laboratory_item where name=$1 limit 1", name)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
		return
	}
	laboratoryItem := FormatSQLRowToMap(lrow)
	_, lok := laboratoryItem["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "1", "msg": "检验项目名称已存在"})
		return
	}

	laboratoryItemSets := []string{"name", "data_type", "is_special"}
	laboratoryItemValues := []string{"'" + name + "'", dataType, isSpecial}

	var itemReferenceSets []string
	var itemReferenceValues []string

	clinicLaboratoryItemSets := []string{"clinic_id"}
	clinicLaboratoryItemValues := []string{clinicID}

	if enName != "" {
		laboratoryItemSets = append(laboratoryItemSets, "en_name")
		laboratoryItemValues = append(laboratoryItemValues, "'"+enName+"'")
	}
	if instrumentCode != "" {
		laboratoryItemSets = append(laboratoryItemSets, "instrument_code")
		laboratoryItemValues = append(laboratoryItemValues, "'"+instrumentCode+"'")
	}
	if unitName != "" {
		laboratoryItemSets = append(laboratoryItemSets, "unit_name")
		laboratoryItemValues = append(laboratoryItemValues, "'"+unitName+"'")
	}
	if clinicalSignificance != "" {
		laboratoryItemSets = append(laboratoryItemSets, "clinical_significance")
		laboratoryItemValues = append(laboratoryItemValues, "'"+clinicalSignificance+"'")
	}

	if status != "" {
		clinicLaboratoryItemSets = append(clinicLaboratoryItemSets, "status")
		clinicLaboratoryItemValues = append(clinicLaboratoryItemValues, status)
	}
	if isDelivery != "" {
		clinicLaboratoryItemSets = append(clinicLaboratoryItemSets, "is_delivery")
		clinicLaboratoryItemValues = append(clinicLaboratoryItemValues, isDelivery)
	}

	laboratoryItemSetStr := strings.Join(laboratoryItemSets, ",")
	laboratoryItemValueStr := strings.Join(laboratoryItemValues, ",")

	laboratoryItemInsertSQL := "insert into laboratory_item (" + laboratoryItemSetStr + ") values (" + laboratoryItemValueStr + ") RETURNING id;"
	fmt.Println("laboratoryItemInsertSQL==", laboratoryItemInsertSQL)

	tx, err := model.DB.Begin()
	var laboratoryItemID string
	err = tx.QueryRow(laboratoryItemInsertSQL).Scan(&laboratoryItemID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	fmt.Println("laboratoryItemID====", laboratoryItemID)

	itemReferenceSets = append(itemReferenceSets, "laboratory_item_id")
	if isSpecial == "false" {
		itemReferenceValues = append(itemReferenceValues, laboratoryItemID)
		if referenceMax != "" {
			itemReferenceSets = append(itemReferenceSets, "reference_max")
			itemReferenceValues = append(itemReferenceValues, "'"+referenceMax+"'")
		}
		if referenceMin != "" {
			itemReferenceSets = append(itemReferenceSets, "reference_min")
			itemReferenceValues = append(itemReferenceValues, "'"+referenceMin+"'")
		}

	} else if isSpecial == "true" && items != "" {
		itemReferenceSets = append(itemReferenceSets, "reference_sex", "age_max", "age_min", "reference_max", "reference_min", "stomach_status", "is_pregnancy")
		var results []map[string]string
		reErr := json.Unmarshal([]byte(items), &results)
		if reErr != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": reErr.Error()})
			return
		}
		for _, v := range results {
			var s []string
			s = append(s, laboratoryItemID)
			referenceSex := v["reference_sex"]
			ageMax := v["age_max"]
			ageMin := v["age_min"]
			referenceMax := v["reference_max"]
			referenceMin := v["reference_min"]
			stomachStatus := v["stomach_status"]
			isPregnancy := v["is_pregnancy"]
			if referenceSex != "" {
				s = append(s, "'"+referenceSex+"'")
			} else {
				s = append(s, `null`)
			}
			if ageMax != "" {
				s = append(s, ageMax)
			} else {
				s = append(s, `null`)
			}
			if ageMin != "" {
				s = append(s, ageMin)
			} else {
				s = append(s, `null`)
			}
			if referenceMax != "" {
				s = append(s, "'"+referenceMax+"'")
			} else {
				s = append(s, `null`)
			}
			if referenceMin != "" {
				s = append(s, "'"+referenceMin+"'")
			} else {
				s = append(s, `null`)
			}
			if stomachStatus != "" {
				s = append(s, "'"+stomachStatus+"'")
			} else {
				s = append(s, `null`)
			}
			if isPregnancy != "" {
				s = append(s, isPregnancy)
			} else {
				s = append(s, `null`)
			}
			str := strings.Join(s, ",")
			str = "(" + str + ")"
			itemReferenceValues = append(itemReferenceValues, str)
		}
	} else {
		ctx.JSON(iris.Map{"code": "1", "msg": "参考值是否特殊数据格式错误"})
		return
	}

	itemReferenceSetStr := strings.Join(itemReferenceSets, ",")
	itemReferenceValueStr := strings.Join(itemReferenceValues, ",")

	clinicLaboratoryItemSets = append(clinicLaboratoryItemSets, "laboratory_item_id")
	clinicLaboratoryItemValues = append(clinicLaboratoryItemValues, laboratoryItemID)

	clinicLaboratoryItemSetStr := strings.Join(clinicLaboratoryItemSets, ",")
	clinicLaboratoryItemValueStr := strings.Join(clinicLaboratoryItemValues, ",")

	itemReferenceInsertSQL := "insert into laboratory_item_reference (" + itemReferenceSetStr + ") values (" + itemReferenceValueStr + ")"
	if isSpecial == "true" {
		itemReferenceInsertSQL = "insert into laboratory_item_reference (" + itemReferenceSetStr + ") values " + itemReferenceValueStr
	}
	fmt.Println("itemReferenceInsertSQL==", itemReferenceInsertSQL)

	clinicLaboratoryItemInsertSQL := "insert into clinic_laboratory_item (" + clinicLaboratoryItemSetStr + ") values (" + clinicLaboratoryItemValueStr + ")"
	fmt.Println("clinicLaboratoryItemInsertSQL==", clinicLaboratoryItemInsertSQL)

	_, err1 := tx.Exec(itemReferenceInsertSQL)
	if err1 != nil {
		fmt.Println(" err1====", err1)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	_, err2 := tx.Exec(clinicLaboratoryItemInsertSQL)
	if err2 != nil {
		fmt.Println(" err2====", err2)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}

	err3 := tx.Commit()
	if err3 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": laboratoryItemID})
}

//LaboratoryItemList 检验项目列表
func LaboratoryItemList(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	keyword := ctx.PostValue("name")
	status := ctx.PostValue("status")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if clinicID == "" {
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

	row := model.DB.QueryRowx("select id from clinic where id=$1 limit 1", clinicID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "查询失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)

	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "所在诊所不存在"})
		return
	}

	countSQL := `select count(cli.id) as total from clinic_laboratory_item cli
		left join laboratory_item li on cli.laboratory_item_id = li.id
		where cli.clinic_id=$1`
	selectSQL := `select li.id as laboratory_item_id,cli.id as clinic_laboratory_item_id,li.name,li.en_name,li.unit_name,li.is_special,li.instrument_code,
		li.data_type,lir.reference_sex,lir.stomach_status,lir.is_pregnancy,lir.reference_max,lir.reference_min,cli.status,cli.is_delivery
		from clinic_laboratory_item cli
		left join laboratory_item li on cli.laboratory_item_id = li.id
		left join laboratory_item_reference lir on lir.laboratory_item_id = li.id
		where cli.clinic_id=$1`

	if keyword != "" {
		countSQL += " and li.name ~'" + keyword + "'"
		selectSQL += " and li.name ~'" + keyword + "'"
	}
	if status != "" {
		countSQL += " and cli.status=" + status
		selectSQL += " and cli.status=" + status
	}

	fmt.Println("countSQL===", countSQL)
	fmt.Println("selectSQL===", selectSQL)
	total := model.DB.QueryRowx(countSQL, clinicID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL+" offset $2 limit $3", clinicID, offset, limit)
	results = FormatSQLRowsToMapArray(rows)

	laboratoryItems := FormatLaboratoryItem(results)

	ctx.JSON(iris.Map{"code": "200", "data": laboratoryItems, "page_info": pageInfo})
}

//LaboratoryItemUpdate 检验项目修改
func LaboratoryItemUpdate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	clinicLaboratoryItemID := ctx.PostValue("clinic_laboratory_item_id")
	laboratoryItemID := ctx.PostValue("laboratory_item_id")
	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	instrumentCode := ctx.PostValue("instrument_code")
	unitName := ctx.PostValue("unit_name")
	clinicalSignificance := ctx.PostValue("clinical_significance")
	dataType := ctx.PostValue("data_type")

	isSpecial := ctx.PostValue("is_special")
	referenceMax := ctx.PostValue("reference_max")
	referenceMin := ctx.PostValue("reference_min")
	items := ctx.PostValue("items")

	status := ctx.PostValue("status")
	isDelivery := ctx.PostValue("is_delivery")

	if clinicID == "" || name == "" || dataType == "" || isSpecial == "" || clinicLaboratoryItemID == "" || laboratoryItemID == "" || status == "" || isDelivery == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from clinic where id=$1 limit 1", clinicID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)
	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据错误"})
		return
	}

	crow := model.DB.QueryRowx("select id,clinic_id,laboratory_item_id from clinic_laboratory_item where id=$1 limit 1", clinicLaboratoryItemID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicLaboratoryItem := FormatSQLRowToMap(crow)
	_, rok := clinicLaboratoryItem["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所检验项数据错误"})
		return
	}
	slaboratoryItemID := strconv.FormatInt(clinicLaboratoryItem["laboratory_item_id"].(int64), 10)
	fmt.Println("laboratoryItemID====", slaboratoryItemID)

	if clinicID != strconv.FormatInt(clinicLaboratoryItem["clinic_id"].(int64), 10) {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据不匹配"})
		return
	}

	if slaboratoryItemID != laboratoryItemID {
		ctx.JSON(iris.Map{"code": "1", "msg": "检验项目数据id不匹配"})
		return
	}

	lrow := model.DB.QueryRowx("select id from laboratory_item where name=$1 and id!=$2 limit 1", name, laboratoryItemID)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	laboratoryItem := FormatSQLRowToMap(lrow)
	_, lok := laboratoryItem["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "1", "msg": "检验项目名称已存在"})
		return
	}

	laboratoryItemSets := []string{"name='" + name + "'", "data_type=" + dataType, "is_special=" + isSpecial}
	var itemReferenceSets []string
	var itemReferenceValues []string
	var clinicLaboratoryItemSets []string

	if enName != "" {
		laboratoryItemSets = append(laboratoryItemSets, "en_name='"+enName+"'")
	}
	if instrumentCode != "" {
		laboratoryItemSets = append(laboratoryItemSets, "instrument_code='"+instrumentCode+"'")
	}
	if unitName != "" {
		laboratoryItemSets = append(laboratoryItemSets, "unit_name='"+unitName+"'")
	}
	if clinicalSignificance != "" {
		laboratoryItemSets = append(laboratoryItemSets, "clinical_significance='"+clinicalSignificance+"'")
	}

	if status != "" {
		clinicLaboratoryItemSets = append(clinicLaboratoryItemSets, "status="+status)
	}
	if isDelivery != "" {
		clinicLaboratoryItemSets = append(clinicLaboratoryItemSets, "is_delivery="+isDelivery)
	}

	laboratoryItemSets = append(laboratoryItemSets, "updated_time=LOCALTIMESTAMP")
	laboratoryItemSetStr := strings.Join(laboratoryItemSets, ",")

	laboratoryItemUpdateSQL := "update laboratory_item set " + laboratoryItemSetStr + " where id=$1"
	fmt.Println("laboratoryItemInsertSQL==", laboratoryItemUpdateSQL)

	tx, err := model.DB.Begin()
	_, err = tx.Exec(laboratoryItemUpdateSQL, laboratoryItemID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}

	itemReferenceSets = append(itemReferenceSets, "laboratory_item_id")
	if isSpecial == "false" {
		itemReferenceValues = append(itemReferenceValues, laboratoryItemID)
		if referenceMax != "" {
			itemReferenceSets = append(itemReferenceSets, "reference_max")
			itemReferenceValues = append(itemReferenceValues, "'"+referenceMax+"'")
		}
		if referenceMin != "" {
			itemReferenceSets = append(itemReferenceSets, "reference_min")
			itemReferenceValues = append(itemReferenceValues, "'"+referenceMin+"'")
		}
	} else if isSpecial == "true" && items != "" {
		itemReferenceSets = append(itemReferenceSets, "reference_sex", "age_max", "age_min", "reference_max", "reference_min", "stomach_status", "is_pregnancy")
		var results []map[string]string
		reErr := json.Unmarshal([]byte(items), &results)
		if reErr != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": reErr.Error()})
			return
		}
		for _, v := range results {
			var s []string
			s = append(s, laboratoryItemID)
			referenceSex := v["reference_sex"]
			ageMax := v["age_max"]
			ageMin := v["age_min"]
			referenceMax := v["reference_max"]
			referenceMin := v["reference_min"]
			stomachStatus := v["stomach_status"]
			isPregnancy := v["is_pregnancy"]
			if referenceSex != "" {
				s = append(s, "'"+referenceSex+"'")
			} else {
				s = append(s, `null`)
			}
			if ageMax != "" {
				s = append(s, ageMax)
			} else {
				s = append(s, `null`)
			}
			if ageMin != "" {
				s = append(s, ageMin)
			} else {
				s = append(s, `null`)
			}
			if referenceMax != "" {
				s = append(s, "'"+referenceMax+"'")
			} else {
				s = append(s, `null`)
			}
			if referenceMin != "" {
				s = append(s, "'"+referenceMin+"'")
			} else {
				s = append(s, `null`)
			}
			if stomachStatus != "" {
				s = append(s, "'"+stomachStatus+"'")
			} else {
				s = append(s, `null`)
			}
			if isPregnancy != "" {
				s = append(s, isPregnancy)
			} else {
				s = append(s, `null`)
			}
			str := strings.Join(s, ",")
			str = "(" + str + ")"
			itemReferenceValues = append(itemReferenceValues, str)
		}
	} else {
		ctx.JSON(iris.Map{"code": "1", "msg": "参考值是否特殊数据格式错误"})
		return
	}

	itemReferenceSetStr := strings.Join(itemReferenceSets, ",")
	itemReferenceValueStr := strings.Join(itemReferenceValues, ",")

	clinicLaboratoryItemSets = append(clinicLaboratoryItemSets, "updated_time=LOCALTIMESTAMP")
	clinicLaboratoryItemSetStr := strings.Join(clinicLaboratoryItemSets, ",")

	itemReferenceInsertSQL := "insert into laboratory_item_reference (" + itemReferenceSetStr + ") values (" + itemReferenceValueStr + ")"
	if isSpecial == "true" {
		itemReferenceInsertSQL = "insert into laboratory_item_reference (" + itemReferenceSetStr + ") values " + itemReferenceValueStr
	}
	fmt.Println("itemReferenceInsertSQL==", itemReferenceInsertSQL)

	clinicLaboratoryItemUpdateSQL := "update clinic_laboratory_item set " + clinicLaboratoryItemSetStr + " where id=$1"
	fmt.Println("clinicLaboratoryItemUpdateSQL==", clinicLaboratoryItemUpdateSQL)

	_, errd := tx.Exec("delete from laboratory_item_reference where laboratory_item_id=$1", laboratoryItemID)
	if errd != nil {
		fmt.Println(" errd====", errd)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errd.Error()})
		return
	}
	_, err1 := tx.Exec(itemReferenceInsertSQL)
	if err1 != nil {
		fmt.Println(" err1====", err1)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	_, err2 := tx.Exec(clinicLaboratoryItemUpdateSQL, clinicLaboratoryItemID)
	if err2 != nil {
		fmt.Println(" err2====", err2)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}

	err3 := tx.Commit()
	if err3 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": laboratoryItemID})
}

//LaboratoryItemStatus 启用 停用
func LaboratoryItemStatus(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	clinicLaboratoryItemID := ctx.PostValue("clinic_laboratory_item_id")
	status := ctx.PostValue("status")

	if clinicID == "" || clinicLaboratoryItemID == "" || status == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from clinic where id=$1 limit 1", clinicID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)
	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据错误"})
		return
	}

	crow := model.DB.QueryRowx("select id,clinic_id from clinic_laboratory_item where id=$1 limit 1", clinicLaboratoryItemID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicLaboratoryItem := FormatSQLRowToMap(crow)
	_, rok := clinicLaboratoryItem["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据错误"})
		return
	}

	if clinicID != strconv.FormatInt(clinicLaboratoryItem["clinic_id"].(int64), 10) {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据不匹配"})
		return
	}
	_, err1 := model.DB.Exec("update clinic_laboratory_item set status=$1 where id=$2", status, clinicLaboratoryItemID)
	if err1 != nil {
		fmt.Println(" err1====", err1)
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//LaboratoryItemSearch 搜索检验项目
func LaboratoryItemSearch(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	keyword := ctx.PostValue("name")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from clinic where id=$1 limit 1", clinicID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "查询失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)
	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "所在诊所不存在"})
		return
	}

	selectSQL := `select li.id as laboratory_item_id,cli.id as clinic_laboratory_item_id,li.name,li.en_name,li.unit_name,li.is_special,li.instrument_code,
		li.data_type,lir.reference_sex,lir.stomach_status,lir.is_pregnancy,lir.reference_max,lir.reference_min,cli.status,cli.is_delivery
		from clinic_laboratory_item cli
		left join laboratory_item li on cli.laboratory_item_id = li.id
		left join laboratory_item_reference lir on lir.laboratory_item_id = li.id
		where cli.clinic_id=$1 and cli.status=true`

	if keyword != "" {
		selectSQL += " and li.name ~'" + keyword + "'"
	}
	fmt.Println("selectSQL===", selectSQL)

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL, clinicID)
	results = FormatSQLRowsToMapArray(rows)

	laboratoryItems := FormatLaboratoryItem(results)

	ctx.JSON(iris.Map{"code": "200", "data": laboratoryItems})
}

//LaboratoryItemDetail 检验项目详情
func LaboratoryItemDetail(ctx iris.Context) {
	clinicLaboratoryItemID := ctx.PostValue("clinic_laboratory_item_id")

	if clinicLaboratoryItemID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select li.id as laboratory_item_id,cli.id as clinic_laboratory_item_id,li.name,li.en_name,li.unit_name,li.is_special,cli.is_delivery,
		li.data_type,lir.reference_sex,lir.stomach_status,lir.is_pregnancy,lir.reference_max,lir.reference_min,cli.status,li.instrument_code
		from clinic_laboratory_item cli
		left join laboratory_item li on cli.laboratory_item_id = li.id
		left join laboratory_item_reference lir on lir.laboratory_item_id = li.id
		where cli.id=$1`

	fmt.Println("selectSQL===", selectSQL)

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL, clinicLaboratoryItemID)
	results = FormatSQLRowsToMapArray(rows)

	laboratoryItems := FormatLaboratoryItem(results)

	ctx.JSON(iris.Map{"code": "200", "data": laboratoryItems})
}
