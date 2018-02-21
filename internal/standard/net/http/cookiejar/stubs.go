package cookiejar

import (
	"net/http/cookiejar"
	"reflect"
)

var New = reflect.ValueOf(cookiejar.New)
