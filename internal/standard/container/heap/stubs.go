package heap

import (
	"container/heap"
	"reflect"
)

var Fix = reflect.ValueOf(heap.Fix)
var Init = reflect.ValueOf(heap.Init)
var Push = reflect.ValueOf(heap.Push)
var Pop = reflect.ValueOf(heap.Pop)
var Remove = reflect.ValueOf(heap.Remove)
