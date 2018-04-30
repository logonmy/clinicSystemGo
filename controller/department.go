package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"
	"time"

	"github.com/kataras/iris"
)

//DepartmentCreate 创建科室
func DepartmentCreate(ctx iris.Context) {
	code := ctx.PostValue("code")
	name := ctx.PostValue("name")
	clinicID := ctx.PostValue("clinic_id")
	weight := ctx.PostValue("weight")
	if code == "" || name == "" || clinicID == "" || weight == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from department where clinic_id = $1 and code=$2 limit 1", clinicID, code)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
		return
	}
	department := FormatSQLRowToMap(row)
	_, ok := department["id"]
	if ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "科室编码已存在"})
		return
	}

	var departmentID int
	err := model.DB.QueryRow("INSERT INTO department (code, name, clinic_id, weight) VALUES ($1, $2, $3, $4) RETURNING id", code, name, clinicID, weight).Scan(&departmentID)
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
	clinicID := ctx.PostValue("clinic_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	if offset == "" {
		offset = "0"
	}

	if limit == "" {
		limit = "10"
	}

	_, err := strconv.Atoi(offset)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "offset 必须为数字"})
		return
	}
	_, err = strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "limit 必须为数字"})
		return
	}

	total := model.DB.QueryRowx("SELECT count(id) as total FROM department WHERE (code=$1 OR (name LIKE '%' || $1 || '%')) AND clinic_id=$2", keyword, clinicID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx("SELECT * FROM department WHERE (code=$1 OR (name LIKE '%' || $1 || '%')) AND clinic_id=$2 offset $3 limit $4", keyword, clinicID, offset, limit)
	results = FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
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
	clinicID := ctx.PostValue("clinic_id")
	weight := ctx.PostValue("weight")
	if departmentID == "" || code == "" || name == "" || clinicID == "" || weight == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	stmt, err := model.DB.Prepare("UPDATE department SET code=$2, name=$3, clinic_id=$4, weight=$5,updated_time=$6 WHERE id=$1")
	if err != nil {
		fmt.Println("Perr ===", err)
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	res, err := stmt.Exec(departmentID, code, name, clinicID, weight, time.Now())
	if err != nil {
		fmt.Println("Eerr ===", err)
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": res})
}
