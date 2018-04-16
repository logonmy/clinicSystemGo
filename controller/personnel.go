package controller

import (
	"clinicSystemGo/model"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/kataras/iris"
)

/**
 * 创建医院管理员
 */
func PersonnelLogin(ctx iris.Context) {
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	if username != "" && password != "" {
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(password))
		passwordMd5 := hex.EncodeToString(md5Ctx.Sum(nil))
		row := model.DB.QueryRowx("select a.id, a.code, a.name, a.username, b.code as clinic_code, b.name as clinic_name from personnel a left join clinic b on a.clinic_code = b.code where  username = $1 and password = $2", username, passwordMd5)
		if row == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "请输入用户名或密码"})
			return
		}
		result := FormatSQLRowToMap(row)
		ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": result})
		return
	}
	ctx.JSON(iris.Map{"code": "-1", "msg": "请输入用户名或密码"})
}

/**
 * 添加人员或医生
 */
func PersonnelAdd(ctx iris.Context) {
	code := ctx.PostValue("code")
	name := ctx.PostValue("name")
	clinicCode := ctx.PostValue("clinic_code")
	departmentID := ctx.PostValue("department_id")
	weight := ctx.PostValue("weight")
	title := ctx.PostValue("title")
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	isClinicAdmin := false

	if code != "" && name != "" && clinicCode != "" && departmentID != "" && weight != "" && title != "" && username != "" && password != "" {
		tx, err := model.DB.Beginx()
		if err != nil {
			fmt.Println("err ===", err.Error())
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
		var personnelID int
		err = model.DB.QueryRow("insert into personnel(code, name, clinic_code, weight, title, username, password, is_clinic_admin) values ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", code, name, clinicCode, weight, title, username, password, isClinicAdmin).Scan(&personnelID)
		if err != nil {
			fmt.Println("11err =======", err)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": "插入失败"})
			return
		}
		fmt.Println("personnelID ======", personnelID)

		var resultID int
		err = model.DB.QueryRow("insert into department_personnel(department_id, personnel_id) values ($1, $2) RETURNING id", departmentID, personnelID).Scan(&resultID)
		fmt.Println("resultID =======", resultID)
		if err != nil {
			fmt.Println("err =======", err)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": "插入失败"})
			return
		}

		err = tx.Commit()
		if err != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
		ctx.JSON(iris.Map{"code": "200", "data": ""})
		return
	}

	ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
}

/**
 * 通过id获取人员（医生）
 */
func PersonnelGetByID(ctx iris.Context) {
	id := ctx.PostValue("id")
	if id != "" {
		row := model.DB.QueryRowx("SELECT * FROM personnel WHERE id=$1 ", id)
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
