package patient

import "errors"

var (
	ErrNoPermission = errors.New("doesnt have permission for action")
)