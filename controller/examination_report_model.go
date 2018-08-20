package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"

	"github.com/kataras/iris"
)

// ExaminationReportModelCreate 创建检查报告医嘱模板
func ExaminationReportModelCreate(ctx iris.Context) {
	modelName := ctx.PostValue("model_name")
	resultExamination := ctx.PostValue("result_examination")
	conclusionExamination := ctx.PostValue("conclusion_examination")
	personnelID := ctx.PostValue("operation_id")

	if modelName == "" || personnelID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from examination_report_model where model_name=$1 limit 1", modelName)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "新增失败"})
		return
	}
	examinationReportModel := FormatSQLRowToMap(row)
	_, ok := examinationReportModel["id"]
	if ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "模板名称已存在"})
		return
	}

	prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", personnelID)
	if prow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "保存模板失败,操作员错误"})
		return
	}
	personnel := FormatSQLRowToMap(prow)
	_, pok := personnel["id"]
	if !pok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "操作员错误"})
		return
	}

	_, err := model.DB.Exec("insert into examination_report_model (model_name,result_examination,conclusion_examination,operation_id) values ($1,$2,$3,$4)", modelName, resultExamination, conclusionExamination, personnelID)
	if err != nil {
		fmt.Println("err ===", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// ExaminationReportModelList 查询检查报告医嘱模板
func ExaminationReportModelList(ctx iris.Context) {
	keyword := ctx.PostValue("keyword")
	operationID := ctx.PostValue("operation_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

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

	countSQL := `select count(id) as total from examination_report_model where id>0 and deleted_time is null`
	selectSQL := `select erm.*,p.name as operation_name from examination_report_model erm
	left join personnel p on erm.operation_id = p.id
	where erm.id >0 and erm.deleted_time is null`

	if keyword != "" {
		countSQL += ` and model_name ~*:keyword`
		selectSQL += ` and erm.model_name ~*:keyword`
	}

	if operationID != "" {
		countSQL += ` and operation_id =:operation_id`
		selectSQL += ` and erm.operation_id=:operation_id`
	}

	var queryOption = map[string]interface{}{
		"operation_id": ToNullInt64(operationID),
		"keyword":      keyword,
		"offset":       ToNullInt64(offset),
		"limit":        ToNullInt64(limit),
	}

	fmt.Println("countSQL===", countSQL)
	fmt.Println("selectSQL===", selectSQL, queryOption)
	total, err := model.DB.NamedQuery(countSQL, queryOption)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err1 := model.DB.NamedQuery(selectSQL+" ORDER BY erm.created_time DESC offset :offset limit :limit", queryOption)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}
	result := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})
}

// ExaminationReportModelDetail 查询检查报告医嘱模板详情
func ExaminationReportModelDetail(ctx iris.Context) {
	examinationReportModelID := ctx.PostValue("examination_report_model_id")

	if examinationReportModelID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	selectmSQL := `select * from examination_report_model where id=$1`
	mrows := model.DB.QueryRowx(selectmSQL, examinationReportModelID)
	if mrows == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询失败"})
		return
	}
	examinationReportModel := FormatSQLRowToMap(mrows)

	ctx.JSON(iris.Map{"code": "200", "data": examinationReportModel})
}

// ExaminationReportModelUpdate 修改检查报告医嘱模板
func ExaminationReportModelUpdate(ctx iris.Context) {
	examinationReportModelID := ctx.PostValue("examination_report_model_id")
	modelName := ctx.PostValue("model_name")
	resultExamination := ctx.PostValue("result_examination")
	conclusionExamination := ctx.PostValue("conclusion_examination")
	personnelID := ctx.PostValue("operation_id")

	if examinationReportModelID == "" || modelName == "" || personnelID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	mrow := model.DB.QueryRowx("select id from examination_report_model where id=$1 limit 1", examinationReportModelID)
	if mrow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	models := FormatSQLRowToMap(mrow)
	_, mok := models["id"]
	if !mok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改的模板不存在"})
		return
	}

	row := model.DB.QueryRowx("select id from examination_report_model where model_name=$1 and id!=$2 limit 1", modelName, examinationReportModelID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	examinationReportModel := FormatSQLRowToMap(row)
	_, ok := examinationReportModel["id"]
	if ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "模板名称已存在"})
		return
	}

	prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", personnelID)
	if prow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改模板失败,操作员错误"})
		return
	}
	personnel := FormatSQLRowToMap(prow)
	_, pok := personnel["id"]
	if !pok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "操作员错误"})
		return
	}

	updateSQL := `update examination_report_model set model_name=$1,result_examination=$2,conclusion_examination=$3,
	operation_id=$4,updated_time=LOCALTIMESTAMP where id=$5`
	_, err := model.DB.Exec(updateSQL, modelName, resultExamination, conclusionExamination, personnelID, examinationReportModelID)
	if err != nil {
		fmt.Println("err ===", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// ExaminationReportModelDelete 删除检查报告医嘱模板
func ExaminationReportModelDelete(ctx iris.Context) {
	examinationReportModelID := ctx.PostValue("examination_report_model_id")

	if examinationReportModelID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	mrow := model.DB.QueryRowx("select id from examination_report_model where id=$1 limit 1", examinationReportModelID)
	if mrow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "修改失败"})
		return
	}
	models := FormatSQLRowToMap(mrow)
	_, mok := models["id"]
	if !mok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "删除的模板不存在"})
		return
	}

	updateSQL := `update examination_report_model set deleted_time=LOCALTIMESTAMP where id=$1`

	_, err := model.DB.Exec(updateSQL, examinationReportModelID)
	if err != nil {
		fmt.Println("err ===", err)
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}
