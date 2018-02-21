package format

import (
	"go/format"
	"reflect"
)

var Node = reflect.ValueOf(format.Node)
var Source = reflect.ValueOf(format.Source)
