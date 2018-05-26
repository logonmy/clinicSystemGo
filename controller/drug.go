package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

//DrugAdd 添加药品
func DrugAdd(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	drugID := ctx.PostValue("drug_Id")
	barcode := ctx.PostValue("barcode")
	name := ctx.PostValue("name")
	pyCode := ctx.PostValue("py_code")
	printName := ctx.PostValue("print_name")
	specification := ctx.PostValue("specification")
	manuFactoryName := ctx.PostValue("manu_factory_name")
	status := ctx.PostValue("status")
	licenseNo := ctx.PostValue("license_no")
	doseFormName := ctx.PostValue("dose_form_name")
	drugTypeID := ctx.PostValue("drug_type_id")

	miniDose := ctx.PostValue("mini_dose")
	miniUnitName := ctx.PostValue("mini_unit_name")
	doseCount := ctx.PostValue("dose_count")
	doseCountUnitName := ctx.PostValue("dose_count_unit_name")
	packingUnitName := ctx.PostValue("packing_unit_name")
	retPrice := ctx.PostValue("ret_price")
	buyPrice := ctx.PostValue("buy_price")
	isDiscount := ctx.PostValue("is_discount")
	isBulkSales := ctx.PostValue("is_bulk_sales")
	bulkSalesPrice := ctx.PostValue("bulk_sales_price")
	fetchAddress := ctx.PostValue("fetch_address")

	onceDose := ctx.PostValue("once_dose")
	onceDoseUnitName := ctx.PostValue("once_dose_unit_name")
	routeAdministrationName := ctx.PostValue("route_administration_name")
	frequencyName := ctx.PostValue("frequency_name")
	defaultRemark := ctx.PostValue("default_remark")

	dayWarning := ctx.PostValue("day_warning")
	stockWarning := ctx.PostValue("stock_warning")
	englishName := ctx.PostValue("english_name")
	syCode := ctx.PostValue("sy_code")

	if clinicID == "" || manuFactoryName == "" || name == "" || retPrice == "" || licenseNo == "" || specification == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	sets := []string{"name", "license_no", "py_code", "print_name", "specification",
		"manu_factory_name", "barcode", "dose_form_name", "drug_type_id", "once_dose", "once_dose_unit_name",
		"route_administration_name", "frequency_name", "default_remark", "english_name", "sy_code",
		"dose_count", "dose_count_unit_name", "packing_unit_name"}

	clinicDrugSets := []string{"clinic_id", "ret_price", "status", "mini_dose", "mini_unit_name", "buy_price",
		"is_discount", "is_bulk_sales", "bulk_sales_price", "fetch_address", "day_warning", "stock_warning", "drug_id"}

	setStr := strings.Join(sets, ",")
	clinicDrugSetStr := strings.Join(clinicDrugSets, ",")

	insertSQL := `insert into drug (` + setStr + `) values ($1, $2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,
		$16,$17,$18,$19) RETURNING id;`
	clinicDrugSQL := `insert into clinic_drug (` + clinicDrugSetStr + `) values ($1, $2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`

	tx, err := model.DB.Begin()

	if drugID == "" {
		err = tx.QueryRow(insertSQL,
			ToNullString(name),
			ToNullString(licenseNo),
			ToNullString(pyCode),
			ToNullString(printName),
			ToNullString(specification),
			ToNullString(manuFactoryName),
			ToNullString(barcode),
			ToNullString(doseFormName),
			ToNullInt64(drugTypeID),
			ToNullInt64(onceDose),
			ToNullString(onceDoseUnitName),
			ToNullString(routeAdministrationName),
			ToNullString(frequencyName),
			ToNullString(defaultRemark),
			ToNullString(englishName),
			ToNullString(syCode),
			ToNullInt64(doseCount),
			ToNullString(doseCountUnitName),
			ToNullString(packingUnitName)).Scan(&drugID)
		if err != nil {
			fmt.Println("err ===", err)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "1", "msg": err})
			return
		}
	}

	fmt.Println("drugID====", drugID)

	_, err1 := tx.Exec(clinicDrugSQL,
		ToNullInt64(clinicID),
		ToNullInt64(retPrice),
		NullBool(status),
		ToNullInt64(miniDose),
		ToNullString(miniUnitName),
		ToNullInt64(buyPrice),
		NullBool(isDiscount),
		NullBool(isBulkSales),
		ToNullInt64(bulkSalesPrice),
		ToNullString(fetchAddress),
		ToNullInt64(dayWarning),
		ToNullInt64(stockWarning),
		ToNullInt64(drugID))
	if err1 != nil {
		fmt.Println(" err1====", err1)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}
	err2 := tx.Commit()
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": drugID})
}

//DrugList 药品列表
func DrugList(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	drugType := ctx.PostValue("type")
	keyword := ctx.PostValue("keyword")
	drugTypeID := ctx.PostValue("drug_type_id")
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

	countSQL := `select count(cd.id) as total from clinic_drug cd
	left join drug d on cd.drug_id = d.id
	where cd.clinic_id=$1`
	selectSQL := `select cd.id as clinic_drug_id,d.name as drug_name,d.specification,d.packing_unit_name,d.type,
			cd.ret_price,d.py_code,cd.is_discount,d.default_remark,cd.status, d.once_dose_unit_name, 
			d.route_administration_name, d.frequency_name, cd.clinic_id,sum(ds.stock_amount) as stock_amount
			from clinic_drug cd
			left join drug d on cd.drug_id = d.id
			left join drug_stock ds on ds.clinic_drug_id = cd.id
			where cd.clinic_id=$1`

	if drugType != "" {
		countSQL += " and d.type=" + drugType
		selectSQL += " and d.type=" + drugType
	}
	if keyword != "" {
		countSQL += " and (d.barcode ~'" + keyword + "' or d.name ~'" + keyword + "')"
		selectSQL += " and (d.barcode ~'" + keyword + "' or d.name ~'" + keyword + "')"
	}
	if status != "" {
		countSQL += " and cd.status=" + status
		selectSQL += " and cd.status=" + status
	}
	if drugTypeID != "" {
		countSQL += " and d.drug_type_id=" + drugTypeID
		selectSQL += " and d.drug_type_id=" + drugTypeID
	}

	selectSQL += ` group by cd.id,d.name,d.specification,d.packing_unit_name,d.type,
	cd.ret_price,d.py_code,cd.is_discount,d.default_remark,cd.status, d.once_dose_unit_name, 
	d.route_administration_name, d.frequency_name, cd.clinic_id`

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

//DrugUpdate 修改药品
func DrugUpdate(ctx iris.Context) {
	clinicDrugID := ctx.PostValue("clinic_drug_id")
	drugType := ctx.PostValue("type")
	barcode := ctx.PostValue("barcode")
	name := ctx.PostValue("name")
	pyCode := ctx.PostValue("py_code")
	printName := ctx.PostValue("print_name")
	specification := ctx.PostValue("specification")
	manuFactoryName := ctx.PostValue("manu_factory_name")
	status := ctx.PostValue("status")
	licenseNo := ctx.PostValue("license_no")
	doseFormName := ctx.PostValue("dose_form_name")
	drugTypeID := ctx.PostValue("drug_type_id")

	miniDose := ctx.PostValue("mini_dose")
	miniUnitName := ctx.PostValue("mini_unit_name")
	doseCount := ctx.PostValue("dose_count")
	doseCountUnitName := ctx.PostValue("dose_count_unit_name")
	packingUnitName := ctx.PostValue("packing_unit_name")
	retPrice := ctx.PostValue("ret_price")
	buyPrice := ctx.PostValue("buy_price")
	isDiscount := ctx.PostValue("is_discount")
	isBulkSales := ctx.PostValue("is_bulk_sales")
	bulkSalesPrice := ctx.PostValue("bulk_sales_price")
	fetchAddress := ctx.PostValue("fetch_address")

	onceDose := ctx.PostValue("once_dose")
	onceDoseUnitName := ctx.PostValue("once_dose_unit_name")
	routeAdministrationName := ctx.PostValue("route_administration_name")
	frequencyName := ctx.PostValue("frequency_name")
	defaultRemark := ctx.PostValue("default_remark")

	dayWarning := ctx.PostValue("day_warning")
	stockWarning := ctx.PostValue("stock_warning")
	englishName := ctx.PostValue("english_name")
	syCode := ctx.PostValue("sy_code")

	if clinicDrugID == "" || barcode == "" || name == "" || retPrice == "" || packingUnitName == "" || drugType == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	var drugID int
	errs := model.DB.QueryRow("select drug_id from clinic_drug where id=$1", clinicDrugID).Scan(&drugID)
	if errs != nil {
		fmt.Println("errs ===", errs)
		ctx.JSON(iris.Map{"code": "1", "msg": "修改的药品不存在"})
		return
	}

	drow := model.DB.QueryRowx("select id from drug where barcode=$1 and id!=$2 limit 1", barcode, drugID)
	if drow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	ddrug := FormatSQLRowToMap(drow)
	_, dok := ddrug["id"]
	if dok {
		ctx.JSON(iris.Map{"code": "1", "msg": "药品已存在"})
		return
	}

	sets := []string{"name=" + "'" + name + "'", "type=" + drugType, "barcode=" + "'" + barcode + "'", "updated_time=LOCALTIMESTAMP"}
	clinicDrugSets := []string{"ret_price=" + retPrice, "updated_time=LOCALTIMESTAMP"}
	if pyCode != "" {
		sets = append(sets, "py_code="+"'"+pyCode+"'")
	}
	if printName != "" {
		sets = append(sets, "print_name="+"'"+printName+"'")
	}
	if specification != "" {
		sets = append(sets, "specification="+"'"+specification+"'")
	}
	if manuFactoryName != "" {
		sets = append(sets, "manu_factory_name='"+manuFactoryName+"'")
	}
	if licenseNo != "" {
		sets = append(sets, "license_no="+"'"+licenseNo+"'")
	}
	if doseFormName != "" {
		sets = append(sets, "dose_form_name='"+doseFormName+"'")
	}
	if drugTypeID != "" {
		sets = append(sets, "drug_type_id="+drugTypeID)
	}
	if onceDose != "" {
		sets = append(sets, "once_dose="+onceDose)
	}
	if onceDoseUnitName != "" {
		sets = append(sets, "once_dose_unit_name='"+onceDoseUnitName+"'")
	}
	if routeAdministrationName != "" {
		sets = append(sets, "route_administration_name='"+routeAdministrationName+"'")
	}
	if frequencyName != "" {
		sets = append(sets, "frequency_name='"+frequencyName+"'")
	}
	if defaultRemark != "" {
		sets = append(sets, "default_remark='"+defaultRemark+"'")
	}
	if englishName != "" {
		sets = append(sets, "english_name="+"'"+englishName+"'")
	}
	if syCode != "" {
		sets = append(sets, "sy_code="+"'"+syCode+"'")
	}

	if doseCount != "" {
		sets = append(sets, "dose_count="+doseCount)
	}
	if doseCountUnitName != "" {
		sets = append(sets, "dose_count_unit_name='"+doseCountUnitName+"'")
	}
	if packingUnitName != "" {
		sets = append(sets, "packing_unit_name='"+packingUnitName+"'")
	}

	if status != "" {
		clinicDrugSets = append(clinicDrugSets, "status="+status)
	}
	if miniDose != "" {
		clinicDrugSets = append(clinicDrugSets, "mini_dose="+miniDose)
	}
	if miniUnitName != "" {
		clinicDrugSets = append(clinicDrugSets, "mini_unit_name='"+miniUnitName+"'")
	}
	if buyPrice != "" {
		clinicDrugSets = append(clinicDrugSets, "buy_price="+buyPrice)
	}
	if isDiscount != "" {
		clinicDrugSets = append(clinicDrugSets, "is_discount="+isDiscount)
	}
	if isBulkSales != "" {
		clinicDrugSets = append(clinicDrugSets, "is_bulk_sales="+isBulkSales)
	}
	if bulkSalesPrice != "" {
		clinicDrugSets = append(clinicDrugSets, "bulk_sales_price="+bulkSalesPrice)
	}
	if fetchAddress != "" {
		clinicDrugSets = append(clinicDrugSets, "fetch_address="+"'"+fetchAddress+"'")
	}
	if dayWarning != "" {
		clinicDrugSets = append(clinicDrugSets, "day_warning="+dayWarning)
	}
	if stockWarning != "" {
		clinicDrugSets = append(clinicDrugSets, "stock_warning="+stockWarning)
	}

	setStr := strings.Join(sets, ",")
	clinicDrugStr := strings.Join(clinicDrugSets, ",")
	fmt.Println("setStr==", setStr)
	fmt.Println("clinicDrugStr==", clinicDrugStr)

	tx, err := model.DB.Begin()
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}

	updateSQL := "update drug set " + setStr + "where id=$1"
	stockUpdateSQL := "update clinic_drug set " + clinicDrugStr + "where id=$1"
	fmt.Println("updateSQL===", updateSQL)

	_, err1 := tx.Exec(updateSQL, drugID)
	if err1 != nil {
		fmt.Println("err ===", err1)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err1.Error()})
		return
	}

	_, err2 := tx.Exec(stockUpdateSQL, clinicDrugID)
	if err2 != nil {
		fmt.Println(" err1====", err1)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}
	err3 := tx.Commit()
	if err3 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": drugID})
}

//DrugSearch 搜索药品
func DrugSearch(ctx iris.Context) {
	keyword := ctx.PostValue("keyword")
	clinicID := ctx.PostValue("clinic_id")
	drugType := ctx.PostValue("type")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select cd.id as clinic_drug_id,d.id as drug_id,d.manu_factory_name,d.name as drug_name,d.packing_unit_name,
	cd.ret_price,cd.buy_price,cd.day_warning from drug d
	left join clinic_drug cd on cd.drug_id = d.id
	where cd.clinic_id=$1`

	if keyword != "" {
		selectSQL += " and (d.barcode ~'" + keyword + "' or d.name ~'" + keyword + "')"
	}

	if drugType != "" {
		selectSQL += " and d.type=" + drugType
	}

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL, clinicID)
	results = FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results})
}

//DrugDetail 药品详情
func DrugDetail(ctx iris.Context) {
	clinicDrugID := ctx.PostValue("clinic_drug_id")

	if clinicDrugID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	sql := `select d.name,
		d.specification,
		d.manu_factory_name,
		d.dose_form_name,
		d.print_name,
		d.license_no,
		dt.name as drug_type_name,d.drug_type_id,
		d.py_code,d.barcode,cd.status,
		cd.mini_dose,
		cd.mini_unit_name,
		d.dose_count,d.dose_count_unit_name,
		d.packing_unit_name,cd.ret_price,cd.buy_price,cd.is_discount,cd.is_bulk_sales,cd.bulk_sales_price,cd.fetch_address,
		d.once_dose,d.once_dose_unit_name,d.route_administration_name,
		d.frequency_name,d.default_remark,cd.day_warning,cd.stock_warning,d.english_name,d.sy_code
		from clinic_drug cd
		left join drug d on cd.drug_id = d.id
		left join drug_type dt on d.drug_type_id = dt.id
		where cd.id=$1`
	arows := model.DB.QueryRowx(sql, clinicDrugID)
	if arows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询结果不存在"})
		return
	}
	result := FormatSQLRowToMap(arows)
	ctx.JSON(iris.Map{"code": "200", "data": result})
}

//BatchSetting 批量设置药品
func BatchSetting(ctx iris.Context) {
	dayWarning := ctx.PostValue("day_warning")
	isDiscount := ctx.PostValue("is_discount")
	items := ctx.PostValue("items")
	if items == "" || (dayWarning == "" && isDiscount == "") {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	var results []map[string]string
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	sql := "update clinic_drug set"

	var sets []string
	asValues := []string{"id"}
	var values []string
	for _, v := range results {
		var s []string
		s = append(s, v["clinic_drug_id"])
		if dayWarning != "" {
			s = append(s, dayWarning)
		}
		if isDiscount != "" {
			s = append(s, isDiscount)
		}
		str := strings.Join(s, ",")
		str = "(" + str + ")"
		values = append(values, str)
	}
	valueStr := strings.Join(values, ",")

	if dayWarning != "" {
		sets = append(sets, " day_warning=tmp.day_warning")
		asValues = append(asValues, "day_warning")
	}

	if isDiscount != "" {
		sets = append(sets, " is_discount=tmp.isDiscount")
		asValues = append(asValues, "isDiscount")
	}
	setStr := strings.Join(sets, ",")
	asStr := strings.Join(asValues, ",")
	sql += setStr + " from (values " + valueStr + ") as tmp(" + asStr + ") where clinic_drug.id = tmp.id"

	_, erre := model.DB.Exec(sql)
	if erre != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": erre.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok"})
}
