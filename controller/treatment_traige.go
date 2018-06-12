package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/kataras/iris"
)

// TreatmentTriageList 获取治疗列表
func TreatmentTriageList(ctx iris.Context) {
	status := ctx.PostValue("order_status")
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")

	if clinicTriagePatientID == "" || status == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select ep.id as treatment_patient_id,ep.clinic_triage_patient_id,
	ce.name as clinic_treatment_name,ep.clinic_treatment_id,
	tpr.times,tpr.remark
	FROM treatment_patient ep 
	left join clinic_treatment ce on ce.id = ep.clinic_treatment_id
	left join mz_paid_orders mo on mo.clinic_triage_patient_id = ep.clinic_triage_patient_id and mo.charge_project_type_id=7 and ep.clinic_treatment_id=mo.charge_project_id
	left join treatment_patient_record tpr on tpr.treatment_patient_id = ep.id
	where mo.id is not NULL and ep.clinic_triage_patient_id = $1 and ep.order_status=$2`

	rows, _ := model.DB.Queryx(selectSQL, clinicTriagePatientID, status)
	results := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})
}

// TreatmentTriageWaiting 获取待治疗的分诊记录
func TreatmentTriageWaiting(ctx iris.Context) {
	treatmentTriageList(ctx, "10")
}

// TreatmentTriageChecking 获取治疗中的分诊记录
func TreatmentTriageChecking(ctx iris.Context) {
	treatmentTriageList(ctx, "20")
}

// TreatmentTriageChecked 获取已治疗的分诊记录
func TreatmentTriageChecked(ctx iris.Context) {
	treatmentTriageList(ctx, "30")
}

// 获取各状态的分诊记录
func treatmentTriageList(ctx iris.Context, status string) {
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
	left join (select clinic_triage_patient_id,count(*) as total_count 
		from treatment_patient where order_status = $1 group by(clinic_triage_patient_id)) up on up.clinic_triage_patient_id = ctp.id 
	left join (select clinic_triage_patient_id,count(*) as mz_count 
		from mz_paid_orders where charge_project_type_id = 7 group by(clinic_triage_patient_id)) mzup on mzup.clinic_triage_patient_id = ctp.id
	where up.total_count > 0 AND mzup.mz_count > 0 AND cp.clinic_id=$2 AND ctp.updated_time BETWEEN $3 and $4 AND (p.name ~$5 OR p.cert_no ~$5 OR p.phone ~$5) `

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
	d.name as department_name ` + sql + `offset $6 limit $7`

	total := model.DB.QueryRowx(countsql, status, clinicID, startDate, endDate, keyword)

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.Queryx(querysql, status, clinicID, startDate, endDate, keyword, offset, limit)
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

// TreatmentTriageRecordCreate 创建治疗记录
func TreatmentTriageRecordCreate(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	operationID := ctx.PostValue("operation_id")
	items := ctx.PostValue("items")

	if clinicTriagePatientID == "" || operationID == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
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

	for _, item := range results {
		treatmentPatientID := item["treatment_patient_id"]
		times := item["times"]
		remark := item["remark"]
		if times.(int) <= 0 {
			ctx.JSON(iris.Map{"code": "-1", "msg": "请填写次数"})
			return
		}
		row := model.DB.QueryRowx("select id from treatment_patient where id=$1 and clinic_triage_patient_id=$2 limit 1", treatmentPatientID, clinicTriagePatientID)
		if row == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "创建失败"})
			return
		}
		treatmentPatient := FormatSQLRowToMap(row)

		_, ok := treatmentPatient["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "就诊人与治疗数据不符合"})
			return
		}

		_, err := tx.Exec(`INSERT INTO treatment_patient_record 
			(clinic_triage_patient_id, treatment_patient_id, times, remark,operation_id) 
			VALUES ($1,$2,$3)`, clinicTriagePatientID, treatmentPatientID, times, remark, operationID)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}

		_, err1 := tx.Exec(`UPDATE treatment_patient set left_times=left_times-$2 ,updated_time=LOCALTIMESTAMP where id = $1`, treatmentPatientID, times)
		if err1 != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
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

// TreatmentTriageRecordList 查询治疗记录
func TreatmentTriageRecordList(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")

	querysql := `select ctp.visit_date,epc.id as treatment_patient_record_id,
	doc.name as doctor_name,
	op.name as opration_name,
	epc.created_time,
	epc.picture_treatment,
	epc.result_treatment,
	epc.conclusion_treatment,
	ce.name as clinic_treatment_name 
	from treatment_patient_record epc 
	left join clinic_triage_patient ctp on ctp.id = epc.clinic_triage_patient_id 
	left join personnel doc on doc.id = ctp.doctor_id 
	left join personnel op on op.id = epc.operation_id 
	left join treatment_patient ep on ep.id = epc.treatment_patient_id
	left join clinic_treatment ce on ce.id = ep.clinic_treatment_id
	where epc.clinic_triage_patient_id = $1`

	rows, _ := model.DB.Queryx(querysql, clinicTriagePatientID)

	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results})

}

//TreatmentTriageUpdate 更新治疗状态
func TreatmentTriageUpdate(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	status := ctx.PostValue("order_status")
	items := ctx.PostValue("items")

	if clinicTriagePatientID == "" || status == "" || items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select id,order_status,left_times FROM treatment_patient where clinic_triage_patient_id = $1`

	rows, _ := model.DB.Queryx(selectSQL, clinicTriagePatientID)
	treatmentPatients := FormatSQLRowsToMapArray(rows)

	if status == "30" {
		for _, v := range treatmentPatients {
			leftTimes := v["left_times"]
			if leftTimes.(int64) > 0 {
				ctx.JSON(iris.Map{"code": "-1", "msg": "治疗未完成"})
				return
			}
		}
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
		treatmentPatientID := v["treatment_patient_id"]

		_, err := tx.Exec(`UPDATE treatment_patient set order_status=$1,updated_time=LOCALTIMESTAMP where id=$2 and clinic_triage_patient_id=$3`, status, treatmentPatientID, clinicTriagePatientID)
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