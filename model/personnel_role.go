package model

import "time"

/**
 * 用户角色关联
 */
type PersonnelRole struct {
	PersonnelID int       `db:"personnel_id"`
	RoleID      int       `db:"role_id"`
	Status      bool      `db:"status"`
	CreateTime  time.Time `db:"created_time"`
	UpdateTime  time.Time `db:"updated_time"`
	DeleteTime  time.Time `db:"deleted_time"`
}
