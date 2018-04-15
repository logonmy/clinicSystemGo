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
	apijson := APIJSON{}
	apijson.Code = -1
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	if username != "" && password != "" {
		personnel := model.Personnel{}
		err := model.DB.Get(&personnel, "SELECT * FROM personnel WHERE username=$1 AND password=$2 ", username, password)
		if err != nil {
			apijson.Msg = "用户名或密码不正确"
			fmt.Println("apijson", apijson)
			ctx.JSON(FormatResult(apijson))
			return
		}
		apijson.Code = 200
		apijson.Data = personnel
		ctx.JSON(FormatResult(apijson))
		return
	}
	ctx.JSON(FormatResult(apijson))
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
