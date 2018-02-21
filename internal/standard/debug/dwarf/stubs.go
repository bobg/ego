package dwarf

import (
	"debug/dwarf"
	"reflect"
)

var New = reflect.ValueOf(dwarf.New)
