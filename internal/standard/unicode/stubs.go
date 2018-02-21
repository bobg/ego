package unicode

import (
	"reflect"
	"unicode"
)

var IsDigit = reflect.ValueOf(unicode.IsDigit)
var IsOneOf = reflect.ValueOf(unicode.IsOneOf)
var In = reflect.ValueOf(unicode.In)
var IsSpace = reflect.ValueOf(unicode.IsSpace)
var IsGraphic = reflect.ValueOf(unicode.IsGraphic)
var IsPrint = reflect.ValueOf(unicode.IsPrint)
var IsPunct = reflect.ValueOf(unicode.IsPunct)
var IsMark = reflect.ValueOf(unicode.IsMark)
var IsNumber = reflect.ValueOf(unicode.IsNumber)
var IsSymbol = reflect.ValueOf(unicode.IsSymbol)
var IsControl = reflect.ValueOf(unicode.IsControl)
var IsLetter = reflect.ValueOf(unicode.IsLetter)
var IsTitle = reflect.ValueOf(unicode.IsTitle)
var SimpleFold = reflect.ValueOf(unicode.SimpleFold)
var ToLower = reflect.ValueOf(unicode.ToLower)
var IsLower = reflect.ValueOf(unicode.IsLower)
var To = reflect.ValueOf(unicode.To)
var ToUpper = reflect.ValueOf(unicode.ToUpper)
var ToTitle = reflect.ValueOf(unicode.ToTitle)
var Is = reflect.ValueOf(unicode.Is)
var IsUpper = reflect.ValueOf(unicode.IsUpper)
