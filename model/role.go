package model

import "time"

/**
 * 用户角色
 */
type Role struct {
	ID              int    		  `db:"id"`
	Name            string      `db:"name"`
	Status          bool    		`db:"status"`
	CreateTime      time.Time   `db:"created_time"`
	UpdateTime      time.Time   `db:"updated_time"`
	DeleteTime      time.Time   `db:"deleted_time"`
}