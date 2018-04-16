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
		fmt.Println("aaa", username, passwordMd5)
		row := model.DB.QueryRowx("select a.id, a.code, a.name, a.username, b.code as clinic_code, b.name as clinic_name from personnel a left join clinic b on a.clinic_code = b.code where  username = $1 and password = $2", username, passwordMd5)
		fmt.Println("row ========", row)
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
 * 获取人员（医生）
 */
func PersonnelGetByID(ctx iris.Context) {
	id := ctx.PostValue("id")
	if id != "" {
		personnel := model.Personnel{}
		err := model.DB.Get(&personnel, "SELECT * FROM personnel WHERE id=$1 ", id)
		if err != nil {
			fmt.Println("err ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": "用户名或密码不正确"})
			return
		}
		ctx.JSON(iris.Map{"code": "200", "data": personnel})
		return
	}
	ctx.JSON(iris.Map{"code": "-1", "msg": "请输入用户名或密码"})
}
