package refl

import "reflect"

type Value struct {
	R reflect.Value
}

func (v Value) Type() Type {
	if v.R != nil {
		return Type{R: v.R.Type()}
	}
	// xxx err
}
