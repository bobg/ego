package plan9obj

import (
	"debug/plan9obj"
	"reflect"
)

var Open = reflect.ValueOf(plan9obj.Open)
var NewFile = reflect.ValueOf(plan9obj.NewFile)
