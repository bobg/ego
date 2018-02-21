package ascii85

import (
	"encoding/ascii85"
	"reflect"
)

var Decode = reflect.ValueOf(ascii85.Decode)
var NewDecoder = reflect.ValueOf(ascii85.NewDecoder)
var Encode = reflect.ValueOf(ascii85.Encode)
var MaxEncodedLen = reflect.ValueOf(ascii85.MaxEncodedLen)
var NewEncoder = reflect.ValueOf(ascii85.NewEncoder)
