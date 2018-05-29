package controller

import (
	"clinicSystemGo/model"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"strings"

	"github.com/kataras/iris"
)

// PersonnelLogin 创建医院管理员
func PersonnelLogin(ctx iris.Context) {
	IP := ctx.RemoteAddr()
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	if username != "" && password != "" {
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(password))
		passwordMd5 := hex.EncodeToString(md5Ctx.Sum(nil))
		row := model.DB.QueryRowx("select a.id, a.code, a.name, a.username, b.id as clinic_id, b.name as clinic_name from personnel a left join clinic b on a.clinic_id = b.id where a.username = $1 and a.password = $2", username, passwordMd5)
		if row == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "用户名或密码错误"})
			return
		}
		result := FormatSQLRowToMap(row)
		if _, ok := result["id"]; ok {
			personnelID := result["id"]
			_ = model.DB.MustExec("INSERT INTO personnel_login_record (personnel_id, ip) VALUES ($1, $2) RETURNING id", personnelID, IP)
			countRow := model.DB.QueryRowx("select count(*) as count from personnel_login_record where personnel_id = $1", personnelID)
			count := FormatSQLRowToMap(countRow)
			ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": result, "login_times": count["count"]})
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
	clinicID := ctx.PostValue("clinic_id")
	departmentID := ctx.PostValue("department_id")
	weight := ctx.PostValue("weight")
	title := ctx.PostValue("title")
	personnelType := ctx.PostValue("personnel_type")
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	isClinicAdmin := false

	if code != "" && name != "" && clinicID != "" && departmentID != "" && weight != "" && title != "" && username != "" && password != "" && personnelType != "" {
		tx, err := model.DB.Begin()
		if err != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(password))
		passwordMd5 := hex.EncodeToString(md5Ctx.Sum(nil))
		var personnelID int
		err = tx.QueryRow("insert into personnel(code, name, clinic_id, weight, title, username, password, is_clinic_admin) values ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", code, name, clinicID, weight, title, username, passwordMd5, isClinicAdmin).Scan(&personnelID)
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
			c.id as clinic_id, c.name as clinic_name,
			d.code as department_code, d.name as department_name, d.id as department_id
			from personnel p 
			left join clinic c on p.clinic_id = c.id 
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
	clinicID := ctx.PostValue("clinic_id")
	personnelType := ctx.PostValue("personnel_type")
	deparmentID := ctx.PostValue("department_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	keyword := ctx.PostValue("keyword")
	if clinicID == "" {
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
	left join clinic c on p.clinic_id = c.id 
	left join department_personnel dp on p.id = dp.personnel_id
	left join department d on dp.department_id = d.id
	where p.clinic_id = $1 and (p.code ~$2 or p.name ~$2)`
	if deparmentID != "" {
		jionSQL += " and d.id = " + deparmentID
	}

	if personnelType != "" {
		jionSQL += " and dp.type = " + personnelType
	}

	total := model.DB.QueryRowx(`select count(p.id) as total `+jionSQL, clinicID, keyword)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rowSQL := `select p.id, p.code, p.name,p.weight,p.title,p.username,p.status,p.is_appointment,
	c.id as clinic_id, c.name as clinic_name,
	dp.type as personnel_type,
	d.code as department_code, d.name as department_name, d.id as department_id ` + jionSQL + " offset $3 limit $4"

	rows, err1 := model.DB.Queryx(rowSQL, clinicID, keyword, offset, limit)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})
}

// PersonnelWithAuthorizationList 获取开通了权限的人员列表
func PersonnelWithAuthorizationList(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	keyword := ctx.PostValue("keyword")
	if clinicID == "" {
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

	countSQL := `select count(p.id) as total  from personnel p
	left join clinic c on p.clinic_id = c.id 
	left join department_personnel dp on p.id = dp.personnel_id
	left join department d on dp.department_id = d.id
	left join (select count(pr.personnel_id) as total,pr.personnel_id from personnel_role pr 
		left join personnel p on p.id=pr.personnel_id group by pr.personnel_id) prc on prc.personnel_id=p.id
	where p.clinic_id=:clinic_id and prc.total>0 and (p.code ~:keyword or p.name ~:keyword)`

	var queryOption = map[string]interface{}{
		"clinic_id": ToNullInt64(clinicID),
		"keyword":   keyword,
		"offset":    ToNullInt64(offset),
		"limit":     ToNullInt64(limit),
	}
	fmt.Println("===%&&&&", queryOption)

	total, err := model.DB.NamedQuery(countSQL, queryOption)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rowSQL := `select p.id, p.code, p.name,p.weight,p.title,p.username,prc.status as personnel_role_status,p.is_appointment,
	c.id as clinic_id, c.name as clinic_name,dp.type as personnel_type,d.code as department_code,
	d.name as department_name, d.id as department_id from personnel p 
	left join clinic c on p.clinic_id = c.id 
	left join department_personnel dp on p.id = dp.personnel_id
	left join department d on dp.department_id = d.id
	left join (select count(pr.personnel_id) as total,pr.personnel_id,pr.status from personnel_role pr 
		left join personnel p on p.id=pr.personnel_id group by pr.personnel_id,pr.status) prc on prc.personnel_id=p.id
	where p.clinic_id=:clinic_id and prc.total>0 and (p.code ~:keyword or p.name ~:keyword) offset :offset limit :limit`

	selectSQL := `select pr.personnel_id,pr.role_id,r.name as role_name from personnel_role pr
		left join role r on r.id = pr.role_id`

	rows, err1 := model.DB.NamedQuery(rowSQL, queryOption)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	personnel := FormatSQLRowsToMapArray(rows)

	prows, err2 := model.DB.NamedQuery(selectSQL, queryOption)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err2})
		return
	}
	personnelRole := FormatSQLRowsToMapArray(prows)
	for _, p := range personnel {
		personnelID := p["id"]
		var roles []interface{}
		for _, pr := range personnelRole {
			rolePersonnelID := pr["personnel_id"]
			if personnelID == rolePersonnelID {
				roles = append(roles, pr)
			}
		}
		p["roles"] = roles
	}
	ctx.JSON(iris.Map{"code": "200", "data": personnel, "page_info": pageInfo})
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
		_, err := tx.Exec("update department_personnel set department_id= $1,updated_time=LOCALTIMESTAMP where personnel_id = $2 and type= $3", departmentID, id, personnelType)
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

// PersonnelAuthorizationAllocation 用户权限分配
func PersonnelAuthorizationAllocation(ctx iris.Context) {
	id := ctx.PostValue("id")
	items := ctx.PostValue("items")
	if id == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
		return
	}

	lrow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", id)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	personnel := FormatSQLRowToMap(lrow)
	_, lok := personnel["id"]
	if !lok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "用户不存在"})
		return
	}

	if items == "" {
		ctx.JSON(iris.Map{"code": "200", "data": nil})
		return
	}
	var results []map[string]string
	reErr := json.Unmarshal([]byte(items), &results)
	if reErr != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": reErr.Error()})
		return
	}

	tx, err := model.DB.Begin()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	insertSQL := `insert into personnel_role (personnel_id,role_id) values ($1,$2)`
	for _, v := range results {
		roleID := v["role_id"]
		rrow := model.DB.QueryRowx("select id from role where id=$1 limit 1", roleID)
		if rrow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
			return
		}
		role := FormatSQLRowToMap(rrow)
		_, rok := role["id"]
		if !rok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "所选权限组不存在"})
			return
		}
		_, err3 := tx.Exec(insertSQL,
			ToNullInt64(id),
			ToNullInt64(roleID),
		)
		if err3 != nil {
			fmt.Println(" err3====", err3)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": nil})
}
