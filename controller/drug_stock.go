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

//DrugInstock 入库
func DrugInstock(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	items := ctx.PostValue("items")
	operationID := ctx.PostValue("instock_operation_id")
	instockWayName := ctx.PostValue("instock_way_name")
	supplierName := ctx.PostValue("supplier_name")
	remark := ctx.PostValue("remark")
	instockDate := ctx.PostValue("instock_date")

	if clinicID == "" || instockWayName == "" || supplierName == "" || instockDate == "" || operationID == "" || items == "" {
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
	sets := []string{
		"storehouse_id",
		"order_number",
		"instock_way_name",
		"supplier_name",
		"instock_date",
		"instock_operation_id",
		"remark"}

	itemSets := []string{
		"clinic_drug_id",
		"instock_amount",
		"buy_price",
		"serial",
		"eff_date",
		"drug_instock_record_id"}

	setStr := strings.Join(sets, ",")
	itemSetStr := strings.Join(itemSets, ",")

	insertSQL := "insert into drug_instock_record (" + setStr + ") values ($1,$2,$3,$4,$5,$6,$7) RETURNING id"

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errb})
		return
	}

	var drugInstockRecordID string
	errp := tx.QueryRow(insertSQL,
		ToNullInt64(storehouseID),
		ToNullString(orderNumber),
		ToNullString(instockWayName),
		ToNullString(supplierName),
		ToNullString(instockDate),
		ToNullInt64(operationID),
		ToNullString(remark),
	).Scan(&drugInstockRecordID)
	if errp != nil {
		fmt.Println("errp ===", errp)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": "保存入库错误"})
		return
	}
	fmt.Println("drugInstockRecordID====", drugInstockRecordID)

	for _, v := range results {
		clinicDrugID := v["clinic_drug_id"]
		instockAmount := v["instock_amount"]
		buyPrice := v["buy_price"]
		serial := v["serial"]
		effDate := v["eff_date"]
		if instockAmount == "" {
			ctx.JSON(iris.Map{"code": "-1", "msg": "数量为必填项"})
			return
		}
		row := model.DB.QueryRowx("select id from clinic_drug where id=$1 limit 1", clinicDrugID)
		if row == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "入库保存失败"})
			return
		}
		clinicDrug := FormatSQLRowToMap(row)
		_, ok := clinicDrug["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "入库药品不存在"})
			return
		}

		insertiSQL := "insert into drug_instock_record_item (" + itemSetStr + ") values ($1,$2,$3,$4,$5,$6)"

		_, err := tx.Exec(insertiSQL,
			ToNullInt64(clinicDrugID),
			ToNullInt64(instockAmount),
			ToNullInt64(buyPrice),
			ToNullString(serial),
			ToNullString(effDate),
			ToNullInt64(drugInstockRecordID),
		)
		if err != nil {
			fmt.Println("err ===", err)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "1", "msg": "请检查是否漏填"})
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

//DrugInstockRecord 入库记录列表
func DrugInstockRecord(ctx iris.Context) {
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

	countSQL := `select count(id) as total from drug_instock_record where storehouse_id=:storehouse_id`
	selectSQL := `select ir.id as drug_instock_record_id,ir.instock_date,ir.order_number, ir.instock_way_name,
		vp.name as verify_operation_name,ir.supplier_name,p.name as instock_operation_name,ir.verify_status
		from drug_instock_record ir
		left join personnel p on ir.instock_operation_id = p.id
		left join personnel vp on ir.verify_operation_id = vp.id
		where storehouse_id=:storehouse_id`

	if startDate != "" && endDate != "" {
		if startDate > endDate {
			ctx.JSON(iris.Map{"code": "-1", "msg": "开始日期必须大于结束日期"})
			return
		}
		countSQL += " and instock_date between :start_date and :end_date"
		selectSQL += " and ir.instock_date between :start_date and :end_date"
	}

	if orderNumber != "" {
		countSQL += " and order_number=:order_number"
		selectSQL += " and ir.order_number=:order_number"
	}

	var queryOption = map[string]interface{}{
		"storehouse_id": ToNullInt64(storehouseID),
		"start_date":    ToNullString(startDate),
		"end_date":      ToNullString(endDate),
		"order_number":  ToNullString(orderNumber),
		"offset":        ToNullInt64(offset),
		"limit":         ToNullInt64(limit),
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

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

//DrugInstockRecordDetail 入库记录详情
func DrugInstockRecordDetail(ctx iris.Context) {
	drugInstockRecordID := ctx.PostValue("drug_instock_record_id")
	if drugInstockRecordID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	sql := `select ir.id as drug_instock_record_id,ir.instock_date,ir.order_number,ir.created_time,ir.supplier_name,ir.remark,ir.verify_status,
		ir.verify_operation_id,vp.name as verify_operation_name,ir.instock_way_name,ir.instock_operation_id,p.name as instock_operation_name 
		from drug_instock_record ir
		left join personnel p on ir.instock_operation_id = p.id
		left join personnel vp on ir.verify_operation_id = vp.id
		where ir.id=$1`

	arow := model.DB.QueryRowx(sql, drugInstockRecordID)
	result := FormatSQLRowToMap(arow)

	isql := `select iri.clinic_drug_id,cd.name as drug_name,cd.packing_unit_name,cd.manu_factory_name,iri.instock_amount,
		iri.buy_price,iri.serial,iri.eff_date,cd.ret_price
		from drug_instock_record_item iri
		left join clinic_drug cd on iri.clinic_drug_id = cd.id
		where iri.drug_instock_record_id=$1`

	irows, err := model.DB.Queryx(isql, drugInstockRecordID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	item := FormatSQLRowsToMapArray(irows)
	result["items"] = item
	ctx.JSON(iris.Map{"code": "200", "data": result})
}

//DrugInstockUpdate 入库记录修改
func DrugInstockUpdate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	drugInstockRecordID := ctx.PostValue("drug_instock_record_id")
	items := ctx.PostValue("items")
	operationID := ctx.PostValue("instock_operation_id")
	instockWayName := ctx.PostValue("instock_way_name")
	supplierName := ctx.PostValue("supplier_name")
	remark := ctx.PostValue("remark")
	instockDate := ctx.PostValue("instock_date")

	if drugInstockRecordID == "" || clinicID == "" || instockWayName == "" || supplierName == "" || instockDate == "" || operationID == "" || items == "" {
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

	row := model.DB.QueryRowx("select * from drug_instock_record where id=$1 limit 1", drugInstockRecordID)
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

	sets := []string{
		"clinic_drug_id",
		"instock_amount",
		"buy_price",
		"serial",
		"eff_date",
		"drug_instock_record_id"}

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errb})
		return
	}

	updateSQL := `update drug_instock_record set instock_way_name=$1,supplier_name=$2,instock_date=$3,instock_operation_id=$4,remark=$5 where id=$6`
	_, erru := tx.Exec(updateSQL, instockWayName, supplierName, instockDate, operationID, remark, drugInstockRecordID)
	if erru != nil {
		fmt.Println("erru ===", erru)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": erru.Error()})
		return
	}

	_, errd := tx.Exec("delete from drug_instock_record_item where drug_instock_record_id=$1", drugInstockRecordID)
	if errd != nil {
		fmt.Println("errd ===", errd)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errd.Error()})
		return
	}

	setStr := strings.Join(sets, ",")

	for _, v := range results {
		clinicDrugID := v["clinic_drug_id"]
		instockAmount := v["instock_amount"]
		buyPrice := v["buy_price"]
		serial := v["serial"]
		effDate := v["eff_date"]
		if instockAmount == "" {
			ctx.JSON(iris.Map{"code": "-1", "msg": "数量为必填项"})
			return
		}
		row := model.DB.QueryRowx("select id from clinic_drug where id=$1 limit 1", clinicDrugID)
		if row == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "入库保存失败"})
			return
		}
		clinicDrug := FormatSQLRowToMap(row)
		_, ok := clinicDrug["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "入库药品不存在"})
			return
		}

		insertiSQL := "insert into drug_instock_record_item (" + setStr + ") values ($1,$2,$3,$4,$5,$6)"

		_, err := tx.Exec(insertiSQL,
			ToNullInt64(clinicDrugID),
			ToNullInt64(instockAmount),
			ToNullInt64(buyPrice),
			ToNullString(serial),
			ToNullString(effDate),
			ToNullInt64(drugInstockRecordID),
		)
		if err != nil {
			fmt.Println("err ===", err)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "1", "msg": "请检查是否漏填"})
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

//DrugInstockCheck 入库审核
func DrugInstockCheck(ctx iris.Context) {
	drugInstockRecordID := ctx.PostValue("drug_instock_record_id")
	operationID := ctx.PostValue("verify_operation_id")

	if drugInstockRecordID == "" || operationID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select * from drug_instock_record where id=$1 limit 1", drugInstockRecordID)
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
	storehouseID := instockRecord["storehouse_id"].(int64)
	verifyStatus := instockRecord["verify_status"]
	if verifyStatus.(string) != "01" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "当前状态不能审核"})
		return
	}

	rows, _ := model.DB.Queryx(`select iri.*,ir.supplier_name from drug_instock_record_item iri
		left join drug_instock_record ir on ir.id = iri.drug_instock_record_id
		where iri.drug_instock_record_id=$1 `, drugInstockRecordID)
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

	sets := []string{
		"storehouse_id",
		"clinic_drug_id",
		"supplier_name",
		"serial",
		"eff_date",
		"buy_price",
		"stock_amount"}
	setStr := strings.Join(sets, ",")

	for _, v := range instockRecordItems {
		clinicDrugID := v["clinic_drug_id"].(int64)
		instockAmount := v["instock_amount"].(int64)
		supplierName := v["supplier_name"].(string)
		serial := v["serial"].(string)
		buyPrice := v["buy_price"].(int64)
		effDate := v["eff_date"].(time.Time).Format("2006-01-02")

		drow := model.DB.QueryRowx(`select id,stock_amount from drug_stock where storehouse_id=$1 and clinic_drug_id=$2 
			and supplier_name=$3 and serial=$4 and eff_date=$5 limit 1`, storehouseID, clinicDrugID, supplierName, serial, effDate)
		if drow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "审核失败"})
			return
		}
		drugStock := FormatSQLRowToMap(drow)
		_, dok := drugStock["id"]
		if !dok {
			insertSQL := "insert into drug_stock (" + setStr + ") values ($1,$2,$3,$4,$5,$6,$7)"
			_, err1 := tx.Exec(insertSQL,
				storehouseID,
				clinicDrugID,
				ToNullString(supplierName),
				ToNullString(serial),
				ToNullString(effDate),
				buyPrice,
				instockAmount,
			)
			if err1 != nil {
				fmt.Println("err1 ===", err1)
				tx.Rollback()
				ctx.JSON(iris.Map{"code": "1", "msg": err1.Error()})
				return
			}
		} else {
			drugStockID := drugStock["id"].(int64)
			stockAmount := drugStock["stock_amount"].(int64) + instockAmount
			updateSQL := "update drug_stock set stock_amount=$1,buy_price=$2,updated_time=LOCALTIMESTAMP where id=$3"
			fmt.Println("updateSQL===", updateSQL)
			_, err2 := tx.Exec(updateSQL,
				stockAmount,
				buyPrice,
				drugStockID,
			)
			if err2 != nil {
				fmt.Println("err2 ===", err2)
				tx.Rollback()
				ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
				return
			}
		}
	}

	_, err3 := tx.Exec("update drug_instock_record set verify_status=$1,verify_operation_id=$2,updated_time=LOCALTIMESTAMP where id=$3", "02", operationID, drugInstockRecordID)
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

//DrugInstockRecordDelete 删除入库记录
func DrugInstockRecordDelete(ctx iris.Context) {
	drugInstockRecordID := ctx.PostValue("drug_instock_record_id")
	if drugInstockRecordID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	rows := model.DB.QueryRowx("select id,verify_status from drug_instock_record where id=$1", drugInstockRecordID)
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
	_, erri := tx.Exec("delete from drug_instock_record_item where drug_instock_record_id=$1", drugInstockRecordID)
	if erri != nil {
		fmt.Println("erri ===", erri)
		ctx.JSON(iris.Map{"code": "-1", "msg": erri.Error()})
		return
	}

	_, errd := tx.Exec("delete from drug_instock_record where id=$1", drugInstockRecordID)
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

//DrugOutstock 出库
func DrugOutstock(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	items := ctx.PostValue("items")
	operationID := ctx.PostValue("outstock_operation_id")
	outstockWayName := ctx.PostValue("outstock_way_name")
	departmentID := ctx.PostValue("department_id")
	personnelID := ctx.PostValue("personnel_id")
	remark := ctx.PostValue("remark")
	outstockDate := ctx.PostValue("outstock_date")

	if clinicID == "" || items == "" || outstockWayName == "" || departmentID == "" || operationID == "" || personnelID == "" || outstockDate == "" {
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

	sets := []string{
		"storehouse_id",
		"department_id",
		"personnel_id",
		"order_number",
		"outstock_way_name",
		"outstock_date",
		"outstock_operation_id",
		"remark"}

	itemSets := []string{
		"drug_stock_id",
		"outstock_amount",
		"drug_outstock_record_id"}

	setStr := strings.Join(sets, ",")
	insertSQL := "insert into drug_outstock_record (" + setStr + ") values ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id"

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errb})
		return
	}

	var drugOutstockRecordID string
	errp := tx.QueryRow(insertSQL,
		ToNullInt64(storehouseID),
		ToNullInt64(departmentID),
		ToNullInt64(personnelID),
		ToNullString(orderNumber),
		ToNullString(outstockWayName),
		ToNullString(outstockDate),
		ToNullInt64(operationID),
		ToNullString(remark),
	).Scan(&drugOutstockRecordID)
	if errp != nil {
		fmt.Println("errp ===", errp)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": "保存出库错误"})
		return
	}
	fmt.Println("drugOutstockRecordID====", drugOutstockRecordID)
	itemSetStr := strings.Join(itemSets, ",")

	for _, v := range results {
		drugStockID := v["drug_stock_id"]
		outstockAmount := v["outstock_amount"]
		if outstockAmount == "" {
			ctx.JSON(iris.Map{"code": "-1", "msg": "数量为必填项"})
			return
		}
		row := model.DB.QueryRowx("select id from drug_stock where id=$1 limit 1", drugStockID)
		if row == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "新增出库失败"})
			return
		}
		drugStock := FormatSQLRowToMap(row)
		_, ok := drugStock["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1-", "msg": "新增出库药品不存在"})
			return
		}
		insertiSQL := "insert into drug_outstock_record_item (" + itemSetStr + ") values ($1,$2,$3)"
		fmt.Println("insertiSQL===", insertiSQL)
		_, err := tx.Exec(insertiSQL,
			ToNullInt64(drugStockID),
			ToNullInt64(outstockAmount),
			ToNullInt64(drugOutstockRecordID),
		)
		if err != nil {
			fmt.Println("err ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": "请检查是否漏填"})
			return
		}
	}

	err3 := tx.Commit()
	if err3 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//DrugOutstockRecord 出库记录列表
func DrugOutstockRecord(ctx iris.Context) {
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

	countSQL := `select count(id) as total from drug_outstock_record where storehouse_id=:storehouse_id`
	selectSQL := `select outr.id as drug_outstock_record_id,outr.outstock_date,outr.order_number, outr.outstock_way_name,
		vp.name as verify_operation_name,d.name as department_name,p.name as personnel_name,op.name as outstock_operation_name,outr.verify_status
		from drug_outstock_record outr
		left join department d on outr.department_id = d.id
		left join personnel p on outr.personnel_id = p.id
		left join personnel op on outr.outstock_operation_id = op.id
		left join personnel vp on outr.verify_operation_id = vp.id
		where storehouse_id=:storehouse_id`

	if startDate != "" && endDate != "" {
		if startDate > endDate {
			ctx.JSON(iris.Map{"code": "-1", "msg": "开始日期必须大于结束日期"})
			return
		}
		countSQL += " and outstock_date between :start_date and :end_date"
		selectSQL += " and outr.outstock_date between :start_date and :end_date"
	}

	if orderNumber != "" {
		countSQL += " and order_number=:order_number"
		selectSQL += " and outr.order_number=:order_number"
	}

	var queryOption = map[string]interface{}{
		"storehouse_id": ToNullInt64(storehouseID),
		"start_date":    ToNullString(startDate),
		"end_date":      ToNullString(endDate),
		"order_number":  ToNullString(orderNumber),
		"offset":        ToNullInt64(offset),
		"limit":         ToNullInt64(limit),
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

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

//DrugOutstockRecordDetail 出库记录详情
func DrugOutstockRecordDetail(ctx iris.Context) {
	drugOutstockRecordID := ctx.PostValue("drug_outstock_record_id")
	if drugOutstockRecordID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	sql := `select outr.id as drug_outstock_record_id,outr.outstock_date,outr.order_number,outr.created_time,
		dept.name as department_name,outr.department_id,outr.remark,outr.verify_operation_id,vp.name as verify_operation_name,
		outr.personnel_id,p.name as personnel_name,outr.outstock_way_name,outr.verify_status,
		outr.outstock_operation_id,op.name as outstock_operation_name 
		from drug_outstock_record outr
		left join department dept on outr.department_id = dept.id
		left join personnel p on outr.personnel_id = p.id
		left join personnel op on outr.outstock_operation_id = op.id
		left join personnel vp on outr.verify_operation_id = vp.id
		where outr.id=$1`

	row := model.DB.QueryRowx(sql, drugOutstockRecordID)
	result := FormatSQLRowToMap(row)

	isql := `select ori.drug_stock_id,cd.name as drug_name,ori.drug_stock_id,cd.packing_unit_name,cd.manu_factory_name,ori.outstock_amount,
		cd.ret_price,ds.buy_price,ds.serial,ds.eff_date
		from drug_outstock_record_item ori
		left join drug_stock ds on ori.drug_stock_id = ds.id
		left join clinic_drug cd on ds.clinic_drug_id = cd.id
		where ori.drug_outstock_record_id=$1`

	irows, err := model.DB.Queryx(isql, drugOutstockRecordID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	item := FormatSQLRowsToMapArray(irows)
	result["items"] = item
	ctx.JSON(iris.Map{"code": "200", "data": result})
}

//DrugOutstockUpdate 出库记录修改
func DrugOutstockUpdate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	drugOutstockRecordID := ctx.PostValue("drug_outstock_record_id")
	items := ctx.PostValue("items")
	operationID := ctx.PostValue("outstock_operation_id")
	outstockWayName := ctx.PostValue("outstock_way_name")
	departmentID := ctx.PostValue("department_id")
	personnelID := ctx.PostValue("personnel_id")
	remark := ctx.PostValue("remark")
	outstockDate := ctx.PostValue("outstock_date")

	if clinicID == "" || drugOutstockRecordID == "" || items == "" || outstockWayName == "" || departmentID == "" || operationID == "" || personnelID == "" || outstockDate == "" {
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

	row := model.DB.QueryRowx("select * from drug_outstock_record where id=$1 and storehouse_id=$2 limit 1", drugOutstockRecordID, storehouseID)
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

	sets := []string{
		"drug_stock_id",
		"outstock_amount",
		"drug_outstock_record_id"}

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errb})
		return
	}

	updateSQL := `update drug_outstock_record set department_id=$1,personnel_id=$2,outstock_way_name=$3,outstock_date=$4,outstock_operation_id=$5,remark=$6 where id=$7`
	_, erru := tx.Exec(updateSQL, departmentID, personnelID, outstockWayName, outstockDate, operationID, remark, drugOutstockRecordID)
	if erru != nil {
		fmt.Println("erru ===", erru)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": erru.Error()})
		return
	}
	_, errd := tx.Exec("delete from drug_outstock_record_item where drug_outstock_record_id=$1", drugOutstockRecordID)
	if errd != nil {
		fmt.Println("errd ===", errd)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errd.Error()})
		return
	}

	setStr := strings.Join(sets, ",")

	for _, v := range results {
		drugStockID := v["drug_stock_id"]
		outstockAmount := v["outstock_amount"]
		if outstockAmount == "" {
			ctx.JSON(iris.Map{"code": "-1", "msg": "数量为必填项"})
			return
		}
		row := model.DB.QueryRowx("select id from drug_stock where id=$1 limit 1", drugStockID)
		if row == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "新增出库失败"})
			return
		}
		drugStock := FormatSQLRowToMap(row)
		_, ok := drugStock["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1-", "msg": "新增出库药品不存在"})
			return
		}
		insertiSQL := "insert into drug_outstock_record_item (" + setStr + ") values ($1,$2,$3)"
		fmt.Println("insertiSQL===", insertiSQL)
		_, err := tx.Exec(insertiSQL,
			ToNullInt64(drugStockID),
			ToNullInt64(outstockAmount),
			ToNullInt64(drugOutstockRecordID),
		)
		if err != nil {
			fmt.Println("err ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": "请检查是否漏填"})
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

//DrugOutstockCheck 出库审核
func DrugOutstockCheck(ctx iris.Context) {
	drugOutstockRecordID := ctx.PostValue("drug_outstock_record_id")
	operationID := ctx.PostValue("verify_operation_id")
	if drugOutstockRecordID == "" || operationID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx(`select * from drug_outstock_record where id=$1 limit 1`, drugOutstockRecordID)
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

	rows, _ := model.DB.Queryx(`select * from drug_outstock_record_item	where drug_outstock_record_id=$1`, drugOutstockRecordID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "审核失败"})
		return
	}
	outstockRecordItems := FormatSQLRowsToMapArray(rows)

	tx, err := model.DB.Begin()
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	_, err2 := tx.Exec("update drug_outstock_record set verify_status=$1,verify_operation_id=$2,updated_time=LOCALTIMESTAMP where id=$3", "02", operationID, drugOutstockRecordID)
	if err2 != nil {
		fmt.Println("err2 ===", err2)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}

	sql := "update drug_stock set stock_amount=$1, updated_time=LOCALTIMESTAMP where id=$2"
	for _, v := range outstockRecordItems {
		drugStockID := v["drug_stock_id"].(int64)
		outstockAmount := v["outstock_amount"].(int64)
		drow := model.DB.QueryRowx("select id,stock_amount from drug_stock where id=$1 limit 1", drugStockID)
		if drow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "审核失败"})
			return
		}
		drugStock := FormatSQLRowToMap(drow)
		_, dok := drugStock["id"]
		if !dok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "出库审核失败"})
			return
		}
		stockAmount := drugStock["stock_amount"].(int64) - outstockAmount

		_, err1 := tx.Exec(sql, stockAmount, drugStockID)
		if err1 != nil {
			fmt.Println("err1 ===", err1)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
			return
		}
	}

	err3 := tx.Commit()
	if err3 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "msg": "ok"})
}

//DrugOutstockRecordDelete 删除出库记录
func DrugOutstockRecordDelete(ctx iris.Context) {
	drugOutstockRecordID := ctx.PostValue("drug_outstock_record_id")
	if drugOutstockRecordID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	rows := model.DB.QueryRowx("select id,verify_status from drug_outstock_record where id=$1", drugOutstockRecordID)
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
	_, erri := tx.Exec("delete from drug_outstock_record_item where drug_outstock_record_id=$1", drugOutstockRecordID)
	if erri != nil {
		fmt.Println("erri ===", erri)
		ctx.JSON(iris.Map{"code": "-1", "msg": erri.Error()})
		return
	}

	_, errd := tx.Exec("delete from drug_outstock_record where id=$1", drugOutstockRecordID)
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

//DrugStockList 库存列表
func DrugStockList(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	keyword := ctx.PostValue("keyword")
	supplierName := ctx.PostValue("supplier_name")
	amount := ctx.PostValue("amount")
	dateWarning := ctx.PostValue("date_warning")
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
		ctx.JSON(iris.Map{"code": "-1", "msg": errs.Error()})
		return
	}

	countSQL := `select count(*) as total from drug_stock ds 
		left join clinic_drug cd on ds.clinic_drug_id = cd.id
		where ds.storehouse_id=:storehouse_id`
	selectSQL := `select 
		cd.name,
		cd.specification,
		cd.packing_unit_name,
		cd.manu_factory_name,
		ds.supplier_name,
		cd.ret_price,
		ds.buy_price,
		ds.serial,
		ds.eff_date,
		ds.stock_amount,
		ds.id as drug_stock_id
		from drug_stock ds
		left join clinic_drug cd on ds.clinic_drug_id = cd.id
		where ds.storehouse_id=:storehouse_id`

	if supplierName != "" {
		countSQL += " and ds.supplier_name = :supplier_name"
		selectSQL += " and ds.supplier_name= :supplier_name"
	}
	if keyword != "" {
		countSQL += " and (cd.name ~:keyword or cd.barcode ~:keyword)"
		selectSQL += " and (cd.name ~:keyword or cd.barcode ~:keyword)"
	}

	if amount != "" {
		countSQL += " and ds.stock_amount>0"
		selectSQL += " and ds.stock_amount>0"
	}
	if dateWarning != "" {
		countSQL += " and (ds.eff_date <= (CURRENT_DATE + cd.day_warning))"
		selectSQL += " and (ds.eff_date <= (CURRENT_DATE + cd.day_warning))"
	}

	var queryOption = map[string]interface{}{
		"storehouse_id": ToNullInt64(storehouseID),
		"supplier_name": ToNullString(supplierName),
		"keyword":       ToNullString(keyword),
		"offset":        ToNullInt64(offset),
		"limit":         ToNullInt64(limit),
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
