package template

import (
	"reflect"
	"text/template"
)

var Must = reflect.ValueOf(template.Must)
var ParseFiles = reflect.ValueOf(template.ParseFiles)
var ParseGlob = reflect.ValueOf(template.ParseGlob)
var JSEscape = reflect.ValueOf(template.JSEscape)
var URLQueryEscaper = reflect.ValueOf(template.URLQueryEscaper)
var HTMLEscapeString = reflect.ValueOf(template.HTMLEscapeString)
var HTMLEscaper = reflect.ValueOf(template.HTMLEscaper)
var JSEscapeString = reflect.ValueOf(template.JSEscapeString)
var HTMLEscape = reflect.ValueOf(template.HTMLEscape)
var JSEscaper = reflect.ValueOf(template.JSEscaper)
var IsTrue = reflect.ValueOf(template.IsTrue)
var New = reflect.ValueOf(template.New)
