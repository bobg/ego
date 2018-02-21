package xml

import (
	"encoding/xml"
	"reflect"
)

var Unmarshal = reflect.ValueOf(xml.Unmarshal)
var Marshal = reflect.ValueOf(xml.Marshal)
var NewEncoder = reflect.ValueOf(xml.NewEncoder)
var MarshalIndent = reflect.ValueOf(xml.MarshalIndent)
var EscapeText = reflect.ValueOf(xml.EscapeText)
var CopyToken = reflect.ValueOf(xml.CopyToken)
var Escape = reflect.ValueOf(xml.Escape)
var NewDecoder = reflect.ValueOf(xml.NewDecoder)
