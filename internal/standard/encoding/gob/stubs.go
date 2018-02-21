package gob

import (
	"encoding/gob"
	"reflect"
)

var Debug = reflect.ValueOf(gob.Debug)
var NewEncoder = reflect.ValueOf(gob.NewEncoder)
var NewDecoder = reflect.ValueOf(gob.NewDecoder)
var RegisterName = reflect.ValueOf(gob.RegisterName)
var Register = reflect.ValueOf(gob.Register)
