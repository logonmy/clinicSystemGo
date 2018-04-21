package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strings"

	"github.com/kataras/iris"
)

// ChargeTypeInit 初始化费用类型
func ChargeTypeInit(ctx iris.Context) {
	names := []string{"西/成药处方", "中药处方", "检验医嘱", "检验项目", "检查医嘱", "材料费用", "其他费用", "诊疗项目"}
	sql := "INSERT INTO charge_project_type (name) VALUES "

	var sets []string
	for i := range names {
		sets = append(sets, "('"+names[i]+"')")
	}

	nsql := sql + strings.Join(sets, ",")

	_, err := model.DB.Query(nsql)
	if err != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// ChargeTypeCreate 创建费用类型
func ChargeTypeCreate(ctx iris.Context) {
	name := ctx.PostValue("name")
	if name == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	var ID int
	err := model.DB.QueryRow("INSERT INTO charge_project_type (name) VALUES ($1) RETURNING id", name).Scan(&ID)
	if err != nil {
		fmt.Println("err ===", err)
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": ID})
}
