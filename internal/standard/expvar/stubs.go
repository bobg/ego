package expvar

import (
	"expvar"
	"reflect"
)

var Publish = reflect.ValueOf(expvar.Publish)
var NewMap = reflect.ValueOf(expvar.NewMap)
var Do = reflect.ValueOf(expvar.Do)
var Handler = reflect.ValueOf(expvar.Handler)
var NewInt = reflect.ValueOf(expvar.NewInt)
var NewFloat = reflect.ValueOf(expvar.NewFloat)
var Get = reflect.ValueOf(expvar.Get)
var NewString = reflect.ValueOf(expvar.NewString)
