package refl

import (
	"errors"
	"reflect"
)

var errBadValue = errors.New("bad value")

type Value struct {
	R *reflect.Value
}

func (v *Value) Type() (*Type, error) {
	if v.R != nil {
		return &Type{R: v.R.Type()}, nil
	}
	return nil, errBadValue
}

func (v *Value) Bool() (bool, error) {
	if v.R == nil || v.R.Kind() != reflect.Bool {
		// xxx err
	}
	return v.R.Bool(), nil
}
