package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/kataras/iris"
)

// DrugDeliveryList 获取药品记录（包括 待发药，已发药，已退药）
func DrugDeliveryList(ctx iris.Context) {
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
	left join drug_stock ds on ds.clinic_drug_id = cd.id 
	where mpo.clinic_triage_patient_id = $1 and mpo.order_status = $2 and mpo.charge_project_type_id in (1,2)`
	countsql := "select count(mpo.*) as total,string_agg(cast ( mpo.id as TEXT ),',') as ids " + SQL

	total := model.DB.QueryRowx(countsql, clinicTriagePatientID, status)
	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit
	querysql := "select mpo.id,mpo.name,mpo.amount,mpo.charge_project_type_id,cd.specification,cd.manu_factory_name,cd.dose_form_name,ds.stock_amount " + SQL + " offset $3 limit $4"

	rows, _ := model.DB.Queryx(querysql, clinicTriagePatientID, status, offset, limit)
	allSelectStatus := true
	results := FormatSQLRowsToMapArray(rows)
	for _, item := range results {
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
	doc.name as doctor_name,
	d.name as department_name ` + sql + `offset $5 limit $6`

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

	fmt.Println(triagePatient, operation, items)

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

	fmt.Println(recordID)

	for _, item := range results {
		orderID := item["mz_paid_orders_id"]
		item["drug_delivery_record_id"] = recordID
		fmt.Println(item)
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

		// 这里有个坑 ， 发药没有指定药房。
		_, err2 := tx.Exec(`UPDATE drug_stock set stock_amount = stock_amount - $1 where clinic_drug_id = $2`, rowMap["amount"], rowMap["charge_project_id"])
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
