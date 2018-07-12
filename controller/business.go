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
	level := ctx.PostValue("level")
	icon := ctx.PostValue("icon")
	name := ctx.PostValue("name")
	weight := ctx.PostValue("weight")
	ascription := ctx.PostValue("ascription")
	parentFunctionMenuID := ctx.PostValue("parent_function_menu_id")
	if url == "" || name == "" || ascription == "" || level == "" || weight == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	insertSQL := "insert into function_menu (level, url, ascription, name, icon, weight, parent_function_menu_id) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err := model.DB.Exec(insertSQL, level, url, ascription, name, ToNullString(icon), weight, ToNullInt64(parentFunctionMenuID))
	if err != nil {
		fmt.Println("err1 ===", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": nil})
}

//MenubarList 获取所有菜单项
func MenubarList(ctx iris.Context) {
	ascription := ctx.PostValue("ascription")
	if ascription == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	// selectSQL := `select p.id as parent_id,p.url as parent_url,c.url as menu_url,p.name as parent_name,c.name menu_name,c.id as function_menu_id from children_function_menu c
	// 	left join parent_function_menu p on p.id = c.parent_function_menu_id`
	// if ascription != "" {
	// 	selectSQL += " where p.ascription='" + ascription + "'"
	// }
	// rows, _ := model.DB.Queryx(selectSQL)
	// if rows == nil {
	// 	ctx.JSON(iris.Map{"code": "-1", "msg": "查询失败"})
	// 	return
	// }
	// parentFunctionMenu := FormatSQLRowsToMapArray(rows)
	// menus := FormatFuntionmenus(parentFunctionMenu)

	selectSQL := `select * from function_menu where ascription=$1`

	rows, _ := model.DB.Queryx(selectSQL, ascription)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询失败"})
		return
	}
	functionMenu := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": functionMenu})
}

//MenubarListByClinicID 获取诊所未开通的菜单项
func MenubarListByClinicID(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	selectSQL := `select 
		fm.id as function_menu_id,
		fm.name as menu_name,
		fm.url as menu_url,
		fm.level,
		fm.weight,
		fm.parent_function_menu_id
		from function_menu fm
		left join clinic_function_menu cfm on cfm.function_menu_id = fm.id and cfm.clinic_id = $1 and cfm.status=true
		where cfm.function_menu_id IS NULL order by fm.level asc,fm.weight asc`
	rows, _ := model.DB.Queryx(selectSQL, clinicID)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "查询失败"})
		return
	}
	parentFunctionMenu := FormatSQLRowsToMapArray(rows)
	// menus := FormatFuntionmenus(parentFunctionMenu)

	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": parentFunctionMenu})

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
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)
	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊所不存在"})
		return
	}

	tx, err := model.DB.Beginx()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	var sets []string
	// if len(results) > 0 {
	// 	for _, v := range results {
	// 		childrenFunctionMenuID := v["function_menu_id"]
	// 		crow := model.DB.QueryRowx("select id from children_function_menu where id=$1 limit 1", childrenFunctionMenuID)
	// 		if crow == nil {
	// 			ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
	// 			return
	// 		}
	// 		childrenFunctionMenu := FormatSQLRowToMap(crow)
	// 		_, cok := childrenFunctionMenu["id"]
	// 		if !cok {
	// 			ctx.JSON(iris.Map{"code": "-1", "msg": "菜单项不存在"})
	// 			return
	// 		}
	// 		sets = append(sets, childrenFunctionMenuID)
	// 	}
	// 	setStr := strings.Join(sets, ",")
	// 	insertSQL := `insert into clinic_children_function_menu (clinic_id, children_function_menu_id)
	// 	select ` + clinicID + `, id from children_function_menu
	// 	where id in (` + setStr +
	// 		`)and id not in (select children_function_menu_id from clinic_children_function_menu where clinic_id = ` + clinicID + `);`

	// 	_, errtx1 := tx.Exec(insertSQL)
	// 	if errtx1 != nil {
	// 		tx.Rollback()
	// 		ctx.JSON(iris.Map{"code": "-1", "msg": errtx1.Error()})
	// 		return
	// 	}

	// 	_, errtx2 := tx.Exec("update clinic_children_function_menu set status=false where clinic_id=$1", clinicID)
	// 	if errtx2 != nil {
	// 		tx.Rollback()
	// 		ctx.JSON(iris.Map{"code": "-1", "msg": errtx2.Error()})
	// 		return
	// 	}
	// 	updateSQL := "update clinic_children_function_menu set status=true where clinic_id=$1 and children_function_menu_id in (" + setStr + ")"
	// 	fmt.Println("updateSQL===", updateSQL)
	// 	_, errtx3 := tx.Exec(updateSQL, clinicID)
	// 	if errtx3 != nil {
	// 		tx.Rollback()
	// 		ctx.JSON(iris.Map{"code": "-1", "msg": errtx3.Error()})
	// 		return
	// 	}
	// }

	if len(results) > 0 {
		for _, v := range results {
			functionMenuID := v["function_menu_id"]
			crow := model.DB.QueryRowx("select id from function_menu where id=$1 limit 1", functionMenuID)
			if crow == nil {
				ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
				return
			}
			functionMenu := FormatSQLRowToMap(crow)
			_, cok := functionMenu["id"]
			if !cok {
				ctx.JSON(iris.Map{"code": "-1", "msg": "菜单项不存在"})
				return
			}
			sets = append(sets, functionMenuID)
		}
		setStr := strings.Join(sets, ",")
		insertSQL := `insert into clinic_function_menu (clinic_id, function_menu_id) 
		select ` + clinicID + `, id from function_menu 
		where id in (` + setStr +
			`)and id not in (select function_menu_id from clinic_function_menu where clinic_id = ` + clinicID + `);`

		_, errtx1 := tx.Exec(insertSQL)
		if errtx1 != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": errtx1.Error()})
			return
		}

		_, errtx2 := tx.Exec("update clinic_function_menu set status=false where clinic_id=$1", clinicID)
		if errtx2 != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": errtx2.Error()})
			return
		}
		updateSQL := "update clinic_function_menu set status=true where clinic_id=$1 and function_menu_id in (" + setStr + ")"
		fmt.Println("updateSQL===", updateSQL)
		_, errtx3 := tx.Exec(updateSQL, clinicID)
		if errtx3 != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": errtx3.Error()})
			return
		}
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
		ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
		return
	}
	admin := FormatSQLRowToMap(row)
	_, ok := admin["id"]
	if ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "账号名称已存在"})
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

		for _, v := range results {
			functionMenuID := v["function_menu_id"]
			crow := model.DB.QueryRowx("select id from function_menu where id=$1", functionMenuID)
			if crow == nil {
				ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
				return
			}
			functionMenu := FormatSQLRowToMap(crow)
			_, cok := functionMenu["id"]
			if !cok {
				ctx.JSON(iris.Map{"code": "-1", "msg": "菜单项不存在"})
				return
			}
			sql := "INSERT INTO admin_function_menu (function_menu_id, admin_id) VALUES ($1,$2)"
			_, err = tx.Exec(sql, ToNullInt64(functionMenuID), ToNullInt64(adminID))
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
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	admin := FormatSQLRowToMap(row)
	_, ok := admin["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改的账号不存在"})
		return
	}

	rrow := model.DB.QueryRowx("select id from admin where username=$1 and id!=$2 limit 1", username, adminID)
	if rrow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	radmin := FormatSQLRowToMap(rrow)
	_, rok := radmin["id"]
	if rok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "账号名称已存在"})
		return
	}

	tx, err := model.DB.Begin()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	_, erra := tx.Exec("update admin set name=$1, title=$2, phone=$3, username=$4, password=$5, is_clinic_admin=true, updated_time=LOCALTIMESTAMP where id=$6", name, title, phone, username, passwordMd5, adminID)
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

		sql1 := "delete from admin_function_menu WHERE admin_id=$1"
		_, errtx := tx.Exec(sql1, ToNullInt64(adminID))
		if errtx != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": errtx.Error()})
			return
		}

		for _, v := range results {
			functionMenuID := v["function_menu_id"]
			crow := model.DB.QueryRowx("select id from function_menu where id=$1 limit 1", functionMenuID)
			if crow == nil {
				ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
				return
			}
			functionMenu := FormatSQLRowToMap(crow)
			_, cok := functionMenu["id"]
			if !cok {
				ctx.JSON(iris.Map{"code": "-1", "msg": "菜单项不存在"})
				return
			}
			sql := "INSERT INTO admin_function_menu (function_menu_id, admin_id) VALUES ($1,$2)"
			fmt.Println(sql)
			_, err = tx.Exec(sql, ToNullInt64(functionMenuID), ToNullInt64(adminID))
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

	countSQL := `select count(id) as total from admin where id>0`
	rowSQL := `select id as admin_id,created_time,name,username,phone,status from admin where id>0`
	if keyword != "" {
		countSQL += ` and name ~*:keyword`
		rowSQL += ` and name ~*:keyword`
	}

	var queryOptions = map[string]interface{}{
		"keyword": ToNullString(keyword),
		"offset":  ToNullInt64(offset),
		"limit":   ToNullInt64(limit),
	}

	totalrow, err1 := model.DB.NamedQuery(countSQL, queryOptions)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	totals := FormatSQLRowsToMapArray(totalrow)
	pageInfo := totals[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.NamedQuery(rowSQL+" offset :offset limit :limit", queryOptions)
	results = FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})
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
	selectSQL := `select 
	fm.id as function_menu_id,
	fm.url as menu_url,
	fm.name as menu_name,
	fm.level,
	fm.weight,
	fm.parent_function_menu_id
	from admin_function_menu af
	left join function_menu fm on fm.id = af.function_menu_id
	where af.admin_id=$1 order by fm.level asc,fm.weight asc`
	rows, err2 := model.DB.Queryx(selectSQL, adminID)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err2})
		return
	}
	admin := FormatSQLRowToMap(arows)
	adminFunctionMenu := FormatSQLRowsToMapArray(rows)
	// menus := FormatFuntionmenus(adminFunctionMenu)
	admin["funtionMenus"] = adminFunctionMenu
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": admin})
}

//MenuGetByClinicID 获取诊所业务信息
func MenuGetByClinicID(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	// selectSQL := `select ccf.id as clinic_function_menu_id,ccf.children_function_menu_id as function_menu_id,pf.id as parent_id,pf.url as parent_url,
	// pf.name as parent_name,cf.url as menu_url,cf.name as menu_name from clinic_children_function_menu ccf
	// left join children_function_menu cf on cf.id = ccf.children_function_menu_id
	// left join parent_function_menu pf on pf.id = cf.parent_function_menu_id
	// where ccf.clinic_id=$1 and ccf.status=true`
	// rows, err2 := model.DB.Queryx(selectSQL, clinicID)
	// if err2 != nil {
	// 	ctx.JSON(iris.Map{"code": "-1", "msg": err2})
	// 	return
	// }
	// clinicFunctionMenu := FormatSQLRowsToMapArray(rows)
	// menus := FormatFuntionmenus(clinicFunctionMenu)

	selectSQL := `select 
	cfm.id as clinic_function_menu_id,
	cfm.function_menu_id as function_menu_id,
	fm.name as menu_name,
	fm.url as menu_url,
	fm.level,
	fm.weight,
	fm.parent_function_menu_id
	from clinic_function_menu cfm
	left join function_menu fm on fm.id = cfm.function_menu_id
	where cfm.clinic_id=$1 and cfm.status=true order by fm.level asc,fm.weight asc`
	rows, err2 := model.DB.Queryx(selectSQL, clinicID)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err2})
		return
	}
	clinicFunctionMenu := FormatSQLRowsToMapArray(rows)
	// menus := FormatFuntionmenus(clinicFunctionMenu)

	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": clinicFunctionMenu})
}
