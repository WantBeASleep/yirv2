package slicer

func Convert[T, E any](slice []T, conventer func(v T) E) []E {
	res := make([]E, 0, len(slice))
	for _, e := range slice {
		res = append(res, conventer(e))
	}

	return res
}
