package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

// ExaminationCreate 创建检查缴费项目
func ExaminationCreate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	pyCode := ctx.PostValue("py_code")
	idcCode := ctx.PostValue("idc_code")
	unitID := ctx.PostValue("unit_id")
	organ := ctx.PostValue("organ")
	remark := ctx.PostValue("remark")

	price := ctx.PostValue("price")
	cost := ctx.PostValue("cost")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")

	if clinicID == "" || name == "" || price == "" {
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

	lrow := model.DB.QueryRowx("select id from examination where name=$1 limit 1", name)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
		return
	}
	examination := FormatSQLRowToMap(lrow)
	_, lok := examination["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "1", "msg": "检查名称已存在"})
		return
	}

	examinationSets := []string{"name"}
	examinationValues := []string{"'" + name + "'"}

	clinicExaminationSets := []string{"clinic_id", "price"}
	clinicExaminationValues := []string{clinicID, price}

	if enName != "" {
		examinationSets = append(examinationSets, "en_name")
		examinationValues = append(examinationValues, "'"+enName+"'")
	}
	if pyCode != "" {
		examinationSets = append(examinationSets, "py_code")
		examinationValues = append(examinationValues, "'"+pyCode+"'")
	}
	if idcCode != "" {
		examinationSets = append(examinationSets, "idc_code")
		examinationValues = append(examinationValues, "'"+idcCode+"'")
	}
	if unitID != "" {
		examinationSets = append(examinationSets, "unit_id")
		examinationValues = append(examinationValues, unitID)
	}
	if organ != "" {
		examinationSets = append(examinationSets, "organ")
		examinationValues = append(examinationValues, "'"+organ+"'")
	}
	if remark != "" {
		examinationSets = append(examinationSets, "remark")
		examinationValues = append(examinationValues, "'"+remark+"'")
	}

	if status != "" {
		clinicExaminationSets = append(clinicExaminationSets, "status")
		clinicExaminationValues = append(clinicExaminationValues, status)
	}
	if cost != "" {
		clinicExaminationSets = append(clinicExaminationSets, "cost")
		clinicExaminationValues = append(clinicExaminationValues, cost)
	}
	if isDiscount != "" {
		clinicExaminationSets = append(clinicExaminationSets, "is_discount")
		clinicExaminationValues = append(clinicExaminationValues, isDiscount)
	}

	examinationSetStr := strings.Join(examinationSets, ",")
	examinationValueStr := strings.Join(examinationValues, ",")

	examinationInsertSQL := "insert into examination (" + examinationSetStr + ") values (" + examinationValueStr + ") RETURNING id;"
	fmt.Println("examinationInsertSQL==", examinationInsertSQL)

	tx, err := model.DB.Begin()
	var examinationID string
	err = tx.QueryRow(examinationInsertSQL).Scan(&examinationID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	fmt.Println("examinationID====", examinationID)

	clinicExaminationSets = append(clinicExaminationSets, "examination_id")
	clinicExaminationValues = append(clinicExaminationValues, examinationID)

	clinicExaminationSetStr := strings.Join(clinicExaminationSets, ",")
	clinicExaminationValueStr := strings.Join(clinicExaminationValues, ",")

	clinicExaminationInsertSQL := "insert into clinic_examination (" + clinicExaminationSetStr + ") values (" + clinicExaminationValueStr + ")"
	fmt.Println("clinicExaminationInsertSQL==", clinicExaminationInsertSQL)

	_, err2 := tx.Exec(clinicExaminationInsertSQL)
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

	ctx.JSON(iris.Map{"code": "200", "data": examinationID})

}

// ExaminationUpdate 更新检查缴费项目
func ExaminationUpdate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	clinicExaminationID := ctx.PostValue("clinic_examination_id")
	examinationID := ctx.PostValue("examination_id")

	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	pyCode := ctx.PostValue("py_code")
	idcCode := ctx.PostValue("idc_code")
	unitID := ctx.PostValue("unit_id")
	organ := ctx.PostValue("organ")
	remark := ctx.PostValue("remark")

	price := ctx.PostValue("price")
	cost := ctx.PostValue("cost")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")

	if clinicID == "" || name == "" || clinicExaminationID == "" || price == "" || examinationID == "" {
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

	crow := model.DB.QueryRowx("select id,clinic_id,examination_id from clinic_examination where id=$1 limit 1", clinicExaminationID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicExamination := FormatSQLRowToMap(crow)
	_, rok := clinicExamination["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所检查项目数据错误"})
		return
	}
	sexaminationID := strconv.FormatInt(clinicExamination["examination_id"].(int64), 10)
	fmt.Println("sexaminationID====", sexaminationID)

	if clinicID != strconv.FormatInt(clinicExamination["clinic_id"].(int64), 10) {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据不匹配"})
		return
	}

	if sexaminationID != examinationID {
		ctx.JSON(iris.Map{"code": "1", "msg": "检查项目数据id不匹配"})
		return
	}

	lrow := model.DB.QueryRowx("select id from examination where name=$1 and id!=$2 limit 1", name, examinationID)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	laboratoryItem := FormatSQLRowToMap(lrow)
	_, lok := laboratoryItem["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "1", "msg": "检查项目名称已存在"})
		return
	}

	examinationSets := []string{"name='" + name + "'"}
	clinicExaminationSets := []string{"price=" + price}

	if enName != "" {
		examinationSets = append(examinationSets, "en_name='"+enName+"'")
	}
	if pyCode != "" {
		examinationSets = append(examinationSets, "py_code='"+pyCode+"'")
	}
	if unitID != "" {
		examinationSets = append(examinationSets, "unit_id="+unitID)
	}
	if idcCode != "" {
		examinationSets = append(examinationSets, "idc_code='"+idcCode+"'")
	}
	if organ != "" {
		examinationSets = append(examinationSets, "organ='"+organ+"'")
	}
	if remark != "" {
		examinationSets = append(examinationSets, "remark='"+remark+"'")
	}

	if status != "" {
		clinicExaminationSets = append(clinicExaminationSets, "status="+status)
	}
	if isDiscount != "" {
		clinicExaminationSets = append(clinicExaminationSets, "is_discount="+isDiscount)
	}
	if cost != "" {
		clinicExaminationSets = append(clinicExaminationSets, "cost="+cost)
	}

	examinationSets = append(examinationSets, "updated_time=LOCALTIMESTAMP")
	examinationSetStr := strings.Join(examinationSets, ",")

	examinationUpdateSQL := "update examination set " + examinationSetStr + " where id=$1"
	fmt.Println("examinationUpdateSQL==", examinationUpdateSQL)

	tx, err := model.DB.Begin()
	_, err = tx.Exec(examinationUpdateSQL, examinationID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}

	clinicExaminationSets = append(clinicExaminationSets, "updated_time=LOCALTIMESTAMP")
	clinicExaminationSetStr := strings.Join(clinicExaminationSets, ",")

	clinicExaminationUpdateSQL := "update clinic_examination set " + clinicExaminationSetStr + " where id=$1"
	fmt.Println("clinicExaminationUpdateSQL==", clinicExaminationUpdateSQL)

	_, err2 := tx.Exec(clinicExaminationUpdateSQL, clinicExaminationID)
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

// ExaminationOnOff 启用和停用
func ExaminationOnOff(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	clinicExaminationID := ctx.PostValue("clinic_examination_id")
	status := ctx.PostValue("status")
	if clinicID == "" || clinicExaminationID == "" || status == "" {
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

	crow := model.DB.QueryRowx("select id,clinic_id from clinic_examination where id=$1 limit 1", clinicExaminationID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicExamination := FormatSQLRowToMap(crow)
	_, rok := clinicExamination["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据错误"})
		return
	}

	if clinicID != strconv.FormatInt(clinicExamination["clinic_id"].(int64), 10) {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据不匹配"})
		return
	}
	_, err1 := model.DB.Exec("update clinic_examination set status=$1 where id=$2", status, clinicExaminationID)
	if err1 != nil {
		fmt.Println(" err1====", err1)
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// ExaminationList 检查缴费项目列表
func ExaminationList(ctx iris.Context) {
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

	countSQL := `select count(cep.id) as total from clinic_examination cep
		left join examination ep on cep.examination_id = ep.id
		where cep.clinic_id=$1`
	selectSQL := `select cep.examination_id,cep.id as clinic_examination_id,ep.name,ep.unit_id,du.name as unit_name,ep.py_code,ep.remark,ep.idc_code,
		ep.organ,ep.en_name,cep.is_discount,cep.price,cep.status,cep.cost
		from clinic_examination cep
		left join examination ep on cep.examination_id = ep.id
		left join dose_unit du on ep.unit_id = du.id
		where cep.clinic_id=$1`

	if keyword != "" {
		countSQL += " and ep.name ~'" + keyword + "'"
		selectSQL += " and ep.name ~'" + keyword + "'"
	}
	if status != "" {
		countSQL += " and cep.status=" + status
		selectSQL += " and cep.status=" + status
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

//ExaminationDetail 检查项目详情
func ExaminationDetail(ctx iris.Context) {
	clinicExaminationID := ctx.PostValue("clinic_examination_id")

	if clinicExaminationID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select cep.examination_id,cep.id as clinic_examination_id,ep.name,ep.unit_id,du.name as unit_name,ep.py_code,ep.remark,ep.idc_code,
	ep.organ,ep.en_name,cep.is_discount,cep.price,cep.status,cep.cost
		from clinic_examination cep
		left join examination ep on cep.examination_id = ep.id
		left join dose_unit du on ep.unit_id = du.id
		where cep.id=$1`

	fmt.Println("selectSQL===", selectSQL)

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL, clinicExaminationID)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})
}
