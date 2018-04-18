package controller

import (
	"clinicSystemGo/model"
	"fmt"

	"github.com/kataras/iris"
)

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
	clinicCode := ctx.PostValue("clinic_code")
	personnelID := ctx.PostValue("personnel_id")
	if certNo == "" || name == "" || birthday == "" || sex == "" || phone == "" || patientChannelID == "" || clinicCode == "" || personnelID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	row := model.DB.QueryRowx("select * from patient where cert_no = $1", certNo)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "登记失败"})
	}
	tx, err := model.DB.Begin()
	patient := FormatSQLRowToMap(row)
	_, ok := patient["cert_no"]
	if !ok {
		var cardNo string
		err = tx.QueryRow(`INSERT INTO patient (
		cert_no, name, birthday, sex, phone, address, profession, remark, patient_channel_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING cert_no`, certNo, name, birthday, sex, phone, address, profession, remark, patientChannelID).Scan(&cardNo)
		if err != nil {
			tx.Rollback()
			fmt.Println("err2 ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
	}

	row = model.DB.QueryRowx("select * from clinic_patient where patient_cert_no= $1 and clinic_code = $2", certNo, clinicCode)
	if row == nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": "登记失败"})
	}
	clinicPatient := FormatSQLRowToMap(row)
	_, ok = clinicPatient["id"]
	if !ok {
		var clinicPatientID int
		err = tx.QueryRow("INSERT INTO clinic_patient (patient_cert_no, clinic_code, personnel_id) VALUES ($1, $2, $3) RETURNING id", certNo, clinicCode, personnelID).Scan(&clinicPatientID)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
	}

	var resultID int
	err = tx.QueryRow("INSERT INTO clinic_patient (patient_cert_no, clinic_code, personnel_id) VALUES ($1, $2, $3) RETURNING id", certNo, clinicCode, personnelID).Scan(&resultID)
	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

}
