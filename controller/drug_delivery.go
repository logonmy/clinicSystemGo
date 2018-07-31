package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris"
)

// DrugDeliveryList 获取药品记录（包括 待发药，已发药，已退药）
func DrugDeliveryList(ctx iris.Context) {
	status := ctx.PostValue("order_status")
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	if status == "40" || status == "30" {
		drugDeliveryListNew(ctx)
		return
	}
	if clinicTriagePatientID == "" || status == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	timeNow := time.Now().Format("2006-01-02")

	SQL := `FROM mz_paid_orders mpo 
	left join clinic_drug cd on cd.id = mpo.charge_project_id 
	left join prescription_chinese_patient pcp on pcp.order_sn = mpo.order_sn
	left join (select clinic_drug_id, sum(stock_amount) as stock_amount from drug_stock where eff_date > '` + timeNow + `' group by clinic_drug_id ) ds on ds.clinic_drug_id = cd.id 
	left join drug_delivery_record_item ddri on ddri.mz_paid_orders_id = mpo.id 
	where mpo.clinic_triage_patient_id = $1 and mpo.order_status = $2 and mpo.charge_project_type_id in (1,2)`

	querysql := "select pcp.amount as prescription_amount,mpo.order_sn,ddri.remark,mpo.order_status,mpo.id,mpo.name,mpo.amount,mpo.charge_project_type_id,cd.specification,cd.manu_factory_name,cd.dose_form_name,ds.stock_amount " + SQL

	rows, _ := model.DB.Queryx(querysql, clinicTriagePatientID, status)
	results := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})

}

// DrugDeliveryList 获取药品记录（已发药，已退药）
func drugDeliveryListNew(ctx iris.Context) {
	status := ctx.PostValue("order_status")
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if clinicTriagePatientID == "" || status == "" {
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

	SQL := `FROM mz_paid_orders mpo 
	left join clinic_drug cd on cd.id = mpo.charge_project_id 
	left join prescription_chinese_patient pcp on pcp.order_sn = mpo.order_sn 
	left join (select clinic_drug_id, sum(stock_amount) as stock_amount from drug_stock group by clinic_drug_id ) ds on ds.clinic_drug_id = cd.id 
	left join drug_delivery_record_item ddri on ddri.mz_paid_orders_id = mpo.id 
	where mpo.clinic_triage_patient_id = $1 and mpo.order_status = $2 and mpo.charge_project_type_id in (1,2)`
	countsql := "select count(mpo.*) as total,string_agg(cast ( mpo.id as TEXT ),',') as ids " + SQL

	total := model.DB.QueryRowx(countsql, clinicTriagePatientID, status)
	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit
	querysql := "select pcp.amount as prescription_amount,mpo.order_sn,ddri.remark,mpo.order_status,mpo.id,mpo.name,mpo.amount,mpo.charge_project_type_id,cd.specification,cd.manu_factory_name,cd.dose_form_name,ds.stock_amount " + SQL + " offset $3 limit $4"

	rows, _ := model.DB.Queryx(querysql, clinicTriagePatientID, status, offset, limit)
	allSelectStatus := true
	results := FormatSQLRowsToMapArray(rows)
	for _, item := range results {
		if item["stock_amount"] == nil {
			item["stock_amount"] = int64(0)
		}
		if item["amount"] == nil {
			item["amount"] = int64(0)
		}

		allSelectStatus = item["stock_amount"].(int64) >= item["amount"].(int64) && allSelectStatus
	}
	pageInfo["allSelectStatus"] = allSelectStatus
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})

}

// DrugDeliveryWaiting 获取用户待发药的分诊记录
func DrugDeliveryWaiting(ctx iris.Context) {
	drugDeliveryTriageList(ctx, "10")
}

// DrugDeliveryIssued 获取用户已发药的分诊记录
func DrugDeliveryIssued(ctx iris.Context) {
	drugDeliveryTriageList(ctx, "30")
}

// DrugDeliveryRefund 获取用户已退药的分诊记录
func DrugDeliveryRefund(ctx iris.Context) {
	drugDeliveryTriageList(ctx, "40")
}

// 获取用户各状态的分诊记录
func drugDeliveryTriageList(ctx iris.Context, status string) {
	keyword := ctx.PostValue("keyword")
	clinicID := ctx.PostValue("clinic_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")
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

	if startDate == "" || endDate == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请输入正确的时间范围"})
		return
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

	sql := ` from clinic_triage_patient ctp 
	left join clinic_patient cp on cp.id = ctp.clinic_patient_id 
	left join personnel doc on doc.id = ctp.doctor_id 
	left join department d on d.id = ctp.department_id  
	left join patient p on p.id = cp.patient_id 
	left join clinic_triage_patient_operation register on ctp.id = register.clinic_triage_patient_id and register.type = 10
	left join personnel triage_personnel on triage_personnel.id = register.personnel_id 
	left join (select clinic_triage_patient_id,count(*) as total_count from mz_paid_orders where charge_project_type_id in (1,2) and order_status = '` + status + `' group by(clinic_triage_patient_id)) up on up.clinic_triage_patient_id = ctp.id 
	where up.total_count > 0 AND cp.clinic_id=$1 AND ctp.updated_time BETWEEN $2 and $3 AND (p.name ~$4 OR p.cert_no ~$4 OR p.phone ~$4) `

	countsql := `select count(*) as total` + sql
	querysql := `select 
	up.total_count,
	ctp.id as clinic_triage_patient_id,
	ctp.clinic_patient_id as clinic_patient_id,
	ctp.updated_time,
	ctp.created_time as register_time,
	triage_personnel.name as register_personnel_name,
	ctp.status,
	ctp.visit_date,
	ctp.register_type,
	p.name as patient_name,
	p.birthday,
	p.sex,
	p.phone,
	cp.patient_id,
	doc.name as doctor_name,
	d.name as department_name ` + sql + ` order by ctp.updated_time ASC offset $5 limit $6`

	total := model.DB.QueryRowx(countsql, clinicID, startDate, endDate, keyword)

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.Queryx(querysql, clinicID, startDate, endDate, keyword, offset, limit)
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

// DrugDeliveryRecordCreate 创建发药记录
func DrugDeliveryRecordCreate(ctx iris.Context) {
	triagePatient := ctx.PostValue("clinic_triage_patient_id")
	operation := ctx.PostValue("operation_id")
	items := ctx.PostValue("items")

	var results []map[string]interface{}
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	if triagePatient == "" || operation == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-2", "msg": "缺少参数"})
		return
	}

	tx, txErr := model.DB.Beginx()
	if txErr != nil {
		ctx.JSON(iris.Map{"code": "-3", "msg": txErr.Error()})
		return
	}

	var recordID interface{}
	err1 := tx.QueryRow("INSERT INTO drug_delivery_record (clinic_triage_patient_id, operation_id) VALUES ($1, $2) RETURNING id", triagePatient, operation).Scan(&recordID)
	if err1 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-4", "msg": err1.Error()})
		return
	}

	for _, item := range results {
		orderID := item["mz_paid_orders_id"]
		item["drug_delivery_record_id"] = recordID
		_, err := tx.NamedExec(`INSERT INTO drug_delivery_record_item (drug_delivery_record_id, mz_paid_orders_id,remark) VALUES (:drug_delivery_record_id, :mz_paid_orders_id, :remark)`, item)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-5", "msg": err.Error()})
			return
		}
		_, err1 := tx.Exec(`UPDATE mz_paid_orders set order_status = '30' where id = $1`, orderID)
		if err1 != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-6", "msg": err1.Error()})
			return
		}

		row := tx.QueryRowx(`select * from mz_paid_orders where id = $1`, orderID)
		rowMap := FormatSQLRowToMap(row)

		err2 := updateDrugStock2(tx, rowMap["charge_project_id"].(int64), rowMap["amount"].(int64))

		if err2 != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-6", "msg": err2.Error()})
			return
		}
	}

	erre := tx.Commit()
	if erre != nil {
		ctx.JSON(iris.Map{"code": "-7", "msg": erre.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "msg": "操作成功"})

}

func updateDrugStock2(tx *sqlx.Tx, clinicDrugID int64, amount int64) error {
	if amount < 0 {
		return errors.New("库存数量有误")
	}
	if amount == 0 {
		return nil
	}

	timeNow := time.Now().Format("2006-01-02")
	row := model.DB.QueryRowx("select * from drug_stock where clinic_drug_id = $1 and stock_amount > 0 and eff_date > $2 ORDER by created_time asc limit 1", clinicDrugID, timeNow)
	rowMap := FormatSQLRowToMap(row)
	if rowMap["stock_amount"] == nil {
		return errors.New("库存不足")
	}

	stockAmount := rowMap["stock_amount"].(int64)

	if stockAmount >= amount {
		_, err := tx.Exec("update drug_stock set stock_amount = $1 where id = $2", stockAmount-amount, rowMap["id"])
		if err != nil {
			tx.Rollback()
			return err
		}

		return nil
	}
	_, err := tx.Exec("update drug_stock set 0 where id = $1", rowMap["id"])
	if err != nil {
		tx.Rollback()
		return err
	}

	return updateDrugStock2(tx, clinicDrugID, amount-stockAmount)
}

// DrugDeliveryRecordRefund 退药
func DrugDeliveryRecordRefund(ctx iris.Context) {
	triagePatient := ctx.PostValue("clinic_triage_patient_id")
	operation := ctx.PostValue("operation_id")
	items := ctx.PostValue("items")

	var results []map[string]interface{}
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	if triagePatient == "" || operation == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-2", "msg": "缺少参数"})
		return
	}

	tx, txErr := model.DB.Beginx()
	if txErr != nil {
		ctx.JSON(iris.Map{"code": "-3", "msg": txErr.Error()})
		return
	}

	var recordID interface{}
	err1 := tx.QueryRow("INSERT INTO drug_delivery_refund_record (clinic_triage_patient_id, operation_id) VALUES ($1, $2) RETURNING id", triagePatient, operation).Scan(&recordID)
	if err1 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-4", "msg": err1.Error()})
		return
	}

	for _, item := range results {
		orderID := item["mz_paid_orders_id"]
		item["drug_delivery_refund_record_id"] = recordID
		_, err := tx.NamedExec(`INSERT INTO drug_delivery_refund_record_item (drug_delivery_refund_record_id, mz_paid_orders_id) VALUES (:drug_delivery_refund_record_id, :mz_paid_orders_id)`, item)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-5", "msg": err.Error()})
			return
		}
		_, err1 := tx.Exec(`UPDATE mz_paid_orders set order_status = '40' where id = $1`, orderID)
		if err1 != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-6", "msg": err1.Error()})
			return
		}

		row := tx.QueryRowx(`select * from mz_paid_orders where id = $1`, orderID)
		rowMap := FormatSQLRowToMap(row)

		_, err2 := tx.Exec(`UPDATE drug_stock set stock_amount = stock_amount + $1 where id in (select id from drug_stock where clinic_drug_id = $2 order by eff_date DESC limit 1)`, rowMap["amount"], rowMap["charge_project_id"])
		if err2 != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-6", "msg": err2.Error()})
			return
		}
	}

	erre := tx.Commit()
	if erre != nil {
		ctx.JSON(iris.Map{"code": "-7", "msg": erre.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "msg": "操作成功"})

}

// DrugDeliveryRecordList 查询发药记录
func DrugDeliveryRecordList(ctx iris.Context) {
	triagePatient := ctx.PostValue("clinic_triage_patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}

	SQL := ` from drug_delivery_record ddc 
	left join clinic_triage_patient ctp on ctp.id = ddc.clinic_triage_patient_id 
	left join personnel doc on doc.id = ctp.doctor_id 
	left join personnel op on op.id = ddc.operation_id 
	left join (select drug_delivery_record_id,min(mz_paid_orders_id) as order_id from drug_delivery_record_item group by drug_delivery_record_id) mp on mp.drug_delivery_record_id = ddc.id
	left join mz_paid_orders mpo on mpo.id = mp.order_id 
	where ddc.clinic_triage_patient_id = $1`

	countsql := "select count(ddc.*) as total " + SQL
	total := model.DB.QueryRowx(countsql, triagePatient)

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	querysql := `select ctp.visit_date,
	 ddc.id as drug_delivery_record_id,
	 doc.name as doctor_name,
	 op.name as opration_name,
	 ddc.created_time,
	 mpo.name as project_name 
	` + SQL + " offset $2 limit $3"

	rows, _ := model.DB.Queryx(querysql, triagePatient, offset, limit)

	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})

}

// DrugDeliveryRecordRefundList 查询退药记录
func DrugDeliveryRecordRefundList(ctx iris.Context) {
	triagePatient := ctx.PostValue("clinic_triage_patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}

	SQL := ` from drug_delivery_refund_record ddc 
	left join clinic_triage_patient ctp on ctp.id = ddc.clinic_triage_patient_id 
	left join personnel doc on doc.id = ctp.doctor_id 
	left join personnel op on op.id = ddc.operation_id 
	left join (select drug_delivery_refund_record_id,min(mz_paid_orders_id) as order_id from drug_delivery_refund_record_item group by drug_delivery_refund_record_id) mp on mp.drug_delivery_refund_record_id = ddc.id
	left join mz_paid_orders mpo on mpo.id = mp.order_id 
	where ddc.clinic_triage_patient_id = $1`

	countsql := "select count(ddc.*) as total " + SQL
	total := model.DB.QueryRowx(countsql, triagePatient)

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	querysql := `select ctp.visit_date,
	 ddc.id as drug_delivery_refund_record_id,
	 doc.name as doctor_name,
	 op.name as opration_name,
	 ddc.created_time,
	 mpo.name as project_name 
	` + SQL + " offset $2 limit $3"

	rows, _ := model.DB.Queryx(querysql, triagePatient, offset, limit)

	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})

}

// DrugDeliveryRecordDetail 查询发药记录详情
func DrugDeliveryRecordDetail(ctx iris.Context) {
	recordID := ctx.PostValue("drug_delivery_record_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if recordID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}

	SQL := ` from drug_delivery_record_item ddri
	left join mz_paid_orders mpo on mpo.id = ddri.mz_paid_orders_id 
	left join clinic_drug cd on mpo.charge_project_id = cd.id 
	where drug_delivery_record_id = $1`

	query := `select mpo.order_status,mpo.name,mpo.amount,mpo.charge_project_type_id,ddri.remark,
	cd.specification,cd.manu_factory_name 
	` + SQL + ` offset $2 limit $3`
	countsql := `select count(*) as total ` + SQL

	total := model.DB.QueryRowx(countsql, recordID)
	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.Queryx(query, recordID, offset, limit)
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

// DrugDeliveryRecordRefundDetail 查询退药记录详情
func DrugDeliveryRecordRefundDetail(ctx iris.Context) {
	recordID := ctx.PostValue("drug_delivery_refund_record_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if recordID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}

	SQL := ` from drug_delivery_refund_record_item ddri
	left join mz_paid_orders mpo on mpo.id = ddri.mz_paid_orders_id 
	left join clinic_drug cd on mpo.charge_project_id = cd.id 
	where drug_delivery_refund_record_id = $1`

	query := `select mpo.order_status,mpo.name,mpo.amount,mpo.charge_project_type_id,
	cd.specification,cd.manu_factory_name 
	` + SQL + ` offset $2 limit $3`
	countsql := `select count(*) as total ` + SQL

	total := model.DB.QueryRowx(countsql, recordID)
	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.Queryx(query, recordID, offset, limit)
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

// DrugDeliveryStockList 获取药品库存记录
func DrugDeliveryStockList(ctx iris.Context) {
	status := ctx.PostValue("order_status")
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")

	if clinicTriagePatientID == "" || status == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	querysql := `select mpo.id as mz_paid_orders_id,mpo.name,mpo.amount,mpo.charge_project_type_id,
	cd.specification,cd.manu_factory_name,cd.dose_form_name,
	ds.stock_amount,ds.eff_date,ds.serial,ds.id as drug_stock_id
	FROM mz_paid_orders mpo 
	left join clinic_drug cd on cd.id = mpo.charge_project_id 
	left join drug_stock ds on ds.clinic_drug_id = mpo.charge_project_id
	where mpo.clinic_triage_patient_id = $1 and mpo.order_status = $2 and mpo.charge_project_type_id in (1,2)`

	rows, _ := model.DB.Queryx(querysql, clinicTriagePatientID, status)
	results := FormatSQLRowsToMapArray(rows)

	for _, item := range results {
		if item["stock_amount"] == nil {
			item["stock_amount"] = int64(0)
		}
		if item["amount"] == nil {
			item["amount"] = int64(0)
		}
	}
	ctx.JSON(iris.Map{"code": "200", "data": results})

}
