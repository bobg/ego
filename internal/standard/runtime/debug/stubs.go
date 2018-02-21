package debug

import (
	"reflect"
	"runtime/debug"
)

var ReadGCStats = reflect.ValueOf(debug.ReadGCStats)
var SetMaxThreads = reflect.ValueOf(debug.SetMaxThreads)
var WriteHeapDump = reflect.ValueOf(debug.WriteHeapDump)
var SetGCPercent = reflect.ValueOf(debug.SetGCPercent)
var FreeOSMemory = reflect.ValueOf(debug.FreeOSMemory)
var SetMaxStack = reflect.ValueOf(debug.SetMaxStack)
var SetPanicOnFault = reflect.ValueOf(debug.SetPanicOnFault)
var SetTraceback = reflect.ValueOf(debug.SetTraceback)
var Stack = reflect.ValueOf(debug.Stack)
var PrintStack = reflect.ValueOf(debug.PrintStack)
