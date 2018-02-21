package utf16

import (
	"reflect"
	"unicode/utf16"
)

var Decode = reflect.ValueOf(utf16.Decode)
var IsSurrogate = reflect.ValueOf(utf16.IsSurrogate)
var DecodeRune = reflect.ValueOf(utf16.DecodeRune)
var EncodeRune = reflect.ValueOf(utf16.EncodeRune)
var Encode = reflect.ValueOf(utf16.Encode)
