package iotest

import (
	"reflect"
	"testing/iotest"
)

var TruncateWriter = reflect.ValueOf(iotest.TruncateWriter)
var DataErrReader = reflect.ValueOf(iotest.DataErrReader)
var OneByteReader = reflect.ValueOf(iotest.OneByteReader)
var HalfReader = reflect.ValueOf(iotest.HalfReader)
var TimeoutReader = reflect.ValueOf(iotest.TimeoutReader)
var NewWriteLogger = reflect.ValueOf(iotest.NewWriteLogger)
var NewReadLogger = reflect.ValueOf(iotest.NewReadLogger)
