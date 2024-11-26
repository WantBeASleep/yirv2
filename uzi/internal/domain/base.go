package domain

import (
	"github.com/google/uuid"
)

type Uzi struct {
	Id         uuid.UUID
	Projection string
	PatientID  uuid.UUID
	DeviceID   int
}

type Image struct {
	Id   uuid.UUID
	Page int
}

type Formation struct {
	Id       uuid.UUID
	Ai       bool
	Tirads23 float64
	Tirads4  float64
	Tirads5  float64
}

type Segment struct {
	Id          uuid.UUID
	ImageID     uuid.UUID
	FormationID uuid.UUID
	Contor      string
	Tirads23    float64
	Tirads4     float64
	Tirads5     float64
}

type Device struct {
	Id   int
	Name string
}
