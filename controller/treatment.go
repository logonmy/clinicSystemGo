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

	lrow := model.DB.QueryRowx("select id from clinic_treatment where name=$1 and clinic_id=$2 limit 1", name, clinicID)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
		return
	}
	clinicTreatment := FormatSQLRowToMap(lrow)
	_, lok := clinicTreatment["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "1", "msg": "治疗名称已存在"})
		return
	}

	clinicTreatmentSets := []string{
		"clinic_id",
		"name",
		"en_name",
		"py_code",
		"idc_code",
		"unit_name",
		"remark",
		"cost",
		"price",
		"status",
		"is_discount"}
	clinicTreatmentSetstr := strings.Join(clinicTreatmentSets, ",")
	clinicTreatmentInsertSQL := "insert into clinic_treatment (" + clinicTreatmentSetstr + ") values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)"

	_, err := model.DB.Exec(clinicTreatmentInsertSQL,
		ToNullInt64(clinicID),
		ToNullString(name),
		ToNullString(enName),
		ToNullString(pyCode),
		ToNullString(idcCode),
		ToNullString(unitName),
		ToNullString(remark),
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

// TreatmentUpdate 更新治疗项目
func TreatmentUpdate(ctx iris.Context) {
	clinicTreatmentID := ctx.PostValue("clinic_treatment_id")
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

	if name == "" || clinicTreatmentID == "" || price == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	crow := model.DB.QueryRowx("select id,clinic_id from clinic_treatment where id=$1 limit 1", clinicTreatmentID)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicTreatment := FormatSQLRowToMap(crow)
	_, rok := clinicTreatment["id"]
	if !rok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所治疗项目数据错误"})
		return
	}

	clinicID := clinicTreatment["clinic_id"]

	lrow := model.DB.QueryRowx("select id from clinic_treatment where name=$1 and id!=$2 and clinic_id=$3 limit 1", name, clinicTreatmentID, clinicID)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicTreatmentu := FormatSQLRowToMap(lrow)
	_, lok := clinicTreatmentu["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "1", "msg": "治疗项目名称已存在"})
		return
	}

	clinicTreatmentUpdateSQL := `update clinic_treatment set 
		name=$1,
		en_name=$2,
		py_code=$3,
		idc_code=$4,
		unit_name=$5,
		remark=$6,
		cost=$7,
		price=$8,
		status=$9,
		is_discount=$10
		where id=$11`

	_, err2 := model.DB.Exec(clinicTreatmentUpdateSQL,
		ToNullString(name),
		ToNullString(enName),
		ToNullString(pyCode),
		ToNullString(idcCode),
		ToNullString(unitName),
		ToNullString(remark),
		ToNullInt64(cost),
		ToNullInt64(price),
		ToNullBool(status),
		ToNullBool(isDiscount),
		ToNullInt64(clinicTreatmentID),
	)
	if err2 != nil {
		fmt.Println(" err2====", err2)
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
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

	countSQL := `select count(id) as total from clinic_treatment	where clinic_id=:clinic_id and name ~:keyword`
	selectSQL := `select id as clinic_treatment_id,name as treatment_name,unit_name,py_code,remark,idc_code,
		en_name,is_discount,price,status,cost,discount_price from clinic_treatment where clinic_id=:clinic_id and name ~:keyword`

	if status != "" {
		countSQL += " and status=:status"
		selectSQL += " and status=:status"
	}

	var queryOption = map[string]interface{}{
		"clinic_id": ToNullInt64(clinicID),
		"keyword":   keyword,
		"status":    ToNullBool(status),
		"offset":    ToNullInt64(offset),
		"limit":     ToNullInt64(limit),
	}

	fmt.Println("countSQL===", countSQL)
	fmt.Println("selectSQL===", selectSQL)
	total, err := model.DB.NamedQuery(countSQL, queryOption)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.NamedQuery(selectSQL+" offset :offset limit :limit", queryOption)
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

	selectSQL := `select id as clinic_treatment_id,name,unit_name,py_code,remark,idc_code,
		en_name,is_discount,price,status,cost,discount_price
		from clinic_treatment
		where id=$1`

	fmt.Println("selectSQL===", selectSQL)

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL, clinicTreatmentID)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})
}
