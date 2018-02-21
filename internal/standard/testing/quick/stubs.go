package quick

import (
	"reflect"
	"testing/quick"
)

var Check = reflect.ValueOf(quick.Check)
var CheckEqual = reflect.ValueOf(quick.CheckEqual)
var Value = reflect.ValueOf(quick.Value)
