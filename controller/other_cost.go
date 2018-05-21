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

	lrow := model.DB.QueryRowx("select id from other_cost where name=$1 limit 1", name)
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

	otherCostSets := []string{"name", "unit_name"}
	otherCostValues := []string{"'" + name + "'", "'" + unitName + "'"}

	clinicOtherCostSets := []string{"clinic_id", "price"}
	clinicOtherCostValues := []string{clinicID, price}

	if enName != "" {
		otherCostSets = append(otherCostSets, "en_name")
		otherCostValues = append(otherCostValues, "'"+enName+"'")
	}
	if pyCode != "" {
		otherCostSets = append(otherCostSets, "py_code")
		otherCostValues = append(otherCostValues, "'"+pyCode+"'")
	}
	if remark != "" {
		otherCostSets = append(otherCostSets, "remark")
		otherCostValues = append(otherCostValues, "'"+remark+"'")
	}

	if status != "" {
		clinicOtherCostSets = append(clinicOtherCostSets, "status")
		clinicOtherCostValues = append(clinicOtherCostValues, status)
	}
	if cost != "" {
		clinicOtherCostSets = append(clinicOtherCostSets, "cost")
		clinicOtherCostValues = append(clinicOtherCostValues, cost)
	}
	if isDiscount != "" {
		clinicOtherCostSets = append(clinicOtherCostSets, "is_discount")
		clinicOtherCostValues = append(clinicOtherCostValues, isDiscount)
	}

	otherCostSetstr := strings.Join(otherCostSets, ",")
	otherCostValuestr := strings.Join(otherCostValues, ",")

	otherCostInsertSQL := "insert into other_cost (" + otherCostSetstr + ") values (" + otherCostValuestr + ") RETURNING id;"
	fmt.Println("otherCostInsertSQL==", otherCostInsertSQL)

	tx, err := model.DB.Begin()
	var otherCostID string
	err = tx.QueryRow(otherCostInsertSQL).Scan(&otherCostID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	fmt.Println("otherCostID====", otherCostID)

	clinicOtherCostSets = append(clinicOtherCostSets, "other_cost_id")
	clinicOtherCostValues = append(clinicOtherCostValues, otherCostID)

	clinicOtherCostSetstr := strings.Join(clinicOtherCostSets, ",")
	clinicOtherCostValuestr := strings.Join(clinicOtherCostValues, ",")

	clinicotherCostInsertSQL := "insert into clinic_other_cost (" + clinicOtherCostSetstr + ") values (" + clinicOtherCostValuestr + ")"
	fmt.Println("clinicotherCostInsertSQL==", clinicotherCostInsertSQL)

	_, err2 := tx.Exec(clinicotherCostInsertSQL)
	if err2 != nil {
		fmt.Println(" err2====", err2)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}

	err3 := tx.Commit()
	if err3 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": otherCostID})

}

// OtherCostUpdate 更新其它费用项目
func OtherCostUpdate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	clinicOtherCostID := ctx.PostValue("clinic_other_cost_id")
	otherCostID := ctx.PostValue("other_cost_id")

	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	pyCode := ctx.PostValue("py_code")
	unitName := ctx.PostValue("unit_name")
	remark := ctx.PostValue("remark")

	price := ctx.PostValue("price")
	cost := ctx.PostValue("cost")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")

	if clinicID == "" || name == "" || clinicOtherCostID == "" || price == "" || otherCostID == "" || unitName == "" {
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

	crow := model.DB.QueryRowx("select id,clinic_id,other_cost_id from clinic_other_cost where id=$1 limit 1", clinicOtherCostID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicOtherCostProject := FormatSQLRowToMap(crow)
	_, rok := clinicOtherCostProject["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所其它费用项目数据错误"})
		return
	}
	sotherCostID := strconv.FormatInt(clinicOtherCostProject["other_cost_id"].(int64), 10)
	fmt.Println("sotherCostID====", sotherCostID)

	if clinicID != strconv.FormatInt(clinicOtherCostProject["clinic_id"].(int64), 10) {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据不匹配"})
		return
	}

	if sotherCostID != otherCostID {
		ctx.JSON(iris.Map{"code": "1", "msg": "其它费用项目数据id不匹配"})
		return
	}

	lrow := model.DB.QueryRowx("select id from other_cost where name=$1 and id!=$2 limit 1", name, otherCostID)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	laboratoryItem := FormatSQLRowToMap(lrow)
	_, lok := laboratoryItem["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "1", "msg": "其它费用项目名称已存在"})
		return
	}

	otherCostSets := []string{"name='" + name + "'"}
	clinicOtherCostSets := []string{"price=" + price}

	if enName != "" {
		otherCostSets = append(otherCostSets, "en_name='"+enName+"'")
	}
	if pyCode != "" {
		otherCostSets = append(otherCostSets, "py_code='"+pyCode+"'")
	}
	if unitName != "" {
		otherCostSets = append(otherCostSets, "unit_name='"+unitName+"'")
	}
	if remark != "" {
		otherCostSets = append(otherCostSets, "remark='"+remark+"'")
	}

	if status != "" {
		clinicOtherCostSets = append(clinicOtherCostSets, "status="+status)
	}
	if isDiscount != "" {
		clinicOtherCostSets = append(clinicOtherCostSets, "is_discount="+isDiscount)
	}
	if cost != "" {
		clinicOtherCostSets = append(clinicOtherCostSets, "cost="+cost)
	}

	otherCostSets = append(otherCostSets, "updated_time=LOCALTIMESTAMP")
	otherCostSetstr := strings.Join(otherCostSets, ",")

	otherCostUpdateSQL := "update other_cost set " + otherCostSetstr + " where id=$1"
	fmt.Println("otherCostUpdateSQL==", otherCostUpdateSQL)

	tx, err := model.DB.Begin()
	_, err = tx.Exec(otherCostUpdateSQL, otherCostID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}

	clinicOtherCostSets = append(clinicOtherCostSets, "updated_time=LOCALTIMESTAMP")
	clinicOtherCostSetstr := strings.Join(clinicOtherCostSets, ",")

	clinicOtherCostUpdateSQL := "update clinic_other_cost set " + clinicOtherCostSetstr + " where id=$1"
	fmt.Println("clinicOtherCostUpdateSQL==", clinicOtherCostUpdateSQL)

	_, err2 := tx.Exec(clinicOtherCostUpdateSQL, clinicOtherCostID)
	if err2 != nil {
		fmt.Println(" err2====", err2)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}

	err3 := tx.Commit()
	if err3 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
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

	countSQL := `select count(coc.id) as total from clinic_other_cost coc
		left join other_cost oc on coc.other_cost_id = oc.id
		where coc.clinic_id=$1`
	selectSQL := `select coc.other_cost_id,coc.id as clinic_other_cost_id,oc.name,oc.py_code,oc.remark,
		oc.en_name,coc.is_discount,coc.price,coc.status,coc.cost,oc.unit_name
		from clinic_other_cost coc
		left join other_cost oc on coc.other_cost_id = oc.id
		where coc.clinic_id=$1`

	if keyword != "" {
		countSQL += " and oc.name ~'" + keyword + "'"
		selectSQL += " and oc.name ~'" + keyword + "'"
	}
	if status != "" {
		countSQL += " and coc.status=" + status
		selectSQL += " and coc.status=" + status
	}

	fmt.Println("countSQL===", countSQL)
	fmt.Println("selectSQL===", selectSQL)
	total := model.DB.QueryRowx(countSQL, clinicID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL+" offset $2 limit $3", clinicID, offset, limit)
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

	selectSQL := `select coc.other_cost_id,coc.id as clinic_other_cost_id,oc.name,oc.py_code,oc.remark,
		oc.en_name,coc.is_discount,coc.price,coc.status,coc.cost,oc.unit_name
		from clinic_other_cost coc
		left join other_cost oc on coc.other_cost_id = oc.id
		where coc.id=$1`

	fmt.Println("selectSQL===", selectSQL)

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL, clinicOtherCostID)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})
}
