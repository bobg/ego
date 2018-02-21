package image

import (
	"image"
	"reflect"
)

var NewUniform = reflect.ValueOf(image.NewUniform)
var NewPaletted = reflect.ValueOf(image.NewPaletted)
var NewRGBA = reflect.ValueOf(image.NewRGBA)
var NewRGBA64 = reflect.ValueOf(image.NewRGBA64)
var NewNRGBA = reflect.ValueOf(image.NewNRGBA)
var NewAlpha16 = reflect.ValueOf(image.NewAlpha16)
var NewGray16 = reflect.ValueOf(image.NewGray16)
var NewCMYK = reflect.ValueOf(image.NewCMYK)
var NewNRGBA64 = reflect.ValueOf(image.NewNRGBA64)
var NewAlpha = reflect.ValueOf(image.NewAlpha)
var NewGray = reflect.ValueOf(image.NewGray)
var Pt = reflect.ValueOf(image.Pt)
var Rect = reflect.ValueOf(image.Rect)
var NewNYCbCrA = reflect.ValueOf(image.NewNYCbCrA)
var NewYCbCr = reflect.ValueOf(image.NewYCbCr)
var RegisterFormat = reflect.ValueOf(image.RegisterFormat)
var DecodeConfig = reflect.ValueOf(image.DecodeConfig)
var Decode = reflect.ValueOf(image.Decode)
