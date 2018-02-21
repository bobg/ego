package mime

import (
	"mime"
	"reflect"
)

var FormatMediaType = reflect.ValueOf(mime.FormatMediaType)
var ParseMediaType = reflect.ValueOf(mime.ParseMediaType)
var AddExtensionType = reflect.ValueOf(mime.AddExtensionType)
var TypeByExtension = reflect.ValueOf(mime.TypeByExtension)
var ExtensionsByType = reflect.ValueOf(mime.ExtensionsByType)
