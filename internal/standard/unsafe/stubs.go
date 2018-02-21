package unsafe

import (
	"reflect"
	"unsafe"
)

var Sizeof = reflect.ValueOf(unsafe.Sizeof)
var Offsetof = reflect.ValueOf(unsafe.Offsetof)
var Alignof = reflect.ValueOf(unsafe.Alignof)
