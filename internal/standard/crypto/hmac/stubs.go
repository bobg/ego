package hmac

import (
	"crypto/hmac"
	"reflect"
)

var New = reflect.ValueOf(hmac.New)
var Equal = reflect.ValueOf(hmac.Equal)
