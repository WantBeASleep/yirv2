package gtclib

type _string struct{}

var String _string

func (_string) ValueToPointer(p string) *string {
	if p == "" {
		return nil
	}

	return &p
}
