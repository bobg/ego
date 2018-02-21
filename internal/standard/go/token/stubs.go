package token

import (
	"go/token"
	"reflect"
)

var Lookup = reflect.ValueOf(token.Lookup)
var NewFileSet = reflect.ValueOf(token.NewFileSet)
