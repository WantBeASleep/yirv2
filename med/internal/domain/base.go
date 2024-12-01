package domain

import (
	"time"

	"github.com/google/uuid"
)

type Doctor struct {
	Id       uuid.UUID `db:"id"`
	FullName string `db:"fullName"`
	Org      string `db:"org"`
	Job      string `db:"job"`
	Desc     string `db:"desc"`
}

type Patient struct {
	Id          uuid.UUID `db:"id"`
	FullName    string `db:"fullname"`
	Email       string `db:"email"`
	Policy      string `db:"policy"`
	Active      bool `db:"active"`
	Malignancy  bool `db:"malignancy"`
	LastUziDate *time.Time `db:"last_uzi_date"`
}

type Card struct {
	DoctorID  uuid.UUID `db:"doctor_id"`
	PatientID uuid.UUID `db:"patient_id"`
	Diagnosis string `db:"diagnosis"`
}
