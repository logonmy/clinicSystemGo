package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"

	"github.com/kataras/iris"
)

// TriageRegister 就诊患者登记
func TriageRegister(ctx iris.Context) {
	certNo := ctx.PostValue("cert_no")
	name := ctx.PostValue("name")
	birthday := ctx.PostValue("birthday")
	sex := ctx.PostValue("sex")
	phone := ctx.PostValue("phone")
	address := ctx.PostValue("address")
	profession := ctx.PostValue("profession")
	remark := ctx.PostValue("remark")
	patientChannelID := ctx.PostValue("patient_channel_id")
	clinicID := ctx.PostValue("clinic_id")
	personnelID := ctx.PostValue("personnel_id")
	departmentID := ctx.PostValue("department_id")
	if certNo == "" || name == "" || birthday == "" || sex == "" || phone == "" || patientChannelID == "" || clinicID == "" || personnelID == "" || departmentID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	row := model.DB.QueryRowx("select * from patient where cert_no = $1", certNo)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "登记失败"})
	}
	tx, err := model.DB.Begin()
	patient := FormatSQLRowToMap(row)
	_, ok := patient["id"]
	patientID := patient["id"]
	if !ok {
		err = tx.QueryRow(`INSERT INTO patient (
		cert_no, name, birthday, sex, phone, address, profession, remark, patient_channel_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`, certNo, name, birthday, sex, phone, address, profession, remark, patientChannelID).Scan(&patientID)
		if err != nil {
			tx.Rollback()
			fmt.Println("err2 ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
	}

	row = model.DB.QueryRowx("select * from clinic_patient where patient_id= $1 and clinic_id = $2", patientID, clinicID)
	if row == nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": "登记失败"})
	}
	clinicPatient := FormatSQLRowToMap(row)
	fmt.Println("clinic_triage_patient ======", clinicPatient)
	_, ok = clinicPatient["id"]
	var clinicPatientID interface{}
	if !ok {
		err = tx.QueryRow("INSERT INTO clinic_patient (patient_id, clinic_id, personnel_id) VALUES ($1, $2, $3) RETURNING id", patientID, clinicID, personnelID).Scan(&clinicPatientID)
		if err != nil {
			fmt.Println("clinic_patient ======", err)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
	} else {
		clinicPatientID = clinicPatient["id"]
	}

	var resultID int
	err = tx.QueryRow("INSERT INTO clinic_triage_patient (department_id, clinic_patient_id, register_personnel_id,register_type) VALUES ($1, $2, $3,1) RETURNING id", departmentID, clinicPatientID, personnelID).Scan(&resultID)
	if err != nil {
		fmt.Println("clinic_triage_patient ======", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": nil})
}

// TriagePatientList 当日登记就诊人列表
func TriagePatientList(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	keyword := ctx.PostValue("keyword")
	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	rowSQL := `select 
	ctp.id,ctp.clinic_patient_id,ctp.visit_date, ctp.treat_status, cp.clinic_id, c.name as clinic_name,
	p.id as patient_id, p.name as patient_name, p.birthday, p.sex, p.cert_no, p.phone,
	ctp.created_time as register_time,
	ctp.department_id, d.name as department_name,
	ctp.register_personnel_id, rp.name as register_personnel_name,
	ctp.triage_personnel_id, tp.name as triage_personnel_name
	from clinic_triage_patient ctp left join clinic_patient cp  on ctp.clinic_patient_id = cp.id
	left join clinic c on c.id = cp.clinic_id
	left join department d on ctp.department_id = d.id
	left join personnel rp on ctp.register_personnel_id = rp.id
	left join personnel tp on ctp.triage_personnel_id = tp.id
	left join patient p on cp.patient_id = p.id 
	where cp.clinic_id = $1 and ctp.visit_date = CURRENT_DATE 
	and (p.cert_no like '%' || $2 || '%' or p.name like '%' || $2 || '%' or p.phone like '%' || $2 || '%') `
	rows, err1 := model.DB.Queryx(rowSQL, clinicID, keyword)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result})
}

//PersonnelChoose 线下分诊选择医生
func PersonnelChoose(ctx iris.Context) {
	clinicPatientID := ctx.PostValue("clinic_patient_id")
	deparmentID := ctx.PostValue("department_id")
	clinicTriagePatientID := ctx.PostValue("id")
	doctorID := ctx.PostValue("doctor_id")
	triagePersonnelID := ctx.PostValue("personnel_id")
	amPm := ctx.PostValue("am_pm")
	visitTypeCode := ctx.PostValue("visit_type_code")

	if clinicPatientID == "" || deparmentID == "" || clinicTriagePatientID == "" || doctorID == "" || triagePersonnelID == "" || amPm == "" || visitTypeCode == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	tx, err := model.DB.Begin()
	var resultID int
	err = tx.QueryRow("update clinic_triage_patient set doctor_id=$1,triage_personnel_id=$2,treat_status=true,triage_time=LOCALTIMESTAMP where id=$3 RETURNING id", doctorID, triagePersonnelID, clinicTriagePatientID).Scan(&resultID)
	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	var registrationID int
	err = tx.QueryRow(`INSERT INTO registration (
		clinic_patient_id, department_id, clinic_triage_patient_id,personnel_id,visit_date,am_pm,visit_type_code,operation_id) 
		VALUES ($1, $2, $3, $4, CURRENT_DATE, $5, $6, $7) RETURNING id`, clinicPatientID, deparmentID, clinicTriagePatientID, doctorID, amPm, visitTypeCode, triagePersonnelID).Scan(&registrationID)
	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": registrationID})
}

//TriagePersonnelList 分诊医生列表
func TriagePersonnelList(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	deparmentID := ctx.PostValue("department_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	keyword := ctx.PostValue("keyword")
	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
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

	countSQL := `from doctor_visit_schedule dvs 
	left join department d on dvs.department_id = d.id 
	left join personnel p on dvs.personnel_id = p.id 
	where p.clinic_id = $1 and (p.name like '%' || $2 || '%') and dvs.visit_date=current_date`

	selectSQL := `(select count(ctp.id) from clinic_triage_patient ctp 
		where treat_status=false and visit_date=current_date and doctor_id=dvs.personnel_id) as waitTotal,
	(select count(ctped.id)from clinic_triage_patient ctped where 
		treat_status=true and visit_date=current_date and doctor_id=dvs.personnel_id) as triagedTotal
	from doctor_visit_schedule dvs 
	left join department d on dvs.department_id = d.id 
	left join personnel p on dvs.personnel_id = p.id
	where p.clinic_id = $1 and (p.name like '%' || $2 || '%') and dvs.visit_date=current_date`

	if deparmentID != "" {
		countSQL += "and dvs.department_id=" + deparmentID
		selectSQL += "and dvs.department_id=" + deparmentID
	}

	total := model.DB.QueryRowx(`select count(distinct(dvs.personnel_id,dvs.department_id)) as total `+countSQL, clinicID, keyword)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rowSQL := `select distinct(p.name, d.name), p.name as doctor_name, d.name as department_name, ` + selectSQL + " offset $3 limit $4"

	rows, err1 := model.DB.Queryx(rowSQL, clinicID, keyword, offset, limit)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})
}
