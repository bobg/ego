package signal

import (
	"os/signal"
	"reflect"
)

var Reset = reflect.ValueOf(signal.Reset)
var Stop = reflect.ValueOf(signal.Stop)
var Ignore = reflect.ValueOf(signal.Ignore)
var Notify = reflect.ValueOf(signal.Notify)
