package parser

import (
	"go/parser"
	"reflect"
)

var ParseFile = reflect.ValueOf(parser.ParseFile)
var ParseExpr = reflect.ValueOf(parser.ParseExpr)
var ParseDir = reflect.ValueOf(parser.ParseDir)
var ParseExprFrom = reflect.ValueOf(parser.ParseExprFrom)
