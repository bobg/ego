package subtle

import (
	"crypto/subtle"
	"reflect"
)

var ConstantTimeCompare = reflect.ValueOf(subtle.ConstantTimeCompare)
var ConstantTimeSelect = reflect.ValueOf(subtle.ConstantTimeSelect)
var ConstantTimeByteEq = reflect.ValueOf(subtle.ConstantTimeByteEq)
var ConstantTimeEq = reflect.ValueOf(subtle.ConstantTimeEq)
var ConstantTimeCopy = reflect.ValueOf(subtle.ConstantTimeCopy)
var ConstantTimeLessOrEq = reflect.ValueOf(subtle.ConstantTimeLessOrEq)
