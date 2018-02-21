package testprognet

import (
	"reflect"
	"runtime/testdata/testprognet"
)

var NetpollDeadlock = reflect.ValueOf(testprognet.NetpollDeadlock)
var NetpollDeadlockInit = reflect.ValueOf(testprognet.NetpollDeadlockInit)
var SignalIgnoreSIGTRAP = reflect.ValueOf(testprognet.SignalIgnoreSIGTRAP)
var SignalDuringExec = reflect.ValueOf(testprognet.SignalDuringExec)
var Nop = reflect.ValueOf(testprognet.Nop)
