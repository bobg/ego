package tar

import (
	"archive/tar"
	"reflect"
)

var NewWriter = reflect.ValueOf(tar.NewWriter)
var NewReader = reflect.ValueOf(tar.NewReader)
var FileInfoHeader = reflect.ValueOf(tar.FileInfoHeader)
