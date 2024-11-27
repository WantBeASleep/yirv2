package uzi

import "github.com/google/uuid"

type OptionalUzi struct {
	Projection *string
	PatientID  *uuid.UUID
}

func (u OptionalUzi) Map() map[string]any {
	res := map[string]any{}
	if u.Projection != nil {
		res["projection"] = *u.Projection
	}
	if u.PatientID != nil {
		res["patient"] = *u.PatientID
	}

	return res
}
