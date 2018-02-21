package plugin

import (
	"plugin"
	"reflect"
)

var Open = reflect.ValueOf(plugin.Open)
