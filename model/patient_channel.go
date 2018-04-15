package model

import "time"

/**
 * 就诊人来源
 */
type PatientChannel struct {
	ID         int       `db:"id"`
	Name       string    `db:"name"`
	Status     bool      `db:"status"`
	CreateTime time.Time `db:"created_time"`
	UpdateTime time.Time `db:"updated_time"`
	DeleteTime time.Time `db:"deleted_time"`
}
