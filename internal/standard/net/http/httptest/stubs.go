package httptest

import (
	"net/http/httptest"
	"reflect"
)

var NewRecorder = reflect.ValueOf(httptest.NewRecorder)
var NewRequest = reflect.ValueOf(httptest.NewRequest)
var NewUnstartedServer = reflect.ValueOf(httptest.NewUnstartedServer)
var NewTLSServer = reflect.ValueOf(httptest.NewTLSServer)
var NewServer = reflect.ValueOf(httptest.NewServer)
