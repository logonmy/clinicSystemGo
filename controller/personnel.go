package controller

import (
	"clinicSystemGo/model"
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
		fmt.Println("aaa", username)
		// err := model.DB.Get(&personnel, "SELECT * FROM personnel WHERE username=$1", username)
		rows, _ := model.DB.Query("SELECT * FROM personnel")
		results := FormatSQLRowToMapArray(rows)
		ctx.JSON(iris.Map{"code": "-1", "msg": "请输入用户名或密码", "data": results})
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
