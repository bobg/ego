package httptrace

import (
	"net/http/httptrace"
	"reflect"
)

var ContextClientTrace = reflect.ValueOf(httptrace.ContextClientTrace)
var WithClientTrace = reflect.ValueOf(httptrace.WithClientTrace)
