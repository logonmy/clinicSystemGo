package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"

	"github.com/kataras/iris"
)

//PatientAdd 新增就诊人
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

	var resultID int
	err = tx.QueryRow("INSERT INTO clinic_patient (patient_cert_no, clinic_code, personnel_id) VALUES ($1, $2, $3) RETURNING id", certNo, clinicCode, personnelID).Scan(&resultID)
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

//PatientList 就诊人列表
func PatientList(ctx iris.Context) {
	clinicCode := ctx.PostValue("clinic_code")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	keyword := ctx.PostValue("keyword")
	if clinicCode == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
		return
	}
	if keyword == "" {
		keyword = "%"
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
	total := model.DB.QueryRowx(`select count(p.cert_no) as total
		from patient p 
		left join clinic_patient cp on p.cert_no = cp.patient_cert_no 
		left join clinic c on c.code = cp.clinic_code
		where c.code = $1 and (p.name like '%' || $2 || '%' or p.cert_no like '%' || $2 || '%' or p.phone like '%' || $2 || '%')`, clinicCode, keyword)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err1 := model.DB.Queryx(`select p.*,pc.name as channel_name from patient p 
		left join clinic_patient cp on p.cert_no = cp.patient_cert_no 
		left join clinic c on c.code = cp.clinic_code
		left join patient_channel pc on p.patient_channel_id = pc.id
		where c.code = $1 and (p.name like '%' || $2 || '%' or p.cert_no like '%' || $2 || '%' or p.phone like '%' || $2 || '%') ORDER BY p.created_time DESC offset $3 limit $4;`, clinicCode, keyword, offset, limit)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})
}

//PatientGetByID 通过身份证获取就诊人
func PatientGetByID(ctx iris.Context) {
	certNo := ctx.PostValue("cert_no")
	if certNo != "" {
		row := model.DB.QueryRowx(`select p.* from patient p 
			left join clinic_patient cp on p.cert_no = cp.patient_cert_no 
			left join clinic c on c.code = cp.clinic_code where
			p.cert_no = $1;`, certNo)
		if row == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "查询结果不存在"})
			return
		}
		result := FormatSQLRowToMap(row)
		ctx.JSON(iris.Map{"code": "200", "data": result})
		return
	}
	ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
}
