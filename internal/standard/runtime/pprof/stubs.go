package pprof

import (
	"reflect"
	"runtime/pprof"
)

var Lookup = reflect.ValueOf(pprof.Lookup)
var WriteHeapProfile = reflect.ValueOf(pprof.WriteHeapProfile)
var StartCPUProfile = reflect.ValueOf(pprof.StartCPUProfile)
var NewProfile = reflect.ValueOf(pprof.NewProfile)
var Profiles = reflect.ValueOf(pprof.Profiles)
var StopCPUProfile = reflect.ValueOf(pprof.StopCPUProfile)
var WithLabels = reflect.ValueOf(pprof.WithLabels)
var Labels = reflect.ValueOf(pprof.Labels)
var Label = reflect.ValueOf(pprof.Label)
var ForLabels = reflect.ValueOf(pprof.ForLabels)
var SetGoroutineLabels = reflect.ValueOf(pprof.SetGoroutineLabels)
var Do = reflect.ValueOf(pprof.Do)
