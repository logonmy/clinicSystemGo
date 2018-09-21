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

	countSQL := `select count(*) as total
		from (
			select to_char(dir.verify_time,'yyyy-mm-dd') as instock_date,dir.order_number,dir.instock_way_name,dir.supplier_name,
			p.name as instock_operation_name,cd.barcode,cd.name as drug_name,cd.specification,cd.manu_factory_name,
			diri.instock_amount,cd.packing_unit_name,cd.ret_price,cd.type,diri.buy_price,diri.serial,diri.eff_date,p.clinic_id
			from drug_instock_record_item diri
			left join drug_instock_record dir on dir.id = diri.drug_instock_record_id
			left join clinic_drug cd on cd.id = diri.clinic_drug_id
			left join personnel p on p.id = dir.instock_operation_id
			where dir.verify_status='02'
			UNION all
			select to_char(dr.created_time,'yyyy-mm-dd') as instock_date,dr.out_refund_no as order_number,'零售退药' as instock_way_name,ds.supplier_name,
			p.name as instock_operation_name,cd.barcode,cd.name as drug_name,cd.specification,cd.manu_factory_name,
			dr.amount as instock_amount,cd.packing_unit_name,cd.ret_price,cd.type,ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
			from drug_retail dr
			left join drug_stock ds on ds.id = dr.drug_stock_id
			left join clinic_drug cd on cd.id = ds.clinic_drug_id
			left join drug_retail_pay_record drpr on drpr.out_trade_no = dr.out_trade_no
			left join personnel p on p.id = drpr.operation_id
			where dr.amount<0
			UNION all
			select to_char(ddd.created_time,'yyyy-mm-dd') as instock_date,
			'MZ' || to_char(ddd.created_time, 'YYYYMMDDHH24MISS') as order_number,'门诊退药' as instock_way_name,ds.supplier_name,
			p.name as instock_operation_name,cd.barcode,cd.name as drug_name,cd.specification,cd.manu_factory_name,
			ddd.amount as instock_amount,cd.packing_unit_name,cd.ret_price,cd.type,ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
			from drug_delivery_detail ddd
			left join drug_stock ds on ddd.drug_stock_id = ds.id
			left join clinic_drug cd on cd.id = ds.clinic_drug_id
			left join personnel p on p.id = ddd.operation_id
			where ddd.amount<0) as a
		where a.clinic_id=:clinic_id`

	selectSQL := `select instock_date,order_number,instock_way_name,supplier_name,
	instock_operation_name,barcode,drug_name,specification,manu_factory_name,
	instock_amount,packing_unit_name,ret_price,buy_price,serial,eff_date
	from (
		select to_char(dir.verify_time,'yyyy-mm-dd') as instock_date,dir.order_number,dir.instock_way_name,dir.supplier_name,
		p.name as instock_operation_name,cd.barcode,cd.name as drug_name,cd.specification,cd.manu_factory_name,
		diri.instock_amount,cd.packing_unit_name,cd.ret_price,cd.type,diri.buy_price,diri.serial,diri.eff_date,p.clinic_id
		from drug_instock_record_item diri
		left join drug_instock_record dir on dir.id = diri.drug_instock_record_id
		left join clinic_drug cd on cd.id = diri.clinic_drug_id
		left join personnel p on p.id = dir.instock_operation_id
		where dir.verify_status='02'
    UNION all
    select to_char(dr.created_time,'yyyy-mm-dd') as instock_date,dr.out_refund_no as order_number,'零售退药' as instock_way_name,ds.supplier_name,
    p.name as instock_operation_name,cd.barcode,cd.name as drug_name,cd.specification,cd.manu_factory_name,
    dr.amount as instock_amount,cd.packing_unit_name,cd.ret_price,cd.type,ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
    from drug_retail dr
    left join drug_stock ds on ds.id = dr.drug_stock_id
		left join clinic_drug cd on cd.id = ds.clinic_drug_id
		left join drug_retail_pay_record drpr on drpr.out_trade_no = dr.out_trade_no
		left join personnel p on p.id = drpr.operation_id
		where dr.amount<0
		UNION all
    select to_char(ddd.created_time,'yyyy-mm-dd') as instock_date,
    'MZ' || to_char(ddd.created_time, 'YYYYMMDDHH24MISS') as order_number,'门诊退药' as instock_way_name,ds.supplier_name,
    p.name as instock_operation_name,cd.barcode,cd.name as drug_name,cd.specification,cd.manu_factory_name,
    ddd.amount as instock_amount,cd.packing_unit_name,cd.ret_price,cd.type,ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
    from drug_delivery_detail ddd
    left join drug_stock ds on ddd.drug_stock_id = ds.id
    left join clinic_drug cd on cd.id = ds.clinic_drug_id
		left join personnel p on p.id = ddd.operation_id
		where ddd.amount<0) as a
  where a.clinic_id=:clinic_id`

	totalSQL := `select 
	sum(instock_amount) as total_instock_amount,
	sum(instock_amount * buy_price) as total_buy_price
	from (
		select to_char(dir.verify_time,'yyyy-mm-dd') as instock_date,dir.order_number,dir.instock_way_name,dir.supplier_name,
		p.name as instock_operation_name,cd.barcode,cd.name as drug_name,cd.specification,cd.manu_factory_name,
		diri.instock_amount,cd.packing_unit_name,cd.ret_price,cd.type,diri.buy_price,diri.serial,diri.eff_date,p.clinic_id
		from drug_instock_record_item diri
		left join drug_instock_record dir on dir.id = diri.drug_instock_record_id
		left join clinic_drug cd on cd.id = diri.clinic_drug_id
		left join personnel p on p.id = dir.instock_operation_id
		where dir.verify_status='02'
    UNION all
    select to_char(dr.created_time,'yyyy-mm-dd') as instock_date,dr.out_refund_no as order_number,'零售退药' as instock_way_name,ds.supplier_name,
    p.name as instock_operation_name,cd.barcode,cd.name as drug_name,cd.specification,cd.manu_factory_name,
    dr.amount as instock_amount,cd.packing_unit_name,cd.ret_price,cd.type,ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
    from drug_retail dr
    left join drug_stock ds on ds.id = dr.drug_stock_id
		left join clinic_drug cd on cd.id = ds.clinic_drug_id
		left join drug_retail_pay_record drpr on drpr.out_trade_no = dr.out_trade_no
		left join personnel p on p.id = drpr.operation_id
		where dr.amount<0
		UNION all
    select to_char(ddd.created_time,'yyyy-mm-dd') as instock_date,
    'MZ' || to_char(ddd.created_time, 'YYYYMMDDHH24MISS') as order_number,'门诊退药' as instock_way_name,ds.supplier_name,
    p.name as instock_operation_name,cd.barcode,cd.name as drug_name,cd.specification,cd.manu_factory_name,
    ddd.amount as instock_amount,cd.packing_unit_name,cd.ret_price,cd.type,ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
    from drug_delivery_detail ddd
    left join drug_stock ds on ddd.drug_stock_id = ds.id
    left join clinic_drug cd on cd.id = ds.clinic_drug_id
		left join personnel p on p.id = ddd.operation_id
		where ddd.amount<0) as a
  where a.clinic_id=:clinic_id`

	if startDate != "" && endDate != "" {
		if startDate > endDate {
			ctx.JSON(iris.Map{"code": "-1", "msg": "开始日期必须大于结束日期"})
			return
		}

		countSQL += " and a.instock_date between :start_date and :end_date"
		selectSQL += " and a.instock_date between :start_date and :end_date"
		totalSQL += " and a.instock_date between :start_date and :end_date"
	}

	if instockWayName != "" {
		countSQL += " and a.instock_way_name =:instock_way_name"
		selectSQL += " and a.instock_way_name =:instock_way_name"
		totalSQL += " and a.instock_way_name =:instock_way_name"
	}

	if supplierName != "" {
		countSQL += " and a.supplier_name =:supplier_name"
		selectSQL += " and a.supplier_name =:supplier_name"
		totalSQL += " and a.supplier_name =:supplier_name"
	}

	if drugType != "" {
		countSQL += " and a.type =:drug_type"
		selectSQL += " and a.type =:drug_type"
		totalSQL += " and a.type =:drug_type"
	}

	if drugName != "" {
		countSQL += " and (a.name ~*:drug_name or a.barcode ~*:drug_name)"
		selectSQL += " and (a.name ~*:drug_name or a.barcode ~*:drug_name)"
		totalSQL += " and (a.name ~*:drug_name or a.barcode ~*:drug_name)"
	}

	if serial != "" {
		countSQL += " and a.serial =:serial"
		selectSQL += " and a.serial =:serial"
		totalSQL += " and a.serial =:serial"
	}

	var queryOption = map[string]interface{}{
		"clinic_id":        ToNullInt64(clinicID),
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
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.NamedQuery(selectSQL+" order by a.instock_date desc offset :offset limit :limit", queryOption)
	results = FormatSQLRowsToMapArray(rows)

	countRows, _ := model.DB.NamedQuery(totalSQL, queryOption)
	countResults := FormatSQLRowsToMapArray(countRows)

	pageInfo["total_instock_amount"] = countResults[0]["total_instock_amount"]
	pageInfo["total_buy_price"] = countResults[0]["total_buy_price"]

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

	countSQL := `select count(*) as total
	from (
		select to_char(dir.verify_time,'yyyy-mm-dd') as outstock_date,dir.order_number,dir.outstock_way_name,ds.supplier_name,
		p.name as outstock_operation_name,doc.name as personnel_name,cd.barcode,cd.name as drug_name,
		cd.specification,cd.manu_factory_name,diri.outstock_amount,cd.packing_unit_name,cd.ret_price,cd.type,
		ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
		from  drug_outstock_record_item diri
		left join drug_outstock_record dir on dir.id = diri.drug_outstock_record_id
		left join drug_stock ds on ds.id = diri.drug_stock_id
		left join clinic_drug cd on cd.id = ds.clinic_drug_id
		left join personnel p on p.id = dir.outstock_operation_id
		left join personnel doc on doc.id = dir.personnel_id
		where dir.verify_status='02'
		UNION all
		select to_char(dr.created_time,'yyyy-mm-dd') as outstock_date,dr.out_trade_no as order_number,'零售发药' as outstock_way_name,ds.supplier_name,
		p.name as outstock_operation_name,p.name as personnel_name,cd.barcode,cd.name as drug_name,cd.specification,cd.manu_factory_name,
		dr.amount as outstock_amount,cd.packing_unit_name,cd.ret_price,cd.type,ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
		from drug_retail dr
		left join drug_stock ds on ds.id = dr.drug_stock_id
		left join clinic_drug cd on cd.id = ds.clinic_drug_id
		left join drug_retail_pay_record drpr on drpr.out_trade_no = dr.out_trade_no
		left join personnel p on p.id = drpr.operation_id
		where dr.amount>0
		UNION all
		select to_char(ddd.created_time,'yyyy-mm-dd') as outstock_date,
		'MZ' || to_char(ddd.created_time, 'YYYYMMDDHH24MISS') as order_number,'门诊发药' as outstock_way_name,ds.supplier_name,
		p.name as outstock_operation_name,p.name as personnel_name,cd.barcode,cd.name as drug_name,cd.specification,cd.manu_factory_name,
		ddd.amount as instock_amount,cd.packing_unit_name,cd.ret_price,cd.type,ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
		from drug_delivery_detail ddd
		left join drug_stock ds on ddd.drug_stock_id = ds.id
		left join clinic_drug cd on cd.id = ds.clinic_drug_id
		left join personnel p on p.id = ddd.operation_id
		where ddd.amount>0) as a
	where a.clinic_id=:clinic_id`

	selectSQL := `select 	
	outstock_date,order_number,outstock_way_name,supplier_name,
	outstock_operation_name,personnel_name,barcode,drug_name,
	specification,manu_factory_name,outstock_amount,packing_unit_name,ret_price,
	buy_price,serial,eff_date
	from (
		select to_char(dir.verify_time,'yyyy-mm-dd') as outstock_date,dir.order_number,dir.outstock_way_name,ds.supplier_name,
		p.name as outstock_operation_name,doc.name as personnel_name,cd.barcode,cd.name as drug_name,
		cd.specification,cd.manu_factory_name,diri.outstock_amount,cd.packing_unit_name,cd.ret_price,cd.type,
		ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
		from  drug_outstock_record_item diri
		left join drug_outstock_record dir on dir.id = diri.drug_outstock_record_id
		left join drug_stock ds on ds.id = diri.drug_stock_id
		left join clinic_drug cd on cd.id = ds.clinic_drug_id
		left join personnel p on p.id = dir.outstock_operation_id
		left join personnel doc on doc.id = dir.personnel_id
		where dir.verify_status='02'
		UNION all
		select to_char(dr.created_time,'yyyy-mm-dd') as outstock_date,dr.out_trade_no as order_number,'零售发药' as outstock_way_name,ds.supplier_name,
		p.name as outstock_operation_name,p.name as personnel_name,cd.barcode,cd.name as drug_name,cd.specification,cd.manu_factory_name,
		dr.amount as outstock_amount,cd.packing_unit_name,cd.ret_price,cd.type,ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
		from drug_retail dr
		left join drug_stock ds on ds.id = dr.drug_stock_id
		left join clinic_drug cd on cd.id = ds.clinic_drug_id
		left join drug_retail_pay_record drpr on drpr.out_trade_no = dr.out_trade_no
		left join personnel p on p.id = drpr.operation_id
		where dr.amount>0
		UNION all
		select to_char(ddd.created_time,'yyyy-mm-dd') as outstock_date,
		'MZ' || to_char(ddd.created_time, 'YYYYMMDDHH24MISS') as order_number,'门诊发药' as outstock_way_name,ds.supplier_name,
		p.name as outstock_operation_name,pa.name as personnel_name,cd.barcode,cd.name as drug_name,cd.specification,cd.manu_factory_name,
		ddd.amount as instock_amount,cd.packing_unit_name,cd.ret_price,cd.type,ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
		from drug_delivery_detail ddd
		left join mz_paid_orders mpo on mpo.id = ddd.mz_paid_orders_id
		left join clinic_triage_patient ctp on ctp.id = mpo.clinic_triage_patient_id
		left join clinic_patient cp on cp.id = ctp.clinic_patient_id
		left join patient pa on pa.id = cp.patient_id
		left join drug_stock ds on ddd.drug_stock_id = ds.id
		left join clinic_drug cd on cd.id = ds.clinic_drug_id
		left join personnel p on p.id = ddd.operation_id
		where ddd.amount>0) as a
	where a.clinic_id=:clinic_id`

	totalSQL := `select 
	sum(outstock_amount) as total_outstock_amount,
	sum(outstock_amount * buy_price) as total_buy_price
	from (
		select to_char(dir.verify_time,'yyyy-mm-dd') as outstock_date,dir.order_number,dir.outstock_way_name,ds.supplier_name,
		p.name as outstock_operation_name,doc.name as personnel_name,cd.barcode,cd.name as drug_name,
		cd.specification,cd.manu_factory_name,diri.outstock_amount,cd.packing_unit_name,cd.ret_price,cd.type,
		ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
		from  drug_outstock_record_item diri
		left join drug_outstock_record dir on dir.id = diri.drug_outstock_record_id
		left join drug_stock ds on ds.id = diri.drug_stock_id
		left join clinic_drug cd on cd.id = ds.clinic_drug_id
		left join personnel p on p.id = dir.outstock_operation_id
		left join personnel doc on doc.id = dir.personnel_id
		where dir.verify_status='02'
		UNION all
		select to_char(dr.created_time,'yyyy-mm-dd') as outstock_date,dr.out_trade_no as order_number,'零售发药' as outstock_way_name,ds.supplier_name,
		p.name as outstock_operation_name,p.name as personnel_name,cd.barcode,cd.name as drug_name,cd.specification,cd.manu_factory_name,
		dr.amount as outstock_amount,cd.packing_unit_name,cd.ret_price,cd.type,ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
		from drug_retail dr
		left join drug_stock ds on ds.id = dr.drug_stock_id
		left join clinic_drug cd on cd.id = ds.clinic_drug_id
		left join drug_retail_pay_record drpr on drpr.out_trade_no = dr.out_trade_no
		left join personnel p on p.id = drpr.operation_id
		where dr.amount>0
		UNION all
		select to_char(ddd.created_time,'yyyy-mm-dd') as outstock_date,
		'MZ' || to_char(ddd.created_time, 'YYYYMMDDHH24MISS') as order_number,'门诊发药' as outstock_way_name,ds.supplier_name,
		p.name as outstock_operation_name,p.name as personnel_name,cd.barcode,cd.name as drug_name,cd.specification,cd.manu_factory_name,
		ddd.amount as instock_amount,cd.packing_unit_name,cd.ret_price,cd.type,ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
		from drug_delivery_detail ddd
		left join drug_stock ds on ddd.drug_stock_id = ds.id
		left join clinic_drug cd on cd.id = ds.clinic_drug_id
		left join personnel p on p.id = ddd.operation_id
		where ddd.amount>0) as a
	where a.clinic_id=:clinic_id`

	if startDate != "" && endDate != "" {
		if startDate > endDate {
			ctx.JSON(iris.Map{"code": "-1", "msg": "开始日期必须大于结束日期"})
			return
		}

		countSQL += " and a.outstock_date between :start_date and :end_date"
		selectSQL += " and a.outstock_date between :start_date and :end_date"
		totalSQL += " and a.outstock_date between :start_date and :end_date"
	}

	if outstockWayName != "" {
		countSQL += " and a.outstock_way_name =:outstock_way_name"
		selectSQL += " and a.outstock_way_name =:outstock_way_name"
		totalSQL += " and a.outstock_way_name =:outstock_way_name"
	}

	if supplierName != "" {
		countSQL += " and a.supplier_name =:supplier_name"
		selectSQL += " and a.supplier_name =:supplier_name"
		totalSQL += " and a.supplier_name =:supplier_name"
	}

	if drugType != "" {
		countSQL += " and a.type =:drug_type"
		selectSQL += " and a.type =:drug_type"
		totalSQL += " and a.type =:drug_type"
	}

	if drugName != "" {
		countSQL += " and (a.drug_name ~*:drug_name or a.barcode ~*:drug_name)"
		selectSQL += " and (a.drug_name ~*:drug_name or a.barcode ~*:drug_name)"
		totalSQL += " and (a.drug_name ~*:drug_name or a.barcode ~*:drug_name)"
	}

	if serial != "" {
		countSQL += " and a.serial =:serial"
		selectSQL += " and a.serial =:serial"
		totalSQL += " and a.serial =:serial"
	}

	if personnelName != "" {
		countSQL += " and a.personnel_name ~*:personnel_name"
		selectSQL += " and a.personnel_name ~*:personnel_name"
		totalSQL += " and a.personnel_name ~*:personnel_name"
	}

	var queryOption = map[string]interface{}{
		"clinic_id":         ToNullInt64(clinicID),
		"outstock_way_name": ToNullString(outstockWayName),
		"supplier_name":     ToNullString(supplierName),
		"drug_name":         ToNullString(drugName),
		"drug_type":         ToNullString(drugType),
		"serial":            ToNullString(serial),
		"personnel_name":    ToNullString(personnelName),
		"start_date":        ToNullString(startDate),
		"end_date":          ToNullString(endDate),
		"offset":            ToNullInt64(offset),
		"limit":             ToNullInt64(limit),
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
	rows, _ := model.DB.NamedQuery(selectSQL+" order by a.outstock_date desc offset :offset limit :limit", queryOption)
	results = FormatSQLRowsToMapArray(rows)

	totalRows, _ := model.DB.NamedQuery(totalSQL, queryOption)
	totalResults := FormatSQLRowsToMapArray(totalRows)

	pageInfo["total_outstock_amount"] = totalResults[0]["total_outstock_amount"]
	pageInfo["total_buy_price"] = totalResults[0]["total_buy_price"]

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
	ds.stock_amount,
	(select sum(in_stock.instock_amount) as total_instock_amount
	from 
	(
		select diri.instock_amount,dir.verify_time,diri.serial,diri.eff_date,dir.supplier_name
		from drug_instock_record_item diri 
		left join drug_instock_record dir on dir.id = diri.drug_instock_record_id
		where dir.verify_status='02'
		UNION all
		select dr.amount as instock_amount,dr.created_time as verify_time,ds.serial,ds.eff_date,ds.supplier_name 
		from drug_retail dr
		left join drug_stock ds on ds.id = dr.drug_stock_id
		where dr.amount<0
		UNION all
		select ddd.amount as instock_amount,ddd.created_time as verify_time,ds.serial,ds.eff_date,ds.supplier_name 
		from drug_delivery_detail ddd
		left join drug_stock ds on ds.id = ddd.drug_stock_id
		where ddd.amount<0
		) as in_stock where in_stock.verify_time BETWEEN :start_date and :end_date
		and in_stock.serial = ds.serial and in_stock.eff_date = ds.eff_date 
		and in_stock.supplier_name = ds.supplier_name),
    (select sum(out_stock.outstock_amount) as total_outstock_amount 
    from
    (
		select dori.outstock_amount,dor.verify_time,dori.drug_stock_id
		from drug_outstock_record_item dori 
		left join drug_outstock_record dor on dor.id = dori.drug_outstock_record_id
		where dor.verify_status='02'
		UNION all
		select amount as outstock_amount,created_time as verify_time,drug_stock_id from drug_retail
		where amount>0
		UNION all
		select amount as outstock_amount,created_time as verify_time,drug_stock_id from drug_delivery_detail
		where amount>0
		) as out_stock where out_stock.drug_stock_id = ds.id and out_stock.verify_time BETWEEN :start_date and :end_date)
	from drug_stock ds
	left join clinic_drug cd on cd.id = ds.clinic_drug_id
	where ds.storehouse_id=:storehouse_id`

	if supplierName != "" {
		selectSQL += " and ds.supplier_name =:supplier_name"
	}

	if drugType != "" {
		selectSQL += " and cd.type =:drug_type"
	}

	if drugName != "" {
		selectSQL += " and (cd.name ~*:drug_name or cd.barcode ~*:drug_name)"
	}

	if serial != "" {
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

	var results []map[string]interface{}

	totalRows, _ := model.DB.NamedQuery(selectSQL+` group by 
	ds.id,
	cd.name,
	cd.barcode,
	cd.specification,
	cd.manu_factory_name,
	ds.supplier_name,
	cd.packing_unit_name,
	ds.serial,
	ds.eff_date,
	ds.buy_price`, queryOption)
	totalResults := FormatSQLRowsToMapArray(totalRows)

	rows, _ := model.DB.NamedQuery(selectSQL+` group by 
	ds.id,
	cd.name,
	cd.barcode,
	cd.specification,
	cd.manu_factory_name,
	ds.supplier_name,
	cd.packing_unit_name,
	ds.serial,
	ds.eff_date,
	ds.buy_price
	offset :offset limit :limit`, queryOption)
	results = FormatSQLRowsToMapArray(rows)

	pageInfo := map[string]interface{}{
		"total":  len(totalResults),
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

	countSQL := `select count(*) as total 	
	from (
		select to_char(mir.verify_time,'yyyy-mm-dd') as instock_date,mir.order_number,mir.instock_way_name,mir.supplier_name,
		p.name as instock_operation_name,cm.idc_code,cm.name as material_name,cm.specification,cm.manu_factory_name,
		miri.instock_amount,cm.unit_name,cm.ret_price,miri.buy_price,miri.serial,miri.eff_date,p.clinic_id
		from material_instock_record_item miri
		left join material_instock_record mir on mir.id = miri.material_instock_record_id
		left join clinic_material cm on cm.id = miri.clinic_material_id
		left join personnel p on p.id = mir.instock_operation_id
		where mir.verify_status='02'
    UNION all
    select to_char(ddd.created_time,'yyyy-mm-dd') as instock_date,
    'MZ' || to_char(ddd.created_time, 'YYYYMMDDHH24MISS') as order_number,'门诊退费' as instock_way_name,ds.supplier_name,
    p.name as instock_operation_name,cd.idc_code,cd.name as material_name,cd.specification,cd.manu_factory_name,
    ddd.amount as instock_amount,cd.unit_name,cd.ret_price,ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
    from material_delivery_detail ddd
    left join material_stock ds on ddd.material_stock_id = ds.id
    left join clinic_material cd on cd.id = ds.clinic_material_id
		left join personnel p on p.id = ddd.operation_id
		where ddd.amount<0) as a
  where a.clinic_id=:clinic_id`

	selectSQL := `select instock_date,order_number,instock_way_name,supplier_name,
	instock_operation_name,idc_code,material_name,specification,manu_factory_name,
	instock_amount,unit_name,ret_price,buy_price,serial,eff_date
	from (
		select to_char(mir.verify_time,'yyyy-mm-dd') as instock_date,mir.order_number,mir.instock_way_name,mir.supplier_name,
		p.name as instock_operation_name,cm.idc_code,cm.name as material_name,cm.specification,cm.manu_factory_name,
		miri.instock_amount,cm.unit_name,cm.ret_price,miri.buy_price,miri.serial,miri.eff_date,p.clinic_id
		from material_instock_record_item miri
		left join material_instock_record mir on mir.id = miri.material_instock_record_id
		left join clinic_material cm on cm.id = miri.clinic_material_id
		left join personnel p on p.id = mir.instock_operation_id
		where mir.verify_status='02'
    UNION all
    select to_char(ddd.created_time,'yyyy-mm-dd') as instock_date,
    'MZ' || to_char(ddd.created_time, 'YYYYMMDDHH24MISS') as order_number,'门诊退费' as instock_way_name,ds.supplier_name,
    p.name as instock_operation_name,cd.idc_code,cd.name as material_name,cd.specification,cd.manu_factory_name,
    ddd.amount as instock_amount,cd.unit_name,cd.ret_price,ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
    from material_delivery_detail ddd
    left join material_stock ds on ddd.material_stock_id = ds.id
    left join clinic_material cd on cd.id = ds.clinic_material_id
		left join personnel p on p.id = ddd.operation_id
		where ddd.amount<0) as a
  where a.clinic_id=:clinic_id`

	totalSQL := `select 
	sum(instock_amount) as total_instock_amount,
	sum(instock_amount * buy_price) as total_buy_price
	from (
		select to_char(mir.verify_time,'yyyy-mm-dd') as instock_date,mir.order_number,mir.instock_way_name,mir.supplier_name,
		p.name as instock_operation_name,cm.idc_code,cm.name as material_name,cm.specification,cm.manu_factory_name,
		miri.instock_amount,cm.unit_name,cm.ret_price,miri.buy_price,miri.serial,miri.eff_date,p.clinic_id
		from material_instock_record_item miri
		left join material_instock_record mir on mir.id = miri.material_instock_record_id
		left join clinic_material cm on cm.id = miri.clinic_material_id
		left join personnel p on p.id = mir.instock_operation_id
		where mir.verify_status='02'
    UNION all
    select to_char(ddd.created_time,'yyyy-mm-dd') as instock_date,
    'MZ' || to_char(ddd.created_time, 'YYYYMMDDHH24MISS') as order_number,'门诊退费' as instock_way_name,ds.supplier_name,
    p.name as instock_operation_name,cd.idc_code,cd.name as material_name,cd.specification,cd.manu_factory_name,
    ddd.amount as instock_amount,cd.unit_name,cd.ret_price,ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
    from material_delivery_detail ddd
    left join material_stock ds on ddd.material_stock_id = ds.id
    left join clinic_material cd on cd.id = ds.clinic_material_id
		left join personnel p on p.id = ddd.operation_id
		where ddd.amount<0) as a
  where a.clinic_id=:clinic_id`

	if startDate != "" && endDate != "" {
		if startDate > endDate {
			ctx.JSON(iris.Map{"code": "-1", "msg": "开始日期必须大于结束日期"})
			return
		}

		countSQL += " and a.instock_date between :start_date and :end_date"
		selectSQL += " and a.instock_date between :start_date and :end_date"
		totalSQL += " and a.instock_date between :start_date and :end_date"
	}

	if instockWayName != "" {
		countSQL += " and a.instock_way_name =:instock_way_name"
		selectSQL += " and a.instock_way_name =:instock_way_name"
		totalSQL += " and a.instock_way_name =:instock_way_name"
	}

	if supplierName != "" {
		countSQL += " and a.supplier_name =:supplier_name"
		selectSQL += " and a.supplier_name =:supplier_name"
		totalSQL += " and a.supplier_name =:supplier_name"
	}

	if materialName != "" {
		countSQL += " and (a.name ~*:material_name or a.py_code ~*:material_name)"
		selectSQL += " and (a.name ~*:material_name or a.py_code ~*:material_name)"
		totalSQL += " and (a.name ~*:material_name or a.py_code ~*:material_name)"
	}

	if serial != "" {
		countSQL += " and a.serial =:serial"
		selectSQL += " and a.serial =:serial"
		totalSQL += " and a.serial =:serial"
	}

	var queryOption = map[string]interface{}{
		"clinic_id":        ToNullInt64(clinicID),
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
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.NamedQuery(selectSQL+" order by a.instock_date desc offset :offset limit :limit", queryOption)
	results = FormatSQLRowsToMapArray(rows)

	totalRows, _ := model.DB.NamedQuery(totalSQL, queryOption)
	totalResults := FormatSQLRowsToMapArray(totalRows)

	pageInfo["total_instock_amount"] = totalResults[0]["total_instock_amount"]
	pageInfo["total_buy_price"] = totalResults[0]["total_buy_price"]

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

	countSQL := `select count(*) as total 	
	from (
		select to_char(mir.verify_time,'yyyy-mm-dd') as outstock_date,mir.order_number,mir.outstock_way_name,ms.supplier_name,
		p.name as outstock_operation_name,doc.name as personnel_name,cm.idc_code,cm.name as material_name,
		cm.specification,cm.manu_factory_name,miri.outstock_amount,cm.unit_name,cm.ret_price,
		ms.buy_price,ms.serial,ms.eff_date,p.clinic_id
		from  material_outstock_record_item miri
		left join material_outstock_record mir on mir.id = miri.material_outstock_record_id
		left join material_stock ms on ms.id = miri.material_stock_id
		left join clinic_material cm on cm.id = ms.clinic_material_id
		left join personnel p on p.id = mir.outstock_operation_id
		left join personnel doc on doc.id = mir.personnel_id
		where mir.verify_status='02'
		UNION all
    select to_char(ddd.created_time,'yyyy-mm-dd') as outstock_date,
    'MZ' || to_char(ddd.created_time, 'YYYYMMDDHH24MISS') as order_number,'门诊收费' as outstock_way_name,ds.supplier_name,
    p.name as outstock_operation_name,pa.name as personnel_name,cd.idc_code,cd.name as material_name,cd.specification,cd.manu_factory_name,
    ddd.amount as outstock_amount,cd.unit_name,cd.ret_price,ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
    from material_delivery_detail ddd
    left join mz_paid_orders mpo on mpo.id = ddd.mz_paid_orders_id
    left join clinic_triage_patient ctp on ctp.id = mpo.clinic_triage_patient_id
    left join clinic_patient cp on cp.id = ctp.clinic_patient_id
    left join patient pa on pa.id = cp.patient_id
    left join material_stock ds on ddd.material_stock_id = ds.id
    left join clinic_material cd on cd.id = ds.clinic_material_id
		left join personnel p on p.id = ddd.operation_id
		where ddd.amount>0) as a
	where a.clinic_id=:clinic_id`
	selectSQL := `select 	
	outstock_date,order_number,outstock_way_name,supplier_name,
	outstock_operation_name,personnel_name,idc_code,material_name,
	specification,manu_factory_name,outstock_amount,unit_name,ret_price,
	buy_price,serial,eff_date
	from (
		select to_char(mir.verify_time,'yyyy-mm-dd') as outstock_date,mir.order_number,mir.outstock_way_name,ms.supplier_name,
		p.name as outstock_operation_name,doc.name as personnel_name,cm.idc_code,cm.name as material_name,
		cm.specification,cm.manu_factory_name,miri.outstock_amount,cm.unit_name,cm.ret_price,
		ms.buy_price,ms.serial,ms.eff_date,p.clinic_id
		from  material_outstock_record_item miri
		left join material_outstock_record mir on mir.id = miri.material_outstock_record_id
		left join material_stock ms on ms.id = miri.material_stock_id
		left join clinic_material cm on cm.id = ms.clinic_material_id
		left join personnel p on p.id = mir.outstock_operation_id
		left join personnel doc on doc.id = mir.personnel_id
		where mir.verify_status='02'
		UNION all
    select to_char(ddd.created_time,'yyyy-mm-dd') as outstock_date,
    'MZ' || to_char(ddd.created_time, 'YYYYMMDDHH24MISS') as order_number,'门诊收费' as outstock_way_name,ds.supplier_name,
    p.name as outstock_operation_name,pa.name as personnel_name,cd.idc_code,cd.name as material_name,cd.specification,cd.manu_factory_name,
    ddd.amount as outstock_amount,cd.unit_name,cd.ret_price,ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
    from material_delivery_detail ddd
    left join mz_paid_orders mpo on mpo.id = ddd.mz_paid_orders_id
    left join clinic_triage_patient ctp on ctp.id = mpo.clinic_triage_patient_id
    left join clinic_patient cp on cp.id = ctp.clinic_patient_id
    left join patient pa on pa.id = cp.patient_id
    left join material_stock ds on ddd.material_stock_id = ds.id
    left join clinic_material cd on cd.id = ds.clinic_material_id
		left join personnel p on p.id = ddd.operation_id
		where ddd.amount>0) as a
	where a.clinic_id=:clinic_id`
	totalSQL := `select 
	sum(outstock_amount) as total_outstock_amount,
	sum(outstock_amount * buy_price) as total_buy_price
	from (
		select to_char(mir.verify_time,'yyyy-mm-dd') as outstock_date,mir.order_number,mir.outstock_way_name,ms.supplier_name,
		p.name as outstock_operation_name,doc.name as personnel_name,cm.idc_code,cm.name as material_name,
		cm.specification,cm.manu_factory_name,miri.outstock_amount,cm.unit_name,cm.ret_price,
		ms.buy_price,ms.serial,ms.eff_date,p.clinic_id
		from  material_outstock_record_item miri
		left join material_outstock_record mir on mir.id = miri.material_outstock_record_id
		left join material_stock ms on ms.id = miri.material_stock_id
		left join clinic_material cm on cm.id = ms.clinic_material_id
		left join personnel p on p.id = mir.outstock_operation_id
		left join personnel doc on doc.id = mir.personnel_id
		where mir.verify_status='02'
		UNION all
    select to_char(ddd.created_time,'yyyy-mm-dd') as outstock_date,
    'MZ' || to_char(ddd.created_time, 'YYYYMMDDHH24MISS') as order_number,'门诊收费' as outstock_way_name,ds.supplier_name,
    p.name as outstock_operation_name,pa.name as personnel_name,cd.idc_code,cd.name as material_name,cd.specification,cd.manu_factory_name,
    ddd.amount as outstock_amount,cd.unit_name,cd.ret_price,ds.buy_price,ds.serial,ds.eff_date,p.clinic_id
    from material_delivery_detail ddd
    left join mz_paid_orders mpo on mpo.id = ddd.mz_paid_orders_id
    left join clinic_triage_patient ctp on ctp.id = mpo.clinic_triage_patient_id
    left join clinic_patient cp on cp.id = ctp.clinic_patient_id
    left join patient pa on pa.id = cp.patient_id
    left join material_stock ds on ddd.material_stock_id = ds.id
    left join clinic_material cd on cd.id = ds.clinic_material_id
		left join personnel p on p.id = ddd.operation_id
		where ddd.amount>0) as a
	where a.clinic_id=:clinic_id`

	if startDate != "" && endDate != "" {
		if startDate > endDate {
			ctx.JSON(iris.Map{"code": "-1", "msg": "开始日期必须大于结束日期"})
			return
		}

		countSQL += " and a.outstock_date between :start_date and :end_date"
		selectSQL += " and a.outstock_date between :start_date and :end_date"
		totalSQL += " and a.outstock_date between :start_date and :end_date"
	}

	if outstockWayName != "" {
		countSQL += " and a.outstock_way_name =:outstock_way_name"
		selectSQL += " and a.outstock_way_name =:outstock_way_name"
		totalSQL += " and a.outstock_way_name =:outstock_way_name"
	}

	if supplierName != "" {
		countSQL += " and a.supplier_name =:supplier_name"
		selectSQL += " and a.supplier_name =:supplier_name"
		totalSQL += " and a.supplier_name =:supplier_name"
	}

	if materialName != "" {
		countSQL += " and (a.material_name ~*:material_name or a.py_code ~*:material_name)"
		selectSQL += " and (a.material_name ~*:material_name or a.py_code ~*:material_name)"
		totalSQL += " and (a.material_name ~*:material_name or a.py_code ~*:material_name)"
	}

	if serial != "" {
		countSQL += " and a.serial =:serial"
		selectSQL += " and a.serial =:serial"
		totalSQL += " and a.serial =:serial"
	}

	if personnelName != "" {
		countSQL += " and a.personnel_name ~*:personnel_name"
		selectSQL += " and a.personnel_name ~*:personnel_name"
		totalSQL += " and a.personnel_name ~*:personnel_name"
	}

	var queryOption = map[string]interface{}{
		"clinic_id":        ToNullInt64(clinicID),
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
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.NamedQuery(selectSQL+" order by a.outstock_date desc offset :offset limit :limit", queryOption)
	results = FormatSQLRowsToMapArray(rows)

	totalRows, _ := model.DB.NamedQuery(totalSQL, queryOption)
	totalResults := FormatSQLRowsToMapArray(totalRows)

	pageInfo["total_outstock_amount"] = totalResults[0]["total_outstock_amount"]
	pageInfo["total_buy_price"] = totalResults[0]["total_buy_price"]

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
	ms.stock_amount,
	(select sum(in_stock.instock_amount) as total_instock_amount
	FROM
	(
	select miri.instock_amount,mir.verify_time,miri.serial,miri.eff_date,mir.supplier_name
	from material_instock_record_item miri 
	left join material_instock_record mir on mir.id = miri.material_instock_record_id
	where mir.verify_status='02'
	UNION all
	select ddd.amount as instock_amount,ddd.created_time as verify_time,ds.serial,ds.eff_date,ds.supplier_name 
	from material_delivery_detail ddd
	left join material_stock ds on ds.id = ddd.material_stock_id
	where ddd.amount<0
	) as in_stock where in_stock.verify_time BETWEEN :start_date and :end_date
	and in_stock.serial = ms.serial and in_stock.eff_date = ms.eff_date 
	and in_stock.supplier_name = ms.supplier_name),
	(select sum(out_stock.outstock_amount) as total_outstock_amount
	from
	(
	select dori.outstock_amount,dor.verify_time,dori.material_stock_id
	from material_outstock_record_item dori 
	left join material_outstock_record dor on dor.id = dori.material_outstock_record_id
	where dor.verify_status='02'
	UNION all
	select amount as outstock_amount,created_time as verify_time,material_stock_id from material_delivery_detail
	where amount>0
	) as out_stock where out_stock.verify_time BETWEEN :start_date and :end_date and out_stock.material_stock_id = ms.id)
	from material_stock ms
	left join clinic_material cm on cm.id = ms.clinic_material_id
	where ms.storehouse_id=1`

	if supplierName != "" {
		selectSQL += " and ms.supplier_name =:supplier_name"
	}

	if materialName != "" {
		selectSQL += " and (cm.name ~*:material_name or cm.barcode ~*:material_name)"
	}

	if serial != "" {
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

	var results []map[string]interface{}

	totalRows, _ := model.DB.NamedQuery(selectSQL+` group by 
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
	ms.buy_price`, queryOption)
	totalResults := FormatSQLRowsToMapArray(totalRows)

	rows, _ := model.DB.NamedQuery(selectSQL+` group by 
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
	ms.buy_price offset :offset limit :limit`, queryOption)
	results = FormatSQLRowsToMapArray(rows)

	pageInfo := map[string]interface{}{
		"total":  len(totalResults),
		"offset": offset,
		"limit":  limit,
	}

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}
