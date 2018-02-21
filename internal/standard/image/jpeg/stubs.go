package jpeg

import (
	"image/jpeg"
	"reflect"
)

var Encode = reflect.ValueOf(jpeg.Encode)
var DecodeConfig = reflect.ValueOf(jpeg.DecodeConfig)
var Decode = reflect.ValueOf(jpeg.Decode)
