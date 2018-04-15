package model

import (
	"database/sql"
	"time"
)

/**
 * 科室
 */
type Department struct {
	ID            int            `db:"id"`
	Code          string         `db:"code"`
	Name          string         `db:"name"`
	ClinicCode    string         `db:"clinic_code"`
	Status        bool           `db:"status"`
	Weight        int            `db:"weight"`
	IsAppointment bool           `db:"is_appointment"`
	CreateTime    time.Time      `db:"created_time"`
	UpdateTime    time.Time      `db:"updated_time"`
	DeleteTime    sql.NullString `db:"deleted_time"`
}
