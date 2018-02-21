package macho

import (
	"debug/macho"
	"reflect"
)

var NewFile = reflect.ValueOf(macho.NewFile)
var Open = reflect.ValueOf(macho.Open)
var NewFatFile = reflect.ValueOf(macho.NewFatFile)
var OpenFat = reflect.ValueOf(macho.OpenFat)
