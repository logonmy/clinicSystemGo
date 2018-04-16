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
func PersonnelCreate(ctx iris.Context) {
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
		tx, err := model.DB.Begin()
		if err != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
		var personnelID int
		err = tx.QueryRow("insert into personnel(code, name, clinic_code, weight, title, username, password, is_clinic_admin) values ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", code, name, clinicCode, weight, title, username, password, isClinicAdmin).Scan(&personnelID)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
		var resultID int
		err = tx.QueryRow("insert into department_personnel(department_id, personnel_id, type) values ($1, $2, 2) RETURNING id", departmentID, personnelID).Scan(&resultID)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}

		err = tx.Commit()
		if err != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
		ctx.JSON(iris.Map{"code": "200", "data": iris.Map{"id": personnelID}})
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
		row := model.DB.QueryRowx("select p.id, p.name,p.weight,p.title,p.username,p.status,c.code as clinic_code, c.name as clinic_name,d.code as department_code, d.name as department_name, d.id as department_id from personnel p left join clinic c on p.clinic_code = c.code left join department_personnel dp on p.id = dp.personnel_id left join department d on dp.department_id = d.id where dp.type = 2 and p.id = 7;", id)
		fmt.Println("row =======", row)
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
