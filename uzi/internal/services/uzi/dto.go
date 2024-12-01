package uzi

import "github.com/google/uuid"

type OptionalUzi struct {
	Projection *string
	Checked    *bool
	PatientID  *uuid.UUID
}

func (u OptionalUzi) Map() map[string]any {
	res := map[string]any{}
	if u.Projection != nil {
		res["projection"] = *u.Projection
	}
	if u.Checked != nil {
		res["checked"] = *u.Checked
	}
	if u.PatientID != nil {
		res["patient"] = *u.PatientID
	}

	return res
}

type OptionalEchographic struct {
	Contors         *string
	LeftLobeLength  *float64
	LeftLobeWidth   *float64
	LeftLobeThick   *float64
	LeftLobeVolum   *float64
	RightLobeLength *float64
	RightLobeWidth  *float64
	RightLobeThick  *float64
	RightLobeVolum  *float64
	GlandVolum      *float64
	Isthmus         *float64
	Struct          *string
	Echogenicity    *string
	RegionalLymph   *string
	Vascularization *string
	Location        *string
	Additional      *string
	Conclusion      *string
}

// TODO: пойти написать рефлексию
func (u OptionalEchographic) Map() map[string]any {
	res := map[string]any{}

	if u.Contors != nil {
		res["contors"] = *u.Contors
	}
	if u.LeftLobeLength != nil {
		res["left_lobe_length"] = *u.LeftLobeLength
	}
	if u.LeftLobeWidth != nil {
		res["left_lobe_width"] = *u.LeftLobeWidth
	}
	if u.LeftLobeThick != nil {
		res["left_lobe_thick"] = *u.LeftLobeThick
	}
	if u.LeftLobeVolum != nil {
		res["left_lobe_volum"] = *u.LeftLobeVolum
	}
	if u.RightLobeLength != nil {
		res["right_lobe_length"] = *u.RightLobeLength
	}
	if u.RightLobeWidth != nil {
		res["right_lobe_width"] = *u.RightLobeWidth
	}
	if u.RightLobeThick != nil {
		res["right_lobe_thick"] = *u.RightLobeThick
	}
	if u.RightLobeVolum != nil {
		res["right_lobe_volum"] = *u.RightLobeVolum
	}
	if u.GlandVolum != nil {
		res["gland_volum"] = *u.GlandVolum
	}
	if u.Isthmus != nil {
		res["isthmus"] = *u.Isthmus
	}
	if u.Struct != nil {
		res["struct"] = *u.Struct
	}
	if u.Echogenicity != nil {
		res["echogenicity"] = *u.Echogenicity
	}
	if u.RegionalLymph != nil {
		res["regional_lymph"] = *u.RegionalLymph
	}
	if u.Vascularization != nil {
		res["vascularization"] = *u.Vascularization
	}
	if u.Location != nil {
		res["location"] = *u.Location
	}
	if u.Additional != nil {
		res["additional"] = *u.Additional
	}
	if u.Conclusion != nil {
		res["conclusion"] = *u.Conclusion
	}

	return res
}
