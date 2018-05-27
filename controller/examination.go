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

	cerow := model.DB.QueryRowx("select id from clinic_examination where clinic_id=$1 and name=$2 limit 1", clinicID, name)
	if cerow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
		return
	}
	clinicExamination := FormatSQLRowToMap(cerow)
	_, lok := clinicExamination["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "1", "msg": "检查医嘱名称在该诊所已存在"})
		return
	}

	clinicExaminationSets := []string{
		"clinic_id",
		"name",
		"en_name",
		"py_code",
		"idc_code",
		"unit_name",
		"organ",
		"status",
		"remark",
		"price",
		"cost",
		"is_discount",
	}
	clinicExaminationSetStr := strings.Join(clinicExaminationSets, ",")

	clinicExaminationInsertSQL := "insert into clinic_examination (" + clinicExaminationSetStr + ") values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING id;"

	var clinicExaminationID int
	err2 := model.DB.QueryRow(clinicExaminationInsertSQL,
		ToNullInt64(clinicID),
		ToNullString(name),
		ToNullString(enName),
		ToNullString(pyCode),
		ToNullString(idcCode),
		ToNullString(unitName),
		ToNullString(organ),
		ToNullBool(status),
		ToNullString(remark),
		ToNullInt64(price),
		ToNullInt64(cost),
		ToNullBool(isDiscount)).Scan(&clinicExaminationID)
	if err2 != nil {
		fmt.Println(" err2====", err2)
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": clinicExaminationID})
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

	crow := model.DB.QueryRowx("select id,clinic_id from clinic_examination where id=$1 limit 1", clinicExaminationID)
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
	clinicID := clinicExamination["clinic_id"]

	lrow := model.DB.QueryRowx("select id from clinic_examination where name=$1 and id!=$2 and clinic_id=$3 limit 1", name, clinicExaminationID, clinicID)
	if lrow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "修改失败"})
		return
	}
	clinicExaminationu := FormatSQLRowToMap(lrow)
	_, lok := clinicExaminationu["id"]
	if lok {
		ctx.JSON(iris.Map{"code": "1", "msg": "检查医嘱名称已存在"})
		return
	}

	clinicExaminationUpdateSQL := `update clinic_examination set 
		name=$1,
		en_name=$2,
		py_code=$3,
		idc_code=$4,
		unit_name=$5,
		organ=$6,
		status=$7,
		remark=$8,
		price=$9,
		cost=$10,
		is_discount=$11 
		where id=$12`

	_, err2 := model.DB.Exec(clinicExaminationUpdateSQL,
		ToNullString(name),
		ToNullString(enName),
		ToNullString(pyCode),
		ToNullString(idcCode),
		ToNullString(unitName),
		ToNullString(organ),
		ToNullBool(status),
		ToNullString(remark),
		ToNullInt64(price),
		ToNullInt64(cost),
		ToNullBool(isDiscount),
		ToNullInt64(clinicExaminationID),
	)
	if err2 != nil {
		fmt.Println(" err2====", err2)
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
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

	countSQL := `select count(id) as total from clinic_examination where clinic_id=:clinic_id`
	selectSQL := `select id as clinic_examination_id,name,unit_name,py_code,remark,idc_code,
		organ,en_name,is_discount,price,status,cost from clinic_examination where clinic_id=:clinic_id`

	if keyword != "" {
		countSQL += " and name ~:keyword"
		selectSQL += " and name ~:keyword"
	}
	if status != "" {
		countSQL += " and status=:status"
		selectSQL += " and status=:status"
	}

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

//ExaminationDetail 检查项目详情
func ExaminationDetail(ctx iris.Context) {
	clinicExaminationID := ctx.PostValue("clinic_examination_id")

	if clinicExaminationID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectSQL := `select id as clinic_examination_id,name,unit_name,py_code,remark,idc_code,
	organ,en_name,is_discount,price,status,cost from clinic_examination	where id=$1`

	fmt.Println("selectSQL===", selectSQL)

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL, clinicExaminationID)
	results = FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results})
}
