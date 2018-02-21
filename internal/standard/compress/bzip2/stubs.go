package bzip2

import (
	"compress/bzip2"
	"reflect"
)

var NewReader = reflect.ValueOf(bzip2.NewReader)
