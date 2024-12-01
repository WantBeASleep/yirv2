package patient

import "time"

type OptionalPatient struct {
	Active      *bool
	Malignancy  *bool
	LastUziDate *time.Time
}

func (u OptionalPatient) Map() map[string]any {
	res := map[string]any{}
	if u.Active != nil {
		res["active"] = *u.Active
	}
	if u.Malignancy != nil {
		res["malignancy"] = *u.Malignancy
	}
	if u.LastUziDate != nil {
		res["last_uzi_date"] = *u.LastUziDate
	}

	return res
}
