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

//DrugAdd 添加药品
func DrugAdd(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	DrugType := ctx.PostValue("type")
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
	doseCountUnitID := ctx.PostValue("dose_count_unit_id")
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

	if clinicID == "" || barcode == "" || name == "" || retPrice == "" || DrugType == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from drug where barcode=$1 limit 1", barcode)
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
	var storehouseID string
	errs := model.DB.QueryRow("select id from storehouse where clinic_id=$1 limit 1", clinicID).Scan(&storehouseID)
	if errs != nil {
		fmt.Println("errs ===", errs)
		ctx.JSON(iris.Map{"code": "1", "msg": errs.Error()})
		return
	}
	fmt.Println("storehouseID==", storehouseID)

	sets := []string{"name", "barcode", "type"}
	values := []string{"'" + name + "'", "'" + barcode + "'", DrugType}
	stockSets := []string{"storehouse_id", "ret_price"}
	stockValues := []string{storehouseID, retPrice}
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
	if englishName != "" {
		sets = append(sets, "english_name")
		values = append(values, "'"+englishName+"'")
	}
	if syCode != "" {
		sets = append(sets, "sy_code")
		values = append(values, "'"+syCode+"'")
	}

	if doseCount != "" {
		sets = append(sets, "dose_count")
		values = append(values, doseCount)
	}
	if doseCountUnitID != "" {
		sets = append(sets, "dose_count_unit_id")
		values = append(values, doseCountUnitID)
	}
	if packingUnitID != "" {
		sets = append(sets, "packing_unit_id")
		values = append(values, packingUnitID)
	}

	if status != "" {
		stockSets = append(stockSets, "status")
		stockValues = append(stockValues, status)
	}
	if miniDose != "" {
		stockSets = append(stockSets, "mini_dose")
		stockValues = append(stockValues, miniDose)
	}
	if miniUnitID != "" {
		stockSets = append(stockSets, "mini_unit_id")
		stockValues = append(stockValues, miniUnitID)
	}
	if buyPrice != "" {
		stockSets = append(stockSets, "buy_price")
		stockValues = append(stockValues, buyPrice)
	}
	if isDiscount != "" {
		stockSets = append(stockSets, "is_discount")
		stockValues = append(stockValues, isDiscount)
	}
	if isBulkSales != "" {
		stockSets = append(stockSets, "is_bulk_sales")
		stockValues = append(stockValues, isBulkSales)
	}
	if bulkSalesPrice != "" {
		stockSets = append(stockSets, "bulk_sales_price")
		stockValues = append(stockValues, bulkSalesPrice)
	}
	if fetchAddress != "" {
		stockSets = append(stockSets, "fetch_address")
		stockValues = append(stockValues, "'"+fetchAddress+"'")
	}
	if effDay != "" {
		stockSets = append(stockSets, "eff_day")
		stockValues = append(stockValues, effDay)
	}
	if stockWarning != "" {
		stockSets = append(stockSets, "stock_warning")
		stockValues = append(stockValues, stockWarning)
	}

	setStr := strings.Join(sets, ",")
	valueStr := strings.Join(values, ",")
	insertSQL := "insert into drug (" + setStr + ") values (" + valueStr + ") RETURNING id;"
	fmt.Println("insertSQL===", insertSQL)

	tx, err := model.DB.Begin()
	var drugID string
	err = tx.QueryRow(insertSQL).Scan(&drugID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	stockSets = append(stockSets, "drug_id")
	stockValues = append(stockValues, drugID)

	stockSetStr := strings.Join(stockSets, ",")
	stockValueStr := strings.Join(stockValues, ",")
	stockSQL := "insert into drug_stock (" + stockSetStr + ") values (" + stockValueStr + ")"
	fmt.Println("stockSQL===", stockSQL)

	_, err1 := tx.Exec(stockSQL)
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
	DrugType := ctx.PostValue("type")
	keyword := ctx.PostValue("keyword")
	drugClassID := ctx.PostValue("drug_class_id")
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

	var storehouseID string
	errs := model.DB.QueryRow("select id from storehouse where clinic_id=$1 limit 1", clinicID).Scan(&storehouseID)
	if errs != nil {
		fmt.Println("errs ===", errs)
		ctx.JSON(iris.Map{"code": "1", "msg": errs.Error()})
		return
	}
	fmt.Println("storehouseID==", storehouseID)

	countSQL := `select count(ds.id) as total from drug_stock ds
	left join drug d on ds.drug_id = d.id
	where storehouse_id=$1`
	selectSQL := `select ds.id as drug_stock_id,d.name as drug_name,d.specification,
	d.packing_unit_id, du.name as packing_unit_name,
			ds.ret_price,d.py_code,ds.is_discount,d.default_remark,ds.status, ds.stock_amount, d.once_dose_unit_id, odu.name as once_dose_unit_name, d.route_administration_id, ra.name as route_administration_name, d.frequency_id, f.name as frequency_name, ds.storehouse_id, s.name as storehouse_name
			from drug_stock ds
			left join drug d on ds.drug_id = d.id
			left join dose_unit du on d.packing_unit_id = du.id
			left join dose_unit odu on d.once_dose_unit_id = odu.id
			left join route_administration ra on d.route_administration_id = ra.id
			left join frequency f on d.frequency_id = f.id
			left join storehouse s on ds.storehouse_id = s.id
			where storehouse_id=$1`

	// var countSet []string
	// var selectSet []string
	if DrugType != "" {
		countSQL += " and d.type=" + DrugType
		selectSQL += " and d.type=" + DrugType
	}
	if keyword != "" {
		countSQL += " and (d.barcode ~'" + keyword + "' or d.name ~'" + keyword + "')"
		selectSQL += " and (d.barcode ~'" + keyword + "' or d.name ~'" + keyword + "')"
		// countSet = append(countSet, "(barcode ~'"+keyword+"' or name ~'"+keyword+"')")
		// selectSet = append(selectSet, "(d.barcode ~'"+keyword+"' or d.name ~'"+keyword+"')")
	}
	if status != "" {
		countSQL += " and ds.status=" + status
		selectSQL += " and ds.status=" + status
		// countSet = append(countSet, "status="+status)
		// selectSet = append(selectSet, "d.status="+status)
	}
	if drugClassID != "" {
		countSQL += " and d.drug_class_id=" + drugClassID
		selectSQL += " and d.drug_class_id=" + drugClassID
		// countSet = append(countSet, "drug_class_id="+status)
		// selectSet = append(selectSet, "d.drug_class_id="+status)
	}

	// if len(countSet) > 0 {
	// 	countStr := strings.Join(countSet, " and ")
	// 	selectStr := strings.Join(selectSet, " and ")
	// 	countSQL = countSQL + " where " + countStr
	// 	selectSQL = selectSQL + " where " + selectStr
	// }
	total := model.DB.QueryRowx(countSQL, storehouseID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL+" offset $2 limit $3", storehouseID, offset, limit)
	results = FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

//DrugUpdate 修改药品
func DrugUpdate(ctx iris.Context) {
	drugStockID := ctx.PostValue("drug_stock_id")
	DrugType := ctx.PostValue("type")
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
	doseCountUnitID := ctx.PostValue("dose_count_unit_id")
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

	if drugStockID == "" || barcode == "" || name == "" || retPrice == "" || packingUnitID == "" || DrugType == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	var drugID int
	errs := model.DB.QueryRow("select drug_id from drug_stock where id=$1", drugStockID).Scan(&drugID)
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

	sets := []string{"name=" + "'" + name + "'", "type=" + DrugType, "barcode=" + "'" + barcode + "'", "updated_time=LOCALTIMESTAMP"}
	stockSets := []string{"ret_price=" + retPrice, "updated_time=LOCALTIMESTAMP"}
	if pyCode != "" {
		sets = append(sets, "py_code="+"'"+pyCode+"'")
	}
	if printName != "" {
		sets = append(sets, "print_name="+"'"+printName+"'")
	}
	if specification != "" {
		sets = append(sets, "specification="+"'"+specification+"'")
	}
	if manuFactory != "" {
		sets = append(sets, "manu_factory="+"'"+manuFactory+"'")
	}
	if licenseNo != "" {
		sets = append(sets, "license_no="+"'"+licenseNo+"'")
	}
	if doseFormID != "" {
		sets = append(sets, "dose_form_id="+doseFormID)
	}
	if drugClassID != "" {
		sets = append(sets, "drug_class_id="+drugClassID)
	}
	if onceDose != "" {
		sets = append(sets, "once_dose="+onceDose)
	}
	if onceDoseUnitID != "" {
		sets = append(sets, "once_dose_unit_id="+onceDoseUnitID)
	}
	if routeAdministrationID != "" {
		sets = append(sets, "route_administration_id="+routeAdministrationID)
	}
	if frequencyID != "" {
		sets = append(sets, "frequency_id="+frequencyID)
	}
	if defaultRemark != "" {
		sets = append(sets, "default_remark="+"'"+defaultRemark+"'")
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
	if doseCountUnitID != "" {
		sets = append(sets, "dose_count_unit_id="+doseCountUnitID)
	}
	if packingUnitID != "" {
		sets = append(sets, "packing_unit_id="+packingUnitID)
	}

	if status != "" {
		stockSets = append(stockSets, "status="+status)
	}

	if miniDose != "" {
		stockSets = append(stockSets, "miniDose="+miniDose)
	}
	if miniUnitID != "" {
		stockSets = append(stockSets, "mini_unit_id="+miniUnitID)
	}
	if buyPrice != "" {
		stockSets = append(stockSets, "buy_price="+buyPrice)
	}
	if isDiscount != "" {
		stockSets = append(stockSets, "is_discount="+isDiscount)
	}
	if isBulkSales != "" {
		stockSets = append(stockSets, "is_bulk_sales="+isBulkSales)
	}
	if bulkSalesPrice != "" {
		stockSets = append(stockSets, "bulk_sales_price="+bulkSalesPrice)
	}
	if fetchAddress != "" {
		stockSets = append(stockSets, "fetch_address="+"'"+fetchAddress+"'")
	}
	if effDay != "" {
		stockSets = append(stockSets, "eff_day="+effDay)
	}
	if stockWarning != "" {
		stockSets = append(stockSets, "stock_warning="+stockWarning)
	}

	setStr := strings.Join(sets, ",")
	stockStr := strings.Join(stockSets, ",")
	fmt.Println("setStr==", setStr)
	fmt.Println("stockStr==", stockStr)

	tx, err := model.DB.Begin()
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}

	updateSQL := "update drug set " + setStr + "where id=$1"
	stockUpdateSQL := "update drug_stock set " + stockStr + "where id=$1"
	fmt.Println("updateSQL===", updateSQL)

	_, err1 := tx.Exec(updateSQL, drugID)
	if err1 != nil {
		fmt.Println("err ===", err1)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err1.Error()})
		return
	}

	_, err2 := tx.Exec(stockUpdateSQL, drugStockID)
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
	DrugType := ctx.PostValue("type")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	var storehouseID string
	errs := model.DB.QueryRow("select id from storehouse where clinic_id=$1 limit 1", clinicID).Scan(&storehouseID)
	if errs != nil {
		fmt.Println("errs ===", errs)
		ctx.JSON(iris.Map{"code": "1", "msg": errs.Error()})
		return
	}

	selectSQL := `select ds.id as drug_stock_id,d.id as drug_id,d.manu_factory,d.name as drug_name,du.name as packing_unit_name,
	ds.ret_price,ds.buy_price,ds.eff_day,ds.stock_amount from drug d
	left join drug_stock ds on ds.drug_id = d.id
	left join dose_unit du on ds.packing_unit_id = du.id
	where ds.storehouse_id=$1`

	if keyword != "" {
		selectSQL += " and (d.barcode ~'" + keyword + "' or d.name ~'" + keyword + "')"
	}

	if DrugType != "" {
		selectSQL += " and d.type=" + DrugType
	}

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL, storehouseID)
	results = FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results})
}

//DrugDetail 药品详情
func DrugDetail(ctx iris.Context) {
	drugStockID := ctx.PostValue("drug_stock_id")

	if drugStockID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	sql := `select d.name,d.specification,d.manu_factory,df.name as dose_form_name,d.dose_form_id,
		d.print_name,d.license_no,dc.name as drug_class_name,d.drug_class_id,d.py_code,d.barcode,ds.status,
		ds.mini_dose,mdu.name as mini_unit_name,ds.mini_unit_id,ds.dose_count,ds.dose_count_unit_id,cdu.name as dose_count_name,
		ds.packing_unit_id,pdu.name as packing_unit_name,ds.ret_price,ds.buy_price,ds.is_discount,ds.is_bulk_sales,ds.bulk_sales_price,ds.fetch_address,
		d.once_dose,d.once_dose_unit_id,odu.name as once_dose_unit_name,d.route_administration_id,ra.name as route_administration_name,
		d.frequency_id,f.name as frequency_name,d.default_remark,ds.eff_day,ds.stock_warning,d.english_name,d.sy_code
		from drug_stock ds
		left join drug d on ds.drug_id = d.id
		left join dose_form df on d.dose_form_id = df.id
		left join drug_class dc on d.drug_class_id = dc.id
		left join dose_unit mdu on ds.mini_unit_id = mdu.id
		left join dose_unit cdu on ds.dose_count_unit_id = cdu.id
		left join dose_unit pdu on ds.packing_unit_id = pdu.id
		left join dose_unit odu on d.once_dose_unit_id = odu.id
		left join route_administration ra on d.route_administration_id = ra.id
		left join frequency f on d.frequency_id = f.id
		where ds.id=$1`
	arows := model.DB.QueryRowx(sql, drugStockID)
	if arows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询结果不存在"})
		return
	}
	result := FormatSQLRowToMap(arows)
	ctx.JSON(iris.Map{"code": "200", "data": result})
}

//BatchSetting 批量设置药品
func BatchSetting(ctx iris.Context) {
	effDay := ctx.PostValue("eff_day")
	isDiscount := ctx.PostValue("is_discount")
	items := ctx.PostValue("items")
	if items == "" || (effDay == "" && isDiscount == "") {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	var results []map[string]string
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	sql := "update drug_stock set"

	var sets []string
	asValues := []string{"id"}
	var values []string
	for _, v := range results {
		var s []string
		s = append(s, v["drug_stock_id"])
		if effDay != "" {
			s = append(s, effDay)
		}
		if isDiscount != "" {
			s = append(s, isDiscount)
		}
		str := strings.Join(s, ",")
		str = "(" + str + ")"
		values = append(values, str)
	}
	valueStr := strings.Join(values, ",")

	if effDay != "" {
		sets = append(sets, " eff_day=tmp.eff_day")
		asValues = append(asValues, "eff_day")
	}

	if isDiscount != "" {
		sets = append(sets, " is_discount=tmp.isDiscount")
		asValues = append(asValues, "isDiscount")
	}
	setStr := strings.Join(sets, ",")
	asStr := strings.Join(asValues, ",")
	sql += setStr + " from (values " + valueStr + ") as tmp(" + asStr + ") where drug_stock.id = tmp.id"

	_, erre := model.DB.Exec(sql)
	if erre != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": erre.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok"})
}

//DrugInstock 入库
func DrugInstock(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	items := ctx.PostValue("items")
	operationID := ctx.PostValue("instock_operation_id")
	instockWayID := ctx.PostValue("instock_way_id")
	supplierID := ctx.PostValue("supplier_id")
	remark := ctx.PostValue("remark")
	instockDate := ctx.PostValue("instock_date")

	if clinicID == "" || instockWayID == "" || supplierID == "" || instockDate == "" || operationID == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	var results []map[string]string
	errj := json.Unmarshal([]byte(items), &results)
	fmt.Println("===", results)

	if errj != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": errj.Error()})
		return
	}

	var storehouseID string
	errs := model.DB.QueryRow("select id from storehouse where clinic_id=$1 limit 1", clinicID).Scan(&storehouseID)
	if errs != nil {
		fmt.Println("errs ===", errs)
		ctx.JSON(iris.Map{"code": "1", "msg": errs.Error()})
		return
	}

	var values []string
	orderNumber := "RKD-" + strconv.FormatInt(time.Now().Unix(), 10)

	sets := []string{
		"storehouse_id",
		"drug_id",
		"packing_unit_id",
		"manu_factory",
		"instock_amount",
		"ret_price",
		"buy_price",
		"serial",
		"eff_day",
		"order_number",
		"instock_way_id",
		"supplier_id",
		"instock_date",
		"instock_operation_id"}

	if remark != "" {
		sets = append(sets, "remark")
	}
	for _, v := range results {
		drugID := v["drug_id"]
		var s []string
		row := model.DB.QueryRowx("select id from drug where id=$1 limit 1", drugID)
		if row == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
			return
		}
		drug := FormatSQLRowToMap(row)
		_, ok := drug["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "新增药品不存在"})
			return
		}
		s = append(s, storehouseID, v["drug_id"], v["packing_unit_id"], "'"+v["manu_factory"]+"'", v["instock_amount"],
			v["ret_price"], v["buy_price"], "'"+v["serial"]+"'", v["eff_day"], "'"+orderNumber+"'", instockWayID,
			supplierID, "date '"+instockDate+"'", operationID)
		if remark != "" {
			s = append(s, "'"+remark+"'")
		}
		str := strings.Join(s, ",")
		str = "(" + str + ")"
		values = append(values, str)
	}

	setStr := strings.Join(sets, ",")
	valueStr := strings.Join(values, ",")
	insertSQL := "insert into instock_record (" + setStr + ") values " + valueStr
	fmt.Println("insertSQL===", insertSQL)
	_, err := model.DB.Exec(insertSQL)
	if err != nil {
		fmt.Println("err ===", err)
		ctx.JSON(iris.Map{"code": "1", "msg": "请检查是否漏填"})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//InstockRecord 入库记录列表
func InstockRecord(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")
	orderNumber := ctx.PostValue("order_number")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if startDate != "" && endDate == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请选择结束日期"})
		return
	}
	if startDate == "" && endDate != "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请选择开始日期"})
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

	var storehouseID string
	errs := model.DB.QueryRow("select id from storehouse where clinic_id=$1 limit 1", clinicID).Scan(&storehouseID)
	if errs != nil {
		fmt.Println("errs ===", errs)
		ctx.JSON(iris.Map{"code": "1", "msg": errs.Error()})
		return
	}

	countSQL := `select count(DISTINCT order_number) as total from instock_record where storehouse_id=$1`
	selectSQL := `select ir.id as instock_record_id,ir.instock_date,ir.order_number, iw.name as instock_way_name,
		vp.name as verify_operation_name,s.name as supplier_name,p.name as instock_operation_name,ir.verify_status
		from instock_record ir
		left join supplier s on ir.supplier_id = s.id
		left join instock_way iw on ir.instock_way_id = iw.id
		left join personnel p on ir.instock_operation_id = p.id
		left join personnel vp on ir.verify_operation_id = vp.id
		where storehouse_id=$1`

	if startDate != "" && endDate != "" {
		if startDate > endDate {
			ctx.JSON(iris.Map{"code": "-1", "msg": "开始日期必须大于结束日期"})
			return
		}
		countSQL += " and instock_date between date'" + startDate + "' and date '" + endDate + "'"
		selectSQL += " and ir.instock_date between date'" + startDate + "' and date '" + endDate + "'"
	}

	if orderNumber != "" {
		countSQL += " and order_number='" + orderNumber + "'"
		selectSQL += " and ir.order_number='" + orderNumber + "'"
	}

	fmt.Println("countSQL===", countSQL)
	fmt.Println("selectSQL===", selectSQL)
	total := model.DB.QueryRowx(countSQL, storehouseID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL+" offset $2 limit $3", storehouseID, offset, limit)
	results = FormatSQLRowsToMapArray(rows)

	var instockRecord []map[string]interface{}
	for _, v := range results {
		has := false
		for _, vRes := range instockRecord {
			if vRes["order_number"].(string) == v["order_number"].(string) {
				has = true
				continue
			}
		}
		if !has {
			instockRecord = append(instockRecord, v)
		}
	}
	ctx.JSON(iris.Map{"code": "200", "data": instockRecord, "page_info": pageInfo})
}

//InstockRecordDetail 入库记录详情
func InstockRecordDetail(ctx iris.Context) {
	orderNumber := ctx.PostValue("order_number")
	if orderNumber == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	sql := `select ir.id as instock_record_id,d.name as drug_name,ir.drug_id,ir.packing_unit_id,du.name as packing_unit_name,ir.manu_factory,ir.instock_amount,
		ir.ret_price,ir.buy_price,ir.serial,ir.instock_date,ir.order_number,ir.created_time,s.name as supplier_name,ir.supplier_id,ir.remark,
		ir.verify_operation_id,vp.name as verify_operation_name,ir.instock_way_id,iw.name as instock_way_name,ir.instock_operation_id,p.name as instock_operation_name 
		from instock_record ir
		left join drug d on ir.drug_id = d.id
		left join supplier s on ir.supplier_id = s.id
		left join instock_way iw on ir.instock_way_id = iw.id
		left join personnel p on ir.instock_operation_id = p.id
		left join personnel vp on ir.verify_operation_id = vp.id
		left join dose_unit du on ir.packing_unit_id = du.id
		where order_number=$1`

	arows, err := model.DB.Queryx(sql, orderNumber)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	result := FormatSQLRowsToMapArray(arows)
	ctx.JSON(iris.Map{"code": "200", "data": result})
}

//InstockUpdate 入库记录修改
func InstockUpdate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	orderNumber := ctx.PostValue("order_number")
	items := ctx.PostValue("items")
	operationID := ctx.PostValue("instock_operation_id")
	instockWayID := ctx.PostValue("instock_way_id")
	supplierID := ctx.PostValue("supplier_id")
	remark := ctx.PostValue("remark")
	instockDate := ctx.PostValue("instock_date")

	if orderNumber == "" || clinicID == "" || instockWayID == "" || supplierID == "" || instockDate == "" || operationID == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	var results []map[string]string
	errj := json.Unmarshal([]byte(items), &results)
	fmt.Println("===", results)

	if errj != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": errj.Error()})
		return
	}

	var storehouseID string
	errs := model.DB.QueryRow("select id from storehouse where clinic_id=$1 limit 1", clinicID).Scan(&storehouseID)
	if errs != nil {
		fmt.Println("errs ===", errs)
		ctx.JSON(iris.Map{"code": "1", "msg": errs.Error()})
		return
	}

	row := model.DB.QueryRowx("select * from instock_record where order_number=$1 and storehouse_id=$2 limit 1", orderNumber, storehouseID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	instockRecord := FormatSQLRowToMap(row)
	_, ok := instockRecord["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "入库记录不存在"})
		return
	}
	verifyStatus := instockRecord["verify_status"]
	createdTime := instockRecord["created_time"].(time.Time)
	fmt.Println("created_time====", createdTime)
	if verifyStatus.(string) == "02" {
		ctx.JSON(iris.Map{"code": "1", "msg": "入库记录已审核，不能修改"})
		return
	}

	tx, errb := model.DB.Begin()

	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errb})
		return
	}

	_, errd := tx.Exec("delete from instock_record where order_number=$1 and storehouse_id=$2", orderNumber, storehouseID)
	if errd != nil {
		fmt.Println("errd ===", errd)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errd.Error()})
		return
	}

	var values []string
	sets := []string{
		"storehouse_id",
		"drug_id",
		"packing_unit_id",
		"manu_factory",
		"instock_amount",
		"ret_price",
		"buy_price",
		"serial",
		"eff_day",
		"order_number",
		"instock_way_id",
		"supplier_id",
		"instock_date",
		"instock_operation_id",
		"created_time"}

	if remark != "" {
		sets = append(sets, "remark")
	}
	for _, v := range results {
		drugID := v["drug_id"]
		var s []string
		row := model.DB.QueryRowx("select id from drug where id=$1 limit 1", drugID)
		if row == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
			return
		}
		drug := FormatSQLRowToMap(row)
		_, ok := drug["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "新增药品不存在"})
			return
		}
		s = append(s, storehouseID, v["drug_id"], v["packing_unit_id"], "'"+v["manu_factory"]+"'", v["instock_amount"],
			v["ret_price"], v["buy_price"], "'"+v["serial"]+"'", v["eff_day"], "'"+orderNumber+"'", instockWayID,
			supplierID, "date '"+instockDate+"'", operationID, "timestamp '"+createdTime.Format("2006-01-02 15:04:05")+"'")
		if remark != "" {
			s = append(s, "'"+remark+"'")
		}
		str := strings.Join(s, ",")
		str = "(" + str + ")"
		values = append(values, str)
	}

	setStr := strings.Join(sets, ",")
	valueStr := strings.Join(values, ",")
	insertSQL := "insert into instock_record (" + setStr + ") values " + valueStr
	fmt.Println("insertSQL===", insertSQL)
	_, err := tx.Exec(insertSQL)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": "请检查是否漏填"})
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

//InstockCheck 入库审核
func InstockCheck(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	orderNumber := ctx.PostValue("order_number")
	operationID := ctx.PostValue("verify_operation_id")
	items := ctx.PostValue("items")
	if orderNumber == "" || clinicID == "" || operationID == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	var results []map[string]string
	errj := json.Unmarshal([]byte(items), &results)
	fmt.Println("===", results)

	if errj != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": errj.Error()})
		return
	}
	var storehouseID string
	errs := model.DB.QueryRow("select id from storehouse where clinic_id=$1 limit 1", clinicID).Scan(&storehouseID)
	if errs != nil {
		fmt.Println("errs ===", errs)
		ctx.JSON(iris.Map{"code": "1", "msg": errs.Error()})
		return
	}

	sql := "update drug_stock set"

	var sets []string
	var asValues []string
	var values []string
	for _, v := range results {
		instockRecordID := v["instock_record_id"]
		var s []string
		row := model.DB.QueryRowx("select * from instock_record where id=$1 and order_number=$2 limit 1", instockRecordID, orderNumber)
		if row == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "审核失败"})
			return
		}
		instockRecord := FormatSQLRowToMap(row)
		_, ok := instockRecord["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "入库记录不存在"})
			return
		}
		verifyStatus := instockRecord["verify_status"]
		if verifyStatus.(string) != "01" {
			ctx.JSON(iris.Map{"code": "1", "msg": "当前状态不能审核"})
			return
		}

		drow := model.DB.QueryRowx("select id,stock_amount from drug_stock where storehouse_id=$1 and drug_id=$2 limit 1", strconv.FormatInt(instockRecord["storehouse_id"].(int64), 10), strconv.FormatInt(instockRecord["drug_id"].(int64), 10))
		if drow == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "审核失败"})
			return
		}
		drugStock := FormatSQLRowToMap(drow)
		_, dok := drugStock["id"]
		if !dok {
			ctx.JSON(iris.Map{"code": "1", "msg": "入库失败"})
			return
		}
		stockAmount := drugStock["stock_amount"].(int64) + instockRecord["instock_amount"].(int64)

		s = append(s, strconv.FormatInt(instockRecord["storehouse_id"].(int64), 10), strconv.FormatInt(instockRecord["drug_id"].(int64), 10), strconv.FormatInt(stockAmount, 10), strconv.FormatInt(instockRecord["ret_price"].(int64), 10),
			strconv.FormatInt(instockRecord["buy_price"].(int64), 10), strconv.FormatInt(instockRecord["packing_unit_id"].(int64), 10), strconv.FormatInt(instockRecord["eff_day"].(int64), 10))
		str := strings.Join(s, ",")
		str = "(" + str + ")"
		values = append(values, str)
	}
	valueStr := strings.Join(values, ",")

	sets = append(sets, " stock_amount=tmp.stockAmount", " ret_price=tmp.ret_price",
		" buy_price=tmp.buy_price", " packing_unit_id=tmp.packing_unit_id", " eff_day=tmp.eff_day", " updated_time=LOCALTIMESTAMP")

	asValues = append(asValues, "storehouse_id", "drug_id", "stockAmount", "ret_price", "buy_price", "packing_unit_id", "eff_day")

	setStr := strings.Join(sets, ",")
	asStr := strings.Join(asValues, ",")
	sql += setStr + " from (values " + valueStr + ") as tmp(" + asStr + ") where drug_stock.storehouse_id = tmp.storehouse_id and drug_stock.drug_id = tmp.drug_id"
	fmt.Println("sql===", sql)

	tx, err := model.DB.Begin()

	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}

	_, err1 := tx.Exec(sql)
	if err1 != nil {
		fmt.Println("err1 ===", err1)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err1.Error()})
		return
	}

	_, err2 := tx.Exec("update instock_record set verify_status=$1,verify_operation_id=$2,updated_time=LOCALTIMESTAMP", "02", operationID)
	if err2 != nil {
		fmt.Println("err2 ===", err2)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err2.Error()})
		return
	}

	err3 := tx.Commit()
	if err3 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "msg": "ok"})
}

//InstockRecordDelete 删除入库记录
func InstockRecordDelete(ctx iris.Context) {
	orderNumber := ctx.PostValue("order_number")
	if orderNumber == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	rows, _ := model.DB.Queryx("select id from instock_record where order_number=$1 and verify_status='01'", orderNumber)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "删除失败"})
		return
	}
	instockRecord := FormatSQLRowsToMapArray(rows)
	len := len(instockRecord)
	if len == 0 {
		ctx.JSON(iris.Map{"code": "1", "msg": "入库记录不存在"})
		return
	}

	_, err := model.DB.Exec("delete from instock_record where order_number=$1 and verify_status='01'", orderNumber)

	if err != nil {
		fmt.Println("err ===", err)
		ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "msg": "ok"})
}

//DrugOutstock 出库
func DrugOutstock(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	items := ctx.PostValue("items")
	operationID := ctx.PostValue("outstock_operation_id")
	outstockWayID := ctx.PostValue("outstock_way_id")
	departmentID := ctx.PostValue("department_id")
	personnelID := ctx.PostValue("personnel_id")
	remark := ctx.PostValue("remark")
	outstockDate := ctx.PostValue("outstock_date")

	if clinicID == "" || items == "" || outstockWayID == "" || departmentID == "" || operationID == "" || personnelID == "" || outstockDate == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	dprow := model.DB.QueryRowx("select id from department_personnel where department_id=$1 and personnel_id!=$2 limit 1", departmentID, personnelID)
	if dprow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "出库失败"})
		return
	}
	departmentPersonnel := FormatSQLRowToMap(dprow)
	_, dok := departmentPersonnel["id"]
	if !dok {
		ctx.JSON(iris.Map{"code": "1", "msg": "领用科室与领用人员不符"})
		return
	}

	var results []map[string]string
	errj := json.Unmarshal([]byte(items), &results)
	fmt.Println("===", results)

	if errj != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": errj.Error()})
		return
	}

	var storehouseID string
	errs := model.DB.QueryRow("select id from storehouse where clinic_id=$1 limit 1", clinicID).Scan(&storehouseID)
	if errs != nil {
		fmt.Println("errs ===", errs)
		ctx.JSON(iris.Map{"code": "1", "msg": errs.Error()})
		return
	}

	var values []string
	orderNumber := "CKD-" + strconv.FormatInt(time.Now().Unix(), 10)

	sets := []string{
		"storehouse_id",
		"drug_id",
		"department_id",
		"personnel_id",
		"packing_unit_id",
		"manu_factory",
		"outstock_amount",
		"ret_price",
		"buy_price",
		"serial",
		"eff_day",
		"order_number",
		"outstock_way_id",
		"outstock_date",
		"outstock_operation_id"}

	if remark != "" {
		sets = append(sets, "remark")
	}
	for _, v := range results {
		drugID := v["drug_id"]
		var s []string
		row := model.DB.QueryRowx("select id from drug where id=$1 limit 1", drugID)
		if row == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "新增出库失败"})
			return
		}
		drug := FormatSQLRowToMap(row)
		_, ok := drug["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "新增出库药品不存在"})
			return
		}
		s = append(s, storehouseID, v["drug_id"], departmentID, personnelID, v["packing_unit_id"], "'"+v["manu_factory"]+"'", v["outstock_amount"],
			v["ret_price"], v["buy_price"], "'"+v["serial"]+"'", v["eff_day"], "'"+orderNumber+"'", outstockWayID,
			"date '"+outstockDate+"'", operationID)
		if remark != "" {
			s = append(s, "'"+remark+"'")
		}
		str := strings.Join(s, ",")
		str = "(" + str + ")"
		values = append(values, str)
	}

	setStr := strings.Join(sets, ",")
	valueStr := strings.Join(values, ",")
	insertSQL := "insert into outstock_record (" + setStr + ") values " + valueStr
	fmt.Println("insertSQL===", insertSQL)
	_, err := model.DB.Exec(insertSQL)
	if err != nil {
		fmt.Println("err ===", err)
		ctx.JSON(iris.Map{"code": "1", "msg": "请检查是否漏填"})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//OutstockRecord 出库记录列表
func OutstockRecord(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")
	orderNumber := ctx.PostValue("order_number")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if startDate != "" && endDate == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请选择结束日期"})
		return
	}
	if startDate == "" && endDate != "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请选择开始日期"})
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

	var storehouseID string
	errs := model.DB.QueryRow("select id from storehouse where clinic_id=$1 limit 1", clinicID).Scan(&storehouseID)
	if errs != nil {
		fmt.Println("errs ===", errs)
		ctx.JSON(iris.Map{"code": "1", "msg": errs.Error()})
		return
	}

	countSQL := `select count(DISTINCT order_number) as total from outstock_record where storehouse_id=$1`
	selectSQL := `select outr.id as outstock_record_id,outr.outstock_date,outr.order_number, ow.name as outstock_way_name,
		vp.name as verify_operation_name,d.name as department_name,p.name as personnel_name,op.name as outstock_operation_name,outr.verify_status
		from outstock_record outr
		left join department d on outr.department_id = d.id
		left join personnel p on outr.personnel_id = p.id
		left join outstock_way ow on outr.outstock_way_id = ow.id
		left join personnel op on outr.outstock_operation_id = op.id
		left join personnel vp on outr.verify_operation_id = vp.id
		where storehouse_id=$1`

	if startDate != "" && endDate != "" {
		if startDate > endDate {
			ctx.JSON(iris.Map{"code": "-1", "msg": "开始日期必须大于结束日期"})
			return
		}
		countSQL += " and outstock_date between date'" + startDate + "' and date '" + endDate + "'"
		selectSQL += " and outr.outstock_date between date'" + startDate + "' and date '" + endDate + "'"
	}

	if orderNumber != "" {
		countSQL += " and order_number='" + orderNumber + "'"
		selectSQL += " and outr.order_number='" + orderNumber + "'"
	}

	fmt.Println("countSQL===", countSQL)
	fmt.Println("selectSQL===", selectSQL)
	total := model.DB.QueryRowx(countSQL, storehouseID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL+" offset $2 limit $3", storehouseID, offset, limit)
	results = FormatSQLRowsToMapArray(rows)

	var outstockRecord []map[string]interface{}
	for _, v := range results {
		has := false
		for _, vRes := range outstockRecord {
			if vRes["order_number"].(string) == v["order_number"].(string) {
				has = true
				continue
			}
		}
		if !has {
			outstockRecord = append(outstockRecord, v)
		}
	}
	ctx.JSON(iris.Map{"code": "200", "data": outstockRecord, "page_info": pageInfo})
}

//OutstockRecordDetail 出库记录详情
func OutstockRecordDetail(ctx iris.Context) {
	orderNumber := ctx.PostValue("order_number")
	if orderNumber == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	sql := `select outr.id as outstock_record_id,d.name as drug_name,outr.drug_id,outr.packing_unit_id,du.name as packing_unit_name,outr.manu_factory,outr.outstock_amount,
		outr.ret_price,outr.buy_price,outr.serial,outr.outstock_date,outr.order_number,outr.created_time,dept.name as department_name,outr.department_id,outr.remark,
		outr.verify_operation_id,vp.name as verify_operation_name,outr.personnel_id,p.name as personnel_name,outr.outstock_way_id,ow.name as outstock_way_name,
		outr.outstock_operation_id,op.name as outstock_operation_name 
		from outstock_record outr
		left join drug d on outr.drug_id = d.id
		left join department dept on outr.department_id = dept.id
		left join personnel p on outr.personnel_id = p.id
		left join outstock_way ow on outr.outstock_way_id = ow.id
		left join personnel op on outr.outstock_operation_id = op.id
		left join personnel vp on outr.verify_operation_id = vp.id
		left join dose_unit du on outr.packing_unit_id = du.id
		where order_number=$1`

	arows, err := model.DB.Queryx(sql, orderNumber)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	result := FormatSQLRowsToMapArray(arows)
	ctx.JSON(iris.Map{"code": "200", "data": result})
}

//OutstockUpdate 出库记录修改
func OutstockUpdate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	orderNumber := ctx.PostValue("order_number")
	items := ctx.PostValue("items")
	operationID := ctx.PostValue("outstock_operation_id")
	outstockWayID := ctx.PostValue("outstock_way_id")
	departmentID := ctx.PostValue("department_id")
	personnelID := ctx.PostValue("personnel_id")
	remark := ctx.PostValue("remark")
	outstockDate := ctx.PostValue("outstock_date")

	if clinicID == "" || orderNumber == "" || items == "" || outstockWayID == "" || departmentID == "" || operationID == "" || personnelID == "" || outstockDate == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	dprow := model.DB.QueryRowx("select id from department_personnel where department_id=$1 and personnel_id!=$2 limit 1", departmentID, personnelID)
	if dprow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "出库失败"})
		return
	}
	departmentPersonnel := FormatSQLRowToMap(dprow)
	_, dok := departmentPersonnel["id"]
	if !dok {
		ctx.JSON(iris.Map{"code": "1", "msg": "领用科室与领用人员不符"})
		return
	}

	var results []map[string]string
	errj := json.Unmarshal([]byte(items), &results)
	fmt.Println("===", results)

	if errj != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": errj.Error()})
		return
	}

	var storehouseID string
	errs := model.DB.QueryRow("select id from storehouse where clinic_id=$1 limit 1", clinicID).Scan(&storehouseID)
	if errs != nil {
		fmt.Println("errs ===", errs)
		ctx.JSON(iris.Map{"code": "1", "msg": errs.Error()})
		return
	}

	row := model.DB.QueryRowx("select * from outstock_record where order_number=$1 and storehouse_id=$2 limit 1", orderNumber, storehouseID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	outstockRecord := FormatSQLRowToMap(row)
	_, ok := outstockRecord["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "出库记录不存在"})
		return
	}
	verifyStatus := outstockRecord["verify_status"]
	createdTime := outstockRecord["created_time"].(time.Time)
	fmt.Println("created_time====", createdTime)
	if verifyStatus.(string) == "02" {
		ctx.JSON(iris.Map{"code": "1", "msg": "出库记录已审核，不能修改"})
		return
	}

	tx, errb := model.DB.Begin()

	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errb})
		return
	}

	_, errd := tx.Exec("delete from outstock_record where order_number=$1 and storehouse_id=$2", orderNumber, storehouseID)
	if errd != nil {
		fmt.Println("errd ===", errd)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errd.Error()})
		return
	}

	var values []string
	sets := []string{
		"storehouse_id",
		"drug_id",
		"department_id",
		"personnel_id",
		"packing_unit_id",
		"manu_factory",
		"outstock_amount",
		"ret_price",
		"buy_price",
		"serial",
		"eff_day",
		"order_number",
		"outstock_way_id",
		"outstock_date",
		"outstock_operation_id",
		"created_time"}

	if remark != "" {
		sets = append(sets, "remark")
	}
	for _, v := range results {
		drugID := v["drug_id"]
		var s []string
		row := model.DB.QueryRowx("select id from drug where id=$1 limit 1", drugID)
		if row == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
			return
		}
		drug := FormatSQLRowToMap(row)
		_, ok := drug["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "修改的出库药品不存在"})
			return
		}
		s = append(s, storehouseID, v["drug_id"], departmentID, personnelID, v["packing_unit_id"], "'"+v["manu_factory"]+"'", v["outstock_amount"],
			v["ret_price"], v["buy_price"], "'"+v["serial"]+"'", v["eff_day"], "'"+orderNumber+"'", outstockWayID,
			"date '"+outstockDate+"'", operationID, "timestamp '"+createdTime.Format("2006-01-02 15:04:05")+"'")
		if remark != "" {
			s = append(s, "'"+remark+"'")
		}
		str := strings.Join(s, ",")
		str = "(" + str + ")"
		values = append(values, str)
	}

	setStr := strings.Join(sets, ",")
	valueStr := strings.Join(values, ",")
	insertSQL := "insert into outstock_record (" + setStr + ") values " + valueStr
	fmt.Println("insertSQL===", insertSQL)
	_, err := tx.Exec(insertSQL)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": "请检查是否漏填"})
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

//OutstockCheck 出库审核
func OutstockCheck(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	orderNumber := ctx.PostValue("order_number")
	operationID := ctx.PostValue("verify_operation_id")
	items := ctx.PostValue("items")
	if orderNumber == "" || clinicID == "" || operationID == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	var results []map[string]string
	errj := json.Unmarshal([]byte(items), &results)
	fmt.Println("===", results)

	if errj != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": errj.Error()})
		return
	}
	var storehouseID string
	errs := model.DB.QueryRow("select id from storehouse where clinic_id=$1 limit 1", clinicID).Scan(&storehouseID)
	if errs != nil {
		fmt.Println("errs ===", errs)
		ctx.JSON(iris.Map{"code": "1", "msg": errs.Error()})
		return
	}

	sql := "update drug_stock set"

	var sets []string
	var asValues []string
	var values []string
	for _, v := range results {
		outstockRecordID := v["outstock_record_id"]
		var s []string
		row := model.DB.QueryRowx("select * from outstock_record where id=$1 and order_number=$2 limit 1", outstockRecordID, orderNumber)
		if row == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "审核失败"})
			return
		}
		outstockRecord := FormatSQLRowToMap(row)
		_, ok := outstockRecord["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "出库记录不存在"})
			return
		}

		verifyStatus := outstockRecord["verify_status"]
		if verifyStatus.(string) != "01" {
			ctx.JSON(iris.Map{"code": "1", "msg": "当前状态不能审核"})
			return
		}

		drow := model.DB.QueryRowx("select id,stock_amount from drug_stock where storehouse_id=$1 and drug_id=$2 limit 1", strconv.FormatInt(outstockRecord["storehouse_id"].(int64), 10), strconv.FormatInt(outstockRecord["drug_id"].(int64), 10))
		if drow == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "审核失败"})
			return
		}
		drugStock := FormatSQLRowToMap(drow)
		_, dok := drugStock["id"]
		if !dok {
			ctx.JSON(iris.Map{"code": "1", "msg": "入库失败"})
			return
		}
		stockAmount := drugStock["stock_amount"].(int64) - outstockRecord["outstock_amount"].(int64)

		s = append(s, strconv.FormatInt(outstockRecord["storehouse_id"].(int64), 10),
			strconv.FormatInt(outstockRecord["drug_id"].(int64), 10), strconv.FormatInt(stockAmount, 10))
		str := strings.Join(s, ",")
		str = "(" + str + ")"
		values = append(values, str)
	}
	valueStr := strings.Join(values, ",")

	sets = append(sets, " stock_amount=tmp.stock_amount", " updated_time=LOCALTIMESTAMP")

	asValues = append(asValues, "storehouse_id", "drug_id", "stock_amount")

	setStr := strings.Join(sets, ",")
	asStr := strings.Join(asValues, ",")
	sql += setStr + " from (values " + valueStr + ") as tmp(" + asStr + ") where drug_stock.storehouse_id = tmp.storehouse_id and drug_stock.drug_id = tmp.drug_id"
	fmt.Println("sql===", sql)

	tx, err := model.DB.Begin()

	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}

	_, err1 := tx.Exec(sql)
	if err1 != nil {
		fmt.Println("err1 ===", err1)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err1.Error()})
		return
	}

	_, err2 := tx.Exec("update outstock_record set verify_status=$1,verify_operation_id=$2,updated_time=LOCALTIMESTAMP", "02", operationID)
	if err2 != nil {
		fmt.Println("err2 ===", err2)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err2.Error()})
		return
	}

	err3 := tx.Commit()
	if err3 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "msg": "ok"})
}

//OutstockRecordDelete 删除出库记录
func OutstockRecordDelete(ctx iris.Context) {
	orderNumber := ctx.PostValue("order_number")
	if orderNumber == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	rows, _ := model.DB.Queryx("select id from outstock_record where order_number=$1 and verify_status='01'", orderNumber)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "删除失败"})
		return
	}
	instockRecord := FormatSQLRowsToMapArray(rows)
	len := len(instockRecord)
	if len == 0 {
		ctx.JSON(iris.Map{"code": "1", "msg": "出库记录不存在"})
		return
	}

	_, err := model.DB.Exec("delete from outstock_record where order_number=$1 and verify_status='01'", orderNumber)

	if err != nil {
		fmt.Println("err ===", err)
		ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "msg": "ok"})
}
