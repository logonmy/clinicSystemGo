package controller

import (
	"clinicSystemGo/model"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/kataras/iris"
)

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
	err = tx.QueryRow("insert into admin(name, title, phone, username, password) values ($1, $2, $3, $4, $5) RETURNING id", name, title, phone, username, passwordMd5).Scan(&adminID)
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
	if adminID == "" || name == "" || title == "" || phone == "" || username == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

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

	updateSQL := `update admin set name=$1, title=$2, phone=$3, username=$4, is_clinic_admin=true, updated_time=LOCALTIMESTAMP where id=$5`

	if password != "" {
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(password))
		passwordMd5 := hex.EncodeToString(md5Ctx.Sum(nil))
		updateSQL = `update admin set name=$1, title=$2, phone=$3, username=$4, password=$5, is_clinic_admin=true, updated_time=LOCALTIMESTAMP where id=$6`

		_, erru := tx.Exec(updateSQL, name, title, phone, username, passwordMd5, adminID)
		if erru != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": erru.Error()})
			return
		}
	} else {
		_, erra := tx.Exec(updateSQL, name, title, phone, username, adminID)
		if erra != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": erra.Error()})
			return
		}
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

// AdminOnOff 启用和停用
func AdminOnOff(ctx iris.Context) {
	adminID := ctx.PostValue("admin_id")
	status := ctx.PostValue("status")
	if adminID == "" || status == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from admin where id=$1 limit 1", adminID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	admin := FormatSQLRowToMap(row)
	_, ok := admin["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "账号不存在"})
		return
	}

	_, err1 := model.DB.Exec("update admin set status=$1 where id=$2", status, adminID)
	if err1 != nil {
		fmt.Println(" err1====", err1)
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//AdminList 平台账号列表
func AdminList(ctx iris.Context) {
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	keyword := ctx.PostValue("keyword")
	startDateStr := ctx.PostValue("start_date")
	endDateStr := ctx.PostValue("end_date")
	status := ctx.PostValue("status")

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

	if startDateStr != "" && endDateStr == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请选择结束日期"})
		return
	}
	if startDateStr == "" && endDateStr != "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请选择开始日期"})
		return
	}

	countSQL := `select count(id) as total from admin where id>0`
	rowSQL := `select id as admin_id,created_time,name,username,phone,status,title from admin where id>0`

	if keyword != "" {
		countSQL += ` and name ~*:keyword`
		rowSQL += ` and name ~*:keyword`
	}

	if startDateStr != "" && endDateStr != "" {
		if startDateStr > endDateStr {
			ctx.JSON(iris.Map{"code": "-1", "msg": "开始日期必须大于结束日期"})
			return
		}

		endDate, erre := time.Parse("2006-01-02", endDateStr)
		if erre != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "end_date 必须为 YYYY-MM-DD 的 有效日期格式"})
			return
		}
		endDateStr = endDate.AddDate(0, 0, 1).Format("2006-01-02")

		countSQL += " and created_time between :start_date and :end_date"
		rowSQL += " and created_time between :start_date and :end_date"
	}

	if status != "" {
		countSQL += " and status =:status"
		rowSQL += " and status =:status"
	}

	var queryOptions = map[string]interface{}{
		"keyword":    ToNullString(keyword),
		"offset":     ToNullInt64(offset),
		"limit":      ToNullInt64(limit),
		"start_date": ToNullString(startDateStr),
		"end_date":   ToNullString(endDateStr),
		"status":     ToNullString(status),
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

	arows := model.DB.QueryRowx("select id as admin_id,created_time,name,username,phone,status,title from admin where id=$1", adminID)
	if arows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询结果不存在"})
		return
	}
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
	from admin_function_menu af
	left join function_menu fm on fm.id = af.function_menu_id
	where af.admin_id=$1 order by fm.level asc,fm.weight asc`
	rows, err2 := model.DB.Queryx(selectSQL, adminID)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err2})
		return
	}
	admin := FormatSQLRowToMap(arows)

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

	// adminFunctionMenu := FormatMenu(funtionmenu)
	adminFunctionMenu := FormatSQLRowsToMapArray(rows)

	admin["funtionMenus"] = adminFunctionMenu
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": admin})
}

// AdminLogin 登录
func AdminLogin(ctx iris.Context) {
	IP := ctx.RemoteAddr()
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	if username != "" && password != "" {
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(password))
		passwordMd5 := hex.EncodeToString(md5Ctx.Sum(nil))
		row := model.DB.QueryRowx(`select id, phone, name, title, username, is_clinic_admin
			from admin
			where username = $1 and password = $2 and status=true`, username, passwordMd5)
		if row == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "用户名或密码错误或账号未启用"})
			return
		}
		result := FormatSQLRowToMap(row)
		if _, ok := result["id"]; ok {
			adminID := result["id"]
			_ = model.DB.MustExec("INSERT INTO admin_login_record (admin_id, ip) VALUES ($1, $2) RETURNING id", adminID, IP)
			countRow := model.DB.QueryRowx("select count(*) as count from admin_login_record where admin_id = $1", adminID)
			count := FormatSQLRowToMap(countRow)

			// token := jwt.New(jwt.SigningMethodHS256)
			// claims := make(jwt.MapClaims)
			// claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
			// claims["iat"] = time.Now().Unix()
			// claims["admin_id"] = adminID
			// token.Claims = claims
			// tokenString, err := token.SignedString([]byte(secretKey))

			// if err != nil {
			// 	fmt.Println("Error while signing the token", err)
			// 	ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			// 	return
			// }

			ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": result, "login_times": count["count"]})
			// ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": tokenString})
			return
		}
		ctx.JSON(iris.Map{"code": "-1", "msg": "用户名或密码错误或账号未启用"})
		return
	}
	ctx.JSON(iris.Map{"code": "-1", "msg": "请输入用户名或密码"})
}

//MenubarUnsetByAdminID 获取平台未开通的菜单项
func MenubarUnsetByAdminID(ctx iris.Context) {
	adminID := ctx.PostValue("admin_id")

	if adminID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

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
		left join admin_function_menu afm on afm.function_menu_id = fm.id and afm.admin_id = $1 and afm.status=true
		where afm.function_menu_id IS NULL and fm.ascription='02' order by fm.level asc,fm.weight asc`
	rows, err := model.DB.Queryx(selectSQL, adminID)
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

//MenuGetByAdminID 获取平台开通菜单项
func MenuGetByAdminID(ctx iris.Context) {
	adminID := ctx.PostValue("admin_id")
	if adminID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx(`select * from admin where id = $1`, adminID)

	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "人员不存在"})
		return
	}

	adminMap := FormatSQLRowToMap(row)

	isClinicAdmin, ok := adminMap["is_clinic_admin"]

	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "人员不存在"})
		return
	}

	if !isClinicAdmin.(bool) {
		selectSQL := `select 
		afm.id as admin_function_menu_id,
		afm.function_menu_id as function_menu_id,
		fm.name as menu_name,
		fm.url as menu_url,
		fm.level,
		fm.weight,
		fm.ascription,
		fm.status,
		fm.icon,
		fm.parent_function_menu_id
		from admin_function_menu afm
		left join function_menu fm on fm.id = afm.function_menu_id
		where afm.admin_id=$1 and afm.status=true order by fm.level asc,fm.weight asc`

		rows, err2 := model.DB.Queryx(selectSQL, adminID)
		if err2 != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
			return
		}

		result := FormatSQLRowsToMapArray(rows)

		ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": result})
	} else {
		selectSQL := `select 
		id as function_menu_id,
		name as menu_name,
		url as menu_url,
		level,
		weight,
		ascription,
		status,
		icon,
		parent_function_menu_id
		from function_menu
		where ascription = '02' order by level asc,weight asc`

		rows, err2 := model.DB.Queryx(selectSQL)
		if err2 != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
			return
		}

		result := FormatSQLRowsToMapArray(rows)

		ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": result})
	}
}
