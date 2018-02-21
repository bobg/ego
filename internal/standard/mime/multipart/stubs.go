package multipart

import (
	"mime/multipart"
	"reflect"
)

var NewReader = reflect.ValueOf(multipart.NewReader)
var NewWriter = reflect.ValueOf(multipart.NewWriter)
