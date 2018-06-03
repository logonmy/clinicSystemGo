package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

//Department 科室
type Department struct {
	DepartmentID  interface{} `json:"department_id" db:"department_id"`
	Code          interface{} `json:"code" db:"code"`
	Name          interface{} `json:"name" db:"name"`
	ClinicID      interface{} `json:"clinic_id" db:"clinic_id"`
	Weight        interface{} `json:"weight" db:"weight"`
	IsAppointment interface{} `json:"is_appointment" db:"is_appointment"`
}

//DepartmentCreate 创建科室
func DepartmentCreate(ctx iris.Context) {
	code := ctx.PostValue("code")
	name := ctx.PostValue("name")
	clinicID := ctx.PostValue("clinic_id")
	weight := ctx.PostValue("weight")
	if code == "" || name == "" || clinicID == "" || weight == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from clinic where id=$1 limit 1", clinicID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)
	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊所数据错误"})
		return
	}

	drow := model.DB.QueryRowx("select id from department where clinic_id = $1 and code=$2 limit 1", clinicID, code)
	if drow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
		return
	}
	department := FormatSQLRowToMap(drow)
	_, dok := department["id"]
	if dok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "科室编码已存在"})
		return
	}

	dnrow := model.DB.QueryRowx("select id from department where clinic_id = $1 and name=$2 limit 1", clinicID, name)
	if dnrow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
		return
	}
	departmentn := FormatSQLRowToMap(dnrow)
	_, dnok := departmentn["id"]
	if dnok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "科室名称已存在"})
		return
	}

	var departmentID int
	err := model.DB.QueryRow("INSERT INTO department (code, name, clinic_id, weight) VALUES ($1, $2, $3, $4) RETURNING id", code, name, clinicID, weight).Scan(&departmentID)
	if err != nil {
		fmt.Println("err ===", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
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
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
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

	total := model.DB.QueryRowx("SELECT count(id) as total FROM department WHERE (code ~*$1 OR name ~*$1) AND clinic_id=$2 and deleted_time is null", keyword, clinicID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx("SELECT * FROM department WHERE (code ~*$1 OR name ~*$1) AND clinic_id=$2 and deleted_time is null order by weight DESC offset $3 limit $4", keyword, clinicID, offset, limit)
	results = FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

//DepartmentDelete 删除科室
func DepartmentDelete(ctx iris.Context) {
	departmentID := ctx.PostValue("department_id")
	if departmentID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	crow := model.DB.QueryRowx("select id,code from department where id=$1 limit 1", departmentID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "删除失败"})
		return
	}
	department := FormatSQLRowToMap(crow)
	_, rok := department["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "科室数据错误"})
		return
	}
	code := department["code"]
	code = code.(string) + "#del"

	stmt, err := model.DB.Prepare("update department set code=$1,deleted_time=LOCALTIMESTAMP WHERE id=$2")
	if err != nil {
		fmt.Println("Perr ===", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	res, err := stmt.Exec(code, departmentID)
	if err != nil {
		fmt.Println("Eerr ===", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
	}
	ctx.JSON(iris.Map{"code": "200", "data": res})
}

//DepartmentUpdate 编辑科室
func DepartmentUpdate(ctx iris.Context) {
	var department Department
	err := ctx.ReadJSON(&department)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	if department.DepartmentID == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	var s []string

	crow := model.DB.QueryRowx("select id,clinic_id from department where id=$1 limit 1", department.DepartmentID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	sdepartment := FormatSQLRowToMap(crow)
	_, rok := sdepartment["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "科室数据错误"})
		return
	}
	clinicID := sdepartment["clinic_id"]

	if department.Code != nil {
		lrow := model.DB.QueryRowx("select id from department where code=$1 and id!=$2 and clinic_id=$3 limit 1", department.Code, department.DepartmentID, clinicID)
		if lrow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
			return
		}
		udepartment := FormatSQLRowToMap(lrow)
		_, dok := udepartment["id"]
		if dok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "科室编码已存在"})
			return
		}
		s = append(s, "code=:code")
	}
	if department.Name != nil {
		dnrow := model.DB.QueryRowx("select id from department where clinic_id = $1 and name=$2 and id!=$3 limit 1", clinicID, department.Name, department.DepartmentID)
		if dnrow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
			return
		}
		departmentn := FormatSQLRowToMap(dnrow)
		_, dnok := departmentn["id"]
		if dnok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "科室名称已存在"})
			return
		}
		s = append(s, "name=:name")
	}
	if department.IsAppointment != nil {
		s = append(s, "is_appointment=:is_appointment")
	}
	if department.Weight != nil {
		s = append(s, "weight=:weight")
	}
	s = append(s, "updated_time=LOCALTIMESTAMP")
	joinSQL := strings.Join(s, ",")
	updateSQL := "update department set " + joinSQL + " WHERE id=:department_id"

	_, errn := model.DB.NamedExec(updateSQL, department)

	if errn != nil {
		fmt.Println("Eerr ===", errn)
		ctx.JSON(iris.Map{"code": "-1", "msg": errn.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": nil})
}
