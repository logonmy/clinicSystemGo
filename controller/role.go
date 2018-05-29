package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"

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
	row := model.DB.QueryRowx("select id from role name=$1 and clinic_id=$2 limit 1", name, clinicID)
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
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	err = tx.QueryRow("insert into role(name,clinic_id) values ($1,$2) RETURNING id", name, clinicID).Scan(&roleID)
	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	if items != "" {
		var results []map[string]string
		err := json.Unmarshal([]byte(items), &results)

		if err != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}

		for _, v := range results {
			clinicChildrenFunctionMenuID := v["functionMenu_id"]
			crow := model.DB.QueryRowx("select id from clinic_children_functionMenu where id=$1 limit 1", clinicChildrenFunctionMenuID)
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

			sql := "INSERT INTO role_clinic_functionMenu (clinic_children_functionMenu_id, role_id) VALUES ($1,$2)"
			_, err = tx.Exec(sql, ToNullInt64(clinicChildrenFunctionMenuID), ToNullInt64(roleID))
			if err != nil {
				tx.Rollback()
				ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
				return
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
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

		sql1 := "delete from role_clinic_functionMenu WHERE role_id=" + roleID
		_, errtx := tx.Exec(sql1)
		if errtx != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": errtx.Error()})
			return
		}

		for _, v := range results {
			clinicChildrenFunctionMenuID := v["functionMenu_id"]
			crow := model.DB.QueryRowx("select id from clinic_children_functionMenu where id=$1 limit 1", clinicChildrenFunctionMenuID)
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

			sql := "INSERT INTO role_clinic_functionMenu (clinic_children_functionMenu_id, role_id) VALUES ($1,$2)"
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

	sql := `SELECT id as role_id,name,status,created_time FROM role where clinic_id=$1`

	if keyword != "" {
		sql += " and name ~'" + keyword + "'"
	}

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(sql, clinicID)
	results = FormatSQLRowsToMapArray(rows)

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
	selectSQL := `select rcf.clinic_children_functionMenu_id as functionMenu_id,pf.id as parent_id,pf.url as parent_url,
	pf.name as parent_name,cf.url as menu_url,cf.name as menu_name from role_clinic_functionMenu rcf
	left join clinic_children_functionMenu ccf on ccf.id = rcf.clinic_children_functionMenu_id
	left join children_functionMenu cf on cf.id = ccf.children_functionMenu_id
	left join parent_functionMenu pf on pf.id = cf.parent_functionMenu_id
	where rcf.role_id=$1`
	rows, err2 := model.DB.Queryx(selectSQL, roleID)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err2})
		return
	}
	role := FormatSQLRowToMap(arows)
	clinicFunctionMenu := FormatSQLRowsToMapArray(rows)
	var menus []Funtionmenus
	for _, v := range clinicFunctionMenu {
		childenURL := v["menu_url"]
		childenName := v["menu_name"]
		functionmenuID := v["functionmenu_id"]
		parentID := v["parent_id"]
		parentURL := v["parent_url"]
		parentName := v["parent_name"]
		has := false
		for k, menu := range menus {
			parentMenuID := menu.ParentID
			childrenMenus := menu.ChildrensMenus
			if strconv.FormatInt(parentID.(int64), 10) == parentMenuID {
				childrens := Menu{
					FunctionmenuID: strconv.FormatInt(functionmenuID.(int64), 10),
					MenuName:       childenName.(string),
					MenuURL:        childenURL.(string),
				}
				menus[k].ChildrensMenus = append(childrenMenus, childrens)
				has = true
			}
		}
		if !has {
			var childrens []Menu
			children := Menu{
				FunctionmenuID: strconv.FormatInt(functionmenuID.(int64), 10),
				MenuName:       childenName.(string),
				MenuURL:        childenURL.(string),
			}
			childrens = append(childrens, children)

			functionMenu := Funtionmenus{
				ParentID:       strconv.FormatInt(parentID.(int64), 10),
				ParentName:     parentName.(string),
				ParentURL:      parentURL.(string),
				ChildrensMenus: childrens,
			}
			menus = append(menus, functionMenu)
		}
	}
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
