package base64

import (
	"encoding/base64"
	"reflect"
)

var NewEncoding = reflect.ValueOf(base64.NewEncoding)
var NewDecoder = reflect.ValueOf(base64.NewDecoder)
var NewEncoder = reflect.ValueOf(base64.NewEncoder)
