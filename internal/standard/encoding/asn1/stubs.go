package asn1

import (
	"encoding/asn1"
	"reflect"
)

var Marshal = reflect.ValueOf(asn1.Marshal)
var UnmarshalWithParams = reflect.ValueOf(asn1.UnmarshalWithParams)
var Unmarshal = reflect.ValueOf(asn1.Unmarshal)
