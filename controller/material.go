package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
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
	unitID := ctx.PostValue("unit_id")
	remark := ctx.PostValue("remark")
	manuFactory := ctx.PostValue("manu_factory")
	specification := ctx.PostValue("specification")

	price := ctx.PostValue("price")
	cost := ctx.PostValue("cost")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")
	effDay := ctx.PostValue("eff_day")
	stockWarning := ctx.PostValue("stock_warning")

	if clinicID == "" || name == "" || price == "" || unitID == "" {
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

	var storehouseID string
	errs := model.DB.QueryRow("select id from storehouse where clinic_id=$1 limit 1", clinicID).Scan(&storehouseID)
	if errs != nil {
		fmt.Println("errs ===", errs)
		ctx.JSON(iris.Map{"code": "1", "msg": errs.Error()})
		return
	}
	fmt.Println("storehouseID==", storehouseID)

	if manuFactory != "" {
		lrow := model.DB.QueryRowx("select id from material where name=$1 and manu_factory=$2 limit 1", name, manuFactory)
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

	materialSets := []string{"name", "unit_id"}
	materialValues := []string{"'" + name + "'", unitID}

	clinicMaterialSets := []string{"storehouse_id", "price"}
	clinicMaterialValues := []string{storehouseID, price}

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
	if manuFactory != "" {
		materialSets = append(materialSets, "manu_factory")
		materialValues = append(materialValues, "'"+manuFactory+"'")
	}
	if specification != "" {
		materialSets = append(materialSets, "specification")
		materialValues = append(materialValues, "'"+specification+"'")
	}

	if status != "" {
		clinicMaterialSets = append(clinicMaterialSets, "status")
		clinicMaterialValues = append(clinicMaterialValues, status)
	}
	if cost != "" {
		clinicMaterialSets = append(clinicMaterialSets, "cost")
		clinicMaterialValues = append(clinicMaterialValues, cost)
	}
	if isDiscount != "" {
		clinicMaterialSets = append(clinicMaterialSets, "is_discount")
		clinicMaterialValues = append(clinicMaterialValues, isDiscount)
	}
	if effDay != "" {
		clinicMaterialSets = append(clinicMaterialSets, "eff_day")
		clinicMaterialValues = append(clinicMaterialValues, effDay)
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

	clinicMaterialInsertSQL := "insert into material_stock (" + clinicMaterialSetStr + ") values (" + clinicMaterialValueStr + ")"
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
	materialStockID := ctx.PostValue("material_stock_id")
	materialID := ctx.PostValue("material_id")

	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	pyCode := ctx.PostValue("py_code")
	idcCode := ctx.PostValue("idc_code")
	unitID := ctx.PostValue("unit_id")
	remark := ctx.PostValue("remark")
	manuFactory := ctx.PostValue("manu_factory")
	specification := ctx.PostValue("specification")

	price := ctx.PostValue("price")
	cost := ctx.PostValue("cost")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")
	effDay := ctx.PostValue("eff_day")
	stockWarning := ctx.PostValue("stock_warning")

	if clinicID == "" || name == "" || materialStockID == "" || price == "" || materialID == "" {
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

	crow := model.DB.QueryRowx("select id,material_id from material_stock where id=$1 limit 1", materialStockID)
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
	clinicMaterialSets := []string{"price=" + price}

	if enName != "" {
		materialSets = append(materialSets, "en_name='"+enName+"'")
	}
	if pyCode != "" {
		materialSets = append(materialSets, "py_code='"+pyCode+"'")
	}
	if unitID != "" {
		materialSets = append(materialSets, "unit_id="+unitID)
	}
	if idcCode != "" {
		materialSets = append(materialSets, "idc_code='"+idcCode+"'")
	}
	if remark != "" {
		materialSets = append(materialSets, "remark='"+remark+"'")
	}
	if manuFactory != "" {
		materialSets = append(materialSets, "manu_factory='"+manuFactory+"'")
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
	if cost != "" {
		clinicMaterialSets = append(clinicMaterialSets, "cost="+cost)
	}
	if effDay != "" {
		clinicMaterialSets = append(clinicMaterialSets, "eff_day="+effDay)
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

	clinicMaterialUpdateSQL := "update material_stock set " + clinicMaterialSetStr + " where id=$1"
	fmt.Println("clinicMaterialUpdateSQL==", clinicMaterialUpdateSQL)

	_, err2 := tx.Exec(clinicMaterialUpdateSQL, materialStockID)
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
	materialStockID := ctx.PostValue("material_stock_id")
	status := ctx.PostValue("status")
	if clinicID == "" || materialStockID == "" || status == "" {
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

	crow := model.DB.QueryRowx("select id from material_stock where id=$1 limit 1", materialStockID)
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

	_, err1 := model.DB.Exec("update material_stock set status=$1 where id=$2", status, materialStockID)
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

	var storehouseID string
	errs := model.DB.QueryRow("select id from storehouse where clinic_id=$1 limit 1", clinicID).Scan(&storehouseID)
	if errs != nil {
		fmt.Println("errs ===", errs)
		ctx.JSON(iris.Map{"code": "1", "msg": errs.Error()})
		return
	}
	fmt.Println("storehouseID==", storehouseID)

	countSQL := `select count(ms.id) as total from material_stock ms
		left join material m on ms.material_id = m.id
		where ms.storehouse_id=$1`
	selectSQL := `select ms.material_id,ms.id as material_stock_id,m.name,m.unit_id,du.name as unit_name,m.py_code,m.remark,m.idc_code,m.manu_factory,m.specification,
		m.en_name,ms.is_discount,ms.price,ms.status,ms.cost,ms.eff_day,ms.stock_warning,ms.stock_amount
		from material_stock ms
		left join material m on ms.material_id = m.id
		left join dose_unit du on m.unit_id = du.id
		where ms.storehouse_id=$1`

	if keyword != "" {
		countSQL += " and m.name ~'" + keyword + "'"
		selectSQL += " and m.name ~'" + keyword + "'"
	}
	if status != "" {
		countSQL += " and ms.status=" + status
		selectSQL += " and ms.status=" + status
	}

	fmt.Println("countSQL===", countSQL)
	fmt.Println("selectSQL===", selectSQL)
	total := model.DB.QueryRowx(countSQL, storehouseID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL+" offset $2 limit $3", storehouseID, offset, limit)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})

}

//MaterialDetail 材料项目详情
func MaterialDetail(ctx iris.Context) {
	materialStockID := ctx.PostValue("material_stock_id")

	if materialStockID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select ms.material_id,ms.id as material_stock_id,m.name,m.unit_id,du.name as unit_name,m.py_code,m.remark,m.idc_code,
		m.manu_factory,m.specification,m.en_name,ms.is_discount,ms.price,ms.status,ms.cost,ms.eff_day,ms.stock_warning,ms.stock_amount
		from material_stock ms
		left join material m on ms.material_id = m.id
		left join dose_unit du on m.unit_id = du.id
		where ms.id=$1`

	fmt.Println("selectSQL===", selectSQL)

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL, materialStockID)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})
}

//MaterialPatientCreate 开材料费
func MaterialPatientCreate(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	personnelID := ctx.PostValue("personnel_id")
	items := ctx.PostValue("items")

	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	if items == "" {
		ctx.JSON(iris.Map{"code": "200", "data": nil})
		return
	}

	var results []map[string]string
	errj := json.Unmarshal([]byte(items), &results)
	fmt.Println("results===", results)
	if errj != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": errj.Error()})
		return
	}
	row := model.DB.QueryRowx(`select id,status from clinic_triage_patient where id=$1 limit 1`, clinicTriagePatientID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "保存检查失败,分诊记录错误"})
		return
	}

	prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", personnelID)
	if prow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "保存检查失败,操作员错误"})
		return
	}
	clinicTriagePatient := FormatSQLRowToMap(row)
	personnel := FormatSQLRowToMap(prow)

	_, ok := clinicTriagePatient["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "分诊记录不存在"})
		return
	}
	status := clinicTriagePatient["status"]
	if status.(int64) != 30 {
		ctx.JSON(iris.Map{"code": "1", "msg": "分诊记录当前状态错误"})
		return
	}
	_, pok := personnel["id"]
	if !pok {
		ctx.JSON(iris.Map{"code": "1", "msg": "操作员错误"})
		return
	}

	var mzUnpaidOrdersValues []string
	mzUnpaidOrdersSets := []string{
		"clinic_triage_patient_id",
		"charge_project_type_id",
		"charge_project_id",
		"order_sn",
		"soft_sn",
		"name",
		"price",
		"amount",
		"unit",
		"total",
		"fee",
		"operation_id",
	}

	var materialStockValues []string
	materialStockSets := []string{
		"clinic_triage_patient_id",
		"clinic_examination_id",
		"order_sn",
		"soft_sn",
		"times",
		"operation_id",
		"illustration",
	}

	orderSn := FormatPayOrderSn(clinicTriagePatientID, "5")

	for index, v := range results {
		clinicExaminationID := v["clinic_examination_id"]
		times := v["times"]
		illustration := v["illustration"]
		fmt.Println("clinicExaminationID====", clinicExaminationID)
		var sl []string
		var sm []string
		materialStockSQL := `select ms.id as material_stock_id,ms.price,ms.is_discount,m.name,du.name as dose_unit_name from material_stock ms
		left join material m on m.id = ms.material_id
		left join dose_unit du on du.id = m.unit_id
		where ms.id=$1`
		trow := model.DB.QueryRowx(materialStockSQL, clinicExaminationID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "检查项错误"})
			return
		}
		materialStock := FormatSQLRowToMap(trow)
		fmt.Println("====", materialStock)
		_, ok := materialStock["clinic_examination_id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "选择的检查项错误"})
			return
		}
		price := materialStock["price"].(int64)
		name := materialStock["name"].(string)
		unitName := materialStock["dose_unit_name"].(string)
		amount, _ := strconv.Atoi(times)
		total := int(price) * amount

		sl = append(sl, clinicTriagePatientID, clinicExaminationID, "'"+orderSn+"'", strconv.Itoa(index), times, personnelID)
		sm = append(sm, clinicTriagePatientID, "4", clinicExaminationID, "'"+orderSn+"'", strconv.Itoa(index), "'"+name+"'", strconv.FormatInt(price, 10), strconv.Itoa(amount), "'"+unitName+"'", strconv.Itoa(total), strconv.Itoa(total), personnelID)

		if illustration == "" {
			sl = append(sl, `null`)
		} else {
			sl = append(sl, "'"+illustration+"'")
		}

		tstr := "(" + strings.Join(sl, ",") + ")"
		materialStockValues = append(materialStockValues, tstr)
		mstr := "(" + strings.Join(sm, ",") + ")"
		mzUnpaidOrdersValues = append(mzUnpaidOrdersValues, mstr)
	}
	tSetStr := strings.Join(materialStockSets, ",")
	tValueStr := strings.Join(materialStockValues, ",")

	mSetStr := strings.Join(mzUnpaidOrdersSets, ",")
	mvValueStr := strings.Join(mzUnpaidOrdersValues, ",")

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errb})
		return
	}
	_, errdlp := tx.Exec("delete from material_patient where clinic_triage_patient_id=$1", clinicTriagePatientID)
	if errdlp != nil {
		fmt.Println("errdlp ===", errdlp)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errdlp.Error()})
		return
	}
	_, errdm := tx.Exec("delete from mz_unpaid_orders where clinic_triage_patient_id=$1 and charge_project_type_id=5", clinicTriagePatientID)
	if errdm != nil {
		fmt.Println("errdm ===", errdm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errdm.Error()})
		return
	}

	inserttSQL := "insert into material_patient (" + tSetStr + ") values " + tValueStr
	fmt.Println("inserttSQL===", inserttSQL)

	_, errt := tx.Exec(inserttSQL)
	if errt != nil {
		fmt.Println("errt ===", errt)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errt.Error()})
		return
	}

	insertmSQL := "insert into mz_unpaid_orders (" + mSetStr + ") values " + mvValueStr
	fmt.Println("insertmSQL===", insertmSQL)
	_, errm := tx.Exec(insertmSQL)
	if errm != nil {
		fmt.Println("errm ===", errm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": "请检查是否漏填"})
		return
	}
	errc := tx.Commit()
	if errc != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errc.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//MaterialPatientGet 获取材料费
func MaterialPatientGet(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	rows, err := model.DB.Queryx(`select mp.*, m.name from material_patient mp left join material_stock ms on material_stock_id = ms.id 
		left join material m on ms.material_id = m.id
		where mp.clinic_triage_patient_id = $1`, clinicTriagePatientID)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
	}

	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result})
}
