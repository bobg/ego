package rc4

import (
	"crypto/rc4"
	"reflect"
)

var NewCipher = reflect.ValueOf(rc4.NewCipher)
