package testprogcgo

import (
	"reflect"
	"runtime/testdata/testprogcgo"
)

var CgoRaceprof = reflect.ValueOf(testprogcgo.CgoRaceprof)
var CgoPprofThread = reflect.ValueOf(testprogcgo.CgoPprofThread)
var CgoPprofThreadNoTraceback = reflect.ValueOf(testprogcgo.CgoPprofThreadNoTraceback)
var GoNop = reflect.ValueOf(testprogcgo.GoNop)
var CgoCCodeSIGPROF = reflect.ValueOf(testprogcgo.CgoCCodeSIGPROF)
var CgoCallbackGC = reflect.ValueOf(testprogcgo.CgoCallbackGC)
var GoCheckM = reflect.ValueOf(testprogcgo.GoCheckM)
var EnsureDropM = reflect.ValueOf(testprogcgo.EnsureDropM)
var CgoPprof = reflect.ValueOf(testprogcgo.CgoPprof)
var NumGoroutine = reflect.ValueOf(testprogcgo.NumGoroutine)
var CallbackNumGoroutine = reflect.ValueOf(testprogcgo.CallbackNumGoroutine)
var CgoRaceSignal = reflect.ValueOf(testprogcgo.CgoRaceSignal)
var CgoCheckBytes = reflect.ValueOf(testprogcgo.CgoCheckBytes)
var CgoSignalDeadlock = reflect.ValueOf(testprogcgo.CgoSignalDeadlock)
var CgoTraceback = reflect.ValueOf(testprogcgo.CgoTraceback)
var CrashTraceback = reflect.ValueOf(testprogcgo.CrashTraceback)
var CgoPanicDeadlock = reflect.ValueOf(testprogcgo.CgoPanicDeadlock)
var CgoExternalThreadPanic = reflect.ValueOf(testprogcgo.CgoExternalThreadPanic)
var Crash = reflect.ValueOf(testprogcgo.Crash)
var CgoExternalThreadSIGPROF = reflect.ValueOf(testprogcgo.CgoExternalThreadSIGPROF)
var CgoExternalThreadSignal = reflect.ValueOf(testprogcgo.CgoExternalThreadSignal)
var G2 = reflect.ValueOf(testprogcgo.G2)
var TracebackContext = reflect.ValueOf(testprogcgo.TracebackContext)
var G1 = reflect.ValueOf(testprogcgo.G1)
var CgoExecSignalMask = reflect.ValueOf(testprogcgo.CgoExecSignalMask)
var CgoDLLImportsMain = reflect.ValueOf(testprogcgo.CgoDLLImportsMain)
