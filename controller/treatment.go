package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

// TreatmentCreate 创建治疗缴费项目
func TreatmentCreate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	pyCode := ctx.PostValue("py_code")
	idcCode := ctx.PostValue("idc_code")
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

	lrow := model.DB.QueryRowx("select id from treatment where name=$1 limit 1", name)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
		return
	}
	treatment := FormatSQLRowToMap(lrow)
	_, lok := treatment["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "1", "msg": "治疗名称已存在"})
		return
	}

	treatmentSets := []string{"name", "unit_name"}
	treatmentValues := []string{"'" + name + "'", "'" + unitName + "'"}

	clinictreatmentSets := []string{"clinic_id", "price"}
	clinictreatmentValues := []string{clinicID, price}

	if enName != "" {
		treatmentSets = append(treatmentSets, "en_name")
		treatmentValues = append(treatmentValues, "'"+enName+"'")
	}
	if pyCode != "" {
		treatmentSets = append(treatmentSets, "py_code")
		treatmentValues = append(treatmentValues, "'"+pyCode+"'")
	}
	if idcCode != "" {
		treatmentSets = append(treatmentSets, "idc_code")
		treatmentValues = append(treatmentValues, "'"+idcCode+"'")
	}
	if remark != "" {
		treatmentSets = append(treatmentSets, "remark")
		treatmentValues = append(treatmentValues, "'"+remark+"'")
	}

	if status != "" {
		clinictreatmentSets = append(clinictreatmentSets, "status")
		clinictreatmentValues = append(clinictreatmentValues, status)
	}
	if cost != "" {
		clinictreatmentSets = append(clinictreatmentSets, "cost")
		clinictreatmentValues = append(clinictreatmentValues, cost)
	}
	if isDiscount != "" {
		clinictreatmentSets = append(clinictreatmentSets, "is_discount")
		clinictreatmentValues = append(clinictreatmentValues, isDiscount)
	}

	treatmentSetstr := strings.Join(treatmentSets, ",")
	treatmentValuestr := strings.Join(treatmentValues, ",")

	treatmentInsertSQL := "insert into treatment (" + treatmentSetstr + ") values (" + treatmentValuestr + ") RETURNING id;"
	fmt.Println("treatmentInsertSQL==", treatmentInsertSQL)

	tx, err := model.DB.Begin()
	var treatmentID string
	err = tx.QueryRow(treatmentInsertSQL).Scan(&treatmentID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	fmt.Println("treatmentID====", treatmentID)

	clinictreatmentSets = append(clinictreatmentSets, "treatment_id")
	clinictreatmentValues = append(clinictreatmentValues, treatmentID)

	clinictreatmentSetstr := strings.Join(clinictreatmentSets, ",")
	clinictreatmentValuestr := strings.Join(clinictreatmentValues, ",")

	clinictreatmentInsertSQL := "insert into clinic_treatment (" + clinictreatmentSetstr + ") values (" + clinictreatmentValuestr + ")"
	fmt.Println("clinictreatmentInsertSQL==", clinictreatmentInsertSQL)

	_, err2 := tx.Exec(clinictreatmentInsertSQL)
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

	ctx.JSON(iris.Map{"code": "200", "data": treatmentID})

}

// TreatmentUpdate 更新治疗项目
func TreatmentUpdate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	clinicTreatmentID := ctx.PostValue("clinic_treatment_id")
	treatmentID := ctx.PostValue("treatment_id")

	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	pyCode := ctx.PostValue("py_code")
	idcCode := ctx.PostValue("idc_code")
	unitName := ctx.PostValue("unit_name")
	remark := ctx.PostValue("remark")

	price := ctx.PostValue("price")
	cost := ctx.PostValue("cost")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")

	if clinicID == "" || name == "" || clinicTreatmentID == "" || price == "" || treatmentID == "" {
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

	crow := model.DB.QueryRowx("select id,clinic_id,treatment_id from clinic_treatment where id=$1 limit 1", clinicTreatmentID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicTreatmentProject := FormatSQLRowToMap(crow)
	_, rok := clinicTreatmentProject["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所治疗项目数据错误"})
		return
	}
	streatmentID := strconv.FormatInt(clinicTreatmentProject["treatment_id"].(int64), 10)
	fmt.Println("streatmentID====", streatmentID)

	if clinicID != strconv.FormatInt(clinicTreatmentProject["clinic_id"].(int64), 10) {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据不匹配"})
		return
	}

	if streatmentID != treatmentID {
		ctx.JSON(iris.Map{"code": "1", "msg": "治疗项目数据id不匹配"})
		return
	}

	lrow := model.DB.QueryRowx("select id from treatment where name=$1 and id!=$2 limit 1", name, treatmentID)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	laboratoryItem := FormatSQLRowToMap(lrow)
	_, lok := laboratoryItem["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "1", "msg": "治疗项目名称已存在"})
		return
	}

	treatmentSets := []string{"name='" + name + "'"}
	clinictreatmentSets := []string{"price=" + price}

	if enName != "" {
		treatmentSets = append(treatmentSets, "en_name='"+enName+"'")
	}
	if pyCode != "" {
		treatmentSets = append(treatmentSets, "py_code='"+pyCode+"'")
	}
	if unitName != "" {
		treatmentSets = append(treatmentSets, "unit_name="+"'"+unitName+"'")
	}
	if idcCode != "" {
		treatmentSets = append(treatmentSets, "idc_code='"+idcCode+"'")
	}
	if remark != "" {
		treatmentSets = append(treatmentSets, "remark='"+remark+"'")
	}

	if status != "" {
		clinictreatmentSets = append(clinictreatmentSets, "status="+status)
	}
	if isDiscount != "" {
		clinictreatmentSets = append(clinictreatmentSets, "is_discount="+isDiscount)
	}
	if cost != "" {
		clinictreatmentSets = append(clinictreatmentSets, "cost="+cost)
	}

	treatmentSets = append(treatmentSets, "updated_time=LOCALTIMESTAMP")
	treatmentSetstr := strings.Join(treatmentSets, ",")

	treatmentUpdateSQL := "update treatment set " + treatmentSetstr + " where id=$1"
	fmt.Println("treatmentUpdateSQL==", treatmentUpdateSQL)

	tx, err := model.DB.Begin()
	_, err = tx.Exec(treatmentUpdateSQL, treatmentID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}

	clinictreatmentSets = append(clinictreatmentSets, "updated_time=LOCALTIMESTAMP")
	clinictreatmentSetstr := strings.Join(clinictreatmentSets, ",")

	clinicTreatmentUpdateSQL := "update clinic_treatment set " + clinictreatmentSetstr + " where id=$1"
	fmt.Println("clinicTreatmentUpdateSQL==", clinicTreatmentUpdateSQL)

	_, err2 := tx.Exec(clinicTreatmentUpdateSQL, clinicTreatmentID)
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

// TreatmentOnOff 启用和停用
func TreatmentOnOff(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	clinicTreatmentID := ctx.PostValue("clinic_treatment_id")
	status := ctx.PostValue("status")
	if clinicID == "" || clinicTreatmentID == "" || status == "" {
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

	crow := model.DB.QueryRowx("select id,clinic_id from clinic_treatment where id=$1 limit 1", clinicTreatmentID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicTreatmentProject := FormatSQLRowToMap(crow)
	_, rok := clinicTreatmentProject["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据错误"})
		return
	}

	if clinicID != strconv.FormatInt(clinicTreatmentProject["clinic_id"].(int64), 10) {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据不匹配"})
		return
	}
	_, err1 := model.DB.Exec("update clinic_treatment set status=$1 where id=$2", status, clinicTreatmentID)
	if err1 != nil {
		fmt.Println(" err1====", err1)
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// TreatmentList 治疗缴费项目列表
func TreatmentList(ctx iris.Context) {
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

	countSQL := `select count(ct.id) as total from clinic_treatment ct
		left join treatment t on ct.treatment_id = t.id
		where ct.clinic_id=$1 and t.name ~$2`
	selectSQL := `select ct.treatment_id,ct.id as clinic_treatment_id,t.name as treatment_name,t.unit_name,t.py_code,t.remark,t.idc_code,
		t.en_name,ct.is_discount,ct.price,ct.status,ct.cost
		from clinic_treatment ct
		left join treatment t on ct.treatment_id = t.id
		where ct.clinic_id=$1 and t.name ~$2`

	if status != "" {
		countSQL += " and ct.status=" + status
		selectSQL += " and ct.status=" + status
	}

	fmt.Println("countSQL===", countSQL)
	fmt.Println("selectSQL===", selectSQL)
	total := model.DB.QueryRowx(countSQL, clinicID, keyword)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL+" offset $3 limit $4", clinicID, keyword, offset, limit)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})

}

//TreatmentDetail 治疗项目详情
func TreatmentDetail(ctx iris.Context) {
	clinicTreatmentID := ctx.PostValue("clinic_treatment_id")

	if clinicTreatmentID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select ct.treatment_id,ct.id as clinic_treatment_id,t.name,t.unit_name,t.py_code,t.remark,t.idc_code,
		t.en_name,ct.is_discount,ct.price,ct.status,ct.cost
		from clinic_treatment ct
		left join treatment t on ct.treatment_id = t.id
		where ct.id=$1`

	fmt.Println("selectSQL===", selectSQL)

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL, clinicTreatmentID)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})
}
