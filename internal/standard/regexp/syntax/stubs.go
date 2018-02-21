package syntax

import (
	"reflect"
	"regexp/syntax"
)

var EmptyOpContext = reflect.ValueOf(syntax.EmptyOpContext)
var IsWordChar = reflect.ValueOf(syntax.IsWordChar)
var Compile = reflect.ValueOf(syntax.Compile)
var Parse = reflect.ValueOf(syntax.Parse)
