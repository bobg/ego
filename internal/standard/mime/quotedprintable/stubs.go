package quotedprintable

import (
	"mime/quotedprintable"
	"reflect"
)

var NewWriter = reflect.ValueOf(quotedprintable.NewWriter)
var NewReader = reflect.ValueOf(quotedprintable.NewReader)
