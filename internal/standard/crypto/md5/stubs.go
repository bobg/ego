package md5

import (
	"crypto/md5"
	"reflect"
)

var New = reflect.ValueOf(md5.New)
var Sum = reflect.ValueOf(md5.Sum)
