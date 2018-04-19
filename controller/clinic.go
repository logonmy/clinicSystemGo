package controller

import (
	"clinicSystemGo/model"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/kataras/iris"
)

// GetClinicByCode 更加code查询诊所
func GetClinicByCode(ctx iris.Context) {
	code := ctx.PostValue("code")
	if code == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("SELECT * FROM clinic WHERE code=$1", code)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询失败"})
		return
	}

	clinic := FormatSQLRowToMap(row)
	ctx.JSON(iris.Map{"code": "200", "data": clinic})
}

//ClinicList 获取诊所列表
func ClinicList(ctx iris.Context) {
	keyword := ctx.PostValue("keyword")
	startDate := ctx.PostValue("startDate")
	endDate := ctx.PostValue("endDate")
	status := ctx.PostValue("status")

	if keyword == "" {
		keyword = "%"
	}

	sql := "SELECT * FROM clinic, where (code LIKE '%" + keyword + "%' or name LIKE '%" + keyword + "%')"

	if status != "" {
		sql = sql + " AND status = " + status
	}
	if startDate != "" && endDate != "" {
		sql = sql + " AND created_time between '" + startDate + "' and '" + endDate + "'"
	}

	// sql = "select * from (" + sql + ") a left join (select clinic_id,username,phone from personnel where is_clinic_admin=true) b on a.code = b.clinic_id"

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(sql)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})
}

//ClinicAdd 添加诊所
func ClinicAdd(ctx iris.Context) {
	code := ctx.PostValue("code")
	name := ctx.PostValue("name")
	responsiblePerson := ctx.PostValue("responsible_person")
	area := ctx.PostValue("area")
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
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	var clinicID int
	err = tx.QueryRow(`INSERT INTO clinic(
		code, name, responsible_person, area, status)
		VALUES ($1, $2, $3, $4, $5) RETURNING id`, code, name, responsiblePerson, area, status).Scan(&clinicID)

	if err != nil {
		fmt.Println("err1 ===", err.Error())
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	amap := map[string]interface{}{
		"code":            "10000",
		"name":            "超级管理员",
		"username":        username,
		"password":        passwordMd5,
		"clinic_id":       clinicID,
		"phone":           phone,
		"is_clinic_admin": true,
	}

	_, err = tx.NamedExec(`INSERT INTO personnel(
		code, name, username, password, clinic_id, phone,is_clinic_admin)
		VALUES (:code, :name, :username, :password, :clinic_id, :phone, :is_clinic_admin)`, amap)
	if err != nil {
		fmt.Println("err3 ===", err)
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
	id := ctx.PostValue("id")
	code := ctx.PostValue("code")
	name := ctx.PostValue("name")
	responsiblePerson := ctx.PostValue("responsible_person")
	area := ctx.PostValue("area")
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

	cmap := map[string]interface{}{
		"id":                 id,
		"code":               code,
		"name":               name,
		"responsible_person": responsiblePerson,
		"area":               area,
		"status":             status,
		"updated_time":       time.Now(),
	}

	amap := map[string]interface{}{
		"code":            "10000",
		"name":            "超级管理员",
		"username":        username,
		"password":        passwordMd5,
		"clinic_id":       id,
		"phone":           phone,
		"is_clinic_admin": true,
		"updated_time":    time.Now(),
	}

	tx, err := model.DB.Beginx()

	if err != nil {
		fmt.Println("err ===", err.Error())
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	_, err = tx.NamedExec(`UPDATE clinic SET code=:code, name=:name, responsible_person=:responsible_person, area=:area, status=:status, updated_time=:updated_time WHERE id=:id`, cmap)

	if err != nil {
		fmt.Println("err ===", err.Error())
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	_, err = tx.NamedExec(`UPDATE personnel SET username=:username, password=:password, phone=:phone, updated_time=:updated_time WHERE clinic_id=:clinic_id and is_clinic_admin=true`, amap)
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
