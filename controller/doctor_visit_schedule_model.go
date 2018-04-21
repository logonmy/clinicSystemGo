package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
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

	var departmentID string
	var personnelID string

	var sets []string

	for _, v := range results {
		departmentID = v["department_id"]
		personnelID = v["personnel_id"]
		s := "(" + v["department_id"] + "," + v["personnel_id"] + "," + v["weekday"] + ",'" + v["am_pm"] + "'," + v["tatal_num"] + "," + v["visit_type_code"] + ")"
		sets = append(sets, s)
	}

	setStr := strings.Join(sets, ",")

	sql1 := "DELETE FROM doctor_visit_schedule_model WHERE department_id=" + departmentID + " AND personnel_id=" + personnelID
	sql := "INSERT INTO doctor_visit_schedule_model( department_id, personnel_id, weekday, am_pm, tatal_num, visit_type_code ) VALUES " + setStr
	fmt.Println(sql1, sql)

	tx, err := model.DB.Beginx()

	if err != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
		return
	}

	_, errtx := tx.Query(sql1)
	if errtx != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errtx.Error()})
		return
	}

	_, errtx2 := tx.Query(sql)
	if errtx2 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errtx2.Error()})
		return
	}

	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})

}

// DoctorVisitScheduleModeUpdate 更新排班模板
func DoctorVisitScheduleModeUpdate(ctx iris.Context) {

}
