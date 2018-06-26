package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

// MaterialCreate 创建材料缴费项目
func MaterialCreate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	pyCode := ctx.PostValue("py_code")
	idcCode := ctx.PostValue("idc_code")
	unitName := ctx.PostValue("unit_name")
	remark := ctx.PostValue("remark")
	manuFactoryName := ctx.PostValue("manu_factory_name")
	specification := ctx.PostValue("specification")
	retPrice := ctx.PostValue("ret_price")
	buyPrice := ctx.PostValue("buy_price")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")
	dayWarning := ctx.PostValue("day_warning")
	stockWarning := ctx.PostValue("stock_warning")

	if clinicID == "" || name == "" || retPrice == "" || unitName == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

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

	if manuFactoryName != "" {
		lrow := model.DB.QueryRowx("select id from clinic_material where name=$1 and manu_factory_name=$2 and clinic_id=$3 limit 1", name, manuFactoryName, clinicID)
		if lrow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
			return
		}
		materialProject := FormatSQLRowToMap(lrow)
		_, lok := materialProject["id"]
		if lok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "该材料已存在"})
			return
		}
	}

	clinicMaterialSets := []string{
		"clinic_id",
		"name",
		"en_name",
		"py_code",
		"idc_code",
		"manu_factory_name",
		"specification",
		"unit_name",
		"remark",
		"ret_price",
		"buy_price",
		"is_discount",
		"day_warning",
		"stock_warning",
		"status",
	}
	clinicMaterialSetstr := strings.Join(clinicMaterialSets, ",")
	clinicMaterialInsertSQL := "insert into clinic_material (" + clinicMaterialSetstr + ") values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)"

	_, err := model.DB.Exec(clinicMaterialInsertSQL,
		ToNullInt64(clinicID),
		ToNullString(name),
		ToNullString(enName),
		ToNullString(pyCode),
		ToNullString(idcCode),
		ToNullString(manuFactoryName),
		ToNullString(specification),
		ToNullString(unitName),
		ToNullString(remark),
		ToNullInt64(retPrice),
		ToNullInt64(buyPrice),
		ToNullBool(isDiscount),
		ToNullInt64(dayWarning),
		ToNullInt64(stockWarning),
		ToNullBool(status),
	)
	if err != nil {
		fmt.Println(" err====", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// MaterialUpdate 更新材料项目
func MaterialUpdate(ctx iris.Context) {
	clinicMaterialID := ctx.PostValue("clinic_material_id")
	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	pyCode := ctx.PostValue("py_code")
	idcCode := ctx.PostValue("idc_code")
	unitName := ctx.PostValue("unit_name")
	remark := ctx.PostValue("remark")
	manuFactoryName := ctx.PostValue("manu_factory_name")
	specification := ctx.PostValue("specification")
	retPrice := ctx.PostValue("ret_price")
	buyPrice := ctx.PostValue("buy_price")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")
	dayWarning := ctx.PostValue("day_warning")
	stockWarning := ctx.PostValue("stock_warning")
	fmt.Println("isDiscount====", isDiscount)
	fmt.Println("dayWarning====", dayWarning)
	fmt.Println("stockWarning====", stockWarning)
	fmt.Println("ToNullInt64(dayWarning)====", ToNullInt64(dayWarning))
	fmt.Println("ToNullInt64(stockWarning)====", ToNullInt64(stockWarning))

	if name == "" || clinicMaterialID == "" || retPrice == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	crow := model.DB.QueryRowx("select id,clinic_id from clinic_material where id=$1 limit 1", clinicMaterialID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	clinicMaterial := FormatSQLRowToMap(crow)
	_, rok := clinicMaterial["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊所材料项目数据错误"})
		return
	}
	clinicID := clinicMaterial["clinic_id"]

	lrow := model.DB.QueryRowx("select id from clinic_material where name=$1 and id!=$2 and manu_factory_name=$3 and clinic_id=$4 limit 1", name, clinicMaterialID, manuFactoryName, clinicID)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	clinicMaterialu := FormatSQLRowToMap(lrow)
	_, lok := clinicMaterialu["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "材料项目名称已存在"})
		return
	}

	clinicMaterialUpdateSQL := `update clinic_material set 
		name=$1,
		en_name=$2,
		py_code=$3,
		idc_code=$4,
		manu_factory_name=$5,
		specification=$6,
		unit_name=$7,
		remark=$8,
		ret_price=$9,
		buy_price=$10,
		is_discount=$11
		day_warning=$12,
		stock_warning=$13,
		status=$14,
		where id=$15`

	_, err2 := model.DB.Exec(clinicMaterialUpdateSQL,
		ToNullString(name),
		ToNullString(enName),
		ToNullString(pyCode),
		ToNullString(idcCode),
		ToNullString(manuFactoryName),
		ToNullString(specification),
		ToNullString(unitName),
		ToNullString(remark),
		ToNullInt64(retPrice),
		ToNullInt64(buyPrice),
		ToNullBool(isDiscount),
		ToNullInt64(dayWarning),
		ToNullInt64(stockWarning),
		ToNullBool(status),
		ToNullInt64(clinicMaterialID),
	)
	if err2 != nil {
		fmt.Println(" err2====", err2)
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// MaterialOnOff 启用和停用
func MaterialOnOff(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	clinicMaterialID := ctx.PostValue("clinic_material_id")
	status := ctx.PostValue("status")
	if clinicID == "" || clinicMaterialID == "" || status == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
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
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊所数据错误"})
		return
	}

	crow := model.DB.QueryRowx("select id from clinic_material where id=$1 limit 1", clinicMaterialID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	clinicMaterialProject := FormatSQLRowToMap(crow)
	_, rok := clinicMaterialProject["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊所数据错误"})
		return
	}

	_, err1 := model.DB.Exec("update clinic_material set status=$1 where id=$2", status, clinicMaterialID)
	if err1 != nil {
		fmt.Println(" err1====", err1)
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// MaterialList 材料缴费项目列表
func MaterialList(ctx iris.Context) {
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
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)

	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "所在诊所不存在"})
		return
	}

	countSQL := `select count(id) as total from clinic_material where clinic_id=:clinic_id`
	selectSQL := `select cm.id as clinic_material_id,cm.name, cm.unit_name,cm.py_code,cm.remark,cm.idc_code,cm.manu_factory_name,cm.specification,
		cm.en_name,cm.is_discount,cm.ret_price,cm.status,cm.buy_price,cm.day_warning,cm.discount_price,cm.stock_warning,sum(ms.stock_amount) as stock_amount
		from clinic_material cm
		left join material_stock ms on ms.clinic_material_id = cm.id
		where cm.clinic_id=:clinic_id`

	if keyword != "" {
		countSQL += " and (name ~*:keyword or en_name ~*:keyword or py_code ~*:keyword) "
		selectSQL += " and (name ~*:keyword or en_name ~*:keyword or py_code ~*:keyword) "
	}
	if status != "" {
		countSQL += " and status=:status"
		selectSQL += " and cm.status=:status"
	}

	selectSQL += ` group by cm.id,cm.name,cm.unit_name,cm.py_code,cm.remark,cm.idc_code,cm.manu_factory_name,cm.specification,
	cm.en_name,cm.is_discount,cm.ret_price,cm.status,cm.buy_price,cm.discount_price,cm.day_warning,cm.stock_warning`

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
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, err1 := model.DB.NamedQuery(selectSQL+" offset :offset limit :limit", queryOption)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})

}

//MaterialDetail 材料项目详情
func MaterialDetail(ctx iris.Context) {
	clinicMaterialID := ctx.PostValue("clinic_material_id")

	if clinicMaterialID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select * from clinic_material where id=$1`

	fmt.Println("selectSQL===", selectSQL)

	rows := model.DB.QueryRowx(selectSQL, clinicMaterialID)
	results := FormatSQLRowToMap(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})
}
