package syslog

import (
	"log/syslog"
	"reflect"
)

var NewLogger = reflect.ValueOf(syslog.NewLogger)
var Dial = reflect.ValueOf(syslog.Dial)
var New = reflect.ValueOf(syslog.New)
