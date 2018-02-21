package doc

import (
	"go/doc"
	"reflect"
)

var ToHTML = reflect.ValueOf(doc.ToHTML)
var ToText = reflect.ValueOf(doc.ToText)
var IsPredeclared = reflect.ValueOf(doc.IsPredeclared)
var Synopsis = reflect.ValueOf(doc.Synopsis)
var New = reflect.ValueOf(doc.New)
var Examples = reflect.ValueOf(doc.Examples)
