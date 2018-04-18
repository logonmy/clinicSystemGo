package controller

import (
	"clinicSystemGo/model"
	"fmt"

	"github.com/kataras/iris"
)

//DepartmentCreate 创建科室
func DepartmentCreate(ctx iris.Context) {
	code := ctx.PostValue("code")
	name := ctx.PostValue("name")
	clinicCode := ctx.PostValue("clinicCode")
	weight := ctx.PostValue("weight")
	if code == "" || name == "" || clinicCode == "" || weight == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	var departmentID int
	err := model.DB.QueryRow("INSERT INTO department (code, name, clinic_code, weight) VALUES ($1, $2, $3, $4) RETURNING id", code, name, clinicCode, weight).Scan(&departmentID)
	if err != nil {
		fmt.Println("err ===", err)
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": departmentID})
}

//DepartmentList 获取科室
func DepartmentList(ctx iris.Context) {
	keyword := ctx.PostValue("keyword")
	clinicCode := ctx.PostValue("clinicCode")
	if clinicCode == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	if keyword == "" {
		keyword = "%"
	}
	var results []map[string]interface{}
	rows, _ := model.DB.Queryx("SELECT * FROM department WHERE (code=$1 OR (name LIKE '%' || $1 || '%')) AND clinic_code=$2", keyword, clinicCode)
	results = FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results})
}

//DepartmentDelete 删除科室
func DepartmentDelete(ctx iris.Context) {
	departmentID := ctx.PostValue("departmentID")
	if departmentID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	stmt, err := model.DB.Prepare("DELETE from department WHERE id=$1")
	if err != nil {
		fmt.Println("Perr ===", err)
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	res, err := stmt.Exec(departmentID)
	if err != nil {
		fmt.Println("Eerr ===", err)
		ctx.JSON(iris.Map{"code": "1", "msg": err})
	}
	ctx.JSON(iris.Map{"code": "200", "data": res})
}

//DepartmentUpdate 编辑科室
func DepartmentUpdate(ctx iris.Context) {
	departmentID := ctx.PostValue("departmentID")
	code := ctx.PostValue("code")
	name := ctx.PostValue("name")
	clinicCode := ctx.PostValue("clinicCode")
	weight := ctx.PostValue("weight")
	if departmentID == "" || code == "" || name == "" || clinicCode == "" || weight == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	stmt, err := model.DB.Prepare("UPDATE department SET code=$2, name=$3, clinic_code=$4, weight=$5 WHERE id=$1")
	if err != nil {
		fmt.Println("Perr ===", err)
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	res, err := stmt.Exec(departmentID, code, name, clinicCode, weight)
	if err != nil {
		fmt.Println("Eerr ===", err)
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": res})
}