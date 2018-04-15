package controller

import (
	"clinicSystemGo/model"
	"fmt"

	"github.com/kataras/iris"
)

/**
 * 创建科室
 */
func DepartmentCreate(ctx iris.Context) {
	code := ctx.PostValue("code")
	name := ctx.PostValue("name")
	clinicCode := ctx.PostValue("clinicCode")
	weight := ctx.PostValue("weight")
	if code == "" || name == "" || clinicCode == "" || weight == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	stmt, err := model.DB.Prepare("INSERT INTO department (code, name, clinic_code, weight) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		fmt.Println("Perr ===", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": "error"})
		return
	}
	res, err := stmt.Exec(code, name, clinicCode, weight)
	if err != nil {
		fmt.Println("Eerr ===", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": "error"})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": res})
}

/**
 * 获取科室
 */
func DepartmentGet(ctx iris.Context) {
	keyword := ctx.PostValue("keyword")
	clinicCode := ctx.PostValue("clinicCode")
	if clinicCode == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	var sql string
	departments := []model.Department{}
	if keyword != "" {
		sql = "SELECT * FROM department WHERE (code=$1 OR (name LIKE '%' || $1 || '%')) AND clinic_code=$2"
		err := model.DB.Select(&departments, sql, keyword, clinicCode)
		if err != nil {
			fmt.Println("err ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": "error"})
		}
	} else {
		sql = "SELECT * FROM department WHERE clinic_code=$1"
		err := model.DB.Select(&departments, sql, clinicCode)
		if err != nil {
			fmt.Println("err ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": "error"})
		}
	}
	ctx.JSON(iris.Map{"code": "200", "data": departments})
}
