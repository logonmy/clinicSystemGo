package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/kataras/iris"
)

// LaboratoryTriageList 获取检验列表
func LaboratoryTriageList(ctx iris.Context) {
	status := ctx.PostValue("order_status")
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")

	if clinicTriagePatientID == "" || status == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select lp.id as laboratory_patient_id,lp.clinic_triage_patient_id,
	cl.name as clinic_laboratory_name,lp.clinic_laboratory_id,lpr.id as laboratory_patient_record_id,lpr.remark
	FROM laboratory_patient lp 
	left join clinic_laboratory cl on cl.id = lp.clinic_laboratory_id
	left join mz_paid_orders mo on mo.clinic_triage_patient_id = lp.clinic_triage_patient_id and mo.charge_project_type_id=3 and lp.clinic_laboratory_id=mo.charge_project_id
	left join laboratory_patient_record lpr on lpr.laboratory_patient_id = lp.id
	where lp.clinic_triage_patient_id = $1 and lp.order_status=$2`

	rows, _ := model.DB.Queryx(selectSQL, clinicTriagePatientID, status)
	results := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})
}

// LaboratoryTriageDetail 获取检验详情
func LaboratoryTriageDetail(ctx iris.Context) {
	laboratoryPatientID := ctx.PostValue("laboratory_patient_id")
	clinicLaboratoryID := ctx.PostValue("clinic_laboratory_id")

	if laboratoryPatientID == "" || clinicLaboratoryID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectAssociationSQL := `select cls.clinic_laboratory_item_id,cli.name,cli.en_name,cli.unit_name,cli.is_special,cls.default_result,
	cli.data_type,clir.reference_sex,clir.stomach_status,clir.is_pregnancy,clir.reference_max,clir.reference_min,cli.status
	from clinic_laboratory_association cls
	left join clinic_laboratory_item cli on cls.clinic_laboratory_item_id = cli.id
	left join clinic_laboratory_item_reference clir on clir.clinic_laboratory_item_id = cli.id
	where cls.clinic_laboratory_id=$1`

	arows, _ := model.DB.Queryx(selectAssociationSQL, clinicLaboratoryID)
	aresults := FormatSQLRowsToMapArray(arows)
	laboratoryItems := FormatLaboratoryItem(aresults)

	selectSQL := `select lpri.clinic_laboratory_item_id,lpri.result_inspection,lpri.property_inspection,cli.name,cli.en_name,cli.unit_name,cli.is_special,
	cli.data_type,clir.reference_sex,clir.stomach_status,clir.is_pregnancy,clir.reference_max,clir.reference_min,cli.status
	FROM laboratory_patient_record_item lpri 
	left join clinic_laboratory_item cli on lpri.clinic_laboratory_item_id = cli.id
	left join clinic_laboratory_item_reference clir on clir.clinic_laboratory_item_id = cli.id
	left join laboratory_patient_record lpr on lpr.id = lpri.laboratory_patient_record_id
	where lpr.laboratory_patient_id = $1`

	rows, _ := model.DB.Queryx(selectSQL, laboratoryPatientID)
	results := FormatSQLRowsToMapArray(rows)
	laboratoryResult := FormatLaboratoryItem(results)

	ctx.JSON(iris.Map{"code": "200", "associationsData": laboratoryItems, "resultsData": laboratoryResult})
}

// LaboratoryTriageWaiting 获取待检验的分诊记录
func LaboratoryTriageWaiting(ctx iris.Context) {
	laboratoryTriageList(ctx, "10")
}

// LaboratoryTriageChecked 获取已检验的分诊记录
func LaboratoryTriageChecked(ctx iris.Context) {
	laboratoryTriageList(ctx, "30")
}

// LaboratoryTriageChecking 获取检验中的分诊记录
func LaboratoryTriageChecking(ctx iris.Context) {
	laboratoryTriageList(ctx, "20")
}

// 获取各状态的分诊记录
func laboratoryTriageList(ctx iris.Context, status string) {
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
	left join medical_record mr on mr.clinic_triage_patient_id = ctp.id
	left join clinic_triage_patient_operation register on ctp.id = register.clinic_triage_patient_id and register.type = 10
	left join personnel triage_personnel on triage_personnel.id = register.personnel_id 
	left join (select clinic_triage_patient_id,count(*) as total_count,max(created_time) as order_time
		from laboratory_patient where order_status = $1 group by(clinic_triage_patient_id)) up on up.clinic_triage_patient_id = ctp.id 
	left join (select clinic_triage_patient_id,count(*) as mz_count 
	from mz_paid_orders where charge_project_type_id = 3 group by(clinic_triage_patient_id)) mzup on mzup.clinic_triage_patient_id = ctp.id
	where up.total_count > 0 AND mzup.mz_count > 0 AND cp.clinic_id=$2 AND ctp.updated_time BETWEEN $3 and $4 AND (p.name ~$5 OR p.cert_no ~$5 OR p.phone ~$5) `

	countsql := `select count(*) as total` + sql
	querysql := `select 
	((select count(*) 
	from laboratory_patient where order_status = '10' and clinic_triage_patient_id = ctp.id )) as waiting_total_count,
	((select count(*) 
	from laboratory_patient where order_status = '20' and clinic_triage_patient_id = ctp.id )) as checking_total_count,
	((select count(*) 
	from laboratory_patient where order_status = '30' and clinic_triage_patient_id = ctp.id )) as checked_total_count,
	ctp.id as clinic_triage_patient_id,
	ctp.clinic_patient_id as clinic_patient_id,
	ctp.updated_time,
	ctp.created_time as register_time,
	up.order_time,
	triage_personnel.name as register_personnel_name,
	ctp.status,
	ctp.visit_date,
	ctp.register_type,
	p.name as patient_name,
	p.birthday,
	p.sex,
	p.phone,
	doc.name as doctor_name,
	d.name as department_name,
	mr.diagnosis,
	mr.allergic_history ` + sql + `order by up.order_time desc offset $6 limit $7`

	total := model.DB.QueryRowx(countsql, status, clinicID, startDate, endDate, keyword)

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.Queryx(querysql, status, clinicID, startDate, endDate, keyword, offset, limit)
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

// LaboratoryTriageRecordCreate 创建检验记录
func LaboratoryTriageRecordCreate(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	laboratoryPatientID := ctx.PostValue("laboratory_patient_id")
	operationID := ctx.PostValue("operation_id")
	remark := ctx.PostValue("remark")
	items := ctx.PostValue("items")

	if clinicTriagePatientID == "" || laboratoryPatientID == "" || operationID == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from laboratory_patient where id=$1 and clinic_triage_patient_id=$2 limit 1", laboratoryPatientID, clinicTriagePatientID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "创建失败"})
		return
	}
	laboratoryPatient := FormatSQLRowToMap(row)

	_, ok := laboratoryPatient["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "就诊人与检验数据不符合"})
		return
	}

	var results []map[string]interface{}
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	tx, txErr := model.DB.Beginx()
	if txErr != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": txErr.Error()})
		return
	}

	rrow := model.DB.QueryRowx("select id from laboratory_patient_record where laboratory_patient_id=$1 and clinic_triage_patient_id=$2 limit 1", laboratoryPatientID, clinicTriagePatientID)
	if rrow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "创建失败"})
		return
	}
	laboratoryPatientRecord := FormatSQLRowToMap(rrow)

	var recordID interface{}
	_, rok := laboratoryPatientRecord["id"]
	if !rok {
		err1 := tx.QueryRow(`INSERT INTO laboratory_patient_record 
			(clinic_triage_patient_id, laboratory_patient_id, operation_id, remark) 
			VALUES ($1,$2,$3,$4) RETURNING id`, clinicTriagePatientID, laboratoryPatientID, operationID, remark).Scan(&recordID)
		if err1 != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
			return
		}

		_, err3 := tx.Exec(`UPDATE laboratory_patient set order_status = '30', updated_time=LOCALTIMESTAMP where id = $1`, laboratoryPatientID)
		if err3 != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
			return
		}

	} else {
		recordID = laboratoryPatientRecord["id"]
		_, err2 := tx.Exec(`UPDATE laboratory_patient_record set
			operation_id=$2, remark=$3, updated_time=LOCALTIMESTAMP where id=$1`, recordID, operationID, remark)
		if err2 != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
			return
		}

		_, errd := tx.Exec(`delete from laboratory_patient_record_item where laboratory_patient_record_id=$1`, recordID)
		if errd != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": errd.Error()})
			return
		}
	}

	for _, item := range results {
		clinicLaboratoryItemID := item["clinic_laboratory_item_id"]
		resultInspection := item["result_inspection"]
		referenceMax := item["reference_max"]
		referenceMin := item["reference_min"]
		referenceValue := item["reference_value"]
		dataType := item["data_type"]
		isNormal := item["is_normal"]

		row := model.DB.QueryRowx("select id from clinic_laboratory_item where id=$1 limit 1", clinicLaboratoryItemID)
		if row == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "创建失败"})
			return
		}
		clinicLaboratoryItem := FormatSQLRowToMap(row)

		_, ok := clinicLaboratoryItem["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "所选检验项目不存在"})
			return
		}

		_, err := tx.Exec(`INSERT INTO laboratory_patient_record_item 
			(laboratory_patient_record_id, clinic_laboratory_item_id,
				result_inspection,reference_max,reference_min,
				reference_value,data_type,is_normal) 
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`, recordID, clinicLaboratoryItemID, resultInspection, referenceMax, referenceMin, referenceValue, dataType, isNormal)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	}

	erre := tx.Commit()
	if erre != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": erre.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "msg": "操作成功"})

}

// LaboratoryTriageRecordList 查询检验记录
func LaboratoryTriageRecordList(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")

	querysql := `select ctp.visit_date,lpc.id as laboratory_patient_record_id,
	doc.name as doctor_name,
	op.name as opration_name,
	lpc.created_time,lpc.remark
	cl.name as clinic_laboratory_name from laboratory_patient_record lpc 
	left join clinic_triage_patient ctp on ctp.id = lpc.clinic_triage_patient_id 
	left join personnel doc on doc.id = ctp.doctor_id 
	left join personnel op on op.id = lpc.operation_id 
	left join laboratory_patient lp on lp.id = lpc.laboratory_patient_id
	left join clinic_laboratory cl on cl.id = lp.clinic_laboratory_id
	where lpc.clinic_triage_patient_id = $1`

	rows, _ := model.DB.Queryx(querysql, clinicTriagePatientID)

	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results})

}

// LaboratoryTriageRecordDetail 查询检验记录详情
func LaboratoryTriageRecordDetail(ctx iris.Context) {
	laboratoryPatientRecordID := ctx.PostValue("laboratory_patient_record_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}

	SQL := ` FROM laboratory_patient_record_item lpri 
	left join clinic_laboratory_item cli on lpri.clinic_laboratory_item_id = cli.id
	left join clinic_laboratory_item_reference clir on clir.clinic_laboratory_item_id = cli.id
	where lpri.laboratory_patient_record_id = $1`

	countsql := "select count(lpci.*) as total " + SQL
	total := model.DB.QueryRowx(countsql, laboratoryPatientRecordID)

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	querysql := `select cls.clinic_laboratory_item_id,cli.name,cli.en_name,cli.unit_name,cli.is_special,cls.default_result,
	cli.data_type,clir.reference_sex,clir.stomach_status,clir.is_pregnancy,clir.reference_max,clir.reference_min,cli.status
	` + SQL + " offset $2 limit $3"

	rows, _ := model.DB.Queryx(querysql, laboratoryPatientRecordID, offset, limit)
	results := FormatSQLRowsToMapArray(rows)
	laboratoryResult := FormatLaboratoryItem(results)

	ctx.JSON(iris.Map{"code": "200", "data": laboratoryResult, "page_info": pageInfo})

}

//LaboratoryTriageUpdate 更新检验状态
func LaboratoryTriageUpdate(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	status := ctx.PostValue("order_status")
	items := ctx.PostValue("items")

	if clinicTriagePatientID == "" || status == "" || items == "" {
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

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errb})
		return
	}

	for _, v := range results {
		laboratoryPatientID := v["laboratory_patient_id"]

		_, err := tx.Exec(`UPDATE laboratory_patient set order_status=$1,updated_time=LOCALTIMESTAMP where id=$2 and clinic_triage_patient_id=$3`, status, laboratoryPatientID, clinicTriagePatientID)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	}

	errc := tx.Commit()
	if errc != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": errc.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "msg": "操作成功"})
}

// LaboratoryTriagePatientRecordList 患者历史检验记录
func LaboratoryTriagePatientRecordList(ctx iris.Context) {
	patientID := ctx.PostValue("patient_id")
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}
	if patientID == "" && clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	if patientID == "" {
		row := model.DB.QueryRowx(`select cp.patient_id from clinic_triage_patient ctp 
			left join clinic_patient cp on ctp.clinic_patient_id = cp.id where ctp.id = $1`, clinicTriagePatientID)
		if row == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "查询就诊人错误"})
			return
		}
		patient := FormatSQLRowToMap(row)
		pID, ok := patient["patient_id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "查询就诊人错误"})
			return
		}
		patientID = strconv.Itoa(int(pID.(int64)))
	}
	countSQL := `select count (*) as total from (select 
		lp.clinic_triage_patient_id, 
		ctp.clinic_patient_id, 
		cp.patient_id 
		from laboratory_patient lp
		left join clinic_triage_patient ctp on lp.clinic_triage_patient_id = ctp.id
		left join clinic_patient cp on ctp.clinic_patient_id = cp.id
		where cp.patient_id = $1 and lp.order_status = '30'
		group by (lp.clinic_triage_patient_id, ctp.clinic_patient_id, cp.patient_id)) aaa`
	total := model.DB.QueryRowx(countSQL, patientID)
	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	querySQL := `select 
	string_agg (ce.name, '，') as clinic_laboratory_name, 
	max (lp.created_time) as finish_time,
	lp.clinic_triage_patient_id, 
	ctp.clinic_patient_id, 
	cp.patient_id,
	c.name as clinic_name,
	d.name as department_name,
	p.name as doctor_name  
	from laboratory_patient lp
	left join clinic_laboratory ce on ce.id = lp.clinic_laboratory_id
	left join clinic_triage_patient ctp on lp.clinic_triage_patient_id = ctp.id
	left join clinic_patient cp on ctp.clinic_patient_id = cp.id
	left join clinic c on c.id = cp.clinic_id
	left join department d on ctp.department_id = d.id
	left join personnel p on ctp.doctor_id = p.id
	where cp.patient_id = $1 and lp.order_status = '30'
	group by (lp.clinic_triage_patient_id, ctp.clinic_patient_id, cp.patient_id, c.name, d.name, p.name)
	order by finish_time DESC
	offset $2 limit $3`

	rows, err := model.DB.Queryx(querySQL, patientID, offset, limit)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}
