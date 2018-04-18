package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"strings"

	"github.com/kataras/iris"
)

// DoctorVisitScheduleModeAdd 添加排班模板
func DoctorVisitScheduleModeAdd(ctx iris.Context) {
	items := ctx.PostValue("items")
	if items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	var results []map[string]string
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	var sets []string

	for _, v := range results {

		s := "(" + v["department_id"] + "," + v["personnel_id"] + "," + v["weekday"] + "," + v["am_pm"] + "," + v["tatal_num"] + "," + v["visit_type_code"] + ")"
		sets = append(sets, s)
	}

	setStr := strings.Join(sets, ",")

	sql := "INSERT INTO doctor_visit_schedule_model( department_id, personnel_id, weekday, am_pm, tatal_num, visit_type_code ) VALUES " + setStr

	_, errs := model.DB.Query(sql)
	if errs != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": errs.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})

}

// DoctorVisitScheduleModeUpdate 更新排班模板
func DoctorVisitScheduleModeUpdate(ctx iris.Context) {

}
