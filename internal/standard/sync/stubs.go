package sync

import (
	"reflect"
	"sync"
)

var NewCond = reflect.ValueOf(sync.NewCond)
