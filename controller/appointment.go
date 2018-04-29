package controller

import (
	"clinicSystemGo/model"
	"fmt"

	"github.com/kataras/iris"
)

// AppointmentCreate 预约预约
func AppointmentCreate(ctx iris.Context) {
	qPatientID := ctx.PostValue("paient_id")
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
	doctorVisitScheduleID := ctx.PostValue("doctor_visit_schedule_id")
	personnelID := ctx.PostValue("personnel_id")
	if name == "" || birthday == "" || sex == "" || phone == "" || patientChannelID == "" || clinicID == "" || personnelID == "" || doctorVisitScheduleID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	row := model.DB.QueryRowx("select * from doctor_visit_schedule where id=$1 and visit_date > CURRENT_DATE", doctorVisitScheduleID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "预约失败"})
		return
	}
	schedule := FormatSQLRowToMap(row)
	_, ok := schedule["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "号源不存在"})
		return
	}

	// 查询就诊人 begin
	if qPatientID != "" {
		row = model.DB.QueryRowx("select * from patient where id = $1", qPatientID)
	} else if certNo != "" {
		row = model.DB.QueryRowx("select * from patient where cert_no = $1", certNo)
	} else {
		row = model.DB.QueryRowx("select * from patient where name = $1 and birthday = $2 and phone = $3", name, birthday, phone)
	}

	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "登记失败"})
		return
	}
	tx, err := model.DB.Begin()
	patient := FormatSQLRowToMap(row)
	_, ok = patient["id"]
	patientID := patient["id"]
	if !ok {
		insertKeys := `name, birthday, sex, phone, address, profession, remark, patient_channel_id`
		insertValues := `$1, $2, $3, $4, $5, $6, $7, $8`
		if certNo != "" {
			insertKeys += ", cert_no"
			insertValues += ", " + certNo
		}
		err = tx.QueryRow(`INSERT INTO patient (insertKeys) 
		VALUES (insertValues) RETURNING id`, name, birthday, sex, phone, address, profession, remark, patientChannelID).Scan(&patientID)
		if err != nil {
			tx.Rollback()
			fmt.Println("err2 ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
	} else {
		updateSQL := `update patient set name= $1,birthday=$2,sex=$3, phone=$4, address=$5,profession = $6,remark= $7 ,patient_channel_id = $8  where id = $9`
		if certNo != "" {
			updateSQL = `update patient set cert_no = ` + certNo + `, name= $1,birthday=$2,sex=$3, phone=$4, address=$5,profession = $6,remark= $7 ,patient_channel_id = $8  where id = $9`
		}
		_, err = tx.Exec(updateSQL, name, birthday, sex, phone, address, profession, remark, patientChannelID, patientID)
		if err != nil {
			tx.Rollback()
			fmt.Println("err2 ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
	}

	// 查询就诊人 end

	row = model.DB.QueryRowx("select * from clinic_patient where patient_id= $1 and clinic_id = $2", patientID, clinicID)
	if row == nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": "预约失败"})
		return
	}
	clinicPatient := FormatSQLRowToMap(row)
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

	departmentID := schedule["department_id"]
	visitDate := schedule["visit_date"]
	doctorID := schedule["personnel_id"]

	var resultID int
	err = tx.QueryRow("INSERT INTO clinic_triage_patient (department_id, clinic_patient_id, register_personnel_id, register_type, visit_date, doctor_id) VALUES ($1, $2, $3,1,$4,$5) RETURNING id", departmentID, clinicPatientID, personnelID, visitDate, doctorID).Scan(&resultID)
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
