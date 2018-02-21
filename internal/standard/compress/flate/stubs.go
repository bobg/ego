package flate

import (
	"compress/flate"
	"reflect"
)

var NewWriterDict = reflect.ValueOf(flate.NewWriterDict)
var NewWriter = reflect.ValueOf(flate.NewWriter)
var NewReaderDict = reflect.ValueOf(flate.NewReaderDict)
var NewReader = reflect.ValueOf(flate.NewReader)
