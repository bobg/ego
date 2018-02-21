package parse

import (
	"reflect"
	"text/template/parse"
)

var Parse = reflect.ValueOf(parse.Parse)
var New = reflect.ValueOf(parse.New)
var IsEmptyTree = reflect.ValueOf(parse.IsEmptyTree)
var NewIdentifier = reflect.ValueOf(parse.NewIdentifier)
