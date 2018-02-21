package cipher

import (
	"crypto/cipher"
	"reflect"
)

var NewCFBEncrypter = reflect.ValueOf(cipher.NewCFBEncrypter)
var NewCFBDecrypter = reflect.ValueOf(cipher.NewCFBDecrypter)
var NewOFB = reflect.ValueOf(cipher.NewOFB)
var NewGCM = reflect.ValueOf(cipher.NewGCM)
var NewGCMWithNonceSize = reflect.ValueOf(cipher.NewGCMWithNonceSize)
var NewCTR = reflect.ValueOf(cipher.NewCTR)
var NewCBCEncrypter = reflect.ValueOf(cipher.NewCBCEncrypter)
var NewCBCDecrypter = reflect.ValueOf(cipher.NewCBCDecrypter)
