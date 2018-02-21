package gosym

import (
	"debug/gosym"
	"reflect"
)

var NewLineTable = reflect.ValueOf(gosym.NewLineTable)
var NewTable = reflect.ValueOf(gosym.NewTable)
