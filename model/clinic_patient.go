package model

import "time"

/**
 * 诊所患者
 */
type ClinicPatient struct {
	ID            int       `db:"id"`
	PatientCertNo string    `db:"patient_cert_no"`
	ClinicCode    string    `db:"clinic_code"`
	PersonnelCode string    `db:"personnel_code"`
	Status        bool      `db:"status"`
	CreateTime    time.Time `db:"created_time"`
	UpdateTime    time.Time `db:"updated_time"`
	DeleteTime    time.Time `db:"deleted_time"`
}
