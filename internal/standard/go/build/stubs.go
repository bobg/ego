package build

import (
	"go/build"
	"reflect"
)

var ImportDir = reflect.ValueOf(build.ImportDir)
var IsLocalImport = reflect.ValueOf(build.IsLocalImport)
var Import = reflect.ValueOf(build.Import)
var ArchChar = reflect.ValueOf(build.ArchChar)
