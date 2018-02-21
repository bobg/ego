package importer

import (
	"go/importer"
	"reflect"
)

var For = reflect.ValueOf(importer.For)
var Default = reflect.ValueOf(importer.Default)
