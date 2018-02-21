package draw

import (
	"image/draw"
	"reflect"
)

var Draw = reflect.ValueOf(draw.Draw)
var DrawMask = reflect.ValueOf(draw.DrawMask)
