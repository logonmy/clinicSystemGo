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
		ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)
	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊所数据错误"})
		return
	}

	clrow := model.DB.QueryRowx("select id from clinic_laboratory where clinic_id=$1 and name=$2 limit 1", clinicID, name)
	if clrow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
		return
	}
	clinicLaboratory := FormatSQLRowToMap(clrow)
	_, lok := clinicLaboratory["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "检查医嘱名称在该诊所已存在"})
		return
	}

	clinicLaboratorySets := []string{
		"clinic_id",
		"name",
		"en_name",
		"py_code",
		"idc_code",
		"unit_name",
		"time_report",
		"clinical_significance",
		"laboratory_sample",
		"cuvette_color_name",
		"status",
		"remark",
		"merge_flag",
		"cost",
		"price",
		"is_delivery",
		"is_discount",
	}

	clinicLaboratorySetStr := strings.Join(clinicLaboratorySets, ",")
	clinicLaboratoryInsertSQL := "insert into clinic_laboratory (" + clinicLaboratorySetStr + ") values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17)"

	_, err2 := model.DB.Exec(clinicLaboratoryInsertSQL,
		ToNullInt64(clinicID),
		ToNullString(name),
		ToNullString(enName),
		ToNullString(pyCode),
		ToNullString(idcCode),
		ToNullString(unitName),
		ToNullString(timeReport),
		ToNullString(clinicalSignificance),
		ToNullString(laboratorySample),
		ToNullString(cuvetteColorName),
		ToNullBool(status),
		ToNullString(remark),
		ToNullInt64(mergeFlag),
		ToNullInt64(cost),
		ToNullInt64(price),
		ToNullBool(isDelivery),
		ToNullBool(isDiscount),
	)
	if err2 != nil {
		fmt.Println(" err2====", err2)
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
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
		ctx.JSON(iris.Map{"code": "-1", "msg": "关联项目失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)

	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "所选诊所检验医嘱不存在"})
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
			ctx.JSON(iris.Map{"code": "-1", "msg": "关联项目失败"})
			return
		}
		clinicLaboratoryItem := FormatSQLRowToMap(row)
		_, ok := clinicLaboratoryItem["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "关联的" + laboratoryItemName + "检验项目不存在"})
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
		ctx.JSON(iris.Map{"code": "-1", "msg": errb})
		return
	}
	_, errd := tx.Exec("delete from clinic_laboratory_association where clinic_laboratory_id=$1", clinicLaboratoryID)
	if errd != nil {
		fmt.Println("errd ===", errd)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errd.Error()})
		return
	}

	insertSQL := "insert into clinic_laboratory_association (" + setStr + ") values " + valueStr
	fmt.Println("insertSQL===", insertSQL)

	_, err := tx.Exec(insertSQL)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": "请检验是否漏填"})
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

	selectSQL := `select cls.clinic_laboratory_item_id,cli.name,cli.en_name,cli.unit_name,cli.is_special,cls.default_result,
	cli.data_type,clir.reference_sex,clir.stomach_status,clir.is_pregnancy,clir.reference_max,clir.reference_min,cli.status
	from clinic_laboratory_association cls
	left join clinic_laboratory_item cli on cls.clinic_laboratory_item_id = cli.id
	left join clinic_laboratory_item_reference clir on clir.clinic_laboratory_item_id = cli.id
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
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)

	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "所在诊所不存在"})
		return
	}

	countSQL := `select count(id) as total from clinic_laboratory where clinic_id=$1`
	selectSQL := `select id as clinic_laboratory_id,name as laboratory_name,unit_name,price,py_code,is_discount,discount_price,
		remark,status from clinic_laboratory where clinic_id=$1`

	if keyword != "" {
		countSQL += " and name ~'" + keyword + "'"
		selectSQL += " and name ~'" + keyword + "'"
	}
	if status != "" {
		countSQL += " and status=" + status
		selectSQL += " and status=" + status
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
	sql := `select id as clinic_laboratory_id,name,en_name,unit_name,py_code,idc_code,remark,
		time_report,clinical_significance,laboratory_sample,cuvette_color_name,discount_price,
		cost,is_discount,status,merge_flag,price,is_delivery
		from clinic_laboratory where id=$1`
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

	crow := model.DB.QueryRowx("select id,clinic_id from clinic_laboratory where id=$1 limit 1", clinicLaboratoryID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	clinicLaboratory := FormatSQLRowToMap(crow)
	_, rok := clinicLaboratory["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊所检验医嘱数据错误"})
		return
	}
	clinicID := clinicLaboratory["clinic_id"]

	lrow := model.DB.QueryRowx("select id from clinic_laboratory where name=$1 and id!=$2 and clinic_id=$3 limit 1", name, clinicLaboratoryID, clinicID)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	clinicLaboratoryItem := FormatSQLRowToMap(lrow)
	_, lok := clinicLaboratoryItem["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "检验医嘱名称已存在"})
		return
	}

	clinicLaboratoryUpdateSQL := `update clinic_laboratory set 
		name=$1,
		en_name=$2,
		py_code=$3,
		idc_code=$4,
		unit_name=$5,
		time_report=$6,
		clinical_significance=$7,
		laboratory_sample=$8,
		cuvette_color_name=$9,
		status=$10,
		remark=$11,
		merge_flag=$12,
		cost=$13,
		price=$14,
		is_delivery=$15,
		is_discount=$16
		where id=$17`

	_, err := model.DB.Exec(clinicLaboratoryUpdateSQL,
		ToNullString(name),
		ToNullString(enName),
		ToNullString(pyCode),
		ToNullString(idcCode),
		ToNullString(unitName),
		ToNullString(timeReport),
		ToNullString(clinicalSignificance),
		ToNullString(laboratorySample),
		ToNullString(cuvetteColorName),
		ToNullBool(status),
		ToNullString(remark),
		ToNullInt64(mergeFlag),
		ToNullInt64(cost),
		ToNullInt64(price),
		ToNullBool(isDelivery),
		ToNullBool(isDiscount),
		ToNullInt64(clinicLaboratoryID),
	)
	if err != nil {
		fmt.Println(" err====", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
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
		ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)
	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊所数据错误"})
		return
	}

	clirow := model.DB.QueryRowx("select id from clinic_laboratory_item where clinic_id=$1 and name=$2 limit 1", clinicID, name)
	if clirow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
		return
	}
	clinicLaboratoryItem := FormatSQLRowToMap(clirow)
	_, lok := clinicLaboratoryItem["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "检验项目名称在该诊所已存在"})
		return
	}

	itemReferenceSets := []string{
		"clinic_laboratory_item_id",
		"reference_max",
		"reference_min",
		"age_max",
		"age_min",
		"reference_sex",
		"stomach_status",
		"is_pregnancy"}
	clinicLaboratoryItemSets := []string{
		"clinic_id",
		"name",
		"en_name",
		"instrument_code",
		"unit_name",
		"clinical_significance",
		"status",
		"is_delivery",
		"data_type",
		"is_special"}

	tx, err := model.DB.Begin()

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	clinicLaboratoryItemSetStr := strings.Join(clinicLaboratoryItemSets, ",")
	itemReferenceSetStr := strings.Join(itemReferenceSets, ",")

	clinicLaboratoryItemInsertSQL := "insert into clinic_laboratory_item (" + clinicLaboratoryItemSetStr + ") values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id;"
	itemReferenceInsertSQL := "insert into clinic_laboratory_item_reference (" + itemReferenceSetStr + ") values ($1,$2,$3,$4,$5,$6,$7,$8)"

	var clinicLaboratoryItemID string
	err1 := tx.QueryRow(clinicLaboratoryItemInsertSQL,
		ToNullInt64(clinicID),
		ToNullString(name),
		ToNullString(enName),
		ToNullString(instrumentCode),
		ToNullString(unitName),
		ToNullString(clinicalSignificance),
		ToNullBool(status),
		ToNullBool(isDelivery),
		ToNullInt64(dataType),
		ToNullBool(isSpecial),
	).Scan(&clinicLaboratoryItemID)
	if err1 != nil {
		fmt.Println(" err1====", err1)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	if isSpecial == "false" {
		_, err2 := tx.Exec(itemReferenceInsertSQL,
			ToNullInt64(clinicLaboratoryItemID),
			ToNullString(referenceMax),
			ToNullString(referenceMin),
			ToNullInt64(""),
			ToNullInt64(""),
			"通用",
			ToNullBool("false"),
			ToNullBool("false"),
		)
		if err2 != nil {
			fmt.Println(" err2====", err2)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
			return
		}
	} else if isSpecial == "true" && items != "" {
		var results []map[string]string
		reErr := json.Unmarshal([]byte(items), &results)
		if reErr != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": reErr.Error()})
			return
		}
		for _, v := range results {
			referenceSex := v["reference_sex"]
			ageMax := v["age_max"]
			ageMin := v["age_min"]
			sreferenceMax := v["reference_max"]
			sreferenceMin := v["reference_min"]
			stomachStatus := v["stomach_status"]
			isPregnancy := v["is_pregnancy"]
			_, err3 := tx.Exec(itemReferenceInsertSQL,
				ToNullInt64(clinicLaboratoryItemID),
				ToNullString(sreferenceMax),
				ToNullString(sreferenceMin),
				ToNullInt64(ageMax),
				ToNullInt64(ageMin),
				ToNullString(referenceSex),
				ToNullBool(stomachStatus),
				ToNullBool(isPregnancy),
			)
			if err3 != nil {
				fmt.Println(" err3====", err3)
				tx.Rollback()
				ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
				return
			}
		}
	} else {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参考值是否特殊数据格式错误"})
		return
	}

	err4 := tx.Commit()
	if err4 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err4.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": clinicLaboratoryItemID})
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
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)

	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "所在诊所不存在"})
		return
	}

	countSQL := `select count(id) as total from clinic_laboratory_item where clinic_id=:clinic_id`
	selectSQL := `select cli.id as clinic_laboratory_item_id,cli.name,cli.en_name,cli.unit_name,cli.is_special,cli.instrument_code,
		cli.data_type,clir.reference_sex,clir.stomach_status,clir.is_pregnancy,clir.reference_max,clir.reference_min,cli.status,cli.is_delivery
		from clinic_laboratory_item cli
		left join clinic_laboratory_item_reference clir on clir.clinic_laboratory_item_id = cli.id
		where cli.clinic_id=:clinic_id`

	if keyword != "" {
		countSQL += " and cli.name ~:keyword"
		selectSQL += " and cli.name ~:keyword"
	}
	if status != "" {
		countSQL += " and cli.status=:status"
		selectSQL += " and cli.status=:status"
	}

	var queryOption = map[string]interface{}{
		"clinic_id": ToNullInt64(clinicID),
		"keyword":   ToNullString(keyword),
		"status":    ToNullBool(status),
		"offset":    ToNullInt64(offset),
		"limit":     ToNullInt64(limit),
	}

	fmt.Println("countSQL===", countSQL)
	fmt.Println("selectSQL===", selectSQL)
	total, err := model.DB.NamedQuery(countSQL, queryOption)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.NamedQuery(selectSQL+" offset :offset limit :limit", queryOption)
	results = FormatSQLRowsToMapArray(rows)

	laboratoryItems := FormatLaboratoryItem(results)

	ctx.JSON(iris.Map{"code": "200", "data": laboratoryItems, "page_info": pageInfo})
}

//LaboratoryItemUpdate 检验项目修改
func LaboratoryItemUpdate(ctx iris.Context) {
	clinicLaboratoryItemID := ctx.PostValue("clinic_laboratory_item_id")
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

	if name == "" || dataType == "" || isSpecial == "" || clinicLaboratoryItemID == "" || status == "" || isDelivery == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	crow := model.DB.QueryRowx("select id,clinic_id from clinic_laboratory_item where id=$1 limit 1", clinicLaboratoryItemID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	clinicLaboratoryItem := FormatSQLRowToMap(crow)
	_, rok := clinicLaboratoryItem["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊所检验项数据错误"})
		return
	}
	clinicID := clinicLaboratoryItem["clinic_id"]

	lrow := model.DB.QueryRowx("select id from clinic_laboratory_item where name=$1 and id!=$2 and clinic_id=$3 limit 1", name, clinicLaboratoryItemID, clinicID)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	clinicLaboratoryItemu := FormatSQLRowToMap(lrow)
	_, lok := clinicLaboratoryItemu["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "检验项目名称已存在"})
		return
	}

	itemReferenceSets := []string{
		"clinic_laboratory_item_id",
		"reference_max",
		"reference_min",
		"age_max",
		"age_min",
		"reference_sex",
		"stomach_status",
		"is_pregnancy"}

	itemReferenceSetStr := strings.Join(itemReferenceSets, ",")
	itemReferenceInsertSQL := "insert into clinic_laboratory_item_reference (" + itemReferenceSetStr + ") values ($1,$2,$3,$4,$5,$6,$7,$8)"

	tx, err := model.DB.Begin()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	clinicLaboratoryItemUpdateSQL := `update clinic_laboratory_item set 
		name=$1,
		en_name=$2,
		instrument_code=$3,
		unit_name=$4,
		clinical_significance=$5,
		status=$6,
		is_delivery=$7,
		data_type=$8,
		is_special=$9 
		where id=$10`

	_, err1 := tx.Exec(clinicLaboratoryItemUpdateSQL,
		ToNullString(name),
		ToNullString(enName),
		ToNullString(instrumentCode),
		ToNullString(unitName),
		ToNullString(clinicalSignificance),
		ToNullBool(status),
		ToNullBool(isDelivery),
		ToNullInt64(dataType),
		ToNullBool(isSpecial),
		ToNullInt64(clinicLaboratoryItemID),
	)
	if err1 != nil {
		fmt.Println(" err1====", err1)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	_, errd := tx.Exec("delete from clinic_laboratory_item_reference where clinic_laboratory_item_id=$1", clinicLaboratoryItemID)
	if errd != nil {
		fmt.Println(" errd====", errd)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errd.Error()})
		return
	}

	if isSpecial == "false" {
		_, err2 := tx.Exec(itemReferenceInsertSQL,
			ToNullInt64(clinicLaboratoryItemID),
			ToNullString(referenceMax),
			ToNullString(referenceMin),
			ToNullInt64(""),
			ToNullInt64(""),
			"通用",
			ToNullBool("false"),
			ToNullBool("false"),
		)
		if err2 != nil {
			fmt.Println(" err2====", err2)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
			return
		}
	} else if isSpecial == "true" && items != "" {
		var results []map[string]string
		reErr := json.Unmarshal([]byte(items), &results)
		if reErr != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": reErr.Error()})
			return
		}
		for _, v := range results {
			referenceSex := v["reference_sex"]
			ageMax := v["age_max"]
			ageMin := v["age_min"]
			sreferenceMax := v["reference_max"]
			sreferenceMin := v["reference_min"]
			stomachStatus := v["stomach_status"]
			isPregnancy := v["is_pregnancy"]
			_, err3 := tx.Exec(itemReferenceInsertSQL,
				ToNullInt64(clinicLaboratoryItemID),
				ToNullString(sreferenceMax),
				ToNullString(sreferenceMin),
				ToNullInt64(ageMax),
				ToNullInt64(ageMin),
				ToNullString(referenceSex),
				ToNullBool(stomachStatus),
				ToNullBool(isPregnancy),
			)
			if err3 != nil {
				fmt.Println(" err3====", err3)
				tx.Rollback()
				ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
				return
			}
		}
	} else {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参考值是否特殊数据格式错误"})
		return
	}

	err3 := tx.Commit()
	if err3 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": clinicLaboratoryItemID})
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
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)
	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊所数据错误"})
		return
	}

	crow := model.DB.QueryRowx("select id,clinic_id from clinic_laboratory_item where id=$1 limit 1", clinicLaboratoryItemID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	clinicLaboratoryItem := FormatSQLRowToMap(crow)
	_, rok := clinicLaboratoryItem["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊所数据错误"})
		return
	}

	if clinicID != strconv.FormatInt(clinicLaboratoryItem["clinic_id"].(int64), 10) {
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊所数据不匹配"})
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
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)
	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "所在诊所不存在"})
		return
	}

	selectSQL := `select cli.id as clinic_laboratory_item_id,cli.name,cli.en_name,cli.unit_name,cli.is_special,cli.instrument_code,
		cli.data_type,clir.reference_sex,clir.stomach_status,clir.is_pregnancy,clir.reference_max,clir.reference_min,cli.status,cli.is_delivery
		from clinic_laboratory_item cli
		left join clinic_laboratory_item_reference clir on clir.clinic_laboratory_item_id = cli.id
		where cli.clinic_id=$1 and cli.status=true`

	if keyword != "" {
		selectSQL += " and cli.name ~'" + keyword + "'"
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

	selectSQL := `select cli.id as clinic_laboratory_item_id,cli.name,cli.en_name,cli.unit_name,cli.is_special,cli.is_delivery,
		cli.data_type,clir.reference_sex,clir.stomach_status,clir.is_pregnancy,clir.reference_max,clir.reference_min,cli.status,cli.instrument_code
		from clinic_laboratory_item cli
		left join clinic_laboratory_item_reference clir on clir.clinic_laboratory_item_id = cli.id
		where cli.id=$1`

	fmt.Println("selectSQL===", selectSQL)

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL, clinicLaboratoryItemID)
	results = FormatSQLRowsToMapArray(rows)

	laboratoryItems := FormatLaboratoryItem(results)

	ctx.JSON(iris.Map{"code": "200", "data": laboratoryItems})
}
