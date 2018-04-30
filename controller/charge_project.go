package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strings"

	"github.com/kataras/iris"
)

// ChargeProjectTreatmentInit 初始化诊疗项目
func ChargeProjectTreatmentInit(ctx iris.Context) {

	sql := "INSERT INTO charge_project_treatment (project_type_id, name, name_en, cost, fee) VALUES (8,'专家挂号费','ZJGHF',1000,10000),(8,'普通门诊挂号费','PTMZGHF',500,1000)"

	_, err := model.DB.Query(sql)
	if err != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

// ChargeProjectTreatmentCreate 收费项目创建
func ChargeProjectTreatmentCreate(ctx iris.Context) {

	projectTypeID := ctx.PostValue("project_type_id")
	name := ctx.PostValue("name")
	nameEn := ctx.PostValue("name_en")
	cost := ctx.PostValue("cost")
	fee := ctx.PostValue("fee")
	status := ctx.PostValue("status")
	if name == "" || projectTypeID == "" || fee == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	insert := []string{
		"project_type_id",
		"name",
		"fee",
	}

	valus := []string{
		projectTypeID,
		"'" + name + "'",
		fee,
	}

	if nameEn != "" {
		insert = append(insert, "name_en")
		valus = append(valus, "'"+nameEn+"'")
	}

	if cost != "" {
		insert = append(insert, "cost")
		valus = append(valus, cost)
	}

	if status != "" {
		insert = append(insert, "status")
		valus = append(valus, status)
	}

	nStr := "(" + strings.Join(insert, ",") + ")"
	vValus := "(" + strings.Join(valus, ",") + ")"

	sql := "INSERT INTO charge_project_treatment " + nStr + " VALUES " + vValus + " RETURNING id"

	fmt.Println(sql)

	var ID int
	err := model.DB.QueryRow(sql).Scan(&ID)
	if err != nil {
		fmt.Println("err ===", err)
		ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": ID})
}
