package adler32

import (
	"hash/adler32"
	"reflect"
)

var Checksum = reflect.ValueOf(adler32.Checksum)
var New = reflect.ValueOf(adler32.New)
