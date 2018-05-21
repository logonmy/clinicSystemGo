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

	if manuFactoryName != "" {
		lrow := model.DB.QueryRowx("select id from material where name=$1 and manu_factory_name=$2 limit 1", name, manuFactoryName)
		if lrow == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
			return
		}
		materialProject := FormatSQLRowToMap(lrow)
		_, lok := materialProject["id"]
		if lok {
			ctx.JSON(iris.Map{"code": "1", "msg": "该材料已存在"})
			return
		}
	}

	materialSets := []string{"name", "unit_name"}
	materialValues := []string{"'" + name + "'", unitName}

	clinicMaterialSets := []string{"clinic_id", "ret_price"}
	clinicMaterialValues := []string{clinicID, retPrice}

	if enName != "" {
		materialSets = append(materialSets, "en_name")
		materialValues = append(materialValues, "'"+enName+"'")
	}
	if pyCode != "" {
		materialSets = append(materialSets, "py_code")
		materialValues = append(materialValues, "'"+pyCode+"'")
	}
	if idcCode != "" {
		materialSets = append(materialSets, "idc_code")
		materialValues = append(materialValues, "'"+idcCode+"'")
	}
	if remark != "" {
		materialSets = append(materialSets, "remark")
		materialValues = append(materialValues, "'"+remark+"'")
	}
	if manuFactoryName != "" {
		materialSets = append(materialSets, "manu_factory_name")
		materialValues = append(materialValues, "'"+manuFactoryName+"'")
	}
	if specification != "" {
		materialSets = append(materialSets, "specification")
		materialValues = append(materialValues, "'"+specification+"'")
	}

	if status != "" {
		clinicMaterialSets = append(clinicMaterialSets, "status")
		clinicMaterialValues = append(clinicMaterialValues, status)
	}
	if buyPrice != "" {
		clinicMaterialSets = append(clinicMaterialSets, "buy_price")
		clinicMaterialValues = append(clinicMaterialValues, buyPrice)
	}
	if isDiscount != "" {
		clinicMaterialSets = append(clinicMaterialSets, "is_discount")
		clinicMaterialValues = append(clinicMaterialValues, isDiscount)
	}
	if dayWarning != "" {
		clinicMaterialSets = append(clinicMaterialSets, "day_warning")
		clinicMaterialValues = append(clinicMaterialValues, dayWarning)
	}
	if stockWarning != "" {
		clinicMaterialSets = append(clinicMaterialSets, "stock_warning")
		clinicMaterialValues = append(clinicMaterialValues, stockWarning)
	}

	materialSetstr := strings.Join(materialSets, ",")
	materialValuestr := strings.Join(materialValues, ",")

	materialInsertSQL := "insert into material (" + materialSetstr + ") values (" + materialValuestr + ") RETURNING id;"
	fmt.Println("materialInsertSQL==", materialInsertSQL)

	tx, err := model.DB.Begin()
	var materialID string
	err = tx.QueryRow(materialInsertSQL).Scan(&materialID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	fmt.Println("materialID====", materialID)

	clinicMaterialSets = append(clinicMaterialSets, "material_id")
	clinicMaterialValues = append(clinicMaterialValues, materialID)

	clinicMaterialSetStr := strings.Join(clinicMaterialSets, ",")
	clinicMaterialValueStr := strings.Join(clinicMaterialValues, ",")

	clinicMaterialInsertSQL := "insert into clinic_material (" + clinicMaterialSetStr + ") values (" + clinicMaterialValueStr + ")"
	fmt.Println("clinicMaterialInsertSQL==", clinicMaterialInsertSQL)

	_, err2 := tx.Exec(clinicMaterialInsertSQL)
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

	ctx.JSON(iris.Map{"code": "200", "data": materialID})

}

// MaterialUpdate 更新材料项目
func MaterialUpdate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	clinicMaterialID := ctx.PostValue("clinic_material_id")
	materialID := ctx.PostValue("material_id")

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

	if clinicID == "" || name == "" || clinicMaterialID == "" || retPrice == "" || materialID == "" {
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

	crow := model.DB.QueryRowx("select id,material_id from clinic_material where id=$1 limit 1", clinicMaterialID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicMaterialProject := FormatSQLRowToMap(crow)
	_, rok := clinicMaterialProject["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所材料项目数据错误"})
		return
	}
	smaterialID := strconv.FormatInt(clinicMaterialProject["material_id"].(int64), 10)
	fmt.Println("smaterialID====", smaterialID)

	if smaterialID != materialID {
		ctx.JSON(iris.Map{"code": "1", "msg": "材料项目数据id不匹配"})
		return
	}

	lrow := model.DB.QueryRowx("select id from material where name=$1 and id!=$2 limit 1", name, materialID)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	laboratoryItem := FormatSQLRowToMap(lrow)
	_, lok := laboratoryItem["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "1", "msg": "材料项目名称已存在"})
		return
	}

	materialSets := []string{"name='" + name + "'"}
	clinicMaterialSets := []string{"ret_price=" + retPrice}

	if enName != "" {
		materialSets = append(materialSets, "en_name='"+enName+"'")
	}
	if pyCode != "" {
		materialSets = append(materialSets, "py_code='"+pyCode+"'")
	}
	if unitName != "" {
		materialSets = append(materialSets, "unit_name='"+unitName+"'")
	}
	if idcCode != "" {
		materialSets = append(materialSets, "idc_code='"+idcCode+"'")
	}
	if remark != "" {
		materialSets = append(materialSets, "remark='"+remark+"'")
	}
	if manuFactoryName != "" {
		materialSets = append(materialSets, "manu_factory_name='"+manuFactoryName+"'")
	}
	if specification != "" {
		materialSets = append(materialSets, "specification='"+specification+"'")
	}

	if status != "" {
		clinicMaterialSets = append(clinicMaterialSets, "status="+status)
	}
	if isDiscount != "" {
		clinicMaterialSets = append(clinicMaterialSets, "is_discount="+isDiscount)
	}
	if buyPrice != "" {
		clinicMaterialSets = append(clinicMaterialSets, "buy_price="+buyPrice)
	}
	if dayWarning != "" {
		clinicMaterialSets = append(clinicMaterialSets, "day_warning="+dayWarning)
	}
	if stockWarning != "" {
		clinicMaterialSets = append(clinicMaterialSets, "stock_warning="+stockWarning)
	}

	materialSets = append(materialSets, "updated_time=LOCALTIMESTAMP")
	materialSetstr := strings.Join(materialSets, ",")

	materialUpdateSQL := "update material set " + materialSetstr + " where id=$1"
	fmt.Println("materialUpdateSQL==", materialUpdateSQL)

	tx, err := model.DB.Begin()
	_, err = tx.Exec(materialUpdateSQL, materialID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}

	clinicMaterialSets = append(clinicMaterialSets, "updated_time=LOCALTIMESTAMP")
	clinicMaterialSetStr := strings.Join(clinicMaterialSets, ",")

	clinicMaterialUpdateSQL := "update clinic_material set " + clinicMaterialSetStr + " where id=$1"
	fmt.Println("clinicMaterialUpdateSQL==", clinicMaterialUpdateSQL)

	_, err2 := tx.Exec(clinicMaterialUpdateSQL, clinicMaterialID)
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
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)
	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据错误"})
		return
	}

	crow := model.DB.QueryRowx("select id from clinic_material where id=$1 limit 1", clinicMaterialID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicMaterialProject := FormatSQLRowToMap(crow)
	_, rok := clinicMaterialProject["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据错误"})
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
		ctx.JSON(iris.Map{"code": "1", "msg": "查询失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)

	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "所在诊所不存在"})
		return
	}

	countSQL := `select count(cm.id) as total from clinic_material cm
		left join material m on cm.material_id = m.id
		where cm.clinic_id=$1`
	selectSQL := `select cm.material_id,cm.id as clinic_material_id,m.name,m.unit_name,m.py_code,m.remark,m.idc_code,m.manu_factory_name,m.specification,
		m.en_name,cm.is_discount,cm.ret_price,cm.status,cm.buy_price,cm.day_warning,cm.stock_warning,sum(ms.stock_amount) as stock_amount
		from clinic_material cm
		left join material m on cm.material_id = m.id
		left join material_stock ms on ms.clinic_material_id = cm.id
		where cm.clinic_id=$1
		group by cm.material_id,cm.id,m.name,m.unit_name,m.py_code,m.remark,m.idc_code,m.manu_factory_name,m.specification,
		m.en_name,cm.is_discount,cm.ret_price,cm.status,cm.buy_price,cm.day_warning,cm.stock_warning`

	if keyword != "" {
		countSQL += " and m.name ~'" + keyword + "'"
		selectSQL += " and m.name ~'" + keyword + "'"
	}
	if status != "" {
		countSQL += " and cm.status=" + status
		selectSQL += " and cm.status=" + status
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

//MaterialDetail 材料项目详情
func MaterialDetail(ctx iris.Context) {
	clinicMaterialID := ctx.PostValue("clinic_material_id")

	if clinicMaterialID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select cm.material_id,cm.id as clinic_material_id,m.name,m.unit_name,m.py_code,m.remark,m.idc_code,
		m.manu_factory_name,m.specification,m.en_name,cm.is_discount,cm.ret_price,cm.status,cm.buy_price,cm.day_warning,cm.stock_warning,cm.stock_amount
		from clinic_material cm
		left join material m on cm.material_id = m.id
		where cm.id=$1`

	fmt.Println("selectSQL===", selectSQL)

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL, clinicMaterialID)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})
}
