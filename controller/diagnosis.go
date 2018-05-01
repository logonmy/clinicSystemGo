package controller

import (
	"clinicSystemGo/model"
	"fmt"

	"github.com/kataras/iris"
)

// DiagnosisCreate 创建诊断
func DiagnosisCreate(ctx iris.Context) {
	pyCode := ctx.PostValue("py_code")
	name := ctx.PostValue("name")
	icdCode := ctx.PostValue("icd_code")

	if name == "" || pyCode == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	var id int
	err := model.DB.QueryRow("INSERT INTO diagnosis (py_code, name, icd_code) VALUES ($1, $2, $3) RETURNING id", pyCode, name, icdCode).Scan(&id)
	if err != nil {
		fmt.Println("err ===", err)
		ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": id})

}
