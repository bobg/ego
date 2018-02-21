package des

import (
	"crypto/des"
	"reflect"
)

var NewCipher = reflect.ValueOf(des.NewCipher)
var NewTripleDESCipher = reflect.ValueOf(des.NewTripleDESCipher)
