package sha1

import (
	"crypto/sha1"
	"reflect"
)

var New = reflect.ValueOf(sha1.New)
var Sum = reflect.ValueOf(sha1.Sum)
