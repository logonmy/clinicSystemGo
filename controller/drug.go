package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

//DrugAdd 添加药品
func DrugAdd(ctx iris.Context) {
	barcode := ctx.PostValue("barcode")
	name := ctx.PostValue("name")
	pyCode := ctx.PostValue("py_code")
	printName := ctx.PostValue("print_name")
	specification := ctx.PostValue("specification")
	manuFactory := ctx.PostValue("manu_factory")
	status := ctx.PostValue("status")
	licenseNo := ctx.PostValue("license_no")
	doseFormID := ctx.PostValue("dose_form_id")
	drugClassID := ctx.PostValue("drug_class_id")

	miniDose := ctx.PostValue("mini_dose")
	miniUnitID := ctx.PostValue("mini_unit_id")
	doseCount := ctx.PostValue("dose_count")
	doseCountID := ctx.PostValue("dose_count_id")
	packingUnitID := ctx.PostValue("packing_unit_id")
	retPrice := ctx.PostValue("ret_price")
	buyPrice := ctx.PostValue("buy_price")
	isDiscount := ctx.PostValue("is_discount")
	isBulkSales := ctx.PostValue("is_bulk_sales")
	bulkSalesPrice := ctx.PostValue("bulk_sales_price")
	fetchAddress := ctx.PostValue("fetch_address")

	onceDose := ctx.PostValue("once_dose")
	onceDoseUnitID := ctx.PostValue("once_dose_unit_id")
	routeAdministrationID := ctx.PostValue("route_administration_id")
	frequencyID := ctx.PostValue("frequency_id")
	defaultRemark := ctx.PostValue("default_remark")

	effDay := ctx.PostValue("eff_day")
	stockWarning := ctx.PostValue("stock_warning")
	englishName := ctx.PostValue("english_name")
	syCode := ctx.PostValue("sy_code")

	if barcode == "" || name == "" || retPrice == "" || packingUnitID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from drug where name=$1 and barcode=$2 limit 1", name, barcode)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
		return
	}
	drug := FormatSQLRowToMap(row)
	_, ok := drug["id"]
	if ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "药品已存在"})
		return
	}

	sets := []string{"name", "barcode", "ret_price", "packing_unit_id"}
	values := []string{"'" + name + "'", "'" + barcode + "'", retPrice, packingUnitID}
	if pyCode != "" {
		sets = append(sets, "py_code")
		values = append(values, "'"+pyCode+"'")
	}
	if printName != "" {
		sets = append(sets, "print_name")
		values = append(values, "'"+printName+"'")
	}
	if specification != "" {
		sets = append(sets, "specification")
		values = append(values, "'"+specification+"'")
	}
	if manuFactory != "" {
		sets = append(sets, "manu_factory")
		values = append(values, "'"+manuFactory+"'")
	}
	if status != "" {
		sets = append(sets, "status")
		values = append(values, status)
	}
	if licenseNo != "" {
		sets = append(sets, "license_no")
		values = append(values, "'"+licenseNo+"'")
	}
	if doseFormID != "" {
		sets = append(sets, "dose_form_id")
		values = append(values, doseFormID)
	}
	if drugClassID != "" {
		sets = append(sets, "drug_class_id")
		values = append(values, drugClassID)
	}
	if miniDose != "" {
		sets = append(sets, "mini_dose")
		values = append(values, miniDose)
	}
	if miniUnitID != "" {
		sets = append(sets, "mini_unit_id")
		values = append(values, miniUnitID)
	}
	if doseCount != "" {
		sets = append(sets, "dose_count")
		values = append(values, doseCount)
	}
	if doseCountID != "" {
		sets = append(sets, "dose_count_id")
		values = append(values, doseCountID)
	}
	if buyPrice != "" {
		sets = append(sets, "buy_price")
		values = append(values, buyPrice)
	}
	if isDiscount != "" {
		sets = append(sets, "is_discount")
		values = append(values, isDiscount)
	}
	if isBulkSales != "" {
		sets = append(sets, "is_bulk_sales")
		values = append(values, isBulkSales)
	}
	if bulkSalesPrice != "" {
		sets = append(sets, "bulk_sales_price")
		values = append(values, bulkSalesPrice)
	}
	if fetchAddress != "" {
		sets = append(sets, "fetch_address")
		values = append(values, "'"+fetchAddress+"'")
	}
	if onceDose != "" {
		sets = append(sets, "once_dose")
		values = append(values, onceDose)
	}
	if onceDoseUnitID != "" {
		sets = append(sets, "once_dose_unit_id")
		values = append(values, onceDoseUnitID)
	}
	if routeAdministrationID != "" {
		sets = append(sets, "route_administration_id")
		values = append(values, routeAdministrationID)
	}
	if frequencyID != "" {
		sets = append(sets, "frequency_id")
		values = append(values, frequencyID)
	}
	if defaultRemark != "" {
		sets = append(sets, "default_remark")
		values = append(values, "'"+defaultRemark+"'")
	}

	if effDay != "" {
		sets = append(sets, "eff_day")
		values = append(values, effDay)
	}
	if stockWarning != "" {
		sets = append(sets, "stock_warning")
		values = append(values, "'"+stockWarning+"'")
	}
	if englishName != "" {
		sets = append(sets, "english_name")
		values = append(values, "'"+englishName+"'")
	}
	if syCode != "" {
		sets = append(sets, "sy_code")
		values = append(values, "'"+syCode+"'")
	}

	setStr := strings.Join(sets, ",")
	valueStr := strings.Join(values, ",")
	insertSQL := "insert into drug (" + setStr + ") values (" + valueStr + ") RETURNING id;"
	fmt.Println("insertSQL===", insertSQL)

	var drugID int
	err := model.DB.QueryRow(insertSQL).Scan(&drugID)
	if err != nil {
		fmt.Println("err ===", err)
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": drugID})
}

//DrugList 药品列表
func DrugList(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	keyword := ctx.PostValue("keyword")
	drugClassID := ctx.PostValue("drug_class_id")
	status := ctx.PostValue("status")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
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

	countSQL := `select count(id) from drug where clinic_id=$1`
	selectSQL := `select d.name as drug_name,d.specification,du.name as packing_unit_name,d.ret_price,d.py_code,d.is_discount,d.default_remark,d.status from drug d
		left join dose_unit du on du.id = d.packing_unit_id
		where clinic_id=$1`
	if keyword != "" {
		countSQL += " and (barcode ~'" + keyword + "' or name ~'" + keyword + "')"
		selectSQL += " and (d.barcode ~'" + keyword + "' or d.name ~'" + keyword + "')"
	}
	if status != "" {
		countSQL += " and status=" + status
		selectSQL += " and d.status=" + status
	}
	if drugClassID != "" {
		countSQL += " and drug_class_id=" + drugClassID
		selectSQL += " and d.drug_class_id=" + drugClassID
	}

	fmt.Println("countSQL", countSQL)
	fmt.Println("selectSQL", selectSQL)

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
	barcode := ctx.PostValue("barcode")
	name := ctx.PostValue("name")
	pyCode := ctx.PostValue("py_code")
	printName := ctx.PostValue("print_name")
	specification := ctx.PostValue("specification")
	manuFactory := ctx.PostValue("manu_factory")
	status := ctx.PostValue("status")
	licenseNo := ctx.PostValue("license_no")
	doseFormID := ctx.PostValue("dose_form_id")
	drugClassID := ctx.PostValue("drug_class_id")

	miniDose := ctx.PostValue("mini_dose")
	miniUnitID := ctx.PostValue("mini_unit_id")
	doseCount := ctx.PostValue("dose_count")
	doseCountID := ctx.PostValue("dose_count_id")
	packingUnitID := ctx.PostValue("packing_unit_id")
	retPrice := ctx.PostValue("ret_price")
	buyPrice := ctx.PostValue("buy_price")
	isDiscount := ctx.PostValue("is_discount")
	isBulkSales := ctx.PostValue("is_bulk_sales")
	bulkSalesPrice := ctx.PostValue("bulk_sales_price")
	fetchAddress := ctx.PostValue("fetch_address")

	onceDose := ctx.PostValue("once_dose")
	onceDoseUnitID := ctx.PostValue("once_dose_unit_id")
	routeAdministrationID := ctx.PostValue("route_administration_id")
	frequencyID := ctx.PostValue("frequency_id")
	defaultRemark := ctx.PostValue("default_remark")

	effDay := ctx.PostValue("eff_day")
	stockWarning := ctx.PostValue("stock_warning")
	englishName := ctx.PostValue("english_name")
	syCode := ctx.PostValue("sy_code")

	if barcode == "" || name == "" || retPrice == "" || packingUnitID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from drug where name=$1 and barcode=$2 limit 1", name, barcode)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
		return
	}
	drug := FormatSQLRowToMap(row)
	_, ok := drug["id"]
	if ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "药品已存在"})
		return
	}

	sets := []string{"name", "barcode", "ret_price", "packing_unit_id"}
	values := []string{"'" + name + "'", "'" + barcode + "'", retPrice, packingUnitID}
	if pyCode != "" {
		sets = append(sets, "py_code")
		values = append(values, "'"+pyCode+"'")
	}
	if printName != "" {
		sets = append(sets, "print_name")
		values = append(values, "'"+printName+"'")
	}
	if specification != "" {
		sets = append(sets, "specification")
		values = append(values, "'"+specification+"'")
	}
	if manuFactory != "" {
		sets = append(sets, "manu_factory")
		values = append(values, "'"+manuFactory+"'")
	}
	if status != "" {
		sets = append(sets, "status")
		values = append(values, status)
	}
	if licenseNo != "" {
		sets = append(sets, "license_no")
		values = append(values, "'"+licenseNo+"'")
	}
	if doseFormID != "" {
		sets = append(sets, "dose_form_id")
		values = append(values, doseFormID)
	}
	if drugClassID != "" {
		sets = append(sets, "drug_class_id")
		values = append(values, drugClassID)
	}
	if miniDose != "" {
		sets = append(sets, "mini_dose")
		values = append(values, miniDose)
	}
	if miniUnitID != "" {
		sets = append(sets, "mini_unit_id")
		values = append(values, miniUnitID)
	}
	if doseCount != "" {
		sets = append(sets, "dose_count")
		values = append(values, doseCount)
	}
	if doseCountID != "" {
		sets = append(sets, "dose_count_id")
		values = append(values, doseCountID)
	}
	if buyPrice != "" {
		sets = append(sets, "buy_price")
		values = append(values, buyPrice)
	}
	if isDiscount != "" {
		sets = append(sets, "is_discount")
		values = append(values, isDiscount)
	}
	if isBulkSales != "" {
		sets = append(sets, "is_bulk_sales")
		values = append(values, isBulkSales)
	}
	if bulkSalesPrice != "" {
		sets = append(sets, "bulk_sales_price")
		values = append(values, bulkSalesPrice)
	}
	if fetchAddress != "" {
		sets = append(sets, "fetch_address")
		values = append(values, "'"+fetchAddress+"'")
	}
	if onceDose != "" {
		sets = append(sets, "once_dose")
		values = append(values, onceDose)
	}
	if onceDoseUnitID != "" {
		sets = append(sets, "once_dose_unit_id")
		values = append(values, onceDoseUnitID)
	}
	if routeAdministrationID != "" {
		sets = append(sets, "route_administration_id")
		values = append(values, routeAdministrationID)
	}
	if frequencyID != "" {
		sets = append(sets, "frequency_id")
		values = append(values, frequencyID)
	}
	if defaultRemark != "" {
		sets = append(sets, "default_remark")
		values = append(values, "'"+defaultRemark+"'")
	}

	if effDay != "" {
		sets = append(sets, "eff_day")
		values = append(values, effDay)
	}
	if stockWarning != "" {
		sets = append(sets, "stock_warning")
		values = append(values, "'"+stockWarning+"'")
	}
	if englishName != "" {
		sets = append(sets, "english_name")
		values = append(values, "'"+englishName+"'")
	}
	if syCode != "" {
		sets = append(sets, "sy_code")
		values = append(values, "'"+syCode+"'")
	}

	setStr := strings.Join(sets, ",")
	valueStr := strings.Join(values, ",")
	insertSQL := "insert into drug (" + setStr + ") values (" + valueStr + ") RETURNING id;"
	fmt.Println("insertSQL===", insertSQL)

	var drugID int
	err := model.DB.QueryRow(insertSQL).Scan(&drugID)
	if err != nil {
		fmt.Println("err ===", err)
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": drugID})
}
