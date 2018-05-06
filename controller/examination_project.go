package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

// ExaminationProjectCreate 创建检查缴费项目
func ExaminationProjectCreate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	pyCode := ctx.PostValue("py_code")
	idcCode := ctx.PostValue("idc_code")
	unit := ctx.PostValue("unit")
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

	lrow := model.DB.QueryRowx("select id from examination_project where name=$1 limit 1", name)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
		return
	}
	examinationProject := FormatSQLRowToMap(lrow)
	_, lok := examinationProject["id"]
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
	if unit != "" {
		examinationSets = append(examinationSets, "unit")
		examinationValues = append(examinationValues, "'"+unit+"'")
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

	examinationInsertSQL := "insert into examination_project (" + examinationSetStr + ") values (" + examinationValueStr + ") RETURNING id;"
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

	clinicExaminationSets = append(clinicExaminationSets, "examination_project_id")
	clinicExaminationValues = append(clinicExaminationValues, examinationID)

	clinicExaminationSetStr := strings.Join(clinicExaminationSets, ",")
	clinicExaminationValueStr := strings.Join(clinicExaminationValues, ",")

	clinicExaminationInsertSQL := "insert into clinic_examination_project (" + clinicExaminationSetStr + ") values (" + clinicExaminationValueStr + ")"
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

// ExaminationProjectUpdate 更新检查缴费项目
func ExaminationProjectUpdate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	clinicExaminationProjectID := ctx.PostValue("clinic_examination_project_id")
	examinationProjectID := ctx.PostValue("examination_project_id")

	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	pyCode := ctx.PostValue("py_code")
	idcCode := ctx.PostValue("idc_code")
	unit := ctx.PostValue("unit")
	organ := ctx.PostValue("organ")
	remark := ctx.PostValue("remark")

	price := ctx.PostValue("price")
	cost := ctx.PostValue("cost")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")

	if clinicID == "" || name == "" || clinicExaminationProjectID == "" || price == "" || examinationProjectID == "" {
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

	crow := model.DB.QueryRowx("select id,clinic_id,examination_project_id from clinic_examination_project where id=$1 limit 1", clinicExaminationProjectID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicExaminationProject := FormatSQLRowToMap(crow)
	_, rok := clinicExaminationProject["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所检查项目数据错误"})
		return
	}
	sexaminationProjectID := strconv.FormatInt(clinicExaminationProject["examination_project_id"].(int64), 10)
	fmt.Println("sexaminationProjectID====", sexaminationProjectID)

	if clinicID != strconv.FormatInt(clinicExaminationProject["clinic_id"].(int64), 10) {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据不匹配"})
		return
	}

	if sexaminationProjectID != examinationProjectID {
		ctx.JSON(iris.Map{"code": "1", "msg": "检查项目数据id不匹配"})
		return
	}

	lrow := model.DB.QueryRowx("select id from examination_project where name=$1 and id!=$2 limit 1", name, examinationProjectID)
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
	if unit != "" {
		examinationSets = append(examinationSets, "unit='"+unit+"'")
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

	examinationUpdateSQL := "update examination_project set " + examinationSetStr + " where id=$1"
	fmt.Println("examinationUpdateSQL==", examinationUpdateSQL)

	tx, err := model.DB.Begin()
	_, err = tx.Exec(examinationUpdateSQL, examinationProjectID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}

	clinicExaminationSets = append(clinicExaminationSets, "updated_time=LOCALTIMESTAMP")
	clinicExaminationSetStr := strings.Join(clinicExaminationSets, ",")

	clinicExaminationUpdateSQL := "update clinic_examination_project set " + clinicExaminationSetStr + " where id=$1"
	fmt.Println("clinicExaminationUpdateSQL==", clinicExaminationUpdateSQL)

	_, err2 := tx.Exec(clinicExaminationUpdateSQL, clinicExaminationProjectID)
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

// ExaminationProjectOnOff 启用和停用
func ExaminationProjectOnOff(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	clinicExaminationProjectID := ctx.PostValue("clinic_examination_project_id")
	status := ctx.PostValue("status")
	if clinicID == "" || clinicExaminationProjectID == "" || status == "" {
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

	crow := model.DB.QueryRowx("select id,clinic_id from clinic_examination_project where id=$1 limit 1", clinicExaminationProjectID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicExaminationProject := FormatSQLRowToMap(crow)
	_, rok := clinicExaminationProject["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据错误"})
		return
	}

	if clinicID != strconv.FormatInt(clinicExaminationProject["clinic_id"].(int64), 10) {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据不匹配"})
		return
	}
	_, err1 := model.DB.Exec("update clinic_examination_project set status=$1 where id=$2", status, clinicExaminationProjectID)
	if err1 != nil {
		fmt.Println(" err1====", err1)
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// ExaminationProjectList 检查缴费项目列表
func ExaminationProjectList(ctx iris.Context) {
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

	countSQL := `select count(cep.id) as total from clinic_examination_project cep
		left join examination_project ep on cep.examination_project_id = ep.id
		where cep.clinic_id=$1`
	selectSQL := `select cep.examination_project_id,cep.id as clinic_examination_project_id,ep.name,ep.unit,ep.py_code,ep.remark,ep.idc_code,
		ep.organ,ep.en_name,cep.is_discount,cep.price,cep.status,cep.cost
		from clinic_examination_project cep
		left join examination_project ep on cep.examination_project_id = ep.id
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

//ExaminationProjectDetail 检查项目详情
func ExaminationProjectDetail(ctx iris.Context) {
	clinicExaminationProjectID := ctx.PostValue("clinic_examination_project_id")

	if clinicExaminationProjectID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select cep.examination_project_id,cep.id as clinic_examination_project_id,ep.name,ep.unit,ep.py_code,ep.remark,ep.idc_code,
	ep.organ,ep.en_name,cep.is_discount,cep.price,cep.status,cep.cost
		from clinic_examination_project cep
		left join examination_project ep on cep.examination_project_id = ep.id
		where cep.id=$1`

	fmt.Println("selectSQL===", selectSQL)

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL, clinicExaminationProjectID)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})
}
