package model

import "time"

/**
 * 就诊人
 */
type Patient struct {
	CertNo           string    `db:"cert_no"`
	Name             string    `db:"name"`
	Birthday         string    `db:"birthday"`
	Sex              int       `db:"sex"`
	Phone            string    `db:"phone"`
	PatientChannelID int       `db:"patient_channel_id"`
	Address          string    `db:"address"`
	Profession       string    `db:"profession"`
	Remark           string    `db:"remark"`
	Status           bool      `db:"status"`
	CreateTime       time.Time `db:"created_time"`
	UpdateTime       time.Time `db:"updated_time"`
	DeleteTime       time.Time `dlb:"deleted_time"`
}
