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

// MaterialCreate 创建材料缴费项目
func MaterialCreate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	pyCode := ctx.PostValue("py_code")
	idcCode := ctx.PostValue("idc_code")
	unitID := ctx.PostValue("unit_id")
	remark := ctx.PostValue("remark")
	manuFactoryID := ctx.PostValue("manu_factory_id")
	specification := ctx.PostValue("specification")

	price := ctx.PostValue("price")
	cost := ctx.PostValue("cost")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")
	effDay := ctx.PostValue("eff_day")
	stockWarning := ctx.PostValue("stock_warning")

	if clinicID == "" || name == "" || price == "" || unitID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
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

	var storehouseID string
	errs := model.DB.QueryRow("select id from storehouse where clinic_id=$1 limit 1", clinicID).Scan(&storehouseID)
	if errs != nil {
		fmt.Println("errs ===", errs)
		ctx.JSON(iris.Map{"code": "1", "msg": errs.Error()})
		return
	}
	fmt.Println("storehouseID==", storehouseID)

	if manuFactoryID != "" {
		lrow := model.DB.QueryRowx("select id from material where name=$1 and manu_factory_id=$2 limit 1", name, manuFactoryID)
		if lrow == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
			return
		}
		materialProject := FormatSQLRowToMap(lrow)
		_, lok := materialProject["id"]
		if lok {
			ctx.JSON(iris.Map{"code": "1", "msg": "该材料已存在"})
			return
		}
	}

	materialSets := []string{"name", "unit_id"}
	materialValues := []string{"'" + name + "'", unitID}

	clinicMaterialSets := []string{"storehouse_id", "price"}
	clinicMaterialValues := []string{storehouseID, price}

	if enName != "" {
		materialSets = append(materialSets, "en_name")
		materialValues = append(materialValues, "'"+enName+"'")
	}
	if pyCode != "" {
		materialSets = append(materialSets, "py_code")
		materialValues = append(materialValues, "'"+pyCode+"'")
	}
	if idcCode != "" {
		materialSets = append(materialSets, "idc_code")
		materialValues = append(materialValues, "'"+idcCode+"'")
	}
	if remark != "" {
		materialSets = append(materialSets, "remark")
		materialValues = append(materialValues, "'"+remark+"'")
	}
	if manuFactoryID != "" {
		materialSets = append(materialSets, "manu_factory_id")
		materialValues = append(materialValues, manuFactoryID)
	}
	if specification != "" {
		materialSets = append(materialSets, "specification")
		materialValues = append(materialValues, "'"+specification+"'")
	}

	if status != "" {
		clinicMaterialSets = append(clinicMaterialSets, "status")
		clinicMaterialValues = append(clinicMaterialValues, status)
	}
	if cost != "" {
		clinicMaterialSets = append(clinicMaterialSets, "cost")
		clinicMaterialValues = append(clinicMaterialValues, cost)
	}
	if isDiscount != "" {
		clinicMaterialSets = append(clinicMaterialSets, "is_discount")
		clinicMaterialValues = append(clinicMaterialValues, isDiscount)
	}
	if effDay != "" {
		clinicMaterialSets = append(clinicMaterialSets, "eff_day")
		clinicMaterialValues = append(clinicMaterialValues, effDay)
	}
	if stockWarning != "" {
		clinicMaterialSets = append(clinicMaterialSets, "stock_warning")
		clinicMaterialValues = append(clinicMaterialValues, stockWarning)
	}

	materialSetstr := strings.Join(materialSets, ",")
	materialValuestr := strings.Join(materialValues, ",")

	materialInsertSQL := "insert into material (" + materialSetstr + ") values (" + materialValuestr + ") RETURNING id;"
	fmt.Println("materialInsertSQL==", materialInsertSQL)

	tx, err := model.DB.Begin()
	var materialID string
	err = tx.QueryRow(materialInsertSQL).Scan(&materialID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	fmt.Println("materialID====", materialID)

	clinicMaterialSets = append(clinicMaterialSets, "material_id")
	clinicMaterialValues = append(clinicMaterialValues, materialID)

	clinicMaterialSetStr := strings.Join(clinicMaterialSets, ",")
	clinicMaterialValueStr := strings.Join(clinicMaterialValues, ",")

	clinicMaterialInsertSQL := "insert into material_stock (" + clinicMaterialSetStr + ") values (" + clinicMaterialValueStr + ")"
	fmt.Println("clinicMaterialInsertSQL==", clinicMaterialInsertSQL)

	_, err2 := tx.Exec(clinicMaterialInsertSQL)
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

	ctx.JSON(iris.Map{"code": "200", "data": materialID})

}

// MaterialUpdate 更新材料项目
func MaterialUpdate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	materialStockID := ctx.PostValue("material_stock_id")
	materialID := ctx.PostValue("material_id")

	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	pyCode := ctx.PostValue("py_code")
	idcCode := ctx.PostValue("idc_code")
	unitID := ctx.PostValue("unit_id")
	remark := ctx.PostValue("remark")
	manuFactoryID := ctx.PostValue("manu_factory_id")
	specification := ctx.PostValue("specification")

	price := ctx.PostValue("price")
	cost := ctx.PostValue("cost")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")
	effDay := ctx.PostValue("eff_day")
	stockWarning := ctx.PostValue("stock_warning")

	if clinicID == "" || name == "" || materialStockID == "" || price == "" || materialID == "" {
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

	crow := model.DB.QueryRowx("select id,material_id from material_stock where id=$1 limit 1", materialStockID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicMaterialProject := FormatSQLRowToMap(crow)
	_, rok := clinicMaterialProject["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所材料项目数据错误"})
		return
	}
	smaterialID := strconv.FormatInt(clinicMaterialProject["material_id"].(int64), 10)
	fmt.Println("smaterialID====", smaterialID)

	if smaterialID != materialID {
		ctx.JSON(iris.Map{"code": "1", "msg": "材料项目数据id不匹配"})
		return
	}

	lrow := model.DB.QueryRowx("select id from material where name=$1 and id!=$2 limit 1", name, materialID)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	laboratoryItem := FormatSQLRowToMap(lrow)
	_, lok := laboratoryItem["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "1", "msg": "材料项目名称已存在"})
		return
	}

	materialSets := []string{"name='" + name + "'"}
	clinicMaterialSets := []string{"price=" + price}

	if enName != "" {
		materialSets = append(materialSets, "en_name='"+enName+"'")
	}
	if pyCode != "" {
		materialSets = append(materialSets, "py_code='"+pyCode+"'")
	}
	if unitID != "" {
		materialSets = append(materialSets, "unit_id="+unitID)
	}
	if idcCode != "" {
		materialSets = append(materialSets, "idc_code='"+idcCode+"'")
	}
	if remark != "" {
		materialSets = append(materialSets, "remark='"+remark+"'")
	}
	if manuFactoryID != "" {
		materialSets = append(materialSets, "manu_factory_id="+manuFactoryID)
	}
	if specification != "" {
		materialSets = append(materialSets, "specification='"+specification+"'")
	}

	if status != "" {
		clinicMaterialSets = append(clinicMaterialSets, "status="+status)
	}
	if isDiscount != "" {
		clinicMaterialSets = append(clinicMaterialSets, "is_discount="+isDiscount)
	}
	if cost != "" {
		clinicMaterialSets = append(clinicMaterialSets, "cost="+cost)
	}
	if effDay != "" {
		clinicMaterialSets = append(clinicMaterialSets, "eff_day="+effDay)
	}
	if stockWarning != "" {
		clinicMaterialSets = append(clinicMaterialSets, "stock_warning="+stockWarning)
	}

	materialSets = append(materialSets, "updated_time=LOCALTIMESTAMP")
	materialSetstr := strings.Join(materialSets, ",")

	materialUpdateSQL := "update material set " + materialSetstr + " where id=$1"
	fmt.Println("materialUpdateSQL==", materialUpdateSQL)

	tx, err := model.DB.Begin()
	_, err = tx.Exec(materialUpdateSQL, materialID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}

	clinicMaterialSets = append(clinicMaterialSets, "updated_time=LOCALTIMESTAMP")
	clinicMaterialSetStr := strings.Join(clinicMaterialSets, ",")

	clinicMaterialUpdateSQL := "update material_stock set " + clinicMaterialSetStr + " where id=$1"
	fmt.Println("clinicMaterialUpdateSQL==", clinicMaterialUpdateSQL)

	_, err2 := tx.Exec(clinicMaterialUpdateSQL, materialStockID)
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

	ctx.JSON(iris.Map{"code": "200", "data": nil})

}

// MaterialOnOff 启用和停用
func MaterialOnOff(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	materialStockID := ctx.PostValue("material_stock_id")
	status := ctx.PostValue("status")
	if clinicID == "" || materialStockID == "" || status == "" {
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

	crow := model.DB.QueryRowx("select id from material_stock where id=$1 limit 1", materialStockID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicMaterialProject := FormatSQLRowToMap(crow)
	_, rok := clinicMaterialProject["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据错误"})
		return
	}

	_, err1 := model.DB.Exec("update material_stock set status=$1 where id=$2", status, materialStockID)
	if err1 != nil {
		fmt.Println(" err1====", err1)
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// MaterialList 材料缴费项目列表
func MaterialList(ctx iris.Context) {
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

	var storehouseID string
	errs := model.DB.QueryRow("select id from storehouse where clinic_id=$1 limit 1", clinicID).Scan(&storehouseID)
	if errs != nil {
		fmt.Println("errs ===", errs)
		ctx.JSON(iris.Map{"code": "1", "msg": errs.Error()})
		return
	}
	fmt.Println("storehouseID==", storehouseID)

	countSQL := `select count(ms.id) as total from material_stock ms
		left join material m on ms.material_id = m.id
		where ms.storehouse_id=$1`
	selectSQL := `select ms.material_id,ms.id as material_stock_id,m.name,m.unit_id,du.name as unit_name,m.py_code,m.remark,m.idc_code,m.manu_factory_id,m.specification,
		m.en_name,ms.is_discount,ms.price,ms.status,ms.cost,ms.eff_day,ms.stock_warning,ms.stock_amount,mf.name as manu_factory_name
		from material_stock ms
		left join material m on ms.material_id = m.id
		left join dose_unit du on m.unit_id = du.id
		left join manu_factory mf on m.manu_factory_id = mf.id
		where ms.storehouse_id=$1`

	if keyword != "" {
		countSQL += " and m.name ~'" + keyword + "'"
		selectSQL += " and m.name ~'" + keyword + "'"
	}
	if status != "" {
		countSQL += " and ms.status=" + status
		selectSQL += " and ms.status=" + status
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

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})

}

//MaterialDetail 材料项目详情
func MaterialDetail(ctx iris.Context) {
	materialStockID := ctx.PostValue("material_stock_id")

	if materialStockID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select ms.material_id,ms.id as material_stock_id,m.name,m.unit_id,du.name as unit_name,m.py_code,m.remark,m.idc_code,
		m.manu_factory_id,mf.name as manu_factory_name,m.specification,m.en_name,ms.is_discount,ms.price,ms.status,ms.cost,ms.eff_day,ms.stock_warning,ms.stock_amount
		from material_stock ms
		left join material m on ms.material_id = m.id
		left join dose_unit du on m.unit_id = du.id
		left join manu_factory mf on m.manu_factory_id = mf.id
		where ms.id=$1`

	fmt.Println("selectSQL===", selectSQL)

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL, materialStockID)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})
}

//MaterialInstock 入库
func MaterialInstock(ctx iris.Context) {
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
		ctx.JSON(iris.Map{"code": "-1", "msg": errs.Error()})
		return
	}
	orderNumber := "DRKD-" + strconv.FormatInt(time.Now().Unix(), 10)
	values := []string{
		storehouseID,
		"'" + orderNumber + "'",
		instockWayID,
		supplierID,
		"date '" + instockDate + "'",
		operationID,
		"'" + remark + "'"}
	sets := []string{
		"storehouse_id",
		"order_number",
		"instock_way_id",
		"supplier_id",
		"instock_date",
		"instock_operation_id",
		"remark"}

	var itemValues []string
	itemSets := []string{
		"material_id",
		"instock_amount",
		"ret_price",
		"buy_price",
		"serial",
		"eff_day",
		"material_instock_record_id"}

	setStr := strings.Join(sets, ",")
	valueStr := strings.Join(values, ",")
	insertSQL := "insert into material_instock_record (" + setStr + ") values (" + valueStr + ") RETURNING id"
	fmt.Println("insertSQL===", insertSQL)

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errb})
		return
	}

	var materialInstockRecordID string
	errp := tx.QueryRow(insertSQL).Scan(&materialInstockRecordID)
	if errp != nil {
		fmt.Println("errp ===", errp)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": "保存入库错误"})
		return
	}
	fmt.Println("materialInstockRecordID====", materialInstockRecordID)

	for _, v := range results {
		materialID := v["material_id"]
		instockAmount := v["instock_amount"]
		if instockAmount == "" {
			ctx.JSON(iris.Map{"code": "-1", "msg": "数量为必填项"})
			return
		}
		var s []string
		row := model.DB.QueryRowx("select id from material where id=$1 limit 1", materialID)
		if row == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "入库保存失败"})
			return
		}
		material := FormatSQLRowToMap(row)
		_, ok := material["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "入库药品不存在"})
			return
		}
		s = append(s, v["material_id"], v["instock_amount"], v["ret_price"], v["buy_price"], "'"+v["serial"]+"'", v["eff_day"], materialInstockRecordID)

		str := strings.Join(s, ",")
		str = "(" + str + ")"
		itemValues = append(itemValues, str)
	}

	itemSetStr := strings.Join(itemSets, ",")
	itemValueStr := strings.Join(itemValues, ",")
	insertiSQL := "insert into material_instock_record_item (" + itemSetStr + ") values " + itemValueStr
	fmt.Println("insertiSQL====", insertiSQL)

	_, err := tx.Exec(insertiSQL)
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

//MaterialInstockRecord 入库记录列表
func MaterialInstockRecord(ctx iris.Context) {
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

	countSQL := `select count(id) as total from material_instock_record where storehouse_id=$1`
	selectSQL := `select ir.id as material_instock_record_id,ir.instock_date,ir.order_number, iw.name as instock_way_name,
		vp.name as verify_operation_name,s.name as supplier_name,p.name as instock_operation_name,ir.verify_status
		from material_instock_record ir
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

	// var instockRecord []map[string]interface{}
	// for _, v := range results {
	// 	has := false
	// 	for _, vRes := range instockRecord {
	// 		if vRes["order_number"].(string) == v["order_number"].(string) {
	// 			has = true
	// 			continue
	// 		}
	// 	}
	// 	if !has {
	// 		instockRecord = append(instockRecord, v)
	// 	}
	// }
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

//MaterialInstockRecordDetail 入库记录详情
func MaterialInstockRecordDetail(ctx iris.Context) {
	materialInstockRecordID := ctx.PostValue("material_instock_record_id")
	if materialInstockRecordID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	sql := `select ir.id as material_instock_record_id,ir.instock_date,ir.order_number,ir.created_time,s.name as supplier_name,ir.supplier_id,ir.remark,ir.verify_status,
		ir.verify_operation_id,vp.name as verify_operation_name,ir.instock_way_id,iw.name as instock_way_name,ir.instock_operation_id,p.name as instock_operation_name 
		from material_instock_record ir
		left join supplier s on ir.supplier_id = s.id
		left join instock_way iw on ir.instock_way_id = iw.id
		left join personnel p on ir.instock_operation_id = p.id
		left join personnel vp on ir.verify_operation_id = vp.id
		where ir.id=$1`

	arow := model.DB.QueryRowx(sql, materialInstockRecordID)
	result := FormatSQLRowToMap(arow)

	isql := `select d.name as material_name,du.name as unit_name,mf.name as manu_factory_name,iri.instock_amount,
		iri.ret_price,iri.buy_price,iri.serial,iri.eff_day
		from material_instock_record_item iri
		left join material d on iri.material_id = d.id
		left join dose_unit du on d.unit_id = du.id
		left join manu_factory mf on d.manu_factory_id = mf.id
		where iri.material_instock_record_id=$1`

	irows, err := model.DB.Queryx(isql, materialInstockRecordID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	item := FormatSQLRowsToMapArray(irows)
	result["items"] = item
	ctx.JSON(iris.Map{"code": "200", "data": result})
}

//MaterialInstockUpdate 入库记录修改
func MaterialInstockUpdate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	materialInstockRecordID := ctx.PostValue("material_instock_record_id")
	items := ctx.PostValue("items")
	operationID := ctx.PostValue("instock_operation_id")
	instockWayID := ctx.PostValue("instock_way_id")
	supplierID := ctx.PostValue("supplier_id")
	remark := ctx.PostValue("remark")
	instockDate := ctx.PostValue("instock_date")

	if materialInstockRecordID == "" || clinicID == "" || instockWayID == "" || supplierID == "" || instockDate == "" || operationID == "" || items == "" {
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

	row := model.DB.QueryRowx("select * from material_instock_record where id=$1 limit 1", materialInstockRecordID)
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
	if verifyStatus.(string) == "02" {
		ctx.JSON(iris.Map{"code": "1", "msg": "入库记录已审核，不能修改"})
		return
	}

	var values []string
	sets := []string{
		"material_id",
		"instock_amount",
		"ret_price",
		"buy_price",
		"serial",
		"eff_day",
		"material_instock_record_id"}

	for _, v := range results {
		materialID := v["material_id"]
		instockAmount := v["instock_amount"]
		if instockAmount == "" {
			ctx.JSON(iris.Map{"code": "-1", "msg": "数量为必填项"})
			return
		}
		var s []string
		row := model.DB.QueryRowx("select id from material where id=$1 limit 1", materialID)
		if row == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
			return
		}
		material := FormatSQLRowToMap(row)
		_, ok := material["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "修改的药品不存在"})
			return
		}
		s = append(s, v["material_id"], v["instock_amount"], v["ret_price"], v["buy_price"], "'"+v["serial"]+"'", v["eff_day"], materialInstockRecordID)
		str := strings.Join(s, ",")
		str = "(" + str + ")"
		values = append(values, str)
	}
	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errb})
		return
	}

	updateSQL := `update material_instock_record set instock_way_id=$1,supplier_id=$2,instock_date=$3,instock_operation_id=$4,remark=$5 where id=$6`
	_, erru := tx.Exec(updateSQL, instockWayID, supplierID, instockDate, operationID, remark, materialInstockRecordID)
	if erru != nil {
		fmt.Println("erru ===", erru)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": erru.Error()})
		return
	}

	_, errd := tx.Exec("delete from material_instock_record_item where material_instock_record_id=$1", materialInstockRecordID)
	if errd != nil {
		fmt.Println("errd ===", errd)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errd.Error()})
		return
	}

	setStr := strings.Join(sets, ",")
	valueStr := strings.Join(values, ",")
	insertSQL := "insert into material_instock_record_item (" + setStr + ") values " + valueStr
	fmt.Println("insertSQL===", insertSQL)
	_, err := tx.Exec(insertSQL)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": "请检查是否漏填"})
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

//MaterialInstockCheck 入库审核
func MaterialInstockCheck(ctx iris.Context) {
	materialInstockRecordID := ctx.PostValue("material_instock_record_id")
	operationID := ctx.PostValue("verify_operation_id")

	if materialInstockRecordID == "" || operationID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select * from material_instock_record where id=$1 limit 1", materialInstockRecordID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "审核失败"})
		return
	}
	instockRecord := FormatSQLRowToMap(row)

	_, ok := instockRecord["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "入库记录不存在"})
		return
	}
	storehouseID := strconv.FormatInt(instockRecord["storehouse_id"].(int64), 10)
	verifyStatus := instockRecord["verify_status"]
	if verifyStatus.(string) != "01" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "当前状态不能审核"})
		return
	}

	rows, _ := model.DB.Queryx(`select * from material_instock_record_item  where material_instock_record_id=$1 `, materialInstockRecordID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "审核失败"})
		return
	}
	instockRecordItems := FormatSQLRowsToMapArray(rows)
	fmt.Println("====", instockRecordItems)

	sql := "update material_stock set"
	var sets []string
	var asValues []string
	var values []string
	for _, v := range instockRecordItems {
		materialID := strconv.FormatInt(v["material_id"].(int64), 10)
		instockAmount := v["instock_amount"].(int64)
		var s []string
		drow := model.DB.QueryRowx("select id,stock_amount from material_stock where storehouse_id=$1 and material_id=$2 limit 1", storehouseID, materialID)
		if drow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "审核失败"})
			return
		}
		materialStock := FormatSQLRowToMap(drow)
		_, dok := materialStock["id"]
		if !dok {
			ctx.JSON(iris.Map{"code": "1", "msg": "入库失败"})
			return
		}
		stockAmount := materialStock["stock_amount"].(int64) + instockAmount
		s = append(s, storehouseID, materialID, strconv.FormatInt(stockAmount, 10))
		str := strings.Join(s, ",")
		str = "(" + str + ")"
		values = append(values, str)
	}

	valueStr := strings.Join(values, ",")
	sets = append(sets, " stock_amount=tmp.stockAmount", " updated_time=LOCALTIMESTAMP")
	asValues = append(asValues, "storehouse_id", "material_id", "stockAmount")

	setStr := strings.Join(sets, ",")
	asStr := strings.Join(asValues, ",")
	sql += setStr + " from (values " + valueStr + ") as tmp(" + asStr + ") where material_stock.storehouse_id = tmp.storehouse_id and material_stock.material_id = tmp.material_id"
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

	_, err2 := tx.Exec("update material_instock_record set verify_status=$1,verify_operation_id=$2,updated_time=LOCALTIMESTAMP", "02", operationID)
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

//MaterialInstockRecordDelete 删除入库记录
func MaterialInstockRecordDelete(ctx iris.Context) {
	materialInstockRecordID := ctx.PostValue("material_instock_record_id")
	if materialInstockRecordID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	rows := model.DB.QueryRowx("select id,verify_status from material_instock_record where id=$1", materialInstockRecordID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "删除失败"})
		return
	}
	instockRecord := FormatSQLRowToMap(rows)
	_, ok := instockRecord["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "入库记录不存在"})
		return
	}
	verifyStatus := instockRecord["verify_status"]
	if verifyStatus.(string) != "01" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "当前状态不能删除"})
		return
	}

	tx, err := model.DB.Begin()

	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	_, erri := tx.Exec("delete from material_instock_record_item where material_instock_record_id=$1", materialInstockRecordID)
	if erri != nil {
		fmt.Println("erri ===", erri)
		ctx.JSON(iris.Map{"code": "-1", "msg": erri.Error()})
		return
	}

	_, errd := tx.Exec("delete from material_instock_record where id=$1", materialInstockRecordID)
	if errd != nil {
		fmt.Println("errd ===", errd)
		ctx.JSON(iris.Map{"code": "-1", "msg": errd.Error()})
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

//MaterialOutstock 出库
func MaterialOutstock(ctx iris.Context) {
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
		ctx.JSON(iris.Map{"code": "-1", "msg": "出库失败"})
		return
	}
	departmentPersonnel := FormatSQLRowToMap(dprow)
	_, dok := departmentPersonnel["id"]
	if !dok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "领用科室与领用人员不符"})
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
		ctx.JSON(iris.Map{"code": "-1", "msg": errs.Error()})
		return
	}
	orderNumber := "DCKD-" + strconv.FormatInt(time.Now().Unix(), 10)

	values := []string{
		storehouseID,
		departmentID,
		personnelID,
		"'" + orderNumber + "'",
		outstockWayID,
		"date '" + outstockDate + "'",
		operationID,
		"'" + remark + "'"}
	sets := []string{
		"storehouse_id",
		"department_id",
		"personnel_id",
		"order_number",
		"outstock_way_id",
		"outstock_date",
		"outstock_operation_id",
		"remark"}

	var itemValues []string
	itemSets := []string{
		"material_id",
		"outstock_amount",
		"ret_price",
		"buy_price",
		"serial",
		"eff_day",
		"material_outstock_record_id"}

	setStr := strings.Join(sets, ",")
	valueStr := strings.Join(values, ",")
	insertSQL := "insert into material_outstock_record (" + setStr + ") values (" + valueStr + ") RETURNING id"
	fmt.Println("insertSQL===", insertSQL)

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errb})
		return
	}

	var materialOutstockRecordID string
	errp := tx.QueryRow(insertSQL).Scan(&materialOutstockRecordID)
	if errp != nil {
		fmt.Println("errp ===", errp)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": "保存出库错误"})
		return
	}
	fmt.Println("materialOutstockRecordID====", materialOutstockRecordID)

	for _, v := range results {
		materialID := v["material_id"]
		outstockAmount := v["outstock_amount"]
		if outstockAmount == "" {
			ctx.JSON(iris.Map{"code": "-1", "msg": "数量为必填项"})
			return
		}
		var s []string
		row := model.DB.QueryRowx("select id from material where id=$1 limit 1", materialID)
		if row == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "新增出库失败"})
			return
		}
		material := FormatSQLRowToMap(row)
		_, ok := material["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "新增出库药品不存在"})
			return
		}
		s = append(s, v["material_id"], v["outstock_amount"],
			v["ret_price"], v["buy_price"], "'"+v["serial"]+"'", v["eff_day"], materialOutstockRecordID)
		str := strings.Join(s, ",")
		str = "(" + str + ")"
		itemValues = append(itemValues, str)
	}

	itemSetStr := strings.Join(itemSets, ",")
	itemValueStr := strings.Join(itemValues, ",")
	insertiSQL := "insert into material_outstock_record_item (" + itemSetStr + ") values " + itemValueStr
	fmt.Println("insertiSQL===", insertiSQL)
	_, err := tx.Exec(insertiSQL)
	if err != nil {
		fmt.Println("err ===", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": "请检查是否漏填"})
		return
	}

	err3 := tx.Commit()
	if err3 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//MaterialOutstockRecord 出库记录列表
func MaterialOutstockRecord(ctx iris.Context) {
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
		ctx.JSON(iris.Map{"code": "-1", "msg": errs.Error()})
		return
	}

	countSQL := `select count(id) as total from material_outstock_record where storehouse_id=$1`
	selectSQL := `select outr.id as material_outstock_record_id,outr.outstock_date,outr.order_number, ow.name as outstock_way_name,
		vp.name as verify_operation_name,d.name as department_name,p.name as personnel_name,op.name as outstock_operation_name,outr.verify_status
		from material_outstock_record outr
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

	// var outstockRecord []map[string]interface{}
	// for _, v := range results {
	// 	has := false
	// 	for _, vRes := range outstockRecord {
	// 		if vRes["order_number"].(string) == v["order_number"].(string) {
	// 			has = true
	// 			continue
	// 		}
	// 	}
	// 	if !has {
	// 		outstockRecord = append(outstockRecord, v)
	// 	}
	// }
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

//MaterialOutstockRecordDetail 出库记录详情
func MaterialOutstockRecordDetail(ctx iris.Context) {
	materialOutstockRecordID := ctx.PostValue("material_outstock_record_id")
	if materialOutstockRecordID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	sql := `select outr.id as material_outstock_record_id,outr.outstock_date,outr.order_number,outr.created_time,
		dept.name as department_name,outr.department_id,outr.remark,outr.verify_operation_id,vp.name as verify_operation_name,
		outr.personnel_id,p.name as personnel_name,outr.outstock_way_id,ow.name as outstock_way_name,
		outr.outstock_operation_id,op.name as outstock_operation_name 
		from material_outstock_record outr
		left join department dept on outr.department_id = dept.id
		left join personnel p on outr.personnel_id = p.id
		left join outstock_way ow on outr.outstock_way_id = ow.id
		left join personnel op on outr.outstock_operation_id = op.id
		left join personnel vp on outr.verify_operation_id = vp.id
		where outr.id=$1`

	row := model.DB.QueryRowx(sql, materialOutstockRecordID)
	result := FormatSQLRowToMap(row)

	isql := `select d.name as material_name,ori.material_id,du.name as unit_name,mf.name as manu_factory_name,ori.outstock_amount,
		ori.ret_price,ori.buy_price,ori.serial,ori.eff_day
		from material_outstock_record_item ori
		left join material d on ori.material_id = d.id
		left join dose_unit du on d.unit_id = du.id
		left join manu_factory mf on d.manu_factory_id = mf.id
		where ori.material_outstock_record_id=$1`

	irows, err := model.DB.Queryx(isql, materialOutstockRecordID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	item := FormatSQLRowsToMapArray(irows)
	result["items"] = item
	ctx.JSON(iris.Map{"code": "200", "data": result})
}

//MaterialOutstockUpdate 出库记录修改
func MaterialOutstockUpdate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	materialOutstockRecordID := ctx.PostValue("material_outstock_record_id")
	items := ctx.PostValue("items")
	operationID := ctx.PostValue("outstock_operation_id")
	outstockWayID := ctx.PostValue("outstock_way_id")
	departmentID := ctx.PostValue("department_id")
	personnelID := ctx.PostValue("personnel_id")
	remark := ctx.PostValue("remark")
	outstockDate := ctx.PostValue("outstock_date")

	if clinicID == "" || materialOutstockRecordID == "" || items == "" || outstockWayID == "" || departmentID == "" || operationID == "" || personnelID == "" || outstockDate == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	dprow := model.DB.QueryRowx("select id from department_personnel where department_id=$1 and personnel_id!=$2 limit 1", departmentID, personnelID)
	if dprow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "出库失败"})
		return
	}
	departmentPersonnel := FormatSQLRowToMap(dprow)
	_, dok := departmentPersonnel["id"]
	if !dok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "领用科室与领用人员不符"})
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

	row := model.DB.QueryRowx("select * from material_outstock_record where id=$1 and storehouse_id=$2 limit 1", materialOutstockRecordID, storehouseID)
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
	if verifyStatus.(string) == "02" {
		ctx.JSON(iris.Map{"code": "1", "msg": "出库记录已审核，不能修改"})
		return
	}

	var values []string
	sets := []string{
		"material_id",
		"outstock_amount",
		"ret_price",
		"buy_price",
		"serial",
		"eff_day",
		"material_outstock_record_id"}

	for _, v := range results {
		materialID := v["material_id"]
		var s []string
		row := model.DB.QueryRowx("select id from material where id=$1 limit 1", materialID)
		if row == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
			return
		}
		material := FormatSQLRowToMap(row)
		_, ok := material["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "修改的出库药品不存在"})
			return
		}
		s = append(s, v["material_id"], v["outstock_amount"], v["ret_price"], v["buy_price"], "'"+v["serial"]+"'", v["eff_day"], materialOutstockRecordID)
		str := strings.Join(s, ",")
		str = "(" + str + ")"
		values = append(values, str)
	}

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errb})
		return
	}

	updateSQL := `update material_outstock_record set department_id=$1,personnel_id=$2,outstock_way_id=$3,outstock_date=$4,outstock_operation_id=$5,remark=$6 where id=$7`
	_, erru := tx.Exec(updateSQL, departmentID, personnelID, outstockWayID, outstockDate, operationID, remark, materialOutstockRecordID)
	if erru != nil {
		fmt.Println("erru ===", erru)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": erru.Error()})
		return
	}
	_, errd := tx.Exec("delete from material_outstock_record_item where material_outstock_record_id=$1", materialOutstockRecordID)
	if errd != nil {
		fmt.Println("errd ===", errd)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errd.Error()})
		return
	}

	setStr := strings.Join(sets, ",")
	valueStr := strings.Join(values, ",")
	insertSQL := "insert into material_outstock_record_item (" + setStr + ") values " + valueStr
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

//MaterialOutstockCheck 出库审核
func MaterialOutstockCheck(ctx iris.Context) {
	materialOutstockRecordID := ctx.PostValue("material_outstock_record_id")
	operationID := ctx.PostValue("verify_operation_id")
	if materialOutstockRecordID == "" || operationID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select * from material_outstock_record where id=$1 limit 1", materialOutstockRecordID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "审核失败"})
		return
	}
	OutstockRecord := FormatSQLRowToMap(row)
	_, ok := OutstockRecord["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "入库记录不存在"})
		return
	}
	storehouseID := strconv.FormatInt(OutstockRecord["storehouse_id"].(int64), 10)
	verifyStatus := OutstockRecord["verify_status"]
	if verifyStatus.(string) != "01" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "当前状态不能审核"})
		return
	}

	rows, _ := model.DB.Queryx(`select * from material_outstock_record_item  where material_outstock_record_id=$1`, materialOutstockRecordID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "审核失败"})
		return
	}
	outstockRecordItems := FormatSQLRowsToMapArray(rows)

	sql := "update material_stock set"

	var sets []string
	var asValues []string
	var values []string
	for _, v := range outstockRecordItems {
		materialID := strconv.FormatInt(v["material_id"].(int64), 10)
		outstockAmount := v["outstock_amount"].(int64)
		var s []string
		drow := model.DB.QueryRowx("select id,stock_amount from material_stock where storehouse_id=$1 and material_id=$2 limit 1", storehouseID, materialID)
		if drow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "审核失败"})
			return
		}
		materialStock := FormatSQLRowToMap(drow)
		_, dok := materialStock["id"]
		if !dok {
			ctx.JSON(iris.Map{"code": "1", "msg": "入库失败"})
			return
		}
		stockAmount := materialStock["stock_amount"].(int64) - outstockAmount

		s = append(s, storehouseID, materialID, strconv.FormatInt(stockAmount, 10))
		str := strings.Join(s, ",")
		str = "(" + str + ")"
		values = append(values, str)
	}
	valueStr := strings.Join(values, ",")

	sets = append(sets, " stock_amount=tmp.stock_amount", " updated_time=LOCALTIMESTAMP")
	asValues = append(asValues, "storehouse_id", "material_id", "stock_amount")

	setStr := strings.Join(sets, ",")
	asStr := strings.Join(asValues, ",")
	sql += setStr + " from (values " + valueStr + ") as tmp(" + asStr + ") where material_stock.storehouse_id = tmp.storehouse_id and material_stock.material_id = tmp.material_id"
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

	_, err2 := tx.Exec("update material_outstock_record set verify_status=$1,verify_operation_id=$2,updated_time=LOCALTIMESTAMP", "02", operationID)
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

//MaterialOutstockRecordDelete 删除出库记录
func MaterialOutstockRecordDelete(ctx iris.Context) {
	materialOutstockRecordID := ctx.PostValue("material_outstock_record_id")
	if materialOutstockRecordID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	rows := model.DB.QueryRowx("select id,verify_status from material_outstock_record where id=$1", materialOutstockRecordID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "删除失败"})
		return
	}
	instockRecord := FormatSQLRowToMap(rows)
	_, ok := instockRecord["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "出库库记录不存在"})
		return
	}
	verifyStatus := instockRecord["verify_status"]
	if verifyStatus.(string) != "01" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "当前状态不能删除"})
		return
	}

	tx, err := model.DB.Begin()

	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	_, erri := tx.Exec("delete from material_outstock_record_item where material_outstock_record_id=$1", materialOutstockRecordID)
	if erri != nil {
		fmt.Println("erri ===", erri)
		ctx.JSON(iris.Map{"code": "-1", "msg": erri.Error()})
		return
	}

	_, errd := tx.Exec("delete from material_outstock_record where id=$1", materialOutstockRecordID)
	if errd != nil {
		fmt.Println("errd ===", errd)
		ctx.JSON(iris.Map{"code": "-1", "msg": errd.Error()})
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
