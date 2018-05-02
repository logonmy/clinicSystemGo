package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

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
			strconv.Itoa(int(v["tatal_num"].(float64))) + ")"
		sets = append(sets, s)
	}

	setStr := strings.Join(sets, ",")

	sql := "INSERT INTO doctor_visit_schedule( department_id, personnel_id, visit_date, am_pm, tatal_num, left_num ) VALUES " + setStr

	row := model.DB.QueryRowx("select * FROM department_personnel WHERE type=2 and department_id=" + departmentID + " AND personnel_id=" + personnelID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "新增号源失败"})
		return
	}
	departmentPersonnel := FormatSQLRowToMap(row)

	fmt.Println("departmentPersonnel =========== ", departmentPersonnel)

	_, ok := departmentPersonnel["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "医生科室不匹配"})
		return
	}

	tx, err := model.DB.Beginx()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	_, err = tx.Exec(sql)

	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
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
	where dvs.open_flag = true and dvs.visit_date > current_date and c.id = $1 and dvs.visit_date BETWEEN date '` + startDate + `' and date '` + endDate + `'`

	if departmentID != "" {
		sql += " and dvs.department_id=" + departmentID
	}

	if personnelID != "" {
		sql += " and dvs.personnel_id=" + personnelID
	}

	rows, err := model.DB.Queryx(sql, clinicID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	schedules := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": schedules})
}

// SchelueDepartments 号源科室列表
func SchelueDepartments(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	sql := `select id as department_id, name 
	from department 
	where id in (select distinct department_id from doctor_visit_schedule 
	where clinic_id = $1 and visit_date > current_date and is_today = false and stop_flag = false and open_flag = true);`
	rows, err := model.DB.Queryx(sql, clinicID)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
	}

	departments := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": departments})
}

// SchelueDoctors 号源医生列表
func SchelueDoctors(ctx iris.Context) {
	departmentID := ctx.PostValue("department_id")
	if departmentID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	sql := `select id as personnel_id, name from personnel
	where id in (select distinct personnel_id from doctor_visit_schedule 
				where department_id = $1 and visit_date > current_date and is_today = false and stop_flag = false and open_flag = true);`
	rows, err := model.DB.Queryx(sql, departmentID)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
	}

	doctors := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": doctors})
}

// DoctorsWithSchedule 获取所有医生的号源信息
func DoctorsWithSchedule(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	departmentID := ctx.PostValue("department_id")
	personnelID := ctx.PostValue("personnel_id")
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	if clinicID == "" || startDate == "" || endDate == "" || offset == "" || limit == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	countSQL := `select count (dp.id) as total from department_personnel dp left join personnel p on p.id = dp.personnel_id where type = 2 and p.clinic_id = $1`
	if departmentID != "" {
		countSQL += " and dp.department_id = " + departmentID
	}
	if personnelID != "" {
		countSQL += " and dp.personnel_id = " + personnelID
	}
	total := model.DB.QueryRowx(countSQL, clinicID)
	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	doctorsSQL := `select p.id as personnel_id, dp.department_id, p.name as personnel_name, d.name as department_name
	from department_personnel dp
	left join department d on dp.department_id = d.id
	left join personnel p on dp.personnel_id = p.id
	where dp.type = 2 and p.clinic_id = $1 `
	if departmentID != "" {
		doctorsSQL += " and dp.department_id = " + departmentID
	}
	if personnelID != "" {
		doctorsSQL += " and dp.personnel_id = " + personnelID
	}
	rows, err := model.DB.Queryx(doctorsSQL+" offset $2 limit $3;", clinicID, offset, limit)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "查询失败，请重试"})
		return
	}
	doctors := FormatSQLRowsToMapArray(rows)

	existSQL := `select * from department_personnel idp left join personnel ip  on idp.personnel_id = ip.id where ip.clinic_id = $3 `

	if departmentID != "" {
		existSQL += " and idp.department_id = " + departmentID
	}
	if personnelID != "" {
		existSQL += " and idp.personnel_id = " + personnelID
	}
	scheduleSQL := `select dvs.id as doctor_visit_schedule_id, 
	dvs.visit_date, 
	dvs.am_pm, 
	dvs.stop_flag,
	dvs.open_flag,
	dvs.personnel_id, 
	dvs.department_id from department_personnel dp, doctor_visit_schedule dvs 
	where dp.personnel_id = dvs.personnel_id and dp.department_id = dvs.department_id 
	and dvs.visit_date BETWEEN $1 AND $2
	and exists(select 1 from (` + existSQL + ` offset $4 limit $5) ldp 
			 where ldp.personnel_id = dp.personnel_id and ldp.department_id = dp.department_id )
			 order by dp.id, dvs.visit_date, dvs.am_pm asc; `

	rows, err = model.DB.Queryx(scheduleSQL, startDate, endDate, clinicID, offset, limit)

	schedules := FormatSQLRowsToMapArray(rows)

	type Schedule struct {
		DoctorVisitScheduleID int    `json:"doctor_visit_schedule_id"`
		AmPm                  string `json:"am_pm"`
		StopFlag              bool   `json:"stop_flag"`
		OpenFlag              bool   `json:"open_flag"`
	}

	type DateSchedule struct {
		VisitDate string     `json:"visit_date"`
		Schedules []Schedule `json:"schedules"`
	}

	type Doctor struct {
		PersonnelID    int            `json:"personnel_id"`
		DepartmentID   int            `json:"department_id"`
		PersonnelName  string         `json:"personnel_name"`
		DepartmentName string         `json:"department_name"`
		DateSchedules  []DateSchedule `json:"date"`
	}

	var results []Doctor

	for _, d := range doctors {
		PersonnelID := int(d["personnel_id"].(int64))
		DepartmentID := int(d["department_id"].(int64))
		lastDate := ""
		var DateSchedules []DateSchedule
		var dateSchedule DateSchedule
		has := false
		for _, s := range schedules {
			if int(s["personnel_id"].(int64)) != PersonnelID && int(s["department_id"].(int64)) != DepartmentID {
				continue
			}
			VisitDate := s["visit_date"].(time.Time).Format("2006-01-02")
			schedule := Schedule{
				DoctorVisitScheduleID: int(s["doctor_visit_schedule_id"].(int64)),
				AmPm:     s["am_pm"].(string),
				StopFlag: s["stop_flag"].(bool),
				OpenFlag: s["open_flag"].(bool),
			}

			if lastDate == VisitDate {
				dateSchedule.Schedules = append(dateSchedule.Schedules, schedule)
				has = true
			} else {
				has = false
				if lastDate != "" {
					DateSchedules = append(DateSchedules, dateSchedule)
				}
				Ss := []Schedule{schedule}
				dateSchedule = DateSchedule{
					VisitDate: VisitDate,
					Schedules: Ss,
				}
				lastDate = VisitDate
			}
		}

		if has {
			DateSchedules = append(DateSchedules, dateSchedule)
		}

		doctor := Doctor{
			PersonnelID:    PersonnelID,
			DepartmentID:   DepartmentID,
			PersonnelName:  d["personnel_name"].(string),
			DepartmentName: d["department_name"].(string),
			DateSchedules:  DateSchedules,
		}
		results = append(results, doctor)
	}

	hasOpenCountSQL := `select count(*) as count from doctor_visit_schedule dvs left join department d on dvs.department_id = d.id
	where d.clinic_id = $1 and dvs.open_flag = true and dvs.visit_date between $2 and $3`

	hasOpenCountRow := model.DB.QueryRowx(hasOpenCountSQL, clinicID, startDate, endDate)

	hasOpenCountMap := FormatSQLRowToMap(hasOpenCountRow)

	_, ok := hasOpenCountMap["count"]

	canOverride := false
	if ok && int(hasOpenCountMap["count"].(int64)) == 0 {
		canOverride = true
	}

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo, "canOverride": canOverride})
}

// CopyScheduleByDate 复制排版
func CopyScheduleByDate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	copyStartDate := ctx.PostValue("copy_start_date")
	insertStartDate := ctx.PostValue("insert_start_date")
	dayLong := ctx.PostValue("day_long")

	if copyStartDate == "" || insertStartDate == "" || dayLong == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	long, err1 := strconv.Atoi(dayLong)

	if err1 != nil || long < 1 {
		ctx.JSON(iris.Map{"code": "-1", "msg": "day_long 必须为大于0 的数字"})
		return
	}

	copyStart, err2 := time.Parse("2006-01-02", copyStartDate)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "copy_start_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}
	insertStart, err3 := time.Parse("2006-01-02", insertStartDate)
	if err3 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "insert_start_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}

	copyEnd := copyStart.AddDate(0, 0, long-1)
	if insertStart.Before(copyEnd) {
		ctx.JSON(iris.Map{"code": "-1", "msg": "插入的结束时间不能大于复制的开始时间"})
		return
	}

	insertEnd := insertStart.AddDate(0, 0, long-1)

	hs := insertStart.Sub(copyStart).Hours()
	ds := int(hs) / 24

	tx, err := model.DB.Beginx()

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	hasOpenCountSQL := `select count(*) as count from doctor_visit_schedule dvs left join department d on dvs.department_id = d.id
	where d.clinic_id = $1 and dvs.open_flag = true and dvs.visit_date between $2 and $3`

	hasOpenCountRow := model.DB.QueryRowx(hasOpenCountSQL, clinicID, insertStart, insertEnd)

	hasOpenCountMap := FormatSQLRowToMap(hasOpenCountRow)

	_, ok := hasOpenCountMap["count"]
	if !ok || int(hasOpenCountMap["count"].(int64)) > 0 {
		ctx.JSON(iris.Map{"code": "-1", "msg": "该时间段内有启用号源，不能覆盖"})
		return
	}

	delSQL := `delete from doctor_visit_schedule where visit_date between $1 and $2 and open_flag = false 
	and exists(select 1 from (
		select * from department_personnel idp left join personnel ip  on idp.personnel_id = ip.id 
		where ip.clinic_id = $3) ldp where ldp.personnel_id = personnel_id and ldp.department_id = department_id)`

	_, err = tx.Exec(delSQL, insertStart, insertEnd, clinicID)

	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	copySQL := `insert into doctor_visit_schedule (department_id, personnel_id, visit_date, am_pm ) 
	select department_id, personnel_id, (date(visit_date + ` + strconv.Itoa(int(ds)) + ` )) as visit_date , am_pm 
	from doctor_visit_schedule dvs left join personnel p on p.id = dvs.personnel_id
	where p.clinic_id = $1 and dvs.visit_date between $2 and $3 `

	_, err = tx.Exec(copySQL, clinicID, copyStart, copyEnd)

	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "msg": "复制排版成功"})

}

// OpenScheduleByDate 开放号源
func OpenScheduleByDate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	startDate := ctx.PostValue("start_date")
	dayLong := ctx.PostValue("day_long")

	if clinicID == "" || startDate == "" || dayLong == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	long, err1 := strconv.Atoi(dayLong)

	if err1 != nil || long < 1 {
		ctx.JSON(iris.Map{"code": "-1", "msg": "day_long 必须为大于0 的数字"})
		return
	}

	openStart, err2 := time.Parse("2006-01-02", startDate)
	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "copy_start_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}
	openEnd := openStart.AddDate(0, 0, long-1)

	tx, err := model.DB.Beginx()

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	openSQL := `update doctor_visit_schedule set open_flag = true where visit_date between $1 and $2 and open_flag = false 
	and exists(select 1 from (
		select * from department_personnel idp left join personnel ip  on idp.personnel_id = ip.id 
		where ip.clinic_id = $3) ldp 
			where ldp.personnel_id = personnel_id and ldp.department_id = department_id)`

	_, err = tx.Exec(openSQL, openStart, openEnd, clinicID)

	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "msg": "开放成功"})
}

// CreateOneSchedule 插入单个号源
func CreateOneSchedule(ctx iris.Context) {
	departmentID := ctx.PostValue("department_id")
	personnelID := ctx.PostValue("personnel_id")
	visitDate := ctx.PostValue("visit_date")
	amPm := ctx.PostValue("am_pm")

	if departmentID == "" || personnelID == "" || visitDate == "" || amPm == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
		return
	}
	checkPersonnelSQL := `select count(*) as count from department_personnel where department_id = $1 and personnel_id= $2 and type = 2`

	row := model.DB.QueryRowx(checkPersonnelSQL, departmentID, personnelID)

	checkCount := FormatSQLRowToMap(row)

	_, ok := checkCount["count"]

	if !ok || int(checkCount["count"].(int64)) < 1 {
		ctx.JSON(iris.Map{"code": "-1", "msg": "科室 医生 不匹配"})
		return
	}

	tx, err := model.DB.Beginx()

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	insertSQL := `insert into doctor_visit_schedule (department_id, personnel_id, visit_date, am_pm ) values ($1, $2, $3, $4)`

	_, err = tx.Exec(insertSQL, departmentID, personnelID, visitDate, amPm)

	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "插入号源成功"})
}

// DeleteOneUnOpenScheduleByID 删除单个未开放号源 byid
func DeleteOneUnOpenScheduleByID(ctx iris.Context) {
	doctorVisitScheduleID := ctx.PostValue("doctor_visit_schedule_id")

	if doctorVisitScheduleID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
		return
	}

	querySQL := `select * from doctor_visit_schedule where id = $1`

	row := model.DB.QueryRowx(querySQL, doctorVisitScheduleID)

	scheuleMap := FormatSQLRowToMap(row)

	_, ok := scheuleMap["open_flag"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "号源不存在"})
		return
	}

	if scheuleMap["open_flag"].(bool) {
		ctx.JSON(iris.Map{"code": "-1", "msg": "已开放号源不能删除"})
		return
	}

	deleteSQL := `delete from doctor_visit_schedule where id = $1 and open_flag = false`

	_, err := model.DB.Exec(deleteSQL, doctorVisitScheduleID)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "-1", "msg": "删除号源成功"})
}

// StopScheduleByID 停诊号源byid
func StopScheduleByID(ctx iris.Context) {
	doctorVisitScheduleID := ctx.PostValue("doctor_visit_schedule_id")

	if doctorVisitScheduleID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
		return
	}

	querySQL := `select * from doctor_visit_schedule where id = $1`

	row := model.DB.QueryRowx(querySQL, doctorVisitScheduleID)

	scheuleMap := FormatSQLRowToMap(row)

	_, ok := scheuleMap["open_flag"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "号源不存在"})
		return
	}

	if !scheuleMap["open_flag"].(bool) {
		ctx.JSON(iris.Map{"code": "-1", "msg": "未开放号源不能停诊"})
		return
	}

	updateSQL := `update doctor_visit_schedule set stop_flag = true where id = $1 and open_flag = true`

	_, err := model.DB.Exec(updateSQL, doctorVisitScheduleID)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "-1", "msg": "停诊成功"})
}
