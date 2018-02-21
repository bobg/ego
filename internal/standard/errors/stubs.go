package errors

import (
	"errors"
	"reflect"
)

var New = reflect.ValueOf(errors.New)
