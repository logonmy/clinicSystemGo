package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"
	"time"

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
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from patient where cert_no = $1 limit 1", certNo)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
		return
	}
	patient := FormatSQLRowToMap(row)
	_, ok := patient["id"]
	if ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "就诊人身份证已存在"})
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
	if startDate != "" && endDate == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请选择结束日期"})
		return
	}
	if startDate == "" && endDate != "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请选择开始日期"})
		return
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
	where c.id = $1 and (p.name ~ $2 or p.cert_no ~ $2 or p.phone ~ $2)`

	selectSQL := `select p.*,pc.name as channel_name from patient p 
	left join clinic_patient cp on p.id = cp.patient_id 
	left join clinic c on c.id = cp.clinic_id
	left join patient_channel pc on p.patient_channel_id = pc.id
	where c.id = $1 and (p.name ~ $2 or p.cert_no ~ $2 or p.phone ~ $2)`

	sql := " ORDER BY p.created_time DESC offset $3 limit $4;"

	if startDate != "" && endDate != "" {
		if startDate > endDate {
			ctx.JSON(iris.Map{"code": "-1", "msg": "开始日期必须大于结束日期"})
			return
		}
		countSQL = countSQL + " AND p.created_time between date'" + startDate + "' - integer '1' and date'" + endDate + "' + integer '1'"
		selectSQL = selectSQL + " AND p.created_time between date'" + startDate + "' - integer '1' and date'" + endDate + "' + integer '1'" + sql
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

// PatientGetByCertNo 通过身份号查就诊人
func PatientGetByCertNo(ctx iris.Context) {
	certNo := ctx.PostValue("cert_no")
	if certNo != "" {
		row := model.DB.QueryRowx(`select * from patient 
			cert_no = $1;`, certNo)
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
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	var patientID string
	err := model.DB.QueryRow(`UPDATE patient set  
		cert_no=$1,name=$2, birthday=$3, sex=$4, phone=$5, address=$6, profession=$7, remark=$8, patient_channel_id=$9, updated_time=$10
		where id=$10 RETURNING id`, certNo, name, birthday, sex, phone, address, profession, remark, patientChannelID, id, time.Now()).Scan(&patientID)
	if err != nil {
		// tx.Rollback()
		fmt.Println("err2 ===", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": patientID})
	return
}

// PatientsGetByKeyword 通过关键字搜索就诊人
func PatientsGetByKeyword(ctx iris.Context) {
	keyword := ctx.PostValue("keyword")
	if keyword == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
		return
	}
	rows, err := model.DB.Queryx(`select * from patient where name like '%' || $1 || '%' or cert_no like '%' || $1 || '%' or phone like '%' || $1 || '%' limit 20`, keyword)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询结果不存在"})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result})
	return
}

// MemberPateintList 会员，就诊人列表
func MemberPateintList(ctx iris.Context) {
	keyword := ctx.PostValue("keyword")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	startDateStr := ctx.PostValue("start_date")
	endDateStr := ctx.PostValue("end_date")

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

	patientSQL := `select p.id, p.name, p.phone, p.birthday, p.sex, p.created_time,max(ctpo.created_time) as visited_time
	from patient p
	left join clinic_patient cp on p.id = cp.patient_id
	left join clinic_triage_patient ctp on cp.id = ctp.clinic_patient_id and ctp.status > 30
	left join clinic_triage_patient_operation ctpo on ctp.id = ctpo.clinic_triage_patient_id and ctpo.type=40  and ctpo.times = 1 where p.deleted_time is null`

	groupBySQL := ` group by p.id, p.phone, p.birthday, p.sex, p.created_time`

	if keyword != "" {
		patientSQL += " and p.name ~:keyword or p.phone ~:keyword or p.cert_no ~:keyword"
	}

	var queryOptions = map[string]interface{}{
		"keyword": ToNullString(keyword),
		"offset":  ToNullInt64(offset),
		"limit":   ToNullInt64(limit),
	}

	if startDateStr != "" {
		startDate, errs := time.Parse("2006-01-02", startDateStr)
		if errs != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "start_date 必须为 YYYY-MM-DD 的 有效日期格式"})
			return
		}
		queryOptions["start_date"] = startDate
		patientSQL += " and ctpo.created_time > :start_date"
	}
	if endDateStr != "" {
		endDate, erre := time.Parse("2006-01-02", endDateStr)
		if erre != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "end_date 必须为 YYYY-MM-DD 的 有效日期格式"})
			return
		}
		endDate = endDate.AddDate(0, 0, 1)
		queryOptions["end_date"] = endDate
		patientSQL += " and ctpo.created_time < :end_date"
	}

	countSQL := "select count(*) as total from (" + patientSQL + groupBySQL + ") counttable"
	pageInfoRows, err := model.DB.NamedQuery(countSQL, queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	pageInfoArray := FormatSQLRowsToMapArray(pageInfoRows)
	pageInfo := pageInfoArray[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err := model.DB.NamedQuery(patientSQL+groupBySQL+" order by p.id desc offset :offset limit :limit", queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	array := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": array, "page_info": pageInfo})
}
