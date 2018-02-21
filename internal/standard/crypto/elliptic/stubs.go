package elliptic

import (
	"crypto/elliptic"
	"reflect"
)

var P224 = reflect.ValueOf(elliptic.P224)
var Unmarshal = reflect.ValueOf(elliptic.Unmarshal)
var P384 = reflect.ValueOf(elliptic.P384)
var GenerateKey = reflect.ValueOf(elliptic.GenerateKey)
var Marshal = reflect.ValueOf(elliptic.Marshal)
var P256 = reflect.ValueOf(elliptic.P256)
var P521 = reflect.ValueOf(elliptic.P521)
