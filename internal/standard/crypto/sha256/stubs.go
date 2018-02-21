package sha256

import (
	"crypto/sha256"
	"reflect"
)

var Sum224 = reflect.ValueOf(sha256.Sum224)
var New = reflect.ValueOf(sha256.New)
var New224 = reflect.ValueOf(sha256.New224)
var Sum256 = reflect.ValueOf(sha256.Sum256)
