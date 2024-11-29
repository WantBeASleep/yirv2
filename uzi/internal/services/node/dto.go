package node

// TODO: починить баг при запросе со всеми полями nil
type OptionalNode struct {
	Tirads23 *float64
	Tirads4  *float64
	Tirads5  *float64
}

func (u OptionalNode) Map() map[string]any {
	res := map[string]any{}
	if u.Tirads23 != nil {
		res["tirads_23"] = *u.Tirads23
	}
	if u.Tirads4 != nil {
		res["tirads_4"] = *u.Tirads4
	}
	if u.Tirads5 != nil {
		res["tirads_5"] = *u.Tirads5
	}

	return res
}
