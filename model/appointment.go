package model

import "time"

/**
 * 挂号记录
 */
type Appointment struct {
	ID              int       `db:"id"`
	ClinicPatientID int       `db:"clinic_patient_id"`
	DepartmentID    int       `db:"department_id"`
	PersonnelID     int       `db:"personnel_id"`
	VisitDate       time.Time `db:"visit_date"`
	AmPm            string    `db:"am_pm"`
	IsToday         bool      `db:"is_today"`
	VisitTypeCode   int       `db:"visit_type_code"`
	Status          bool      `db:"status"`
	VisitPlace      string    `db:"visit_place"`
	SortNo          int       `db:"sort_no"`
	OperationID     int       `db:"operation_id"`
	CreateTime      time.Time `db:"create_time"`
	UpdateTime      time.Time `db:"update_time"`
	DeleteTime      time.Time `db:"deleted_time"`
}
