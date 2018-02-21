package utf8

import (
	"reflect"
	"unicode/utf8"
)

var ValidString = reflect.ValueOf(utf8.ValidString)
var DecodeRune = reflect.ValueOf(utf8.DecodeRune)
var RuneStart = reflect.ValueOf(utf8.RuneStart)
var FullRune = reflect.ValueOf(utf8.FullRune)
var ValidRune = reflect.ValueOf(utf8.ValidRune)
var DecodeRuneInString = reflect.ValueOf(utf8.DecodeRuneInString)
var DecodeLastRune = reflect.ValueOf(utf8.DecodeLastRune)
var EncodeRune = reflect.ValueOf(utf8.EncodeRune)
var RuneCount = reflect.ValueOf(utf8.RuneCount)
var DecodeLastRuneInString = reflect.ValueOf(utf8.DecodeLastRuneInString)
var RuneCountInString = reflect.ValueOf(utf8.RuneCountInString)
var Valid = reflect.ValueOf(utf8.Valid)
var RuneLen = reflect.ValueOf(utf8.RuneLen)
var FullRuneInString = reflect.ValueOf(utf8.FullRuneInString)
