package controller

import (
	"clinicSystemGo/model"
	"fmt"

	"github.com/kataras/iris"
)

// VisitTypeCreate 添加出诊类型
func VisitTypeCreate(ctx iris.Context) {

	code := ctx.PostValue("code")
	name := ctx.PostValue("name")
	openFlag := ctx.PostValue("open_flag")
	fee := ctx.PostValue("fee")

	fmt.Println(code, name, openFlag, fee)

	if code == "" || name == "" || openFlag == "" || fee == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	cmap := map[string]interface{}{
		"code":      code,
		"name":      name,
		"open_flag": openFlag,
		"fee":       fee,
	}
	_, err := model.DB.NamedExec(`INSERT INTO visit_type ( code, name, open_flag, fee) VALUES (:code, :name, :open_flag, :fee)`, cmap)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})

}

// VisitTypeList 获取就诊类型列表
func VisitTypeList(ctx iris.Context) {
	code := ctx.PostValue("code")
	sql := "SELECT * FROM visit_type"
	if code != "" {
		sql = sql + " where code=" + code
	}
	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(sql)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})

}
