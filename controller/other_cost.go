package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

// OtherCostCreate 创建其它费用缴费项目
func OtherCostCreate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	pyCode := ctx.PostValue("py_code")
	unitName := ctx.PostValue("unit_name")
	remark := ctx.PostValue("remark")
	price := ctx.PostValue("price")
	cost := ctx.PostValue("cost")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")

	if clinicID == "" || name == "" || price == "" || unitName == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from clinic where id=$1 limit 1", clinicID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)
	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据错误"})
		return
	}

	lrow := model.DB.QueryRowx("select id from clinic_other_cost where name=$1 and clinic_id=$2 limit 1", name, clinicID)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
		return
	}
	otherCost := FormatSQLRowToMap(lrow)
	_, lok := otherCost["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "1", "msg": "其它费用名称已存在"})
		return
	}

	clinicOtherCostSets := []string{
		"clinic_id",
		"name",
		"en_name",
		"py_code",
		"unit_name",
		"remark",
		"cost",
		"price",
		"status",
		"is_discount"}
	clinicOtherCostSetstr := strings.Join(clinicOtherCostSets, ",")
	clinicOtherCostInsertSQL := "insert into clinic_other_cost (" + clinicOtherCostSetstr + ") values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"

	_, err := model.DB.Exec(clinicOtherCostInsertSQL,
		ToNullInt64(clinicID),
		ToNullString(name),
		ToNullString(enName),
		ToNullString(pyCode),
		ToNullString(unitName),
		ToNullString(remark),
		ToNullInt64(cost),
		ToNullInt64(price),
		ToNullBool(status),
		ToNullBool(isDiscount),
	)
	if err != nil {
		fmt.Println(" err====", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// OtherCostUpdate 更新其它费用项目
func OtherCostUpdate(ctx iris.Context) {
	clinicOtherCostID := ctx.PostValue("clinic_other_cost_id")
	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	pyCode := ctx.PostValue("py_code")
	unitName := ctx.PostValue("unit_name")
	remark := ctx.PostValue("remark")
	price := ctx.PostValue("price")
	cost := ctx.PostValue("cost")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")

	if name == "" || clinicOtherCostID == "" || price == "" || unitName == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	crow := model.DB.QueryRowx("select id,clinic_id from clinic_other_cost where id=$1 limit 1", clinicOtherCostID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicOtherCost := FormatSQLRowToMap(crow)
	_, rok := clinicOtherCost["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所其它费用项目数据错误"})
		return
	}
	clinicID := clinicOtherCost["clinic_id"]

	lrow := model.DB.QueryRowx("select id from clinic_other_cost where name=$1 and id!=$2 and clinic_id=$3 limit 1", name, clinicOtherCostID, clinicID)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicOtherCostu := FormatSQLRowToMap(lrow)
	_, lok := clinicOtherCostu["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "1", "msg": "其它费用项目名称已存在"})
		return
	}

	clinicOtherCostUpdateSQL := `update clinic_other_cost set 
		name=$1,
		en_name=$2,
		py_code=$3,
		unit_name=$4,
		remark=$5,
		cost=$6,
		price=$7,
		status=$8,
		is_discount=$9
		where id=$10`

	_, err2 := model.DB.Exec(clinicOtherCostUpdateSQL,
		ToNullString(name),
		ToNullString(enName),
		ToNullString(pyCode),
		ToNullString(unitName),
		ToNullString(remark),
		ToNullInt64(cost),
		ToNullInt64(price),
		ToNullBool(status),
		ToNullBool(isDiscount),
		ToNullInt64(clinicOtherCostID),
	)
	if err2 != nil {
		fmt.Println(" err2====", err2)
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// OtherCostOnOff 启用和停用
func OtherCostOnOff(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	clinicOtherCostID := ctx.PostValue("clinic_other_cost_id")
	status := ctx.PostValue("status")
	if clinicID == "" || clinicOtherCostID == "" || status == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
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
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据错误"})
		return
	}

	crow := model.DB.QueryRowx("select id,clinic_id from clinic_other_cost where id=$1 limit 1", clinicOtherCostID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicOtherCostProject := FormatSQLRowToMap(crow)
	_, rok := clinicOtherCostProject["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据错误"})
		return
	}

	if clinicID != strconv.FormatInt(clinicOtherCostProject["clinic_id"].(int64), 10) {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据不匹配"})
		return
	}
	_, err1 := model.DB.Exec("update clinic_other_cost set status=$1 where id=$2", status, clinicOtherCostID)
	if err1 != nil {
		fmt.Println(" err1====", err1)
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// OtherCostList 其它费用缴费项目列表
func OtherCostList(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	keyword := ctx.PostValue("keyword")
	status := ctx.PostValue("status")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
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

	row := model.DB.QueryRowx("select id from clinic where id=$1 limit 1", clinicID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "查询失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)

	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "所在诊所不存在"})
		return
	}

	countSQL := `select count(id) as total from clinic_other_cost where clinic_id=:clinic_id`
	selectSQL := `select id as clinic_other_cost_id,name,py_code,remark,
		en_name,is_discount,price,status,cost,unit_name
		from clinic_other_cost coc
		where clinic_id=:clinic_id`

	if keyword != "" {
		countSQL += " and name ~:keyword"
		selectSQL += " and name ~:keyword"
	}
	if status != "" {
		countSQL += " and status=:status"
		selectSQL += " and status=:status"
	}

	var queryOption = map[string]interface{}{
		"clinic_id": ToNullInt64(clinicID),
		"keyword":   ToNullString(keyword),
		"status":    ToNullBool(status),
		"offset":    ToNullInt64(offset),
		"limit":     ToNullInt64(limit),
	}

	fmt.Println("countSQL===", countSQL)
	fmt.Println("selectSQL===", selectSQL)
	total, err := model.DB.NamedQuery(countSQL, queryOption)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL+" offset :offset limit :limit", queryOption)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})

}

//OtherCostDetail 其它费用项目详情
func OtherCostDetail(ctx iris.Context) {
	clinicOtherCostID := ctx.PostValue("clinic_other_cost_id")

	if clinicOtherCostID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select id as clinic_other_cost_id,name,py_code,remark,
		en_name,is_discount,price,status,cost,unit_name
		from clinic_other_cost coc
		left join other_cost oc on other_cost_id = id
		where id=$1`

	fmt.Println("selectSQL===", selectSQL)

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL, clinicOtherCostID)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})
}
