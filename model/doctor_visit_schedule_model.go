package model

import "time"

/**
 * 医生排版模板
 */
type DoctorVisitScheduleModel struct {
	ID            int       `db:"id"`
	DepartmentID  int       `db:"department_id"`
	PersonnelID   int       `db:"personnel_id"`
	Weekday       int       `db:"weekday"`
	AmPm          string    `db:"am_pm"`
	TatalNum      int       `db:"tatal_num"`
	VisitTypeCode int       `db:"visit_type_code"`
	CreateTime    time.Time `db:"create_time"`
	UpdateTime    time.Time `db:"update_time"`
	DeleteTime    time.Time `db:"deleted_time"`
}
