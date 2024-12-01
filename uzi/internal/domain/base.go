package domain

import (
	"time"

	"github.com/google/uuid"
)

type Uzi struct {
	Id         uuid.UUID `db:"id"`
	Projection string    `db:"projection"`
	Checked    bool      `db:"checked"`
	PatientID  uuid.UUID `db:"patient_id"`
	DeviceID   int       `db:"device_id"`
	CreateAt   time.Time `db:"create_at"`
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

type Echographic struct {
	Id              uuid.UUID `db:"id"`
	Contors         string    `db:"contors"`
	LeftLobeLength  float64   `db:"left_lobe_length"`
	LeftLobeWidth   float64   `db:"left_lobe_width"`
	LeftLobeThick   float64   `db:"left_lobe_thick"`
	LeftLobeVolum   float64   `db:"left_lobe_volum"`
	RightLobeLength float64   `db:"right_lobe_length"`
	RightLobeWidth  float64   `db:"right_lobe_width"`
	RightLobeThick  float64   `db:"right_lobe_thick"`
	RightLobeVolum  float64   `db:"right_lobe_volum"`
	GlandVolum      float64   `db:"gland_volum"`
	Isthmus         float64   `db:"isthmus"`
	Struct          string    `db:"struct"`
	Echogenicity    string    `db:"echogenicity"`
	RegionalLymph   string    `db:"regional_lymph"`
	Vascularization string    `db:"vascularization"`
	Location        string    `db:"location"`
	Additional      string    `db:"additional"`
	Conclusion      string    `db:"conclusion"`
}
