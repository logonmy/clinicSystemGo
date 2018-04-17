package controller

import (
	"clinicSystemGo/model"
	"fmt"

	"github.com/kataras/iris"
)

//PatientAdd 就诊人
func PatientAdd(ctx iris.Context) {
	certNo := ctx.PostValue("certNo")
	name := ctx.PostValue("name")
	birthday := ctx.PostValue("birthday")
	sex := ctx.PostValue("sex")
	phone := ctx.PostValue("phone")
	address := ctx.PostValue("address")
	profession := ctx.PostValue("profession")
	remark := ctx.PostValue("remark")
	patientChannelID := ctx.PostValue("patientChannelId")
	clinicCode := ctx.PostValue("clinicCode")
	personnelID := ctx.PostValue("personnelId")
	if certNo == "" || name == "" || birthday == "" || sex == "" || phone == "" || patientChannelID == "" || clinicCode == "" || personnelID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	tx, err := model.DB.Begin()
	if err != nil {
		fmt.Println("err1 ===", err)
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	_, err = tx.Query(`INSERT INTO patient (
		cert_no, name, birthday, sex, phone, address, profession, remark, patient_channel_id) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`, certNo, name, birthday, sex, phone, address, profession, remark, patientChannelID)
	if err != nil {
		tx.Rollback()
		fmt.Println("err2 ===", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	var resultID int
	fmt.Println(" =========", certNo, clinicCode, personnelID)
	err = tx.QueryRow("INSERT into clinic_patient (patient_cert_no, clinic_code, personnel_id) values ($1, $2, $3) RETURNING id", certNo, clinicCode, personnelID).Scan(&resultID)
	if err != nil {
		tx.Rollback()
		fmt.Println("err3 ===", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("err4 ===", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": resultID})
	return
}
