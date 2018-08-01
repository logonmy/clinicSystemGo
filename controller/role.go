package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/kataras/iris"
)

//RoleCreate 权限角色添加
func RoleCreate(ctx iris.Context) {
	name := ctx.PostValue("name")
	clinicID := ctx.PostValue("clinic_id")
	items := ctx.PostValue("items")
	if name == "" || clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	var roleID string
	row := model.DB.QueryRowx("select id from role where name=$1 and clinic_id=$2 limit 1", name, clinicID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
		return
	}
	role := FormatSQLRowToMap(row)
	_, ok := role["id"]
	if ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "权限组名称已存在"})
		return
	}
	tx, err := model.DB.Begin()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	errq := tx.QueryRow("insert into role (name,clinic_id) values ($1,$2) RETURNING id", name, clinicID).Scan(&roleID)
	if errq != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errq.Error()})
		return
	}

	if items != "" {
		var results []map[string]string
		errj := json.Unmarshal([]byte(items), &results)

		if errj != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": errj.Error()})
			return
		}

		for _, v := range results {
			clinicFunctionMenuID := v["clinic_function_menu_id"]
			crow := model.DB.QueryRowx("select id from clinic_function_menu where id=$1 limit 1", clinicFunctionMenuID)
			if crow == nil {
				ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
				return
			}
			clinicFunctionMenu := FormatSQLRowToMap(crow)
			_, cok := clinicFunctionMenu["id"]
			if !cok {
				ctx.JSON(iris.Map{"code": "-1", "msg": "诊所菜单项不存在"})
				return
			}
			sql := "INSERT INTO role_clinic_function_menu (clinic_function_menu_id, role_id) VALUES ($1,$2)"
			_, erre := tx.Exec(sql, ToNullInt64(clinicFunctionMenuID), ToNullInt64(roleID))
			if erre != nil {
				fmt.Println("erre====", erre.Error())
				tx.Rollback()
				ctx.JSON(iris.Map{"code": "-1", "msg": erre.Error()})
				return
			}
		}
	}
	errc := tx.Commit()
	if errc != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errc.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": roleID})
}

//RoleUpdate 权限分组修改
func RoleUpdate(ctx iris.Context) {
	roleID := ctx.PostValue("role_id")
	name := ctx.PostValue("name")
	items := ctx.PostValue("items")
	if roleID == "" || name == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	row := model.DB.QueryRowx("select id from role where id=$1", roleID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	role := FormatSQLRowToMap(row)
	_, ok := role["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改的权限组不存在"})
		return
	}

	rrow := model.DB.QueryRowx("select id from role where name=$1 and id !=$2 limit 1", name, roleID)
	if rrow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	rrole := FormatSQLRowToMap(rrow)
	_, rok := rrole["id"]
	if rok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "权限组名称已存在"})
		return
	}

	tx, err := model.DB.Begin()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	_, erra := tx.Exec("update role set name=$1,updated_time=LOCALTIMESTAMP where id=$2", name, roleID)
	if erra != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": erra})
		return
	}

	if items != "" {
		var results []map[string]string
		err := json.Unmarshal([]byte(items), &results)
		fmt.Println("===", results)
		if err != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}

		sql1 := "delete from role_clinic_function_menu WHERE role_id=" + roleID
		_, errtx := tx.Exec(sql1)
		if errtx != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": errtx.Error()})
			return
		}

		for _, v := range results {
			clinicFunctionMenuID := v["clinic_function_menu_id"]
			crow := model.DB.QueryRowx("select id from clinic_function_menu where id=$1 limit 1", clinicFunctionMenuID)
			if crow == nil {
				ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
				return
			}
			clinicFunctionMenu := FormatSQLRowToMap(crow)
			_, cok := clinicFunctionMenu["id"]
			if !cok {
				ctx.JSON(iris.Map{"code": "-1", "msg": "诊所菜单项不存在"})
				return
			}

			sql := "INSERT INTO role_clinic_function_menu (clinic_function_menu_id, role_id) VALUES ($1,$2)"
			_, err = tx.Exec(sql, ToNullInt64(clinicFunctionMenuID), ToNullInt64(roleID))
			if err != nil {
				tx.Rollback()
				ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
				return
			}
		}
	}
	err3 := tx.Commit()
	if err3 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": roleID})
}

//RoleDelete 删除角色
func RoleDelete(ctx iris.Context) {
	roleID := ctx.PostValue("role_id")
	if roleID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	crow := model.DB.QueryRowx("select id,name from role where id=$1 limit 1", roleID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "删除失败"})
		return
	}
	role := FormatSQLRowToMap(crow)
	_, rok := role["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "删除的角色不存在"})
		return
	}
	name := role["name"]
	name = name.(string) + "#" + strconv.Itoa(int(time.Now().Unix()))

	_, err := model.DB.Exec("update role set name=$1,deleted_time=LOCALTIMESTAMP WHERE id=$2", name, roleID)
	if err != nil {
		fmt.Println("Perr ===", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//RoleList 权限角色列表
func RoleList(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	keyword := ctx.PostValue("keyword")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

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

	sql := `SELECT r.id as role_id,r.name,r.status,r.created_time, string_agg( fm.name,  ',') as function_menu_name FROM role r 
	left join role_clinic_function_menu rcfm on r.id = rcfm.role_id 
	left join clinic_function_menu cfm on rcfm.clinic_function_menu_id = cfm.id
	left join function_menu fm on fm.id = cfm.function_menu_id
	where r.clinic_id=:clinic_id and r.status = true and r.deleted_time is null
	group by (r.id, r.name, r.status, r.created_time)
	order by r.created_time desc`

	countSQL := `SELECT count(*) as total FROM role where clinic_id=:clinic_id and status = true and deleted_time is null`

	if keyword != "" {
		sql += ` and r.name ~*:keyword`
		countSQL += ` and name ~*:keyword`
	}

	var queryOptions = map[string]interface{}{
		"clinic_id": ToNullInt64(clinicID),
		"keyword":   ToNullString(keyword),
		"offset":    ToNullInt64(offset),
		"limit":     ToNullInt64(limit),
	}

	total, err2 := model.DB.NamedQuery(countSQL, queryOptions)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-2", "msg": err2.Error()})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err := model.DB.NamedQuery(sql+" offset :offset limit :limit", queryOptions)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	results := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
}

//RoleDetail 权限角色详情
func RoleDetail(ctx iris.Context) {
	roleID := ctx.PostValue("role_id")
	if roleID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id as role_id,name,status from role where id=$1", roleID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询结果不存在"})
		return
	}
	role := FormatSQLRowToMap(row)
	selectSQL := `select 
	rcf.clinic_function_menu_id,
	cfm.function_menu_id,
	fm.name as menu_name,
	fm.url as menu_url,
	fm.level,
	fm.weight,
	fm.ascription,
	fm.status,
	fm.icon,
	fm.parent_function_menu_id
	from role_clinic_function_menu rcf
	right join clinic_function_menu cfm on cfm.id = rcf.clinic_function_menu_id and cfm.status=true
	left join function_menu fm on fm.id = cfm.function_menu_id
	where rcf.role_id=$1 order by fm.level asc,fm.weight asc`
	rows, err2 := model.DB.Queryx(selectSQL, roleID)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err2})
		return
	}
	roles := FormatSQLRowsToMapArray(rows)

	// var funtionmenu []Funtionmenu
	// for rows.Next() {
	// 	var f Funtionmenu
	// 	err := rows.StructScan(&f)
	// 	if err != nil {
	// 		fmt.Println("err=====", err.Error())
	// 		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
	// 		return
	// 	}
	// 	funtionmenu = append(funtionmenu, f)
	// }

	// clinicFunctionMenu := FormatMenu(funtionmenu)
	// role["funtionMenus"] = clinicFunctionMenu

	role["funtionMenus"] = roles

	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": role})
}

//RoleAllocation 在角色下分配用户
func RoleAllocation(ctx iris.Context) {
	roleID := ctx.PostValue("role_id")
	items := ctx.PostValue("items")

	if roleID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
	}
	if items == "" {
		ctx.JSON(iris.Map{"code": "200", "data": nil})
		return
	}

	row := model.DB.QueryRowx("select id from role where id=$1 limit 1", roleID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "分配失败"})
		return
	}
	role := FormatSQLRowToMap(row)
	_, ok := role["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "权限组不存在"})
		return
	}

	var results []map[string]string
	errj := json.Unmarshal([]byte(items), &results)
	fmt.Println("===", results)
	if errj != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": errj.Error()})
		return
	}

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		ctx.JSON(iris.Map{"code": "-1", "msg": errb})
		return
	}
	_, errd := tx.Exec("delete from personnel_role where role_id=$1", roleID)

	if errd != nil {
		fmt.Println("errd ===", errd)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errd.Error()})
		return
	}

	for _, personnel := range results {
		personnelID := personnel["personnel_id"]
		prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", personnelID)
		if prow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "分配失败"})
			return
		}
		personneli := FormatSQLRowToMap(prow)
		_, pok := personneli["id"]
		if !pok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "所选用户不存在"})
			return
		}

		prrow := model.DB.QueryRowx("select personnel_id from personnel_role where personnel_id=$1 and role_id=$2 limit 1", personnelID, roleID)
		if prrow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "分配失败"})
			return
		}
		personnelRole := FormatSQLRowToMap(prrow)
		_, prok := personnelRole["personnel_id"]
		if prok {
			continue
		}
		_, err := tx.Exec("insert into personnel_role (personnel_id, role_id) values ($1,$2)", personnelID, roleID)

		if err != nil {
			fmt.Println("err ===", err)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
	}
	errc := tx.Commit()
	if errc != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errc.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//RoleFunctionUnset 获取角色未开通的菜单项
func RoleFunctionUnset(ctx iris.Context) {
	roleID := ctx.PostValue("role_id")
	if roleID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id,clinic_id from role where id=$1 limit 1", roleID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询失败"})
		return
	}
	role := FormatSQLRowToMap(row)
	_, ok := role["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询的权限组不存在"})
		return
	}
	clinicID := role["clinic_id"]

	selectSQL := `select 
	cfm.id as clinic_function_menu_id,
	fm.id as function_menu_id,
	fm.name as menu_name,
	fm.url as menu_url,
	fm.level,
	fm.weight,
	fm.ascription,
	fm.status,
	fm.icon,
	fm.parent_function_menu_id
	from clinic_function_menu cfm
	left join function_menu fm on fm.id = cfm.function_menu_id
	left join role_clinic_function_menu rcf on rcf.clinic_function_menu_id = cfm.id and rcf.role_id=$1
	where rcf.clinic_function_menu_id IS NULL and cfm.status=true and cfm.clinic_id=$2`

	rows, _ := model.DB.Queryx(selectSQL, roleID, clinicID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "查询失败"})
		return
	}

	var funtionmenu []Funtionmenu
	for rows.Next() {
		var f Funtionmenu
		err := rows.StructScan(&f)
		if err != nil {
			fmt.Println("err=====", err.Error())
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
		formatUnsetMenu := FormatUnsetMenu(f)
		fmt.Println("formatUnsetMenu=====", formatUnsetMenu)
		funtionmenu = append(funtionmenu, formatUnsetMenu...)
	}
	fmt.Println("funtionmenu=====", funtionmenu)

	result := FormatMenu(funtionmenu)

	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": result})
}

//PersonnelsByRole 角色分配的用户列表
func PersonnelsByRole(ctx iris.Context) {
	roleID := ctx.PostValue("role_id")

	if roleID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select p.id as personnel_id, p.name as personnel_name, d.name as department_name from personnel_role pr
	left join personnel p on p.id = pr.personnel_id
	left join department_personnel dp on dp.personnel_id = p.id
	left join department d on d.id = dp.department_id
	where pr.role_id=$1`

	rows, _ := model.DB.Queryx(selectSQL, roleID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "查询失败"})
		return
	}
	resData := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": resData})
}
