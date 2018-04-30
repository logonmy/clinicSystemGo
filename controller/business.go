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

//MenubarCreate 添加功能菜单栏
func MenubarCreate(ctx iris.Context) {
	url := ctx.PostValue("url")
	name := ctx.PostValue("name")
	ascription := ctx.PostValue("ascription")
	parentFunctionMenuID := ctx.PostValue("parent_functionMenu_id")
	if url == "" || name == "" || ascription == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	var resultDataID int
	insertSQL := "insert into parent_functionMenu (url, ascription, name) VALUES ($1, $2, $3) RETURNING id"
	if parentFunctionMenuID != "" {
		row := model.DB.QueryRowx("select id from parent_functionMenu where id=$1", parentFunctionMenuID)
		if row == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "创建失败"})
			return
		}
		parentFunctionMenu := FormatSQLRowToMap(row)
		_, ok := parentFunctionMenu["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "父级菜单ID不存在"})
			return
		}
		insertSQL = "insert into children_functionMenu (url, name, parent_functionMenu_id) VALUES ($1, $2, $3) RETURNING id"
		err := model.DB.QueryRow(insertSQL, url, name, parentFunctionMenuID).Scan(&resultDataID)
		if err != nil {
			fmt.Println("err1 ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
		ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": resultDataID})
		return
	}
	err := model.DB.QueryRow(insertSQL, url, ascription, name).Scan(&resultDataID)
	if err != nil {
		fmt.Println("err1 ===", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": resultDataID})
}

//MenubarList 获取所有菜单项
func MenubarList(ctx iris.Context) {
	ascription := ctx.PostValue("ascription")
	selectSQL := `select p.id as parent_id,p.url as parent_url,c.url as menu_url,p.name as parent_name,c.name menu_name,c.id as functionMenu_id from children_functionMenu c 
		left join parent_functionMenu p on p.id = c.parent_functionMenu_id`
	if ascription != "" {
		selectSQL += " where p.ascription='" + ascription + "'"
	}
	fmt.Println("====", selectSQL)
	rows, _ := model.DB.Queryx(selectSQL)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "查询失败"})
		return
	}
	parentFunctionMenu := FormatSQLRowsToMapArray(rows)

	var menus []Funtionmenus
	for _, v := range parentFunctionMenu {
		childenURL := v["menu_url"]
		childenName := v["menu_name"]
		functionmenuID := v["functionmenu_id"]
		parentID := v["parent_id"]
		parentURL := v["parent_url"]
		parentName := v["parent_name"]
		has := false
		for k, menu := range menus {
			// parentMenuID := menu["parent_id"]
			// childrenMenus := menu["childrens_menus"]
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
			// childrens := map[string]interface{}{
			// 	"functionmenu_id": functionmenuID,
			// 	"menu_name":       childenName,
			// 	"menu_url":        childenURL,
			// }
			// functionMenu := map[string]interface{}{
			// 	"parent_id":       parentID,
			// 	"parent_name":     parentName,
			// 	"parent_url":      parentURL,
			// 	"childrens_menus": childrens,
			// }
			// menus = append(menus, functionMenu)
		}
	}

	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": menus})
}

//BusinessAssign 诊所分配业务
func BusinessAssign(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	items := ctx.PostValue("items")
	if items == "" || clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	var results []map[string]string
	err := json.Unmarshal([]byte(items), &results)
	fmt.Println("===", results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	row := model.DB.QueryRowx("select id from clinic where id=$1 limit 1", clinicID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)
	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所不存在"})
		return
	}

	var sets []string

	for _, v := range results {
		s := "(" + v["functionMenu_id"] + "," + clinicID + ")"
		sets = append(sets, s)
	}

	setStr := strings.Join(sets, ",")

	sql1 := "DELETE FROM clinic_children_functionMenu WHERE clinic_id=" + clinicID
	sql := "INSERT INTO clinic_children_functionMenu( children_functionMenu_id, clinic_id ) VALUES " + setStr
	fmt.Println(sql1, sql)

	tx, err := model.DB.Beginx()

	if err != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
		return
	}

	_, errtx := tx.Exec(sql1)
	if errtx != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errtx.Error()})
		return
	}

	_, errtx2 := tx.Exec(sql)
	if errtx2 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errtx2.Error()})
		return
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//AdminCreate 平台账号添加
func AdminCreate(ctx iris.Context) {
	name := ctx.PostValue("name")
	title := ctx.PostValue("title")
	phone := ctx.PostValue("phone")
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	items := ctx.PostValue("items")
	if name == "" || title == "" || phone == "" || username == "" || password == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	passwordMd5 := hex.EncodeToString(md5Ctx.Sum(nil))
	var adminID string
	row := model.DB.QueryRowx("select id from admin username=$1 limit 1", username)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
		return
	}
	admin := FormatSQLRowToMap(row)
	_, ok := admin["id"]
	if ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "账号名称已存在"})
		return
	}
	tx, err := model.DB.Begin()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	err = tx.QueryRow("insert into admin(name, title, phone, username, password, is_clinic_admin) values ($1, $2, $3, $4, $5, true) RETURNING id", name, title, phone, username, passwordMd5).Scan(&adminID)
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

		var sets []string

		for _, v := range results {
			s := "(" + v["functionMenu_id"] + "," + adminID + ")"
			sets = append(sets, s)
		}

		setStr := strings.Join(sets, ",")

		sql := "INSERT INTO admin_functionMenu (children_functionMenu_id, admin_id) VALUES " + setStr
		fmt.Println(sql)
		_, err = tx.Exec(sql)
		if err != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
			return
		}
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": adminID})
}

//AdminUpdate 平台账号修改
func AdminUpdate(ctx iris.Context) {
	adminID := ctx.PostValue("admin_id")
	name := ctx.PostValue("name")
	title := ctx.PostValue("title")
	phone := ctx.PostValue("phone")
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	items := ctx.PostValue("items")
	if adminID == "" || name == "" || title == "" || phone == "" || username == "" || password == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	passwordMd5 := hex.EncodeToString(md5Ctx.Sum(nil))
	row := model.DB.QueryRowx("select id from admin where id=$1 limit 1", adminID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	admin := FormatSQLRowToMap(row)
	_, ok := admin["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改的账号不存在"})
		return
	}

	rrow := model.DB.QueryRowx("select id from admin where username=$1 and id!=$2 limit 1", username, adminID)
	if rrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	radmin := FormatSQLRowToMap(rrow)
	_, rok := radmin["id"]
	if rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "账号名称已存在"})
		return
	}

	tx, err := model.DB.Begin()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	_, erra := tx.Exec("update admin set name=$1, title=$2, phone=$3, username=$4, password=$5, is_clinic_admin=true where id=$6", name, title, phone, username, passwordMd5, adminID)
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

		var sets []string

		for _, v := range results {
			s := "(" + v["functionMenu_id"] + "," + adminID + ")"
			sets = append(sets, s)
		}

		setStr := strings.Join(sets, ",")

		sql1 := "delete from admin_functionMenu WHERE admin_id=" + adminID
		sql := "insert into admin_functionMenu (children_functionMenu_id, admin_id) VALUES " + setStr
		fmt.Println(sql)

		_, errtx := tx.Exec(sql1)
		if errtx != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "1", "msg": errtx.Error()})
			return
		}

		_, errtx2 := tx.Exec(sql)
		if errtx2 != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "1", "msg": errtx2.Error()})
			return
		}
	}
	err3 := tx.Commit()
	if err3 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": adminID})
}

//AdminList 平台账号列表
func AdminList(ctx iris.Context) {
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

	joinnSQL := `from admin where name ~$1 or username ~$1`

	total := model.DB.QueryRowx(`select count(id) as total `+joinnSQL, keyword)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rowSQL := `select id as admin_id,created_time,name,username,phone,status ` + joinnSQL + " offset $2 limit $3"

	rows, err1 := model.DB.Queryx(rowSQL, keyword, offset, limit)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})
}

//AdminGetByID 获取平台账号信息
func AdminGetByID(ctx iris.Context) {
	adminID := ctx.PostValue("admin_id")
	if adminID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	arows := model.DB.QueryRowx("select id as admin_id,created_time,name,username,phone,status from admin where id=$1", adminID)
	if arows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询结果不存在"})
		return
	}
	selectSQL := `select cf.id as functionMenu_id,pf.id as parent_id,pf.url as parent_url,pf.name as parent_name,cf.url as menu_url,cf.name as menu_name from admin_functionMenu af
	left join children_functionMenu cf on cf.id = af.children_functionMenu_id
	left join parent_functionMenu pf on pf.id = cf.parent_functionMenu_id
	where af.admin_id=$1`
	rows, err2 := model.DB.Queryx(selectSQL, adminID)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err2})
		return
	}
	admin := FormatSQLRowToMap(arows)
	adminFunctionMenu := FormatSQLRowsToMapArray(rows)
	var menus []Funtionmenus
	for _, v := range adminFunctionMenu {
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
	admin["funtionMenus"] = menus
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": admin})
}

//MenuGetByClinicID 获取诊所业务信息
func MenuGetByClinicID(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select ccf.id as functionMenu_id,pf.id as parent_id,pf.url as parent_url,
	pf.name as parent_name,cf.url as menu_url,cf.name as menu_name from clinic_children_functionMenu ccf
	left join children_functionMenu cf on cf.id = ccf.children_functionMenu_id
	left join parent_functionMenu pf on pf.id = cf.parent_functionMenu_id
	where ccf.clinic_id=$1`
	rows, err2 := model.DB.Queryx(selectSQL, clinicID)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err2})
		return
	}
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

	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": menus})
}