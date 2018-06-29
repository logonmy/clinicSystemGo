package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"

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
			clinicChildrenFunctionMenuID := v["clinic_function_menu_id"]
			crow := model.DB.QueryRowx("select id from clinic_children_function_menu where id=$1 limit 1", clinicChildrenFunctionMenuID)
			if crow == nil {
				ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
				return
			}
			clinicChildrenFunctionMenu := FormatSQLRowToMap(crow)
			_, cok := clinicChildrenFunctionMenu["id"]
			if !cok {
				ctx.JSON(iris.Map{"code": "-1", "msg": "诊所菜单项不存在"})
				return
			}
			sql := "INSERT INTO role_clinic_function_menu (clinic_children_function_menu_id, role_id) VALUES ($1,$2)"
			_, erre := tx.Exec(sql, ToNullInt64(clinicChildrenFunctionMenuID), ToNullInt64(roleID))
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
	row := model.DB.QueryRowx("select id from role where id=$1 limit 1", roleID)
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
			clinicChildrenFunctionMenuID := v["clinic_function_menu_id"]
			crow := model.DB.QueryRowx("select id from clinic_children_function_menu where id=$1 limit 1", clinicChildrenFunctionMenuID)
			if crow == nil {
				ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
				return
			}
			clinicChildrenFunctionMenu := FormatSQLRowToMap(crow)
			_, cok := clinicChildrenFunctionMenu["id"]
			if !cok {
				ctx.JSON(iris.Map{"code": "-1", "msg": "诊所菜单项不存在"})
				return
			}

			sql := "INSERT INTO role_clinic_function_menu (clinic_children_function_menu_id, role_id) VALUES ($1,$2)"
			_, err = tx.Exec(sql, ToNullInt64(clinicChildrenFunctionMenuID), ToNullInt64(roleID))
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

//RoleList 权限角色列表
func RoleList(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	keyword := ctx.PostValue("keyword")

	sql := `SELECT id as role_id,name,status,created_time FROM role where clinic_id=:clinic_id`

	if keyword != "" {
		sql += ` and name ~*:keyword`
	}

	var queryOptions = map[string]interface{}{
		"clinic_id": ToNullInt64(clinicID),
		"keyword":   ToNullString(keyword),
	}

	rows, err := model.DB.NamedQuery(sql, queryOptions)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	results := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})
}

//RoleDetail 权限角色详情
func RoleDetail(ctx iris.Context) {
	roleID := ctx.PostValue("role_id")
	if roleID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	arows := model.DB.QueryRowx("select id as role_id,name,status from role where id=$1", roleID)
	if arows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询结果不存在"})
		return
	}
	selectSQL := `select rcf.clinic_children_function_menu_id as clinic_function_menu_id,
	ccf.children_function_menu_id as function_menu_id,pf.id as parent_id,pf.url as parent_url,
	pf.name as parent_name,cf.url as menu_url,cf.name as menu_name from role_clinic_function_menu rcf
	left join clinic_children_function_menu ccf on ccf.id = rcf.clinic_children_function_menu_id and ccf.status=true
	left join children_function_menu cf on cf.id = ccf.children_function_menu_id
	left join parent_function_menu pf on pf.id = cf.parent_function_menu_id
	where rcf.role_id=$1`
	rows, err2 := model.DB.Queryx(selectSQL, roleID)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err2})
		return
	}
	role := FormatSQLRowToMap(arows)
	clinicFunctionMenu := FormatSQLRowsToMapArray(rows)
	menus := FormatFuntionmenus(clinicFunctionMenu)
	role["funtionMenus"] = menus
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
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
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

	selectSQL := `select ccf.id as clinic_function_menu_id,
	ccf.children_function_menu_id as function_menu_id,pf.id as parent_id,pf.url as parent_url,
	pf.name as parent_name,cf.url as menu_url,cf.name as menu_name from clinic_children_function_menu ccf
	left join children_function_menu cf on cf.id = ccf.children_function_menu_id
	left join role_clinic_function_menu rcf on rcf.clinic_children_function_menu_id = ccf.id and rcf.role_id=$1
	left join parent_function_menu pf on pf.id = cf.parent_function_menu_id
	where rcf.clinic_children_function_menu_id IS NULL and ccf.status=true and ccf.clinic_id=$2`

	rows, _ := model.DB.Queryx(selectSQL, roleID, clinicID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "查询失败"})
		return
	}
	parentFunctionMenu := FormatSQLRowsToMapArray(rows)
	menus := FormatFuntionmenus(parentFunctionMenu)

	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": menus})
}
