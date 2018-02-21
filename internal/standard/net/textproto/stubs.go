package textproto

import (
	"net/textproto"
	"reflect"
)

var CanonicalMIMEHeaderKey = reflect.ValueOf(textproto.CanonicalMIMEHeaderKey)
var NewReader = reflect.ValueOf(textproto.NewReader)
var Dial = reflect.ValueOf(textproto.Dial)
var TrimString = reflect.ValueOf(textproto.TrimString)
var NewConn = reflect.ValueOf(textproto.NewConn)
var TrimBytes = reflect.ValueOf(textproto.TrimBytes)
var NewWriter = reflect.ValueOf(textproto.NewWriter)
