package windows

import (
	"reflect"
	"runtime/testdata/testprogcgo/windows"
)

var GetThread = reflect.ValueOf(windows.GetThread)
