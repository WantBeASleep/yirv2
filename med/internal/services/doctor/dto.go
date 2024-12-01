package doctor

type OptionalDoctor struct {
	Org  *string
	Job  *string
	Desc *string
}

func (u OptionalDoctor) Map() map[string]any {
	res := map[string]any{}
	if u.Org != nil {
		res["org"] = *u.Org
	}
	if u.Job != nil {
		res["job"] = *u.Job
	}
	if u.Desc != nil {
		res["desc"] = *u.Desc
	}

	return res
}
