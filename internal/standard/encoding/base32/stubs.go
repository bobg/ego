package base32

import (
	"encoding/base32"
	"reflect"
)

var NewEncoding = reflect.ValueOf(base32.NewEncoding)
var NewDecoder = reflect.ValueOf(base32.NewDecoder)
var NewEncoder = reflect.ValueOf(base32.NewEncoder)
