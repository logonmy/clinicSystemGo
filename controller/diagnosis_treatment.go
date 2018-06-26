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

	lrow := model.DB.QueryRowx("select id from clinic_diagnosis_treatment where name=$1 and clinic_id=$2 limit 1", name, clinicID)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
		return
	}
	diagnosisTreatment := FormatSQLRowToMap(lrow)
	_, lok := diagnosisTreatment["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊疗名称已存在"})
		return
	}

	diagnosisTreatmentSets := []string{
		"clinic_id",
		"name",
		"en_name",
		"cost",
		"price",
		"status",
		"is_discount"}
	diagnosisTreatmentSetstr := strings.Join(diagnosisTreatmentSets, ",")
	diagnosisTreatmentInsertSQL := "insert into clinic_diagnosis_treatment (" + diagnosisTreatmentSetstr + ") values ($1,$2,$3,$4,$5,$6,$7)"

	_, err := model.DB.Exec(diagnosisTreatmentInsertSQL,
		ToNullInt64(clinicID),
		ToNullString(name),
		ToNullString(enName),
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

// DiagnosisTreatmentUpdate 更新诊疗项目
func DiagnosisTreatmentUpdate(ctx iris.Context) {
	clinicDiagnosisTreatmentID := ctx.PostValue("clinic_diagnosis_treatment_id")
	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	price := ctx.PostValue("price")
	cost := ctx.PostValue("cost")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")

	if name == "" || clinicDiagnosisTreatmentID == "" || price == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	crow := model.DB.QueryRowx("select id,clinic_id from clinic_diagnosis_treatment where id=$1 limit 1", clinicDiagnosisTreatmentID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	clinicDiagnosisTreatment := FormatSQLRowToMap(crow)
	_, rok := clinicDiagnosisTreatment["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊所诊疗项目数据错误"})
		return
	}
	clinicID := clinicDiagnosisTreatment["clinic_id"]

	lrow := model.DB.QueryRowx("select id from clinic_diagnosis_treatment where name=$1 and id!=$2 limit 1", name, clinicDiagnosisTreatmentID, clinicID)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	laboratoryItem := FormatSQLRowToMap(lrow)
	_, lok := laboratoryItem["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊疗项目名称已存在"})
		return
	}

	clinicDiagnosisTreatmentUpdateSQL := `update clinic_diagnosis_treatment set 
		name=$1,
		en_name=$2,
		cost=$7,
		price=$8,
		status=$9,
		is_discount=$10
		where id=$11`

	_, err2 := model.DB.Exec(clinicDiagnosisTreatmentUpdateSQL,
		ToNullString(name),
		ToNullString(enName),
		ToNullInt64(cost),
		ToNullInt64(price),
		ToNullBool(status),
		ToNullBool(isDiscount),
		ToNullInt64(clinicDiagnosisTreatmentID),
	)
	if err2 != nil {
		fmt.Println(" err2====", err2)
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
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
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)
	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊所数据错误"})
		return
	}

	crow := model.DB.QueryRowx("select id,clinic_id from clinic_diagnosis_treatment where id=$1 limit 1", clinicDiagnosisTreatmentID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	clinicDiagnosisTreatmentProject := FormatSQLRowToMap(crow)
	_, rok := clinicDiagnosisTreatmentProject["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊所数据错误"})
		return
	}

	if clinicID != strconv.FormatInt(clinicDiagnosisTreatmentProject["clinic_id"].(int64), 10) {
		ctx.JSON(iris.Map{"code": "-1", "msg": "诊所数据不匹配"})
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
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)

	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "所在诊所不存在"})
		return
	}

	countSQL := `select count(id) as total from clinic_diagnosis_treatment where clinic_id=:clinic_id`
	selectSQL := `select id as clinic_diagnosis_treatment_id,name,
		en_name,is_discount,price,status,cost
		from clinic_diagnosis_treatment cdt
		where clinic_id=:clinic_id`

	if keyword != "" {
		countSQL += ` and name ~*:keyword`
		selectSQL += ` and name ~*:keyword`
	}
	if status != "" {
		countSQL += " and status=:status"
		selectSQL += " and status=:status"
	}

	var queryOptions = map[string]interface{}{
		"clinic_id": ToNullInt64(clinicID),
		"keyword":   ToNullString(keyword),
		"offset":    ToNullInt64(offset),
		"limit":     ToNullInt64(limit),
		"status":    ToNullBool(status),
	}

	fmt.Println("countSQL===", countSQL)
	fmt.Println("selectSQL===", selectSQL)
	pageInfoRows, err := model.DB.NamedQuery(countSQL, queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	pageInfoArray := FormatSQLRowsToMapArray(pageInfoRows)
	pageInfo := pageInfoArray[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, _ := model.DB.NamedQuery(selectSQL+" offset :offset limit :limit", queryOptions)
	results := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})

}

//DiagnosisTreatmentDetail 诊疗项目详情
func DiagnosisTreatmentDetail(ctx iris.Context) {
	clinicDiagnosisTreatmentID := ctx.PostValue("clinic_diagnosis_treatment_id")

	if clinicDiagnosisTreatmentID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select id as clinic_diagnosis_treatment_id,name,
		en_name,is_discount,price,status,cost
		from clinic_diagnosis_treatment
		where id=$1`

	fmt.Println("selectSQL===", selectSQL)

	var results map[string]interface{}
	rows := model.DB.QueryRowx(selectSQL, clinicDiagnosisTreatmentID)
	results = FormatSQLRowToMap(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})
}
