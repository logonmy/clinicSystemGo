package model

import (
	"time"
	"database/sql"
)

/**
 * 诊所
 */
type Clinic struct {
	Code              string    `db:"code"`
	Name              string    `db:"name"`
	ResponsiblePerson string    `db:"responsible_person"`
	Area              string    `db:"area"`
	Status            bool      `db:"status"`
	CreateTime        time.Time `db:"created_time"`
	UpdateTime        time.Time `db:"updated_time"`
	DeleteTime        sql.NullString `db:"deleted_time"`
}
