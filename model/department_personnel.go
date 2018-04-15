package model

import "time"

/**
 * 诊所
 */
type DepartmentPersonnel struct {
	ID           int       `db:"id"`
	DepartmentID int       `db:"department_id"`
	PersonnelID  int       `db:"personnel_id"`
	Type         int       `db:"type"`
	Status       bool      `db:"status"`
	CreateTime   time.Time `db:"create_time"`
	UpdateTime   time.Time `db:"update_time"`
	DeleteTime   time.Time `db:"deleted_time"`
}
