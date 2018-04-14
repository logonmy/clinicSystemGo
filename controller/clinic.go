package controller

import (
	"clinicSystemGo/model"
	"fmt"

	"github.com/kataras/iris"
)

/**
 * 获取诊所
 */
func GetClinicByCode(ctx iris.Context) {
	code := ctx.PostValue("code")
	fmt.Println("code ======= ", code)
	if code != "" {
		clinic := model.Clinic{}
		err := model.DB.Get(&clinic, "SELECT * FROM clinic WHERE code=$1", code)

		if err != nil {
			fmt.Println("err ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": "error"})
			return
		}
		ctx.JSON(iris.Map{"code": "200", "data": clinic})
	}
}
