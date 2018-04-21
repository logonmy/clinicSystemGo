package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strconv"

	"github.com/kataras/iris"
)

// AppointmentCreate 预约预约
func AppointmentCreate(ctx iris.Context) {
	certNo := ctx.PostValue("cert_no")
	name := ctx.PostValue("name")
	birthday := ctx.PostValue("birthday")
	sex := ctx.PostValue("sex")
	phone := ctx.PostValue("phone")
	address := ctx.PostValue("address")
	profession := ctx.PostValue("profession")
	remark := ctx.PostValue("remark")
	patientChannelID := ctx.PostValue("patient_channel_id")
	clinicID := ctx.PostValue("clinic_id")
	doctorVisitScheduleID := ctx.PostValue("doctor_visit_schedule_id")
	personnelID := ctx.PostValue("personnel_id")
	if certNo == "" || name == "" || birthday == "" || sex == "" || phone == "" || patientChannelID == "" || clinicID == "" || personnelID == "" || doctorVisitScheduleID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	row := model.DB.QueryRowx("select * from doctor_visit_schedule where id=$1 and visit_date > CURRENT_DATE", doctorVisitScheduleID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "预约失败"})
		return
	}
	schedule := FormatSQLRowToMap(row)
	_, ok := schedule["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "号源不存在"})
		return
	}
	row = model.DB.QueryRowx("select * from patient where cert_no = $1", certNo)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "预约失败"})
		return
	}
	tx, err := model.DB.Begin()
	patient := FormatSQLRowToMap(row)
	_, ok = patient["id"]
	patientID := patient["id"]
	if !ok {
		err = tx.QueryRow(`INSERT INTO patient (
		cert_no, name, birthday, sex, phone, address, profession, remark, patient_channel_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`, certNo, name, birthday, sex, phone, address, profession, remark, patientChannelID).Scan(&patientID)
		if err != nil {
			tx.Rollback()
			fmt.Println("err2 ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
	} else {
		_, err = tx.Exec(`update patient set cert_no = $1, name= $2,birthday=$3,sex=$4, phone=$5, address=$6,profession = $7,remark= $8 ,patient_channel_id = $9  where id = $10`, certNo, name, birthday, sex, phone, address, profession, remark, patientChannelID, patientID)
		if err != nil {
			tx.Rollback()
			fmt.Println("err2 ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
	}

	row = model.DB.QueryRowx("select * from clinic_patient where patient_id= $1 and clinic_id = $2", patientID, clinicID)
	if row == nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": "预约失败"})
		return
	}
	clinicPatient := FormatSQLRowToMap(row)
	_, ok = clinicPatient["id"]
	var clinicPatientID interface{}
	if !ok {
		err = tx.QueryRow("INSERT INTO clinic_patient (patient_id, clinic_id, personnel_id) VALUES ($1, $2, $3) RETURNING id", patientID, clinicID, personnelID).Scan(&clinicPatientID)
		if err != nil {
			fmt.Println("clinic_patient ======", err)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
	} else {
		clinicPatientID = clinicPatient["id"]
	}

	var resultID int
	departmentID := schedule["department_id"]
	err = tx.QueryRow("INSERT INTO clinic_triage_patient (department_id, clinic_patient_id, register_personnel_id,register_type) VALUES ($1, $2, $3,1) RETURNING id", departmentID, clinicPatientID, personnelID).Scan(&resultID)
	if err != nil {
		fmt.Println("clinic_triage_patient ======", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": nil})
}

// AppointmentList 预约记录列表
func AppointmentList(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	personnelID := ctx.PostValue("personnel_id")
	deparmentID := ctx.PostValue("department_id")
	keyword := ctx.PostValue("keyword")
	startDate := ctx.PostValue("startDate")
	endDate := ctx.PostValue("endDate")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
		return
	}
	if startDate != "" && endDate == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请选择结束日期"})
		return
	}
	if startDate == "" && endDate != "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请选择开始日期"})
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

	registartionSQL := `from appointment a 
	left join department d on a.department_id = d.id 
	left join personnel ps on a.personnel_id = ps.id 
	left join clinic_patient cp on a.clinic_patient_id = cp.id 
	left join patient p on cp.patient_id = p.id 
	where ps.clinic_id = $1`

	if deparmentID != "" {
		registartionSQL += " and a.department_id=" + deparmentID
	}
	if personnelID != "" {
		registartionSQL += " and a.department_id=" + personnelID
	}
	if keyword != "" {
		registartionSQL += " and (p.cert_no like '%" + keyword + "%' or p.name like '%" + keyword + "%' or p.phone like '%" + keyword + "%')"
	}

	if startDate != "" && endDate != "" {
		if startDate > endDate {
			ctx.JSON(iris.Map{"code": "-1", "msg": "开始日期必须大于结束日期"})
			return
		}
		registartionSQL += " and a.created_time between date'" + startDate + "' - integer '1' and '" + endDate + "' + integer '1'"
	}

	total := model.DB.QueryRowx(`select count(a.id) as total `+registartionSQL, clinicID)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rowSQL := `select ps.name as doctor_name, p.sex, p.birthday, d.name as department_name, a.id ` + registartionSQL + " offset $2 limit $3"

	rows, err1 := model.DB.Queryx(rowSQL, clinicID, offset, limit)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})

}
