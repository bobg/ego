package trace

import (
	"reflect"
	"runtime/trace"
)

var Start = reflect.ValueOf(trace.Start)
var Stop = reflect.ValueOf(trace.Stop)
