package cgi

import (
	"net/http/cgi"
	"reflect"
)

var Request = reflect.ValueOf(cgi.Request)
var RequestFromMap = reflect.ValueOf(cgi.RequestFromMap)
var Serve = reflect.ValueOf(cgi.Serve)
