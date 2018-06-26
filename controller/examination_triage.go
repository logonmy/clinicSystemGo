package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/kataras/iris"
)

// ExaminationTriageList 获取检查列表
func ExaminationTriageList(ctx iris.Context) {
	status := ctx.PostValue("order_status")
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")

	if clinicTriagePatientID == "" || status == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select ep.id as examination_patient_id,ep.clinic_triage_patient_id,
	ce.name as clinic_examination_name,ep.clinic_examination_id,mpr.id as examination_patient_record_id,
	mpr.picture_examination,mpr.result_examination,mpr.conclusion_examination
	FROM examination_patient ep 
	left join clinic_examination ce on ce.id = ep.clinic_examination_id
	left join mz_paid_orders mo on mo.clinic_triage_patient_id = ep.clinic_triage_patient_id and mo.charge_project_type_id=4 and ep.clinic_examination_id=mo.charge_project_id
	left join examination_patient_record mpr on mpr.examination_patient_id = ep.id
	where mo.id is not NULL and ep.clinic_triage_patient_id = $1 and ep.order_status=$2`

	rows, _ := model.DB.Queryx(selectSQL, clinicTriagePatientID, status)
	results := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})
}

// ExaminationTriageWaiting 获取待检查的分诊记录
func ExaminationTriageWaiting(ctx iris.Context) {
	examinationTriageList(ctx, "10")
}

// ExaminationTriageChecking 获取检查中的分诊记录
func ExaminationTriageChecking(ctx iris.Context) {
	examinationTriageList(ctx, "20")
}

// ExaminationTriageChecked 获取已检查的分诊记录
func ExaminationTriageChecked(ctx iris.Context) {
	examinationTriageList(ctx, "30")
}

// 获取各状态的分诊记录
func examinationTriageList(ctx iris.Context, status string) {
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
	left join clinic_triage_patient_operation register on ctp.id = register.clinic_triage_patient_id and register.type = 40 and times=1
	left join personnel triage_personnel on triage_personnel.id = register.personnel_id
	left join (select clinic_triage_patient_id,count(*) as total_count,max(created_time) as order_time
		from examination_patient where order_status = $1 group by(clinic_triage_patient_id)) up on up.clinic_triage_patient_id = ctp.id 
	left join (select clinic_triage_patient_id,count(*) as mz_count 
		from mz_paid_orders where charge_project_type_id = 4 group by(clinic_triage_patient_id)) mzup on mzup.clinic_triage_patient_id = ctp.id	
	where up.total_count > 0 AND mzup.mz_count > 0 AND cp.clinic_id=$2 AND ctp.updated_time BETWEEN $3 and $4 AND (p.name ~$5 OR p.cert_no ~$5 OR p.phone ~$5) `

	countsql := `select count(*) as total` + sql
	querysql := `select 
	((select count(*) 
	from examination_patient where order_status = '10' and clinic_triage_patient_id = ctp.id )) as waiting_total_count,
	((select count(*) 
	from examination_patient where order_status = '20' and clinic_triage_patient_id = ctp.id )) as checking_total_count,
	((select count(*) 
	from examination_patient where order_status = '30' and clinic_triage_patient_id = ctp.id )) as checked_total_count,
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
	d.name as department_name ` + sql + `order by up.order_time desc offset $6 limit $7`

	total := model.DB.QueryRowx(countsql, status, clinicID, startDate, endDate, keyword)

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.Queryx(querysql, status, clinicID, startDate, endDate, keyword, offset, limit)
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

// ExaminationTriageRecordCreate 创建检查记录
func ExaminationTriageRecordCreate(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	examinationPatientID := ctx.PostValue("examination_patient_id")
	operationID := ctx.PostValue("operation_id")
	pictureExamination := ctx.PostValue("picture_examination")
	resultExamination := ctx.PostValue("result_examination")
	conclusionExamination := ctx.PostValue("conclusion_examination")

	if clinicTriagePatientID == "" || examinationPatientID == "" || operationID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from examination_patient where id=$1 and clinic_triage_patient_id=$2 limit 1", examinationPatientID, clinicTriagePatientID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "创建失败"})
		return
	}
	examinationPatient := FormatSQLRowToMap(row)

	_, ok := examinationPatient["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "就诊人与检查数据不符合"})
		return
	}

	rrow := model.DB.QueryRowx("select id from examination_patient_record where clinic_triage_patient_id=$1 and examination_patient_id=$2 limit 1", clinicTriagePatientID, examinationPatientID)
	if rrow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "创建失败"})
		return
	}
	examinationPatientRecord := FormatSQLRowToMap(rrow)

	examinationPatientRecordID, rok := examinationPatientRecord["id"]
	if !rok {
		_, err1 := model.DB.Exec(`INSERT INTO examination_patient_record 
			(clinic_triage_patient_id, examination_patient_id, operation_id, picture_examination,result_examination,conclusion_examination) 
			VALUES ($1,$2,$3,$4,$5,$6)`, clinicTriagePatientID, examinationPatientID, operationID, pictureExamination, resultExamination, conclusionExamination)
		if err1 != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
			return
		}

		_, err2 := model.DB.Exec(`UPDATE examination_patient set order_status = '30', updated_time=LOCALTIMESTAMP where id = $1`, examinationPatientID)
		if err2 != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
			return
		}

	} else {
		_, err3 := model.DB.Exec(`update examination_patient_record set
			operation_id=$2, picture_examination=$3,result_examination=$4,conclusion_examination=$5, updated_time=LOCALTIMESTAMP where id=$1`, examinationPatientRecordID, operationID, ToNullString(pictureExamination), ToNullString(resultExamination), ToNullString(conclusionExamination))
		if err3 != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
			return
		}
	}

	ctx.JSON(iris.Map{"code": "200", "msg": "操作成功"})
}

// ExaminationTriageRecordList 查询检查记录
func ExaminationTriageRecordList(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")

	querysql := `select ctp.visit_date,epc.id as examination_patient_record_id,
	doc.name as doctor_name,
	op.name as opration_name,
	epc.created_time,
	epc.picture_examination,
	epc.result_examination,
	epc.conclusion_examination,
	ce.name as clinic_examination_name 
	from examination_patient_record epc 
	left join clinic_triage_patient ctp on ctp.id = epc.clinic_triage_patient_id 
	left join personnel doc on doc.id = ctp.doctor_id 
	left join personnel op on op.id = epc.operation_id 
	left join examination_patient ep on ep.id = epc.examination_patient_id
	left join clinic_examination ce on ce.id = ep.clinic_examination_id
	where epc.clinic_triage_patient_id = $1`

	rows, _ := model.DB.Queryx(querysql, clinicTriagePatientID)

	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results})

}

//ExaminationTriageUpdate 更新检查状态
func ExaminationTriageUpdate(ctx iris.Context) {
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
		examinationPatientID := v["examination_patient_id"]

		_, err := tx.Exec(`UPDATE examination_patient set order_status=$1,updated_time=LOCALTIMESTAMP where id=$2 and clinic_triage_patient_id=$3`, status, examinationPatientID, clinicTriagePatientID)
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

// ExaminationTriagePatientRecordList 患者历史检查记录
func ExaminationTriagePatientRecordList(ctx iris.Context) {
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
		ep.clinic_triage_patient_id, 
		ctp.clinic_patient_id, 
		cp.patient_id 
		from examination_patient ep
		left join clinic_triage_patient ctp on ep.clinic_triage_patient_id = ctp.id
		left join clinic_patient cp on ctp.clinic_patient_id = cp.id
		where cp.patient_id = $1 and ep.order_status = '30'
		group by (ep.clinic_triage_patient_id, ctp.clinic_patient_id, cp.patient_id)) aaa`
	total := model.DB.QueryRowx(countSQL, patientID)
	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	querySQL := `select 
	string_agg (ce.name, '，') as clinic_examination_name, 
	ep.clinic_triage_patient_id, 
	ctp.clinic_patient_id, 
	ctpo.created_time as finish_time,
	cp.patient_id,
	c.name as clinic_name 
	from examination_patient ep
	left join clinic_examination ce on ce.id = ep.clinic_examination_id
	left join clinic_triage_patient ctp on ep.clinic_triage_patient_id = ctp.id
	left join clinic_triage_patient_operation ctpo on ctp.id = ctpo.clinic_triage_patient_id and ctpo.type = 40 and ctpo.times = 1
	left join clinic_patient cp on ctp.clinic_patient_id = cp.id
	left join clinic c on c.id = cp.clinic_id
	where cp.patient_id = $1 and ep.order_status = '30'
	group by (ep.clinic_triage_patient_id, ctp.clinic_patient_id, cp.patient_id, ctpo.created_time, c.name)
	order by ctpo.created_time DESC
	offset $2 limit $3`

	rows, err := model.DB.Queryx(querySQL, patientID, offset, limit)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}
