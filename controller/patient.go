package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"

	"github.com/kataras/iris"
)

//PatientAdd 新增就诊人
func PatientAdd(ctx iris.Context) {
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
	if certNo == "" || name == "" || birthday == "" || sex == "" || phone == "" || patientChannelID == "" || clinicID == "" || personnelID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	tx, err := model.DB.Begin()

	var patientID string
	err = tx.QueryRow(`INSERT INTO patient (
		cert_no, name, birthday, sex, phone, address, profession, remark, patient_channel_id) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`, certNo, name, birthday, sex, phone, address, profession, remark, patientChannelID).Scan(&patientID)
	if err != nil {
		tx.Rollback()
		fmt.Println("err2 ===", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	var resultID int
	err = tx.QueryRow("INSERT INTO clinic_patient (patient_id, clinic_id, personnel_id) VALUES ($1, $2, $3) RETURNING id", patientID, clinicID, personnelID).Scan(&resultID)
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
	clinicID := ctx.PostValue("clinic_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	keyword := ctx.PostValue("keyword")
	startDate := ctx.PostValue("startDate")
	endDate := ctx.PostValue("endDate")
	if clinicID == "" {
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
	countSQL := `select count(p.cert_no) as total
	from patient p 
	left join clinic_patient cp on p.id = cp.patient_id 
	left join clinic c on c.id = cp.clinic_id
	where c.id = $1 and (p.name like '%' || $2 || '%' or p.cert_no like '%' || $2 || '%' or p.phone like '%' || $2 || '%')`

	selectSQL := `select p.*,pc.name as channel_name from patient p 
	left join clinic_patient cp on p.id = cp.patient_id 
	left join clinic c on c.id = cp.clinic_id
	left join patient_channel pc on p.patient_channel_id = pc.id
	where c.id = $1 and (p.name like '%' || $2 || '%' or p.cert_no like '%' || $2 || '%' or p.phone like '%' || $2 || '%')`

	sql := " ORDER BY p.created_time DESC offset $3 limit $4;"

	if startDate != "" && endDate != "" {
		countSQL = countSQL + " AND p.created_time between '" + startDate + "' and '" + endDate + "'"
		selectSQL = selectSQL + " AND p.created_time between '" + startDate + "' and '" + endDate + "'" + sql
	} else if startDate != "" && endDate == "" {
		countSQL = countSQL + " AND p.created_time > '" + startDate + "'"
		selectSQL = selectSQL + " AND p.created_time > '" + startDate + "'" + sql
	} else if endDate != "" && startDate == "" {
		countSQL = countSQL + " AND p.created_time < '" + endDate + "'"
		selectSQL = selectSQL + " AND p.created_time < '" + endDate + "'" + sql
	} else {
		selectSQL = selectSQL + sql
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
	total := model.DB.QueryRowx(countSQL, clinicID, keyword)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err1 := model.DB.Queryx(selectSQL, clinicID, keyword, offset, limit)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})
}

//PatientGetByID 通过id就诊人
func PatientGetByID(ctx iris.Context) {
	id := ctx.PostValue("id")
	if id != "" {
		row := model.DB.QueryRowx(`select p.* from patient p 
			left join clinic_patient cp on p.id = cp.patient_id 
			left join clinic c on c.id = cp.clinic_id where
			p.id = $1;`, id)
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

//PatientUpdate 编辑就诊人
func PatientUpdate(ctx iris.Context) {
	id := ctx.PostValue("id")
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
	if id == "" || certNo == "" || name == "" || birthday == "" || sex == "" || phone == "" || patientChannelID == "" || clinicID == "" || personnelID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	var patientID string
	err := model.DB.QueryRow(`UPDATE patient set  
		cert_no=$1,name=$2, birthday=$3, sex=$4, phone=$5, address=$6, profession=$7, remark=$8, patient_channel_id=$9
		where id=$10 RETURNING id`, certNo, name, birthday, sex, phone, address, profession, remark, patientChannelID, id).Scan(&patientID)
	if err != nil {
		// tx.Rollback()
		fmt.Println("err2 ===", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	// var resultID int
	// err = tx.QueryRow("UPDATE clinic_patient set patient_id=$1, clinic_id=$2, personnel_id=$3 where patient_id=$4 RETURNING id", patientID, clinicID, personnelID, patientID).Scan(&resultID)
	// if err != nil {
	// 	tx.Rollback()
	// 	fmt.Println("err3 ===", err)
	// 	ctx.JSON(iris.Map{"code": "-1", "msg": err})
	// 	return
	// }
	// err = tx.Commit()
	// if err != nil {
	// 	fmt.Println("err4 ===", err)
	// 	ctx.JSON(iris.Map{"code": "-1", "msg": err})
	// 	return
	// }
	ctx.JSON(iris.Map{"code": "200", "data": patientID})
	return
}
