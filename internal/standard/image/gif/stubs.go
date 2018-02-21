package gif

import (
	"image/gif"
	"reflect"
)

var EncodeAll = reflect.ValueOf(gif.EncodeAll)
var Encode = reflect.ValueOf(gif.Encode)
var DecodeAll = reflect.ValueOf(gif.DecodeAll)
var Decode = reflect.ValueOf(gif.Decode)
var DecodeConfig = reflect.ValueOf(gif.DecodeConfig)
