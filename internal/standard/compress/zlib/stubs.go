package zlib

import (
	"compress/zlib"
	"reflect"
)

var NewWriter = reflect.ValueOf(zlib.NewWriter)
var NewWriterLevel = reflect.ValueOf(zlib.NewWriterLevel)
var NewWriterLevelDict = reflect.ValueOf(zlib.NewWriterLevelDict)
var NewReaderDict = reflect.ValueOf(zlib.NewReaderDict)
var NewReader = reflect.ValueOf(zlib.NewReader)
