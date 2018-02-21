package rsa

import (
	"crypto/rsa"
	"reflect"
)

var SignPSS = reflect.ValueOf(rsa.SignPSS)
var VerifyPSS = reflect.ValueOf(rsa.VerifyPSS)
var GenerateMultiPrimeKey = reflect.ValueOf(rsa.GenerateMultiPrimeKey)
var GenerateKey = reflect.ValueOf(rsa.GenerateKey)
var DecryptOAEP = reflect.ValueOf(rsa.DecryptOAEP)
var EncryptOAEP = reflect.ValueOf(rsa.EncryptOAEP)
var SignPKCS1v15 = reflect.ValueOf(rsa.SignPKCS1v15)
var VerifyPKCS1v15 = reflect.ValueOf(rsa.VerifyPKCS1v15)
var EncryptPKCS1v15 = reflect.ValueOf(rsa.EncryptPKCS1v15)
var DecryptPKCS1v15 = reflect.ValueOf(rsa.DecryptPKCS1v15)
var DecryptPKCS1v15SessionKey = reflect.ValueOf(rsa.DecryptPKCS1v15SessionKey)
