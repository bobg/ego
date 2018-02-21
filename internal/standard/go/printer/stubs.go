package printer

import (
	"go/printer"
	"reflect"
)

var Fprint = reflect.ValueOf(printer.Fprint)
