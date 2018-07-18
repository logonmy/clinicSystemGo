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

//Personnel 人员
type Personnel struct {
	PersonnelID   interface{} `json:"personnel_id" db:"personnel_id"`
	DepartmentID  interface{} `json:"department_id" db:"department_id"`
	Title         interface{} `json:"title" db:"title"`
	Code          interface{} `json:"code" db:"code"`
	Name          interface{} `json:"name" db:"name"`
	Username      interface{} `json:"username" db:"username"`
	Password      interface{} `json:"password" db:"password"`
	PersonnelType interface{} `json:"personnel_type" db:"personnel_type"`
	ClinicID      interface{} `json:"clinic_id" db:"clinic_id"`
	Weight        interface{} `json:"weight" db:"weight"`
	IsAppointment interface{} `json:"is_appointment" db:"is_appointment"`
	Status        interface{} `json:"status" db:"status"`
}

// PersonnelLogin 登录
func PersonnelLogin(ctx iris.Context) {
	IP := ctx.RemoteAddr()
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	if username != "" && password != "" {
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(password))
		passwordMd5 := hex.EncodeToString(md5Ctx.Sum(nil))
		row := model.DB.QueryRowx(`select a.id, a.code, a.name, a.username, 
			b.id as clinic_id, b.name as clinic_name 
			from personnel a left join clinic b on a.clinic_id = b.id 
			where a.username = $1 and a.password = $2`, username, passwordMd5)
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

	if code != "" && name != "" && clinicID != "" && departmentID != "" && title != "" && personnelType != "" {
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

		drow := model.DB.QueryRowx("select id from personnel where clinic_id = $1 and code=$2 limit 1", clinicID, code)
		if drow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
			return
		}
		personnel := FormatSQLRowToMap(drow)
		_, pok := personnel["id"]
		if pok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "人员编码已存在"})
			return
		}

		if username != "" {
			urow := model.DB.QueryRowx("select id from personnel where clinic_id = $1 and username=$2 limit 1", clinicID, username)
			if urow == nil {
				ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
				return
			}
			personnel := FormatSQLRowToMap(urow)
			_, uok := personnel["id"]
			if uok {
				ctx.JSON(iris.Map{"code": "-1", "msg": "账号已存在"})
				return
			}
		}

		tx, err := model.DB.Begin()
		if err != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
		passwordMd5 := ""
		if password != "" {
			md5Ctx := md5.New()
			md5Ctx.Write([]byte(password))
			passwordMd5 = hex.EncodeToString(md5Ctx.Sum(nil))
		}
		var personnelID int
		if weight == "" {
			weight = "1"
		}
		err = tx.QueryRow("insert into personnel(code, name, clinic_id, weight, title, username, password, is_clinic_admin) values ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", code, name, clinicID, ToNullInt64(weight), ToNullString(title), ToNullString(username), ToNullString(passwordMd5), isClinicAdmin).Scan(&personnelID)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}

		_, err = tx.Exec("insert into department_personnel(department_id, personnel_id, type) values ($1, $2, $3)", departmentID, personnelID, personnelType)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}

		err = tx.Commit()
		if err != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
		ctx.JSON(iris.Map{"code": "200", "data": nil})
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
	where p.clinic_id = $1 and p.deleted_time is null and (p.code ~$2 or p.name ~$2)`
	if deparmentID != "" {
		jionSQL += " and d.id = " + deparmentID
	}

	if personnelType != "" {
		jionSQL += " and dp.type = " + personnelType
	}

	total := model.DB.QueryRowx(`select count(p.id) as total `+jionSQL, clinicID, keyword)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rowSQL := `select p.id, p.code, p.name,p.weight,p.title,p.username,p.status,p.is_appointment,
	c.id as clinic_id, c.name as clinic_name,
	dp.type as personnel_type,
	d.code as department_code, d.name as department_name, d.id as department_id ` + jionSQL + "order by weight DESC offset $3 limit $4"

	rows, err1 := model.DB.Queryx(rowSQL, clinicID, keyword, offset, limit)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
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
	where p.clinic_id=:clinic_id and p.deleted_time is null and prc.total>0 and (p.code ~:keyword or p.name ~:keyword)`

	var queryOption = map[string]interface{}{
		"clinic_id": ToNullInt64(clinicID),
		"keyword":   keyword,
		"offset":    ToNullInt64(offset),
		"limit":     ToNullInt64(limit),
	}
	fmt.Println("===%&&&&", queryOption)

	total, err := model.DB.NamedQuery(countSQL, queryOption)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
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
	where p.clinic_id=:clinic_id and p.deleted_time is null and prc.total>0 and (p.code ~:keyword or p.name ~:keyword) offset :offset limit :limit`

	selectSQL := `select pr.personnel_id,pr.role_id,r.name as role_name from personnel_role pr
		left join role r on r.id = pr.role_id`

	rows, err1 := model.DB.NamedQuery(rowSQL, queryOption)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}
	personnel := FormatSQLRowsToMapArray(rows)

	prows, err2 := model.DB.NamedQuery(selectSQL, queryOption)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
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
	var personnel Personnel
	err := ctx.ReadJSON(&personnel)
	fmt.Println("personnel===", personnel)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	if personnel.PersonnelID == nil || personnel.PersonnelType == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	var s []string

	crow := model.DB.QueryRowx("select id,clinic_id from personnel where id=$1 limit 1", personnel.PersonnelID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	spersonnel := FormatSQLRowToMap(crow)
	_, rok := spersonnel["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "人员数据错误"})
		return
	}
	clinicID := spersonnel["clinic_id"]

	if personnel.Code != nil {
		lrow := model.DB.QueryRowx("select id from personnel where code=$1 and id!=$2 and clinic_id=$3 limit 1", personnel.Code, personnel.PersonnelID, clinicID)
		if lrow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
			return
		}
		upersonnel := FormatSQLRowToMap(lrow)
		_, dok := upersonnel["id"]
		if dok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "人员编码已存在"})
			return
		}
		s = append(s, "code=:code")
	}
	if personnel.Name != nil {
		s = append(s, "name=:name")
	}
	if personnel.Title != nil {
		s = append(s, "title=:title")
	}
	if personnel.Username != nil {
		s = append(s, "username=:username")
	}
	if personnel.IsAppointment != nil {
		s = append(s, "is_appointment=:is_appointment")
	}
	if personnel.Weight != nil {
		s = append(s, "weight=:weight")
	}
	if personnel.Status != nil {
		s = append(s, "status=:status")
	}
	if personnel.Password != nil {
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(personnel.Password.(string)))
		personnel.Password = hex.EncodeToString(md5Ctx.Sum(nil))
		s = append(s, "password=:password")
	}

	tx, err := model.DB.Beginx()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	if personnel.DepartmentID != nil {
		_, err := tx.NamedExec("update department_personnel set department_id=:department_id,updated_time=LOCALTIMESTAMP where personnel_id=:personnel_id and type=:personnel_type", personnel)
		if err != nil {
			fmt.Println("5")
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	}

	s = append(s, "updated_time=LOCALTIMESTAMP")
	joinSQL := strings.Join(s, ",")
	psql := "update personnel set " + joinSQL + " where id=:personnel_id"
	fmt.Println("psql", psql)
	_, err = tx.NamedExec(psql, personnel)
	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
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
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	_, err1 := tx.Exec("delete from personnel_role where personnel_id=$1", ToNullInt64(id))
	if err1 != nil {
		fmt.Println(" err1====", err1.Error())
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
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
			fmt.Println(" err3====", err3.Error())
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//PersonnelDelete 删除人员
func PersonnelDelete(ctx iris.Context) {
	personnelID := ctx.PostValue("personnel_id")
	if personnelID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	crow := model.DB.QueryRowx("select id,code from personnel where id=$1 limit 1", personnelID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "删除失败"})
		return
	}
	personnel := FormatSQLRowToMap(crow)
	_, rok := personnel["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "科室数据错误"})
		return
	}
	code := personnel["code"]
	code = code.(string) + "#del"

	_, err := model.DB.Exec("update personnel set code=$1,deleted_time=LOCALTIMESTAMP WHERE id=$2", code, personnelID)
	if err != nil {
		fmt.Println("Perr ===", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//FunMenusByPersonnel 通过用户查询用户菜单
func FunMenusByPersonnel(ctx iris.Context) {
	personnelID := ctx.PostValue("id")

	if personnelID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select DISTINCT fm.parent_function_menu_id, fm.id as function_menu_id, fm.url as menu_url, fm.name as menu_name, fm.weight, fm.level, fm.icon from personnel_role pr 
	left join role_clinic_function_menu rcfm on pr.role_id = rcfm.role_id
	left join clinic_function_menu cfm on rcfm.clinic_function_menu_id = cfm.id
	left join function_menu fm on cfm.function_menu_id = fm.id 
	where cfm.status = true and rcfm.status = true and pr.personnel_id = $1
	order by fm.level, fm.weight asc `

	rows, err := model.DB.Queryx(selectSQL, personnelID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	result := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": result})

}

//PersonnelRoles 通过用户查询用户角色
func PersonnelRoles(ctx iris.Context) {
	personnelID := ctx.PostValue("id")

	if personnelID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select r.id as role_id,r.name,r.status,r.created_time from personnel_role pr 
	left join role r on pr.role_id = r.id
	where pr.personnel_id = $1 and r.status=true;`

	rows, err := model.DB.Queryx(selectSQL, personnelID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	result := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": result})

}

// PersonnelWithUsername 有账号的医生礼拜（包含角色）
func PersonnelWithUsername(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	keyword := ctx.PostValue("keyword")
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

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	var queryOptions = map[string]interface{}{
		"clinic_id": ToNullInt64(clinicID),
		"keyword":   ToNullString(keyword),
		"offset":    ToNullInt64(offset),
		"limit":     ToNullInt64(limit),
	}

	countSQL := `select count(*) as total from personnel where username is not null and deleted_time is null and clinic_id = :clinic_id`
	rowSQL := `select p.id as personnel_id, p.username, p.name as personnel_name, d.name as department_name, p.status, dp.type as personnel_type, string_agg (r.name, '，') as role_name from personnel p 
	left join department_personnel dp on dp.personnel_id = p.id
	left join department d on dp.department_id = d.id 
	left join personnel_role pr on p.id = pr.personnel_id
	left join role r on pr.role_id = r.id
	where p.username is not null and p.deleted_time is null and p.clinic_id = :clinic_id `

	if keyword != "" {
		countSQL += ` and (name ~*:keyword or username ~*:keyword) `
		rowSQL += ` and (p.name ~*:keyword or p.username ~*:keyword) `
	}

	total, err2 := model.DB.NamedQuery(countSQL, queryOptions)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-2", "msg": err2.Error()})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err3 := model.DB.NamedQuery(rowSQL+` group by (p.id, p.username, p.name, d.name, p.status, dp.type)
	order by p.id asc offset :offset limit :limit`, queryOptions)
	if err3 != nil {
		ctx.JSON(iris.Map{"code": "-2", "msg": err3.Error()})
		return
	}
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

// UpdatePersonnelStatus 修改账号生效状态
func UpdatePersonnelStatus(ctx iris.Context) {
	personnelID := ctx.PostValue("personnel_id")
	status := ctx.PostValue("status")
	if personnelID == "" || status == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	querySQL := `update personnel set status = $1 where id = $2`

	_, err := model.DB.Exec(querySQL, status, personnelID)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "msg": "修改成功"})
}

// UpdatePersonnelUsername 修改用户名密码
func UpdatePersonnelUsername(ctx iris.Context) {
	personnelID := ctx.PostValue("personnel_id")
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	if personnelID == "" || username == "" || password == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	passwordMd5 := hex.EncodeToString(md5Ctx.Sum(nil))

	querySQL := `update personnel set username = $1, password = $2 where id = $3`

	_, err := model.DB.Exec(querySQL, username, passwordMd5, personnelID)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "msg": "修改成功"})
}
