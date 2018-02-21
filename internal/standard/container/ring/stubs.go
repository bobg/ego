package ring

import (
	"container/ring"
	"reflect"
)

var New = reflect.ValueOf(ring.New)
