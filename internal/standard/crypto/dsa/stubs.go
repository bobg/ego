package dsa

import (
	"crypto/dsa"
	"reflect"
)

var Verify = reflect.ValueOf(dsa.Verify)
var GenerateParameters = reflect.ValueOf(dsa.GenerateParameters)
var GenerateKey = reflect.ValueOf(dsa.GenerateKey)
var Sign = reflect.ValueOf(dsa.Sign)
