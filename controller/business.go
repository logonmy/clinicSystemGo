package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
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

	// selectSQL := `select * from function_menu where ascription=$1 order by level asc,weight asc`

	// rows, _ := model.DB.Queryx(selectSQL, ascription)
	// if rows == nil {
	// 	ctx.JSON(iris.Map{"code": "-1", "msg": "查询失败"})
	// 	return
	// }
	// result := FormatSQLRowsToMapArray(rows)

	selectSQL := `select 
	id as function_menu_id,
	parent_function_menu_id,
	name as menu_name,
	url as menu_url,
	level,
	ascription,
	icon,
	status,
	weight
	from function_menu
	where ascription=$1 order by level asc,weight asc`

	rows, err := model.DB.Queryx(selectSQL, ascription)
	if rows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

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

	// result := FormatMenu(funtionmenu)
	result := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": result})
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
		fm.ascription,
		fm.status,
		fm.icon,
		fm.parent_function_menu_id
		from function_menu fm
		left join clinic_function_menu cfm on cfm.function_menu_id = fm.id and cfm.clinic_id = $1 and cfm.status=true
		where cfm.function_menu_id IS NULL and fm.ascription='01' order by fm.level asc,fm.weight asc`
	rows, err := model.DB.Queryx(selectSQL, clinicID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	// var funtionmenu []Funtionmenu
	// for rows.Next() {
	// 	var f Funtionmenu
	// 	err := rows.StructScan(&f)
	// 	if err != nil {
	// 		fmt.Println("err=====", err.Error())
	// 		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
	// 		return
	// 	}
	// 	formatUnsetMenu := FormatUnsetMenu(f)
	// 	funtionmenu = append(funtionmenu, formatUnsetMenu...)
	// }

	// result := FormatMenu(funtionmenu)
	result := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": result})
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

//MenuGetByClinicID 获取诊所业务信息
func MenuGetByClinicID(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select 
	cfm.id as clinic_function_menu_id,
	cfm.function_menu_id as function_menu_id,
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
	where cfm.clinic_id=$1 and cfm.status=true order by fm.level asc,fm.weight asc`

	rows, err2 := model.DB.Queryx(selectSQL, clinicID)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err2})
		return
	}

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

	// result := FormatMenu(funtionmenu)
	result := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": result})
}
