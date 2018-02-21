package pe

import (
	"debug/pe"
	"reflect"
)

var NewFile = reflect.ValueOf(pe.NewFile)
var Open = reflect.ValueOf(pe.Open)
