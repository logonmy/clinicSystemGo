package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

//ClinicDrugCreate 添加药品
func ClinicDrugCreate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	drugClassID := ctx.PostValue("drug_class_id")
	name := ctx.PostValue("name")
	specification := ctx.PostValue("specification")
	manuFactoryName := ctx.PostValue("manu_factory_name")
	doseFormName := ctx.PostValue("dose_form_name")
	printName := ctx.PostValue("print_name")
	licenseNo := ctx.PostValue("license_no")
	drugType := ctx.PostValue("type")
	pyCode := ctx.PostValue("py_code")
	barCode := ctx.PostValue("barcode")
	status := ctx.PostValue("status")
	dosage := ctx.PostValue("dosage")
	dosageUnitName := ctx.PostValue("dosage_unit_name")
	preparationCount := ctx.PostValue("preparation_count")
	preparationCountUnitName := ctx.PostValue("preparation_count_unit_name")
	packingUnitName := ctx.PostValue("packing_unit_name")
	retPrice := ctx.PostValue("ret_price")
	buyPrice := ctx.PostValue("buy_price")
	miniDose := ctx.PostValue("mini_dose")
	isDiscount := ctx.PostValue("is_discount")
	isBulkSales := ctx.PostValue("is_bulk_sales")
	bulkSalesPrice := ctx.PostValue("bulk_sales_price")
	fetchAddress := ctx.PostValue("fetch_address")
	onceDose := ctx.PostValue("once_dose")
	onceDoseUnitName := ctx.PostValue("once_dose_unit_name")
	routeAdministrationName := ctx.PostValue("route_administration_name")
	frequencyName := ctx.PostValue("frequency_name")
	illustration := ctx.PostValue("illustration")
	dayWarning := ctx.PostValue("day_warning")
	stockWarning := ctx.PostValue("stock_warning")
	englishName := ctx.PostValue("english_name")
	syCode := ctx.PostValue("sy_code")
	countryFlag := ctx.PostValue("country_flag")
	selfSlag := ctx.PostValue("self_flag")
	drugDlag := ctx.PostValue("drug_flag")

	if clinicID == "" || name == "" || retPrice == "" || isBulkSales == "" || drugType == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if drugType == "0" && drugClassID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "药品分类码不能为空"})
		return
	}

	if isBulkSales == "true" && bulkSalesPrice == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "拆零价格必填"})
		return
	}

	if fetchAddress != "0" && fetchAddress != "-1" && fetchAddress != "2" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "取药地点必须为 0， 1， 2"})
		return
	}

	// 判断是否存在
	selectSQL := `select * from clinic_drug where clinic_id = $1 and name = $2 and specification = $3 and manu_factory_name = $4`
	drugRow := model.DB.QueryRowx(selectSQL, ToNullInt64(clinicID), ToNullString(name), ToNullString(specification), ToNullString(manuFactoryName))
	drugMap := FormatSQLRowToMap(drugRow)
	_, ok := drugMap["id"]
	if ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "拆零价格必填"})
		return
	}

	// 插入数据
	insertSQL := `insert into clinic_drug (
		clinic_id,
		type,
		drug_class_id,
		name,
		specification,
		manu_factory_name,
		dose_form_name,
		print_name,
		license_no,
		py_code,
		barcode,
		status,
		dosage,
		dosage_unit_name,
		preparation_count,
		preparation_count_unit_name,
		packing_unit_name,
		ret_price,
		buy_price,
		mini_dose,
		is_discount,
		is_bulk_sales,
		bulk_sales_price,
		fetch_address,
		once_dose,
		once_dose_unit_name,
		route_administration_name,
		frequency_name,
		illustration,
		day_warning,
		stock_warning,
		english_name,
		sy_code,
		country_flag,
		self_flag,
		drug_flag
	)
	values (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10,
		$11,
		$12,
		$13,
		$14,
		$15,
		$16,
		$17,
		$18,
		$19,
		$20,
		$21,
		$22,
		$23,
		$24,
		$25,
		$26,
		$27,
		$28,
		$29,
		$30,
		$31,
		$32,
		$33,
		$34,
		$35,
		$36
	)`

	fmt.Println("insertSQL ====", insertSQL)
	_, err := model.DB.Exec(insertSQL,
		ToNullInt64(clinicID),
		ToNullInt64(drugType),
		ToNullInt64(drugClassID),
		ToNullString(name),
		ToNullString(specification),
		ToNullString(manuFactoryName),
		ToNullString(doseFormName),
		ToNullString(printName),
		ToNullString(licenseNo),
		ToNullString(pyCode),
		ToNullString(barCode),
		ToNullBool(status),
		ToNullInt64(dosage),
		ToNullString(dosageUnitName),
		ToNullInt64(preparationCount),
		ToNullString(preparationCountUnitName),
		ToNullString(packingUnitName),
		ToNullInt64(retPrice),
		ToNullInt64(buyPrice),
		ToNullInt64(miniDose),
		ToNullBool(isDiscount),
		ToNullBool(isBulkSales),
		ToNullInt64(bulkSalesPrice),
		ToNullInt64(fetchAddress),
		ToNullInt64(onceDose),
		ToNullString(onceDoseUnitName),
		ToNullString(routeAdministrationName),
		ToNullString(frequencyName),
		ToNullString(illustration),
		ToNullInt64(dayWarning),
		ToNullInt64(stockWarning),
		ToNullString(englishName),
		ToNullString(syCode),
		ToNullBool(countryFlag),
		ToNullBool(selfSlag),
		ToNullBool(drugDlag),
	)

	if err != nil {
		fmt.Println("err === ", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": "-1"})
}

//ClinicDrugList 药品列表
func ClinicDrugList(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	drugType := ctx.PostValue("type")
	drugClassID := ctx.PostValue("drug_class_id")
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

	countSQL := `select count(*) as total from clinic_drug cd where clinic_id=:clinic_id`
	selectSQL := `select 
		cd.id as clinic_drug_id,
		cd.type,
		cd.manu_factory_name,
		cd.name as drug_name,
		cd.specification,
		cd.packing_unit_name,
		cd.ret_price,
		cd.buy_price,
		cd.py_code,
		cd.is_discount,
		cd.illustration,
		cd.status, 
		cd.once_dose,
		cd.once_dose_unit_name, 
		cd.route_administration_name, 
		cd.frequency_name,
		cd.fetch_address,
		cd.clinic_id,
		sum(ds.stock_amount) as stock_amount
		from clinic_drug cd
		left join drug_stock ds on ds.clinic_drug_id = cd.id
		where cd.clinic_id = :clinic_id`

	if drugType != "" {
		countSQL += " and cd.type = :type"
		selectSQL += " and cd.type= :type"
	}

	if status != "" {
		countSQL += " and cd.status = :status"
		selectSQL += " and cd.status = :status"
	}

	if keyword != "" {
		keyword = strings.ToUpper(keyword)
		countSQL += " and (cd.py_code ~:keyword or cd.name ~:keyword)"
		selectSQL += " and (cd.py_code ~:keyword or cd.name ~:keyword)"
	}

	if drugClassID != "" {
		countSQL += " and (cd.drug_class_id = :drug_class_id or cd.drug_class_id = :drug_class_id)"
		selectSQL += " and (cd.drug_class_id = :drug_class_id or cd.drug_class_id = :drug_class_id)"
	}

	selectSQL += ` group by 
		cd.id,
		cd.type,
		cd.manu_factory_name,
		cd.name,
		cd.specification,
		cd.packing_unit_name,
		cd.ret_price,
		cd.buy_price,
		cd.py_code,
		cd.is_discount,
		cd.illustration,
		cd.status, 
		cd.once_dose,
		cd.once_dose_unit_name, 
		cd.route_administration_name, 
		cd.frequency_name, 
		cd.fetch_address,
		cd.clinic_id `

	var queryOption = map[string]interface{}{
		"clinic_id":     ToNullInt64(clinicID),
		"type":          ToNullString(drugType),
		"status":        ToNullBool(status),
		"keyword":       ToNullString(keyword),
		"offset":        ToNullInt64(offset),
		"limit":         ToNullInt64(limit),
		"drug_class_id": ToNullInt64(drugClassID),
	}
	total, err := model.DB.NamedQuery(countSQL, queryOption)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, err := model.DB.NamedQuery(selectSQL+" offset :offset limit :limit", queryOption)
	if err != nil {
		fmt.Println("err ====", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	results = FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

//ClinicDrugUpdate 修改药品
func ClinicDrugUpdate(ctx iris.Context) {
	clinicDrugID := ctx.PostValue("clinic_drug_id")
	name := ctx.PostValue("name")
	drugClassID := ctx.PostValue("drug_class_id")
	specification := ctx.PostValue("specification")
	manuFactoryName := ctx.PostValue("manu_factory_name")
	doseFormName := ctx.PostValue("dose_form_name")
	printName := ctx.PostValue("print_name")
	licenseNo := ctx.PostValue("license_no")
	pyCode := ctx.PostValue("py_code")
	barCode := ctx.PostValue("barcode")
	status := ctx.PostValue("status")
	dosage := ctx.PostValue("dosage")
	dosageUnitName := ctx.PostValue("dosage_unit_name")
	preparationCount := ctx.PostValue("preparation_count")
	preparationCountUnitName := ctx.PostValue("preparation_count_unit_name")
	packingUnitName := ctx.PostValue("packing_unit_name")
	retPrice := ctx.PostValue("ret_price")
	buyPrice := ctx.PostValue("buy_price")
	miniDose := ctx.PostValue("mini_dose")
	isDiscount := ctx.PostValue("is_discount")
	isBulkSales := ctx.PostValue("is_bulk_sales")
	bulkSalesPrice := ctx.PostValue("bulk_sales_price")
	fetchAddress := ctx.PostValue("fetch_address")
	onceDose := ctx.PostValue("once_dose")
	onceDoseUnitName := ctx.PostValue("once_dose_unit_name")
	routeAdministrationName := ctx.PostValue("route_administration_name")
	frequencyName := ctx.PostValue("frequency_name")
	illustration := ctx.PostValue("illustration")
	dayWarning := ctx.PostValue("day_warning")
	stockWarning := ctx.PostValue("stock_warning")
	englishName := ctx.PostValue("english_name")
	syCode := ctx.PostValue("sy_code")
	countryFlag := ctx.PostValue("country_flag")
	selfSlag := ctx.PostValue("self_flag")
	drugDlag := ctx.PostValue("drug_flag")

	if clinicDrugID == "" || name == "" || retPrice == "" || isBulkSales == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if isBulkSales == "true" && bulkSalesPrice == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "拆零价格必填"})
		return
	}

	if fetchAddress != "0" && fetchAddress != "-1" && fetchAddress != "2" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "取药地点必须为 0， 1， 2"})
		return
	}

	crow := model.DB.QueryRowx("select id,clinic_id from clinic_drug where id=$1 limit 1", clinicDrugID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	clinicDrug := FormatSQLRowToMap(crow)
	_, rok := clinicDrug["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊所材料项目数据错误"})
		return
	}
	clinicID := clinicDrug["clinic_id"]

	lrow := model.DB.QueryRowx("select id from clinic_drug where name=$1 and id!=$2 and manu_factory_name=$3 and clinic_id=$4 and specification=$5 limit 1", name, clinicDrugID, manuFactoryName, clinicID, specification)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	clinicDrugu := FormatSQLRowToMap(lrow)
	_, lok := clinicDrugu["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "改药品已存在"})
		return
	}

	// 修改数据
	updateSQL := `udpate clinic_drug set
		name = $2,
		specification = $3,
		manu_factory_name = $4,
		dose_form_name = $5,
		print_name = $6,
		license_no = $7,
		drug_class_id = $8,
		py_code = $9,
		barcode = $10,
		status = $11,
		dosage = $12,
		dosage_unit_name = $13,
		preparation_count = $14,
		preparation_count_unit_name = $15,
		packing_unit_name = $16,
		ret_price = $17,
		buy_price = $18,
		mini_dose = $19,
		is_discount = $20,
		is_bulk_sales = $21,
		bulk_sales_price = $22,
		fetch_address = $23,
		once_dose = $24,
		once_dose_unit_name = $25,
		route_administration_name = $26,
		frequency_name = $27,
		illustration = $28,
		day_warning = $29,
		stock_warning = $30,
		english_name = $31,
		sy_code = $32,
		country_flag = $33,
		self_flag = $34,
		drug_flag = $35
		where id = $1`
	model.DB.Exec(updateSQL,
		ToNullInt64(clinicDrugID),
		ToNullString(name),
		ToNullString(specification),
		ToNullString(manuFactoryName),
		ToNullString(doseFormName),
		ToNullString(printName),
		ToNullString(licenseNo),
		ToNullInt64(drugClassID),
		ToNullString(pyCode),
		ToNullString(barCode),
		ToNullBool(status),
		ToNullInt64(dosage),
		ToNullString(dosageUnitName),
		ToNullInt64(preparationCount),
		ToNullString(preparationCountUnitName),
		ToNullString(packingUnitName),
		ToNullInt64(retPrice),
		ToNullInt64(buyPrice),
		ToNullInt64(miniDose),
		ToNullBool(isDiscount),
		ToNullBool(isBulkSales),
		ToNullInt64(bulkSalesPrice),
		ToNullInt64(fetchAddress),
		ToNullInt64(onceDose),
		ToNullString(onceDoseUnitName),
		ToNullString(routeAdministrationName),
		ToNullString(frequencyName),
		ToNullString(illustration),
		ToNullInt64(dayWarning),
		ToNullInt64(stockWarning),
		ToNullString(englishName),
		ToNullString(syCode),
		ToNullBool(countryFlag),
		ToNullBool(selfSlag),
		ToNullBool(drugDlag),
	)
	ctx.JSON(iris.Map{"code": "200", "data": clinicDrugID})
}

//ClinicDrugDetail 药品详情
func ClinicDrugDetail(ctx iris.Context) {
	clinicDrugID := ctx.PostValue("clinic_drug_id")
	if clinicDrugID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	sql := `select * clinic_drug from where id = $1`
	arow := model.DB.QueryRowx(sql, clinicDrugID)
	if arow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询结果不存在"})
		return
	}
	result := FormatSQLRowToMap(arow)
	ctx.JSON(iris.Map{"code": "200", "data": result})
}

//ClinicDrugBatchSetting 批量设置药品
func ClinicDrugBatchSetting(ctx iris.Context) {
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
