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
	unitName := ctx.PostValue("unit_name")
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
	if unitName != "" {
		examinationSets = append(examinationSets, "unit_name")
		examinationValues = append(examinationValues, "'"+unitName+"'")
	}
	if organ != "" {
		examinationSets = append(examinationSets, "organ")
		examinationValues = append(examinationValues, "'"+organ+"'")
	}

	if status != "" {
		clinicExaminationSets = append(clinicExaminationSets, "status")
		clinicExaminationValues = append(clinicExaminationValues, status)
	}
	if remark != "" {
		clinicExaminationSets = append(examinationSets, "remark")
		clinicExaminationValues = append(examinationValues, "'"+remark+"'")
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
	lrow := model.DB.QueryRowx("select id from examination where name=$1 limit 1", name)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
		return
	}
	examination := FormatSQLRowToMap(lrow)
	_, lok := examination["id"]
	if !lok {
		err = tx.QueryRow(examinationInsertSQL).Scan(&examinationID)
		if err != nil {
			fmt.Println("err ===", err)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "1", "msg": err})
			return
		}
	} else {
		examinationID = strconv.Itoa(int(examination["id"].(int64)))
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
	clinicExaminationID := ctx.PostValue("clinic_examination_id")

	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	pyCode := ctx.PostValue("py_code")
	idcCode := ctx.PostValue("idc_code")
	unitName := ctx.PostValue("unit_name")
	organ := ctx.PostValue("organ")
	remark := ctx.PostValue("remark")

	price := ctx.PostValue("price")
	cost := ctx.PostValue("cost")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")

	if name == "" || clinicExaminationID == "" || price == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if cost == "" {
		cost = "0"
	}
	if status == "" {
		status = "true"
	}
	if isDiscount == "" {
		isDiscount = "false"
	}

	crow := model.DB.QueryRowx("select id,clinic_id,examination_id from clinic_examination where id=$1 limit 1", clinicExaminationID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicExamination := FormatSQLRowToMap(crow)
	_, rok := clinicExamination["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所检查医嘱数据错误"})
		return
	}
	examinationID := strconv.FormatInt(clinicExamination["examination_id"].(int64), 10)
	clinicID := strconv.FormatInt(clinicExamination["clinic_id"].(int64), 10)
	fmt.Println("examinationID====", examinationID)
	fmt.Println("clinicID====", clinicID)

	lrow := model.DB.QueryRowx("select id from examination where name=$1 and id!=$2 limit 1", name, examinationID)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	examinationItem := FormatSQLRowToMap(lrow)
	_, lok := examinationItem["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "1", "msg": "检查医嘱名称已存在"})
		return
	}

	examinationUpdateSQL := `update examination set name=$1,en_name=$2,py_code=$3,idc_code=$4,
	unit_name=$5,organ=$6,remark=$7 where id=$8`
	fmt.Println("examinationUpdateSQL==", examinationUpdateSQL)

	clinicExaminationUpdateSQL := `update clinic_examination set clinic_id=$1,examination_id=$2,cost=$3,price=$4,status=$5,is_discount=$6 where id=$7`
	fmt.Println("clinicExaminationUpdateSQL==", clinicExaminationUpdateSQL)

	tx, err := model.DB.Begin()
	_, err = tx.Exec(examinationUpdateSQL, name, enName, pyCode, idcCode, unitName, organ, remark, examinationID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}

	_, err2 := tx.Exec(clinicExaminationUpdateSQL, clinicID, examinationID, cost, price, status, isDiscount, clinicExaminationID)
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

	countSQL := `select count(ce.id) as total from clinic_examination ce
		left join examination e on ce.examination_id = e.id
		where ce.clinic_id=$1`
	selectSQL := `select ce.examination_id,ce.id as clinic_examination_id,e.name,e.unit_name,e.py_code,ce.remark,e.idc_code,
		e.organ,e.en_name,ce.is_discount,ce.price,ce.status,ce.cost
		from clinic_examination ce
		left join examination e on ce.examination_id = e.id
		where ce.clinic_id=$1`

	if keyword != "" {
		countSQL += " and e.name ~'" + keyword + "'"
		selectSQL += " and e.name ~'" + keyword + "'"
	}
	if status != "" {
		countSQL += " and ce.status=" + status
		selectSQL += " and ce.status=" + status
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

	selectSQL := `select ce.examination_id,ce.id as clinic_examination_id,e.name,e.unit_name,e.py_code,e.remark,e.idc_code,
	e.organ,e.en_name,ce.is_discount,ce.price,ce.status,ce.cost
		from clinic_examination ce
		left join examination e on ce.examination_id = e.id
		where ce.id=$1`

	fmt.Println("selectSQL===", selectSQL)

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL, clinicExaminationID)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})
}
