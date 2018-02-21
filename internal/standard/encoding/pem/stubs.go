package pem

import (
	"encoding/pem"
	"reflect"
)

var EncodeToMemory = reflect.ValueOf(pem.EncodeToMemory)
var Encode = reflect.ValueOf(pem.Encode)
var Decode = reflect.ValueOf(pem.Decode)
