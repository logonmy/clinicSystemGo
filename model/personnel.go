package model

import "time"

/**
 * 人员
 */
type Personnel struct {
	ID            int       `db:"id"`
	Code          string    `db:"code"`
	Name          string    `db:"name"`
	ClinicCode    string    `db:"clinic_code"`
	Weight        int       `db:"weight"`
	Title         string    `db:"title"`
	UserName      string    `db:"username"`
	Password      string    `db:"password"`
	IsClinicAdmin bool      `db:"is_clinic_admin"`
	status        bool      `db:"status"`
	CreateTime    time.Time `db:"created_time"`
	UpdateTime    time.Time `db:"updated_time"`
	DeleteTime    time.Time `db:"deleted_time"`
}
