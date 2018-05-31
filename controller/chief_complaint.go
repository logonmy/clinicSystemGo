package controller

import (
	"clinicSystemGo/model"

	"github.com/kataras/iris"
)

// ChiefComplaintList 获取主诉列表
func ChiefComplaintList(ctx iris.Context) {
	chiefComplaintList, _ := model.DB.Queryx("select * from chief_complaint")
	results := FormatSQLRowsToMapArray(chiefComplaintList)
	ctx.JSON(iris.Map{"code": "200", "data": results})
}
