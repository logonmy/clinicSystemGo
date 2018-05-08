package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

// DiagnosisTreatmentCreate 创建诊疗缴费项目
func DiagnosisTreatmentCreate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")

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

	lrow := model.DB.QueryRowx("select id from diagnosis_treatment where name=$1 limit 1", name)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
		return
	}
	diagnosisTreatment := FormatSQLRowToMap(lrow)
	_, lok := diagnosisTreatment["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊疗名称已存在"})
		return
	}

	diagnosisTreatmentSets := []string{"name"}
	diagnosisTreatmentValues := []string{"'" + name + "'"}

	clinicDiagnosisTreatmentSets := []string{"clinic_id", "price"}
	clinicDiagnosisTreatmentValues := []string{clinicID, price}

	if enName != "" {
		diagnosisTreatmentSets = append(diagnosisTreatmentSets, "en_name")
		diagnosisTreatmentValues = append(diagnosisTreatmentValues, "'"+enName+"'")
	}

	if status != "" {
		clinicDiagnosisTreatmentSets = append(clinicDiagnosisTreatmentSets, "status")
		clinicDiagnosisTreatmentValues = append(clinicDiagnosisTreatmentValues, status)
	}
	if cost != "" {
		clinicDiagnosisTreatmentSets = append(clinicDiagnosisTreatmentSets, "cost")
		clinicDiagnosisTreatmentValues = append(clinicDiagnosisTreatmentValues, cost)
	}
	if isDiscount != "" {
		clinicDiagnosisTreatmentSets = append(clinicDiagnosisTreatmentSets, "is_discount")
		clinicDiagnosisTreatmentValues = append(clinicDiagnosisTreatmentValues, isDiscount)
	}

	diagnosisTreatmentSetstr := strings.Join(diagnosisTreatmentSets, ",")
	diagnosisTreatmentValuestr := strings.Join(diagnosisTreatmentValues, ",")

	diagnosisTreatmentInsertSQL := "insert into diagnosis_treatment (" + diagnosisTreatmentSetstr + ") values (" + diagnosisTreatmentValuestr + ") RETURNING id;"
	fmt.Println("diagnosisTreatmentInsertSQL==", diagnosisTreatmentInsertSQL)

	tx, err := model.DB.Begin()
	var diagnosisTreatmentID string
	err = tx.QueryRow(diagnosisTreatmentInsertSQL).Scan(&diagnosisTreatmentID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	fmt.Println("diagnosisTreatmentID====", diagnosisTreatmentID)

	clinicDiagnosisTreatmentSets = append(clinicDiagnosisTreatmentSets, "diagnosis_treatment_id")
	clinicDiagnosisTreatmentValues = append(clinicDiagnosisTreatmentValues, diagnosisTreatmentID)

	clinicDiagnosisTreatmentSetstr := strings.Join(clinicDiagnosisTreatmentSets, ",")
	clinicDiagnosisTreatmentValuestr := strings.Join(clinicDiagnosisTreatmentValues, ",")

	clinicDiagnosisTreatmentInsertSQL := "insert into clinic_diagnosis_treatment (" + clinicDiagnosisTreatmentSetstr + ") values (" + clinicDiagnosisTreatmentValuestr + ")"
	fmt.Println("clinicDiagnosisTreatmentInsertSQL==", clinicDiagnosisTreatmentInsertSQL)

	_, err2 := tx.Exec(clinicDiagnosisTreatmentInsertSQL)
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

	ctx.JSON(iris.Map{"code": "200", "data": diagnosisTreatmentID})

}

// DiagnosisTreatmentUpdate 更新诊疗项目
func DiagnosisTreatmentUpdate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	clinicDiagnosisTreatmentID := ctx.PostValue("clinic_diagnosis_treatment_id")
	diagnosisTreatmentID := ctx.PostValue("diagnosis_treatment_id")

	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")

	price := ctx.PostValue("price")
	cost := ctx.PostValue("cost")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")

	if clinicID == "" || name == "" || clinicDiagnosisTreatmentID == "" || price == "" || diagnosisTreatmentID == "" {
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

	crow := model.DB.QueryRowx("select id,clinic_id,diagnosis_treatment_id from clinic_diagnosis_treatment where id=$1 limit 1", clinicDiagnosisTreatmentID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicDiagnosisTreatmentProject := FormatSQLRowToMap(crow)
	_, rok := clinicDiagnosisTreatmentProject["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所诊疗项目数据错误"})
		return
	}
	sdiagnosisTreatmentID := strconv.FormatInt(clinicDiagnosisTreatmentProject["diagnosis_treatment_id"].(int64), 10)
	fmt.Println("sdiagnosisTreatmentID====", sdiagnosisTreatmentID)

	if clinicID != strconv.FormatInt(clinicDiagnosisTreatmentProject["clinic_id"].(int64), 10) {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据不匹配"})
		return
	}

	if sdiagnosisTreatmentID != diagnosisTreatmentID {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊疗项目数据id不匹配"})
		return
	}

	lrow := model.DB.QueryRowx("select id from diagnosis_treatment where name=$1 and id!=$2 limit 1", name, diagnosisTreatmentID)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	laboratoryItem := FormatSQLRowToMap(lrow)
	_, lok := laboratoryItem["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊疗项目名称已存在"})
		return
	}

	diagnosisTreatmentSets := []string{"name='" + name + "'"}
	clinicDiagnosisTreatmentSets := []string{"price=" + price}

	if enName != "" {
		diagnosisTreatmentSets = append(diagnosisTreatmentSets, "en_name='"+enName+"'")
	}

	if status != "" {
		clinicDiagnosisTreatmentSets = append(clinicDiagnosisTreatmentSets, "status="+status)
	}
	if isDiscount != "" {
		clinicDiagnosisTreatmentSets = append(clinicDiagnosisTreatmentSets, "is_discount="+isDiscount)
	}
	if cost != "" {
		clinicDiagnosisTreatmentSets = append(clinicDiagnosisTreatmentSets, "cost="+cost)
	}

	diagnosisTreatmentSets = append(diagnosisTreatmentSets, "updated_time=LOCALTIMESTAMP")
	diagnosisTreatmentSetstr := strings.Join(diagnosisTreatmentSets, ",")

	diagnosisTreatmentUpdateSQL := "update diagnosis_treatment set " + diagnosisTreatmentSetstr + " where id=$1"
	fmt.Println("diagnosisTreatmentUpdateSQL==", diagnosisTreatmentUpdateSQL)

	tx, err := model.DB.Begin()
	_, err = tx.Exec(diagnosisTreatmentUpdateSQL, diagnosisTreatmentID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}

	clinicDiagnosisTreatmentSets = append(clinicDiagnosisTreatmentSets, "updated_time=LOCALTIMESTAMP")
	clinicDiagnosisTreatmentSetstr := strings.Join(clinicDiagnosisTreatmentSets, ",")

	clinicDiagnosisTreatmentUpdateSQL := "update clinic_diagnosis_treatment set " + clinicDiagnosisTreatmentSetstr + " where id=$1"
	fmt.Println("clinicDiagnosisTreatmentUpdateSQL==", clinicDiagnosisTreatmentUpdateSQL)

	_, err2 := tx.Exec(clinicDiagnosisTreatmentUpdateSQL, clinicDiagnosisTreatmentID)
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

// DiagnosisTreatmentOnOff 启用和停用
func DiagnosisTreatmentOnOff(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	clinicDiagnosisTreatmentID := ctx.PostValue("clinic_diagnosis_treatment_id")
	status := ctx.PostValue("status")
	if clinicID == "" || clinicDiagnosisTreatmentID == "" || status == "" {
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

	crow := model.DB.QueryRowx("select id,clinic_id from clinic_diagnosis_treatment where id=$1 limit 1", clinicDiagnosisTreatmentID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicDiagnosisTreatmentProject := FormatSQLRowToMap(crow)
	_, rok := clinicDiagnosisTreatmentProject["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据错误"})
		return
	}

	if clinicID != strconv.FormatInt(clinicDiagnosisTreatmentProject["clinic_id"].(int64), 10) {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据不匹配"})
		return
	}
	_, err1 := model.DB.Exec("update clinic_diagnosis_treatment set status=$1 where id=$2", status, clinicDiagnosisTreatmentID)
	if err1 != nil {
		fmt.Println(" err1====", err1)
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// DiagnosisTreatmentList 诊疗缴费项目列表
func DiagnosisTreatmentList(ctx iris.Context) {
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

	countSQL := `select count(cdt.id) as total from clinic_diagnosis_treatment cdt
		left join diagnosis_treatment dt on cdt.diagnosis_treatment_id = dt.id
		where cdt.clinic_id=$1`
	selectSQL := `select cdt.diagnosis_treatment_id,cdt.id as clinic_diagnosis_treatment_id,dt.name,
		dt.en_name,cdt.is_discount,cdt.price,cdt.status,cdt.cost
		from clinic_diagnosis_treatment cdt
		left join diagnosis_treatment dt on cdt.diagnosis_treatment_id = dt.id
		where cdt.clinic_id=$1`

	if keyword != "" {
		countSQL += " and dt.name ~'" + keyword + "'"
		selectSQL += " and dt.name ~'" + keyword + "'"
	}
	if status != "" {
		countSQL += " and cdt.status=" + status
		selectSQL += " and cdt.status=" + status
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

//DiagnosisTreatmentDetail 诊疗项目详情
func DiagnosisTreatmentDetail(ctx iris.Context) {
	clinicDiagnosisTreatmentID := ctx.PostValue("clinic_diagnosis_treatment_id")

	if clinicDiagnosisTreatmentID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select cdt.diagnosis_treatment_id,cdt.id as clinic_diagnosis_treatment_id,dt.name,
	dt.en_name,cdt.is_discount,cdt.price,cdt.status,cdt.cost
		from clinic_diagnosis_treatment cdt
		left join diagnosis_treatment dt on cdt.diagnosis_treatment_id = dt.id
		where cdt.id=$1`

	fmt.Println("selectSQL===", selectSQL)

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL, clinicDiagnosisTreatmentID)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})
}
