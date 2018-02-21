package sha512

import (
	"crypto/sha512"
	"reflect"
)

var New = reflect.ValueOf(sha512.New)
var Sum512 = reflect.ValueOf(sha512.Sum512)
var Sum384 = reflect.ValueOf(sha512.Sum384)
var Sum512_224 = reflect.ValueOf(sha512.Sum512_224)
var New512_256 = reflect.ValueOf(sha512.New512_256)
var New384 = reflect.ValueOf(sha512.New384)
var New512_224 = reflect.ValueOf(sha512.New512_224)
var Sum512_256 = reflect.ValueOf(sha512.Sum512_256)
