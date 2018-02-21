package color

import (
	"image/color"
	"reflect"
)

var ModelFunc = reflect.ValueOf(color.ModelFunc)
var CMYKToRGB = reflect.ValueOf(color.CMYKToRGB)
var RGBToCMYK = reflect.ValueOf(color.RGBToCMYK)
var RGBToYCbCr = reflect.ValueOf(color.RGBToYCbCr)
var YCbCrToRGB = reflect.ValueOf(color.YCbCrToRGB)
