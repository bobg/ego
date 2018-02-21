package hex

import (
	"encoding/hex"
	"reflect"
)

var Decode = reflect.ValueOf(hex.Decode)
var EncodeToString = reflect.ValueOf(hex.EncodeToString)
var DecodeString = reflect.ValueOf(hex.DecodeString)
var Dump = reflect.ValueOf(hex.Dump)
var Dumper = reflect.ValueOf(hex.Dumper)
var EncodedLen = reflect.ValueOf(hex.EncodedLen)
var Encode = reflect.ValueOf(hex.Encode)
var DecodedLen = reflect.ValueOf(hex.DecodedLen)
