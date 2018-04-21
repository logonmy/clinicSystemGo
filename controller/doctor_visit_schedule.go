package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

// DoctorVistScheduleCreate 新增号源
func DoctorVistScheduleCreate(ctx iris.Context) {
	items := ctx.PostValue("items")
	departmentID := ctx.PostValue("department_id")
	personnelID := ctx.PostValue("personnel_id")
	if items == "" || departmentID == "" || personnelID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	var results []map[string]interface{}
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	var sets []string
	for _, v := range results {
		s := "(" + departmentID + "," + personnelID + ", date '" + v["visit_date"].(string) + "' ,'" + v["am_pm"].(string) + "'," +
			strconv.Itoa(int(v["tatal_num"].(float64))) + "," +
			strconv.Itoa(int(v["tatal_num"].(float64))) + "," +
			strconv.Itoa(int(v["visit_type_code"].(float64))) + ")"
		sets = append(sets, s)
	}

	setStr := strings.Join(sets, ",")

	sql := "INSERT INTO doctor_visit_schedule( department_id, personnel_id, visit_date, am_pm, tatal_num, left_num, visit_type_code ) VALUES " + setStr

	row := model.DB.QueryRowx("select * FROM department_personnel WHERE type=2 and department_id=" + departmentID + " AND personnel_id=" + personnelID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "新增号源失败"})
		return
	}
	departmentPersonnel := FormatSQLRowToMap(row)

	fmt.Println("departmentPersonnel =========== ", departmentPersonnel)

	_, ok := departmentPersonnel["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "医生科室不匹配"})
		return
	}

	tx, err := model.DB.Beginx()
	if err != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
		return
	}

	_, err = tx.Exec(sql)

	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "插入成功"})
}

// DoctorVistScheduleList 获取号源列表
func DoctorVistScheduleList(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	departmentID := ctx.PostValue("department_id")
	personnelID := ctx.PostValue("personnel_id")
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")

	if clinicID == "" || startDate == "" || endDate == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	sql := `select dvs.id, dvs.visit_date, dvs.am_pm, dvs.tatal_num, dvs.left_num,
	d.id as department_id, d.name as department_name,
	p.id as personnel_id, p.name as personnel_name
	from doctor_visit_schedule dvs
	left join department d on dvs.department_id = d.id
	left join personnel p on dvs.personnel_id = p.id
	left join clinic c on p.clinic_id = c.id
	where dvs.visit_date > current_date and c.id = $1 and visit_date BETWEEN time $2 + integer '-1' and $3 + integer '1'`

	if departmentID != "" {
		sql += " and dvs.department_id=" + departmentID
	}

	if personnelID != "" {
		sql += " and dvs.personnel_id=" + personnelID
	}

	rows, err := model.DB.Queryx(sql, clinicID, startDate, endDate)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	schedules := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": schedules})
}
