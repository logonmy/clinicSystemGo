package model

import "time"

/**
 * 出诊类型
 */
type VisitType struct {
	Code              int    		`db:"code"`
	Name              string    `db:"name"`
	OpenFlag 					bool    `db:"responsible_person"`
	Fee               int    		`db:"fee"`
	CreateTime        time.Time `db:"created_time"`
	UpdateTime        time.Time `db:"updated_time"`
	DeleteTime        time.Time `db:"deleted_time"`
}