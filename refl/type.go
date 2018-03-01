package refl

import "reflect"

type Type struct {
	R reflect.Type
}

func (t Type) AssignableTo(other Type) bool {
	if t.R != nil && other.R != nil {
		return t.R.AssignableTo(other.R)
	}
	// xxx other cases
	return false
}
