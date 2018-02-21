package zip

import (
	"archive/zip"
	"reflect"
)

var RegisterCompressor = reflect.ValueOf(zip.RegisterCompressor)
var RegisterDecompressor = reflect.ValueOf(zip.RegisterDecompressor)
var FileInfoHeader = reflect.ValueOf(zip.FileInfoHeader)
var NewWriter = reflect.ValueOf(zip.NewWriter)
var OpenReader = reflect.ValueOf(zip.OpenReader)
var NewReader = reflect.ValueOf(zip.NewReader)
