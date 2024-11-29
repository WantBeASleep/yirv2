package domain

import "github.com/google/uuid"

type Uzi struct {
	Id         uuid.UUID `db:"id"`
	Projection string    `db:"projection"`
	PatientID  uuid.UUID `db:"patient_id"`
	DeviceID   int       `db:"device_id"`
}

type Image struct {
	Id   uuid.UUID `db:"id"`
	Page int       `db:"page"`
}

type Node struct {
	Id       uuid.UUID `db:"id"`
	Ai       bool      `db:"ai"`
	Tirads23 float64   `db:"tirads_23"`
	Tirads4  float64   `db:"tirads_4"`
	Tirads5  float64   `db:"tirads_5"`
}

type Segment struct {
	Id       uuid.UUID `db:"id"`
	ImageID  uuid.UUID `db:"image_id"`
	NodeID   uuid.UUID `db:"node_id"`
	Contor   string    `db:"contor"`
	Tirads23 float64   `db:"tirads_23"`
	Tirads4  float64   `db:"tirads_4"`
	Tirads5  float64   `db:"tirads_5"`
}

type Device struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
