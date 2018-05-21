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
		"clinic_material_id",
		"instock_amount",
		"buy_price",
		"serial",
		"eff_date",
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
		clinicMaterialID := v["clinic_material_id"]
		instockAmount := v["instock_amount"]
		if instockAmount == "" {
			ctx.JSON(iris.Map{"code": "-1", "msg": "数量为必填项"})
			return
		}
		var s []string
		row := model.DB.QueryRowx("select id from clinic_material where id=$1 limit 1", clinicMaterialID)
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
		s = append(s, v["clinic_material_id"], v["instock_amount"], v["buy_price"], "'"+v["serial"]+"'", "date '"+v["eff_date"]+"'", materialInstockRecordID)

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

	isql := `select m.name as material_name,m.packing_unit_name,m.manu_factory_name,iri.instock_amount,
		iri.buy_price,iri.serial,iri.eff_date,cm.ret_price
		from material_instock_record_item iri
		left join clinic_material cm on iri.clinic_material_id = cm.id
		left join material m on cm.material_id = m.id
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
		"clinic_material_id",
		"instock_amount",
		"buy_price",
		"serial",
		"eff_date",
		"material_instock_record_id"}

	for _, v := range results {
		clinicMaterialID := v["clinic_material_id"]
		instockAmount := v["instock_amount"]
		if instockAmount == "" {
			ctx.JSON(iris.Map{"code": "-1", "msg": "数量为必填项"})
			return
		}
		var s []string
		row := model.DB.QueryRowx("select id from clinic_material where id=$1 limit 1", clinicMaterialID)
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
		s = append(s, v["clinic_material_id"], v["instock_amount"], v["buy_price"], "'"+v["serial"]+"'", "date '"+v["eff_date"]+"'", materialInstockRecordID)
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

	rows, _ := model.DB.Queryx(`select iri.*,ir.supplier_id from material_instock_record_item iri
		left join material_instock_record ir on ir.id = iri.material_instock_record_id
		where iri.material_instock_record_id=$1 `, materialInstockRecordID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "审核失败"})
		return
	}
	instockRecordItems := FormatSQLRowsToMapArray(rows)
	fmt.Println("====", instockRecordItems)

	tx, err := model.DB.Begin()

	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}

	for _, v := range instockRecordItems {
		clinicMaterialID := strconv.FormatInt(v["clinic_material_id"].(int64), 10)
		instockAmount := v["instock_amount"].(int64)
		supplierID := strconv.FormatInt(v["supplier_id"].(int64), 10)
		serial := v["serial"].(string)
		buyPrice := strconv.FormatInt(v["buy_price"].(int64), 10)
		effDate := v["eff_date"].(string)
		var s []string
		drow := model.DB.QueryRowx(`select id,stock_amount from material_stock where storehouse_id=$1 and clinic_material_id=$2 
			and supplier_id=$3 and serial=$4 and eff_date=$5 limit 1`, storehouseID, clinicMaterialID, supplierID, serial, effDate)
		if drow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "审核失败"})
			return
		}
		materialStock := FormatSQLRowToMap(drow)
		s = append(s, storehouseID, clinicMaterialID, supplierID, serial, effDate, buyPrice)
		_, dok := materialStock["id"]
		if !dok {
			s = append(s, strconv.FormatInt(instockAmount, 10))
			sets := []string{
				"storehouse_id",
				"clinic_material_id",
				"supplier_id",
				"serial",
				"eff_date",
				"buy_price",
				"stock_amount"}
			setStr := strings.Join(sets, ",")
			valueStr := strings.Join(s, ",")
			insertSQL := "insert into material_stock (" + setStr + ") values (" + valueStr + ")"
			_, err1 := tx.Exec(insertSQL)
			if err1 != nil {
				fmt.Println("err1 ===", err1)
				tx.Rollback()
				ctx.JSON(iris.Map{"code": "1", "msg": err1.Error()})
				return
			}
		} else {
			materialStockID := strconv.FormatInt(materialStock["id"].(int64), 10)
			stockAmount := materialStock["stock_amount"].(int64) + instockAmount
			sets := append(s, "stock_amount="+strconv.FormatInt(stockAmount, 10), "buy_price="+buyPrice, "updated_time=LOCALTIMESTAMP")
			setStr := strings.Join(sets, ",")
			updateSQL := "update material_instock set " + setStr + "where id=$1"
			fmt.Println("updateSQL===", updateSQL)
			_, err2 := tx.Exec(updateSQL, materialStockID)
			if err2 != nil {
				fmt.Println("err2 ===", err2)
				tx.Rollback()
				ctx.JSON(iris.Map{"code": "1", "msg": err2.Error()})
				return
			}
		}
	}

	_, err3 := tx.Exec("update material_instock_record set verify_status=$1,verify_operation_id=$2,updated_time=LOCALTIMESTAMP", "02", operationID)
	if err3 != nil {
		fmt.Println("err3 ===", err3)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err3.Error()})
		return
	}

	err4 := tx.Commit()
	if err4 != nil {
		fmt.Println("err4 ===", err4)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err4.Error()})
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
		"clinic_material_id",
		"outstock_amount",
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
		clinicMaterialID := v["clinic_material_id"]
		outstockAmount := v["outstock_amount"]
		if outstockAmount == "" {
			ctx.JSON(iris.Map{"code": "-1", "msg": "数量为必填项"})
			return
		}
		var s []string
		row := model.DB.QueryRowx("select id from clinic_material where id=$1 limit 1", clinicMaterialID)
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
		s = append(s, v["clinic_material_id"], v["outstock_amount"], materialOutstockRecordID)
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
		vp.name as verify_operation_name,m.name as department_name,p.name as personnel_name,op.name as outstock_operation_name,outr.verify_status
		from material_outstock_record outr
		left join department m on outr.department_id = m.id
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

	isql := `select m.name as material_name,ori.clinic_material_id,m.packing_unit_name,m.manu_factory_name,ori.outstock_amount,
		cm.ret_price,ms.buy_price,ms.serial,ms.eff_date
		from material_outstock_record_item ori
		left join material_stock ms on ori.material_stock_id = ms.id
		left join clinic_material cm on ms.clinic_material_id = cm.id
		left join material m on cm.material_id = m.id
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
		"clinic_material_id",
		"outstock_amount",
		"material_outstock_record_id"}

	for _, v := range results {
		clinicMaterialID := v["clinic_material_id"]
		var s []string
		row := model.DB.QueryRowx("select id from clinic_material where id=$1 limit 1", clinicMaterialID)
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
		s = append(s, v["clinic_material_id"], v["outstock_amount"], materialOutstockRecordID)
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

	row := model.DB.QueryRowx(`select * from material_outstock_record where id=$1 limit 1`, materialOutstockRecordID)
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
	verifyStatus := OutstockRecord["verify_status"]
	if verifyStatus.(string) != "01" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "当前状态不能审核"})
		return
	}

	rows, _ := model.DB.Queryx(`select ori.* from material_outstock_record_item ori
		left join material_outstock_record or on ori.material_outstock_record_id = or.id
		where ori.material_outstock_record_id=$1`, materialOutstockRecordID)
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
		materialStockID := strconv.FormatInt(v["material_stock_id"].(int64), 10)
		outstockAmount := v["outstock_amount"].(int64)
		var s []string
		drow := model.DB.QueryRowx("select id,stock_amount from material_stock where id=$1 limit 1", materialStockID)
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

		s = append(s, materialStockID, strconv.FormatInt(stockAmount, 10))
		str := strings.Join(s, ",")
		str = "(" + str + ")"
		values = append(values, str)
	}
	valueStr := strings.Join(values, ",")

	sets = append(sets, " stock_amount=tmp.stock_amount", " updated_time=LOCALTIMESTAMP")
	asValues = append(asValues, "id", "stock_amount")

	setStr := strings.Join(sets, ",")
	asStr := strings.Join(asValues, ",")
	sql += setStr + " from (values " + valueStr + ") as tmp(" + asStr + ") where material_stock.id = tmp.id"
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
