package ecdsa

import (
	"crypto/ecdsa"
	"reflect"
)

var Verify = reflect.ValueOf(ecdsa.Verify)
var GenerateKey = reflect.ValueOf(ecdsa.GenerateKey)
var Sign = reflect.ValueOf(ecdsa.Sign)
