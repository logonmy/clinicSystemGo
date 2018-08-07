package controller

import (
	"clinicSystemGo/model"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/kataras/iris"
)

//ClinicList 获取诊所列表
func ClinicList(ctx iris.Context) {
	keyword := ctx.PostValue("keyword")
	startDateStr := ctx.PostValue("start_date")
	endDateStr := ctx.PostValue("end_date")
	status := ctx.PostValue("status")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if offset == "" {
		offset = "0"
	}

	if limit == "" {
		limit = "10"
	}

	queryMap := map[string]interface{}{
		"keyword": ToNullString(keyword),
		"status":  ToNullBool(status),
		"offset":  ToNullInt64(offset),
		"limit":   ToNullInt64(limit),
	}

	sql := `FROM clinic c
	left join personnel p on p.clinic_id = c.id and p.is_clinic_admin = true
	where c.deleted_time is null `

	if startDateStr != "" {
		startDate, errs := time.Parse("2006-01-02", startDateStr)
		if errs != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "start_date 必须为 YYYY-MM-DD 的 有效日期格式"})
			return
		}
		queryMap["start_date"] = startDate
		sql += " and c.created_time > :start_date"
	}
	if endDateStr != "" {
		endDate, erre := time.Parse("2006-01-02", endDateStr)
		if erre != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "end_date 必须为 YYYY-MM-DD 的 有效日期格式"})
			return
		}
		endDate = endDate.AddDate(0, 0, 1)
		queryMap["end_date"] = endDate
		sql += " and c.created_time < :end_date"
	}

	if keyword != "" {
		sql = sql + ` AND (c.code ~*:keyword or c.name ~*:keyword) `
	}

	if status != "" {
		sql = sql + " AND c.status = :status"
	}

	var results []map[string]interface{}
	rows, err1 := model.DB.NamedQuery("SELECT c.*,p.phone,p.username,p.clinic_id "+sql+" order by c.created_time DESC offset :offset limit :limit", queryMap)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}
	total, err2 := model.DB.NamedQuery("SELECT COUNT (*) as total "+sql, queryMap)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}
	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

//ClinicAdd 添加诊所
func ClinicAdd(ctx iris.Context) {
	code := ctx.PostValue("code")
	name := ctx.PostValue("name")
	responsiblePerson := ctx.PostValue("responsible_person")
	area := ctx.PostValue("area")
	province := ctx.PostValue("province")
	city := ctx.PostValue("city")
	district := ctx.PostValue("district")
	status := ctx.PostValue("status")

	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	phone := ctx.PostValue("phone")

	if code == "" || name == "" || responsiblePerson == "" || area == "" || status == "" || username == "" || password == "" || phone == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	passwordMd5 := hex.EncodeToString(md5Ctx.Sum(nil))

	tx, err := model.DB.Beginx()

	if err != nil {
		fmt.Println("err ===", err.Error())
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	var clinicID int
	err = tx.QueryRow(`INSERT INTO clinic(
			code, name, responsible_person, province, city, district, area, status)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`, code, name, responsiblePerson, province, city, district, area, status).Scan(&clinicID)
	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	_, er := tx.Exec(`insert into storehouse (name,clinic_id) VALUES ($1,$2)`, name+"-默认", clinicID)

	if er != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": er.Error()})
		return
	}

	if err != nil {
		fmt.Println("err ===", err.Error())
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	amap := map[string]interface{}{
		"code":            string(clinicID) + "001",
		"name":            "超级管理员",
		"username":        username,
		"password":        passwordMd5,
		"clinic_id":       clinicID,
		"phone":           phone,
		"is_clinic_admin": true,
	}

	_, err = tx.NamedExec(`INSERT INTO personnel(
		code, name, username, password, clinic_id, phone, is_clinic_admin)
		VALUES (:code, :name, :username, :password, :clinic_id, :phone, :is_clinic_admin)`, amap)
	if err != nil {
		fmt.Println("err ===", err.Error())
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//ClinicUpdate 更新诊所信息
func ClinicUpdate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	name := ctx.PostValue("name")
	responsiblePerson := ctx.PostValue("responsible_person")
	area := ctx.PostValue("area")
	province := ctx.PostValue("province")
	city := ctx.PostValue("city")
	district := ctx.PostValue("district")
	status := ctx.PostValue("status")

	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	phone := ctx.PostValue("phone")

	if clinicID == "" || name == "" || responsiblePerson == "" || area == "" || status == "" || username == "" || password == "" || phone == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	passwordMd5 := hex.EncodeToString(md5Ctx.Sum(nil))

	cmap := map[string]interface{}{
		"clinicId":           ToNullInt64(clinicID),
		"name":               ToNullString(name),
		"province":           ToNullString(province),
		"city":               ToNullString(city),
		"district":           ToNullString(district),
		"responsible_person": ToNullString(responsiblePerson),
		"area":               ToNullString(area),
		"status":             ToNullBool(status),
		"updated_time":       time.Now(),
	}

	amap := map[string]interface{}{
		"username":     ToNullString(username),
		"password":     ToNullString(passwordMd5),
		"clinicId":     ToNullInt64(clinicID),
		"phone":        ToNullString(phone),
		"updated_time": time.Now(),
	}

	tx, err := model.DB.Beginx()

	if err != nil {
		fmt.Println("err ===", err.Error())
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	_, err = tx.NamedExec(`UPDATE clinic SET name=:name, responsible_person=:responsible_person, province=:province,city=:city,district=:district,area=:area, status=:status, updated_time=:updated_time WHERE id=:clinicId`, cmap)

	if err != nil {
		fmt.Println("err ===", err.Error())
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	_, err = tx.NamedExec(`UPDATE personnel SET username=:username, password=:password, phone=:phone, updated_time=:updated_time WHERE clinic_id=:clinicId and is_clinic_admin=true`, amap)
	if err != nil {
		fmt.Println("err ===", err.Error())
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// ClinicUpdateStatus 打开或关闭诊所
func ClinicUpdateStatus(ctx iris.Context) {
	status := ctx.PostValue("status")
	ID := ctx.PostValue("clinic_id")
	if status == "" || ID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	queryMap := map[string]interface{}{
		"status": status,
		"ID":     ID,
	}

	_, err := model.DB.NamedExec(`UPDATE clinic SET status = :status WHERE id = :ID`, queryMap)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "msg": "执行成功"})

}

//ClinicGetByID 获取诊所详情
func ClinicGetByID(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	row := model.DB.QueryRowx("select id as clinic_id,code,name,phone,area,responsible_person,status,created_time from clinic where id=$1", clinicID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)
	_, ok := clinic["clinic_id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊所不存在"})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": clinic})
}

// GetClinicCode 获取最新的诊所编码
func GetClinicCode(ctx iris.Context) {
	row := model.DB.QueryRowx("select code from clinic order by code DESC limit 1")
	rowMap := FormatSQLRowToMap(row)
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": rowMap})
}
