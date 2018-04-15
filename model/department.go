package model

import "time"

/**
 * 科室
 */
type Department struct {
	ID            int       `db:"id"`
	Code          string    `db:"code"`
	Name          string    `db:"name"`
	ClinicCode    string    `db:"clinic_code"`
	Status        bool      `db:"status"`
	IsAppointment bool      `db:"is_appointment"`
	CreateTime    time.Time `db:"create_time"`
	UpdateTime    time.Time `db:"update_time"`
	DeleteTime    time.Time `db:"deleted_time"`
}
