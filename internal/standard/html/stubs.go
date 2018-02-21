package html

import (
	"html"
	"reflect"
)

var EscapeString = reflect.ValueOf(html.EscapeString)
var UnescapeString = reflect.ValueOf(html.UnescapeString)
