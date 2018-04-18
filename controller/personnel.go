package controller

import (
	"clinicSystemGo/model"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"

	"strings"

	"github.com/kataras/iris"
)

// PersonnelLogin 创建医院管理员
func PersonnelLogin(ctx iris.Context) {
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	if username != "" && password != "" {
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(password))
		passwordMd5 := hex.EncodeToString(md5Ctx.Sum(nil))
		row := model.DB.QueryRowx("select a.id, a.code, a.name, a.username, b.code as clinic_code, b.name as clinic_name from personnel a left join clinic b on a.clinic_code = b.code where a.username = $1 and a.password = $2", username, passwordMd5)
		if row == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "用户名或密码错误"})
			return
		}
		result := FormatSQLRowToMap(row)
		if _, ok := result["id"]; ok {
			ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": result})
			return
		}
		ctx.JSON(iris.Map{"code": "-1", "msg": "用户名或密码错误"})
		return
	}
	ctx.JSON(iris.Map{"code": "-1", "msg": "请输入用户名或密码"})
}

// PersonnelCreate 添加人员或医生
func PersonnelCreate(ctx iris.Context) {
	code := ctx.PostValue("code")
	name := ctx.PostValue("name")
	clinicCode := ctx.PostValue("clinic_code")
	departmentID := ctx.PostValue("department_id")
	weight := ctx.PostValue("weight")
	title := ctx.PostValue("title")
	personnelType := ctx.PostValue("personnel_type")
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	isClinicAdmin := false

	if code != "" && name != "" && clinicCode != "" && departmentID != "" && weight != "" && title != "" && username != "" && password != "" && personnelType != "" {
		tx, err := model.DB.Begin()
		if err != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(password))
		passwordMd5 := hex.EncodeToString(md5Ctx.Sum(nil))
		var personnelID int
		err = tx.QueryRow("insert into personnel(code, name, clinic_code, weight, title, username, password, is_clinic_admin) values ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", code, name, clinicCode, weight, title, username, passwordMd5, isClinicAdmin).Scan(&personnelID)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
		var resultID int
		err = tx.QueryRow("insert into department_personnel(department_id, personnel_id, type) values ($1, $2, $3) RETURNING id", departmentID, personnelID, personnelType).Scan(&resultID)
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

// PersonnelGetByID 通过id获取人员（医生）
func PersonnelGetByID(ctx iris.Context) {
	id := ctx.PostValue("id")
	if id != "" {
		row := model.DB.QueryRowx(`select p.id, p.name,p.weight,p.title,p.username,p.status,p.is_appointment,
			c.code as clinic_code, c.name as clinic_name,
			d.code as department_code, d.name as department_name, d.id as department_id
			from personnel p 
			left join clinic c on p.clinic_code = c.code 
			left join department_personnel dp on p.id = dp.personnel_id
			left join department d on dp.department_id = d.id
			where dp.type = 2 and p.id = $1;`, id)
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

// PersonnelList 获取人员列表
func PersonnelList(ctx iris.Context) {
	clinicCode := ctx.PostValue("clinic_code")
	personnelType := ctx.PostValue("personnel_type")
	deparmentID := ctx.PostValue("department_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	keyword := ctx.PostValue("keyword")
	if clinicCode == "" || personnelType == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
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

	jionSQL := `from personnel p 
	left join clinic c on p.clinic_code = c.code 
	left join department_personnel dp on p.id = dp.personnel_id
	left join department d on dp.department_id = d.id
	where p.clinic_code = $1 and (p.code like '%' || $2 || '%' or p.name like '%' || $2 || '%') and dp.type = $3`
	if deparmentID != "" {
		jionSQL += " and d.id = " + deparmentID
	}

	total := model.DB.QueryRowx(`select count(p.id) as total `+jionSQL, clinicCode, keyword, personnelType)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rowSQL := `select p.id, p.name,p.weight,p.title,p.username,p.status,p.is_appointment,
	c.code as clinic_code, c.name as clinic_name,
	dp.type as personnel_type,
	d.code as department_code, d.name as department_name, d.id as department_id ` + jionSQL + " offset $4 limit $5"

	rows, err1 := model.DB.Queryx(rowSQL, clinicCode, keyword, personnelType, offset, limit)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})
}

// PersonnelUpdate 修改人员
func PersonnelUpdate(ctx iris.Context) {
	id := ctx.PostValue("id")
	departmentID := ctx.PostValue("department_id")
	weight := ctx.PostValue("weight")
	title := ctx.PostValue("title")
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	personnelType := ctx.PostValue("personnel_type")
	isAppointment := ctx.PostValue("is_appointment")
	status := ctx.PostValue("status")
	if id == "" || personnelType == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
		return
	}
	var sets []string
	if weight != "" {
		sets = append(sets, "weight="+weight)
	}
	if title != "" {
		sets = append(sets, "title='"+title+"'")
	}
	if username != "" {
		sets = append(sets, "username='"+username+"'")
	}
	if status != "" {
		sets = append(sets, "status="+status)
	}
	if isAppointment != "" {
		sets = append(sets, "is_appointment="+isAppointment)
	}
	if password != "" {
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(password))
		passwordMd5 := hex.EncodeToString(md5Ctx.Sum(nil))
		sets = append(sets, "password='"+passwordMd5+"'")
	}
	setStr := strings.Join(sets, ",")
	psql := "update personnel set " + setStr + " where id=" + id
	fmt.Println("===", psql)
	tx, err := model.DB.Begin()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	if departmentID != "" {
		_, err := tx.Exec("update department_personnel set department_id= $1 where personnel_id = $2 and type= $3", departmentID, id, personnelType)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
	}
	_, err = tx.Exec(psql)
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
	ctx.JSON(iris.Map{"code": "200", "data": nil})
}
