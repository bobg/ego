package lzw

import (
	"compress/lzw"
	"reflect"
)

var NewWriter = reflect.ValueOf(lzw.NewWriter)
var NewReader = reflect.ValueOf(lzw.NewReader)
