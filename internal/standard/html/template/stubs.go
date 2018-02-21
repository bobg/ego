package template

import (
	"html/template"
	"reflect"
)

var ParseGlob = reflect.ValueOf(template.ParseGlob)
var Must = reflect.ValueOf(template.Must)
var ParseFiles = reflect.ValueOf(template.ParseFiles)
var IsTrue = reflect.ValueOf(template.IsTrue)
var New = reflect.ValueOf(template.New)
var HTMLEscapeString = reflect.ValueOf(template.HTMLEscapeString)
var JSEscapeString = reflect.ValueOf(template.JSEscapeString)
var JSEscaper = reflect.ValueOf(template.JSEscaper)
var JSEscape = reflect.ValueOf(template.JSEscape)
var HTMLEscape = reflect.ValueOf(template.HTMLEscape)
var HTMLEscaper = reflect.ValueOf(template.HTMLEscaper)
var URLQueryEscaper = reflect.ValueOf(template.URLQueryEscaper)
