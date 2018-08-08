package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"
	"time"

	"github.com/kataras/iris"
)

//DrugInstockStatistics 药品入库统计
func DrugInstockStatistics(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	instockWayName := ctx.PostValue("instock_way_name")
	supplierName := ctx.PostValue("supplier_name")
	drugType := ctx.PostValue("drug_type")
	drugName := ctx.PostValue("drug_name")
	serial := ctx.PostValue("serial")
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")
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

	countSQL := `select count(*) as total from drug_instock_record_item diri
	left join drug_instock_record dir on dir.id = diri.drug_instock_record_id
	left join clinic_drug cd on cd.id = diri.clinic_drug_id where dir.storehouse_id=:storehouse_id`
	selectSQL := `select dir.instock_date,dir.order_number,dir.instock_way_name,dir.supplier_name,
	p.name as instock_operation_name,cd.barcode,cd.name as drug_name,cd.specification,cd.manu_factory_name,
	diri.instock_amount,cd.packing_unit_name,cd.ret_price,diri.buy_price,diri.serial,diri.eff_date
	from drug_instock_record_item diri
	left join drug_instock_record dir on dir.id = diri.drug_instock_record_id
	left join clinic_drug cd on cd.id = diri.clinic_drug_id
	left join personnel p on p.id = dir.instock_operation_id
	where dir.storehouse_id=:storehouse_id`

	if startDate != "" && endDate != "" {
		if startDate > endDate {
			ctx.JSON(iris.Map{"code": "-1", "msg": "开始日期必须大于结束日期"})
			return
		}

		countSQL += " and dir.instock_date between :start_date and :end_date"
		selectSQL += " and dir.instock_date between :start_date and :end_date"
	}

	if instockWayName != "" {
		countSQL += " and dir.instock_way_name =:instock_way_name"
		selectSQL += " and dir.instock_way_name =:instock_way_name"
	}

	if supplierName != "" {
		countSQL += " and dir.supplier_name =:supplier_name"
		selectSQL += " and dir.supplier_name =:supplier_name"
	}

	if drugType != "" {
		countSQL += " and cd.type =:drug_type"
		selectSQL += " and cd.type =:drug_type"
	}

	if drugName != "" {
		countSQL += " and (cd.name ~*:drug_name or cd.barcode ~*:drug_name)"
		selectSQL += " and (cd.name ~*:drug_name or cd.barcode ~*:drug_name)"
	}

	if serial != "" {
		countSQL += " and diri.serial =:serial"
		selectSQL += " and diri.serial =:serial"
	}

	var queryOption = map[string]interface{}{
		"storehouse_id":    ToNullInt64(storehouseID),
		"instock_way_name": ToNullString(instockWayName),
		"supplier_name":    ToNullString(supplierName),
		"drug_name":        ToNullString(drugName),
		"drug_type":        ToNullString(drugType),
		"serial":           ToNullString(serial),
		"start_date":       ToNullString(startDate),
		"end_date":         ToNullString(endDate),
		"offset":           ToNullInt64(offset),
		"limit":            ToNullInt64(limit),
	}

	total, err := model.DB.NamedQuery(countSQL, queryOption)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.NamedQuery(selectSQL+" order by dir.instock_date desc offset :offset limit :limit", queryOption)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

//DrugOutstockStatistics 药品出库统计
func DrugOutstockStatistics(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	outstockWayName := ctx.PostValue("outstock_way_name")
	supplierName := ctx.PostValue("supplier_name")
	drugType := ctx.PostValue("drug_type")
	drugName := ctx.PostValue("drug_name")
	serial := ctx.PostValue("serial")
	personnelName := ctx.PostValue("personnel_name")
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")
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

	countSQL := `select count(*) as total from drug_outstock_record_item diri
	left join drug_outstock_record dir on dir.id = diri.drug_outstock_record_id
	left join drug_stock ds on ds.id = diri.drug_stock_id
	left join clinic_drug cd on cd.id = ds.clinic_drug_id 
	left join personnel doc on doc.id = dir.personnel_id
	where dir.storehouse_id=:storehouse_id`
	selectSQL := `select dir.outstock_date,dir.order_number,dir.outstock_way_name,ds.supplier_name,
	p.name as outstock_operation_name,doc.name as personnel_name,cd.barcode,cd.name as drug_name,
	cd.specification,cd.manu_factory_name,diri.outstock_amount,cd.packing_unit_name,cd.ret_price,
	ds.buy_price,ds.serial,ds.eff_date
	from  drug_outstock_record_item diri
	left join drug_outstock_record dir on dir.id = diri.drug_outstock_record_id
	left join drug_stock ds on ds.id = diri.drug_stock_id
	left join clinic_drug cd on cd.id = ds.clinic_drug_id
	left join personnel p on p.id = dir.outstock_operation_id
	left join personnel doc on doc.id = dir.personnel_id
	where dir.storehouse_id=:storehouse_id`

	if startDate != "" && endDate != "" {
		if startDate > endDate {
			ctx.JSON(iris.Map{"code": "-1", "msg": "开始日期必须大于结束日期"})
			return
		}

		countSQL += " and dir.outstock_date between :start_date and :end_date"
		selectSQL += " and dir.outstock_date between :start_date and :end_date"
	}

	if outstockWayName != "" {
		countSQL += " and dir.outstock_way_name =:outstock_way_name"
		selectSQL += " and dir.outstock_way_name =:outstock_way_name"
	}

	if supplierName != "" {
		countSQL += " and ds.supplier_name =:supplier_name"
		selectSQL += " and ds.supplier_name =:supplier_name"
	}

	if drugType != "" {
		countSQL += " and cd.type =:drug_type"
		selectSQL += " and cd.type =:drug_type"
	}

	if drugName != "" {
		countSQL += " and (cd.name ~*:drug_name or cd.barcode ~*:drug_name)"
		selectSQL += " and (cd.name ~*:drug_name or cd.barcode ~*:drug_name)"
	}

	if serial != "" {
		countSQL += " and ds.serial =:serial"
		selectSQL += " and ds.serial =:serial"
	}

	if personnelName != "" {
		countSQL += " and doc.name ~*:personnel_name"
		selectSQL += " and doc.name ~*:personnel_name"
	}

	var queryOption = map[string]interface{}{
		"storehouse_id":    ToNullInt64(storehouseID),
		"instock_way_name": ToNullString(outstockWayName),
		"supplier_name":    ToNullString(supplierName),
		"drug_name":        ToNullString(drugName),
		"drug_type":        ToNullString(drugType),
		"serial":           ToNullString(serial),
		"personnel_name":   ToNullString(personnelName),
		"start_date":       ToNullString(startDate),
		"end_date":         ToNullString(endDate),
		"offset":           ToNullInt64(offset),
		"limit":            ToNullInt64(limit),
	}

	total, err := model.DB.NamedQuery(countSQL, queryOption)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.NamedQuery(selectSQL+" order by dir.outstock_date desc offset :offset limit :limit", queryOption)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

//DrugInvoicingStatistics 药品进存销统计
func DrugInvoicingStatistics(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	supplierName := ctx.PostValue("supplier_name")
	drugType := ctx.PostValue("drug_type")
	drugName := ctx.PostValue("drug_name")
	serial := ctx.PostValue("serial")
	startDateStr := ctx.PostValue("start_date")
	endDateStr := ctx.PostValue("end_date")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if startDateStr == "" || endDateStr == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请选择开始和结束日期"})
		return
	}

	endDate, erre := time.Parse("2006-01-02", endDateStr)
	if erre != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "end_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}
	endDateStr = endDate.AddDate(0, 0, 1).Format("2006-01-02")

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

	selectSQL := `
	select 
	ds.id,
	cd.barcode,
	cd.name,
	cd.specification,
	cd.manu_factory_name,
	ds.supplier_name,
	cd.packing_unit_name,
	ds.serial,
	ds.eff_date,
	ds.buy_price,
	sum(in_stock.instock_amount) as total_instock_amount,
	sum(out_stock.outstock_amount) as total_outstock_amount
	from drug_stock ds
	left join clinic_drug cd on cd.id = ds.clinic_drug_id
	
	left join ((select diri.instock_amount,dir.verify_time,diri.serial,diri.eff_date,dir.supplier_name
	from drug_instock_record_item diri 
	left join drug_instock_record dir on dir.id = diri.drug_instock_record_id
	where dir.verify_status='02')
	UNION all
	(select dr.amount as instock_amount,dr.created_time as verify_time,ds.serial,ds.eff_date,ds.supplier_name 
	from drug_retail dr
	left join drug_stock ds on ds.id = dr.drug_stock_id
	where dr.amount>0)) as in_stock on in_stock.verify_time BETWEEN :start_date and :end_date and in_stock.serial = ds.serial and in_stock.eff_date = ds.eff_date and in_stock.supplier_name = ds.supplier_name
	
	left join ((select dori.outstock_amount,dor.verify_time,dori.drug_stock_id
	from drug_outstock_record_item dori 
	left join drug_outstock_record dor on dor.id = dori.drug_outstock_record_id
	where dor.verify_status='02')
	UNION all
	(select amount as outstock_amount,created_time as verify_time,drug_stock_id from drug_retail
	where amount<0)) as out_stock on out_stock.drug_stock_id = ds.id and out_stock.verify_time BETWEEN :start_date and :end_date
	where ds.storehouse_id=:storehouse_id
	group by 
	ds.id,
	cd.name,
	cd.barcode,
	cd.specification,
	cd.manu_factory_name,
	ds.supplier_name,
	cd.packing_unit_name,
	ds.serial,
	ds.eff_date,
	ds.buy_price`

	if supplierName != "" {
		// countSQL += " and ds.supplier_name =:supplier_name"
		selectSQL += " and ds.supplier_name =:supplier_name"
	}

	if drugType != "" {
		// countSQL += " and cd.type =:drug_type"
		selectSQL += " and cd.type =:drug_type"
	}

	if drugName != "" {
		// countSQL += " and (cd.name ~*:drug_name or cd.barcode ~*:drug_name)"
		selectSQL += " and (cd.name ~*:drug_name or cd.barcode ~*:drug_name)"
	}

	if serial != "" {
		// countSQL += " and ds.serial =:serial"
		selectSQL += " and ds.serial =:serial"
	}

	var queryOption = map[string]interface{}{
		"storehouse_id": ToNullInt64(storehouseID),
		"supplier_name": ToNullString(supplierName),
		"drug_name":     ToNullString(drugName),
		"drug_type":     ToNullString(drugType),
		"serial":        ToNullString(serial),
		"start_date":    ToNullString(startDateStr),
		"end_date":      ToNullString(endDateStr),
		"offset":        ToNullInt64(offset),
		"limit":         ToNullInt64(limit),
	}

	// total, err := model.DB.NamedQuery(countSQL, queryOption)
	// if err != nil {
	// 	ctx.JSON(iris.Map{"code": "-1", "msg": err})
	// 	return
	// }

	// pageInfo := FormatSQLRowsToMapArray(total)[0]

	var results []map[string]interface{}
	rows, _ := model.DB.NamedQuery(selectSQL+" offset :offset limit :limit", queryOption)
	results = FormatSQLRowsToMapArray(rows)

	pageInfo := map[string]interface{}{
		"total":  len(results),
		"offset": offset,
		"limit":  limit,
	}

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

//MaterialInstockStatistics 耗材入库统计
func MaterialInstockStatistics(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	instockWayName := ctx.PostValue("instock_way_name")
	supplierName := ctx.PostValue("supplier_name")
	materialName := ctx.PostValue("material_name")
	serial := ctx.PostValue("serial")
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")
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

	countSQL := `select count(*) as total from material_instock_record_item miri
	left join material_instock_record mir on mir.id = miri.material_instock_record_id
	left join clinic_material cm on cm.id = miri.clinic_material_id where mir.storehouse_id=:storehouse_id`
	selectSQL := `select mir.instock_date,mir.order_number,mir.instock_way_name,mir.supplier_name,
	p.name as instock_operation_name,cm.idc_code,cm.name as material_name,cm.specification,cm.manu_factory_name,
	miri.instock_amount,cm.unit_name,cm.ret_price,miri.buy_price,miri.serial,miri.eff_date
	from material_instock_record_item miri
	left join material_instock_record mir on mir.id = miri.material_instock_record_id
	left join clinic_material cm on cm.id = miri.clinic_material_id
	left join personnel p on p.id = mir.instock_operation_id
	where mir.storehouse_id=:storehouse_id`

	if startDate != "" && endDate != "" {
		if startDate > endDate {
			ctx.JSON(iris.Map{"code": "-1", "msg": "开始日期必须大于结束日期"})
			return
		}

		countSQL += " and mir.instock_date between :start_date and :end_date"
		selectSQL += " and mir.instock_date between :start_date and :end_date"
	}

	if instockWayName != "" {
		countSQL += " and mir.instock_way_name =:instock_way_name"
		selectSQL += " and mir.instock_way_name =:instock_way_name"
	}

	if supplierName != "" {
		countSQL += " and mir.supplier_name =:supplier_name"
		selectSQL += " and mir.supplier_name =:supplier_name"
	}

	if materialName != "" {
		countSQL += " and (cm.name ~*:material_name or cm.py_code ~*:material_name)"
		selectSQL += " and (cm.name ~*:material_name or cm.py_code ~*:material_name)"
	}

	if serial != "" {
		countSQL += " and miri.serial =:serial"
		selectSQL += " and miri.serial =:serial"
	}

	var queryOption = map[string]interface{}{
		"storehouse_id":    ToNullInt64(storehouseID),
		"instock_way_name": ToNullString(instockWayName),
		"supplier_name":    ToNullString(supplierName),
		"material_name":    ToNullString(materialName),
		"serial":           ToNullString(serial),
		"start_date":       ToNullString(startDate),
		"end_date":         ToNullString(endDate),
		"offset":           ToNullInt64(offset),
		"limit":            ToNullInt64(limit),
	}

	total, err := model.DB.NamedQuery(countSQL, queryOption)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.NamedQuery(selectSQL+" order by mir.instock_date desc offset :offset limit :limit", queryOption)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

//MaterialOutstockStatistics 耗材出库统计
func MaterialOutstockStatistics(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	outstockWayName := ctx.PostValue("outstock_way_name")
	supplierName := ctx.PostValue("supplier_name")
	materialName := ctx.PostValue("material_name")
	serial := ctx.PostValue("serial")
	personnelName := ctx.PostValue("personnel_name")
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")
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

	countSQL := `select count(*) as total from material_outstock_record_item miri
	left join material_outstock_record mir on mir.id = miri.material_outstock_record_id
	left join material_stock ms on ms.id = miri.material_stock_id
	left join clinic_material cm on cm.id = ms.clinic_material_id 
	left join personnel doc on doc.id = mir.personnel_id
	where mir.storehouse_id=:storehouse_id`
	selectSQL := `select mir.outstock_date,mir.order_number,mir.outstock_way_name,ms.supplier_name,
	p.name as outstock_operation_name,doc.name as personnel_name,cm.idc_code,cm.name as material_name,
	cm.specification,cm.manu_factory_name,miri.outstock_amount,cm.unit_name,cm.ret_price,
	ms.buy_price,ms.serial,ms.eff_date
	from  material_outstock_record_item miri
	left join material_outstock_record mir on mir.id = miri.material_outstock_record_id
	left join material_stock ms on ms.id = miri.material_stock_id
	left join clinic_material cm on cm.id = ms.clinic_material_id
	left join personnel p on p.id = mir.outstock_operation_id
	left join personnel doc on doc.id = mir.personnel_id
	where mir.storehouse_id=:storehouse_id`

	if startDate != "" && endDate != "" {
		if startDate > endDate {
			ctx.JSON(iris.Map{"code": "-1", "msg": "开始日期必须大于结束日期"})
			return
		}

		countSQL += " and mir.outstock_date between :start_date and :end_date"
		selectSQL += " and mir.outstock_date between :start_date and :end_date"
	}

	if outstockWayName != "" {
		countSQL += " and mir.outstock_way_name =:outstock_way_name"
		selectSQL += " and mir.outstock_way_name =:outstock_way_name"
	}

	if supplierName != "" {
		countSQL += " and ms.supplier_name =:supplier_name"
		selectSQL += " and ms.supplier_name =:supplier_name"
	}

	if materialName != "" {
		countSQL += " and (cm.name ~*:material_name or cm.py_code ~*:material_name)"
		selectSQL += " and (cm.name ~*:material_name or cm.py_code ~*:material_name)"
	}

	if serial != "" {
		countSQL += " and ms.serial =:serial"
		selectSQL += " and ms.serial =:serial"
	}

	if personnelName != "" {
		countSQL += " and doc.name ~*:personnel_name"
		selectSQL += " and doc.name ~*:personnel_name"
	}

	var queryOption = map[string]interface{}{
		"storehouse_id":    ToNullInt64(storehouseID),
		"instock_way_name": ToNullString(outstockWayName),
		"supplier_name":    ToNullString(supplierName),
		"material_name":    ToNullString(materialName),
		"serial":           ToNullString(serial),
		"personnel_name":   ToNullString(personnelName),
		"start_date":       ToNullString(startDate),
		"end_date":         ToNullString(endDate),
		"offset":           ToNullInt64(offset),
		"limit":            ToNullInt64(limit),
	}

	total, err := model.DB.NamedQuery(countSQL, queryOption)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.NamedQuery(selectSQL+" order by mir.outstock_date desc offset :offset limit :limit", queryOption)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

//MaterialInvoicingStatistics 耗材进存销统计
func MaterialInvoicingStatistics(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	supplierName := ctx.PostValue("supplier_name")
	materialName := ctx.PostValue("material_name")
	serial := ctx.PostValue("serial")
	startDateStr := ctx.PostValue("start_date")
	endDateStr := ctx.PostValue("end_date")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if startDateStr == "" || endDateStr == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请选择开始和结束日期"})
		return
	}

	endDate, erre := time.Parse("2006-01-02", endDateStr)
	if erre != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "end_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}
	endDateStr = endDate.AddDate(0, 0, 1).Format("2006-01-02")

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

	selectSQL := `
	select 
	ms.id,
	cm.name,
	cm.py_code,
	cm.idc_code,
	cm.unit_name,
	cm.specification,
	cm.manu_factory_name,
	ms.supplier_name,
	ms.serial,
	ms.eff_date,
	ms.buy_price,
	sum(in_stock.instock_amount) as total_instock_amount,
	sum(out_stock.outstock_amount) as total_outstock_amount
	from material_stock ms
	left join clinic_material cm on cm.id = ms.clinic_material_id
	
	left join (select miri.instock_amount,mir.verify_time,miri.serial,miri.eff_date,mir.supplier_name
	from material_instock_record_item miri 
	left join material_instock_record mir on mir.id = miri.material_instock_record_id
	where mir.verify_status='02') as in_stock on in_stock.verify_time BETWEEN :start_date and :end_date and in_stock.serial = ms.serial and in_stock.eff_date = ms.eff_date and in_stock.supplier_name = ms.supplier_name
	
	left join (select dori.outstock_amount,dor.verify_time,dori.material_stock_id
	from material_outstock_record_item dori 
	left join material_outstock_record dor on dor.id = dori.material_outstock_record_id
	where dor.verify_status='02') as out_stock on out_stock.material_stock_id = ms.id and out_stock.verify_time BETWEEN :start_date and :end_date
	where ms.storehouse_id=:storehouse_id
	group by 
	ms.id,
	cm.name,
	cm.py_code,
	cm.idc_code,
	cm.unit_name,
	cm.specification,
	cm.manu_factory_name,
	ms.supplier_name,
	ms.serial,
	ms.eff_date,
	ms.buy_price`

	if supplierName != "" {
		// countSQL += " and ms.supplier_name =:supplier_name"
		selectSQL += " and ms.supplier_name =:supplier_name"
	}

	if materialName != "" {
		// countSQL += " and (cm.name ~*:material_name or cm.barcode ~*:material_name)"
		selectSQL += " and (cm.name ~*:material_name or cm.barcode ~*:material_name)"
	}

	if serial != "" {
		// countSQL += " and ms.serial =:serial"
		selectSQL += " and ms.serial =:serial"
	}

	var queryOption = map[string]interface{}{
		"storehouse_id": ToNullInt64(storehouseID),
		"supplier_name": ToNullString(supplierName),
		"material_name": ToNullString(materialName),
		"serial":        ToNullString(serial),
		"start_date":    ToNullString(startDateStr),
		"end_date":      ToNullString(endDateStr),
		"offset":        ToNullInt64(offset),
		"limit":         ToNullInt64(limit),
	}

	// total, err := model.DB.NamedQuery(countSQL, queryOption)
	// if err != nil {
	// 	ctx.JSON(iris.Map{"code": "-1", "msg": err})
	// 	return
	// }

	// pageInfo := FormatSQLRowsToMapArray(total)[0]

	var results []map[string]interface{}
	rows, _ := model.DB.NamedQuery(selectSQL+" offset :offset limit :limit", queryOption)
	results = FormatSQLRowsToMapArray(rows)

	pageInfo := map[string]interface{}{
		"total":  len(results),
		"offset": offset,
		"limit":  limit,
	}

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}
