package pprof

import (
	"net/http/pprof"
	"reflect"
)

var Handler = reflect.ValueOf(pprof.Handler)
var Index = reflect.ValueOf(pprof.Index)
var Cmdline = reflect.ValueOf(pprof.Cmdline)
var Profile = reflect.ValueOf(pprof.Profile)
var Trace = reflect.ValueOf(pprof.Trace)
var Symbol = reflect.ValueOf(pprof.Symbol)
