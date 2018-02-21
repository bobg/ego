package png

import (
	"image/png"
	"reflect"
)

var Encode = reflect.ValueOf(png.Encode)
var DecodeConfig = reflect.ValueOf(png.DecodeConfig)
var Decode = reflect.ValueOf(png.Decode)
