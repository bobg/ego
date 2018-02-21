package fnv

import (
	"hash/fnv"
	"reflect"
)

var New64 = reflect.ValueOf(fnv.New64)
var New32 = reflect.ValueOf(fnv.New32)
var New128 = reflect.ValueOf(fnv.New128)
var New128a = reflect.ValueOf(fnv.New128a)
var New32a = reflect.ValueOf(fnv.New32a)
var New64a = reflect.ValueOf(fnv.New64a)
