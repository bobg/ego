package gzip

import (
	"compress/gzip"
	"reflect"
)

var NewWriter = reflect.ValueOf(gzip.NewWriter)
var NewWriterLevel = reflect.ValueOf(gzip.NewWriterLevel)
var NewReader = reflect.ValueOf(gzip.NewReader)
