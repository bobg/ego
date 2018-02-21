package aes

import (
	"crypto/aes"
	"reflect"
)

var NewCipher = reflect.ValueOf(aes.NewCipher)
