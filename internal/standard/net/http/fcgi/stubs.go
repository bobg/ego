package fcgi

import (
	"net/http/fcgi"
	"reflect"
)

var ProcessEnv = reflect.ValueOf(fcgi.ProcessEnv)
var Serve = reflect.ValueOf(fcgi.Serve)
