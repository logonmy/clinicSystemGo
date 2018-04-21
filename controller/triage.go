package controller

import (
	"clinicSystemGo/model"
	"fmt"

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
	err = tx.QueryRow("INSERT INTO clinic_triage_patient (department_id, clinic_patient_id, register_personnel_id) VALUES ($1, $2, $3) RETURNING id", departmentID, clinicPatientID, personnelID).Scan(&resultID)
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
	ctp.id,ctp.visit_date, ctp.treat_status, cp.clinic_id, c.name as clinic_name,
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
