package rand

import (
	"crypto/rand"
	"reflect"
)

var Prime = reflect.ValueOf(rand.Prime)
var Int = reflect.ValueOf(rand.Int)
var Read = reflect.ValueOf(rand.Read)
