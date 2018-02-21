package json

import (
	"encoding/json"
	"reflect"
)

var Valid = reflect.ValueOf(json.Valid)
var Unmarshal = reflect.ValueOf(json.Unmarshal)
var HTMLEscape = reflect.ValueOf(json.HTMLEscape)
var Marshal = reflect.ValueOf(json.Marshal)
var MarshalIndent = reflect.ValueOf(json.MarshalIndent)
var NewEncoder = reflect.ValueOf(json.NewEncoder)
var NewDecoder = reflect.ValueOf(json.NewDecoder)
var Indent = reflect.ValueOf(json.Indent)
var Compact = reflect.ValueOf(json.Compact)
