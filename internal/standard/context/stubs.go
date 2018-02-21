package context

import (
	"context"
	"reflect"
)

var WithDeadline = reflect.ValueOf(context.WithDeadline)
var WithTimeout = reflect.ValueOf(context.WithTimeout)
var Background = reflect.ValueOf(context.Background)
var TODO = reflect.ValueOf(context.TODO)
var WithCancel = reflect.ValueOf(context.WithCancel)
var WithValue = reflect.ValueOf(context.WithValue)
