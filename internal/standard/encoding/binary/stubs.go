package binary

import (
	"encoding/binary"
	"reflect"
)

var Write = reflect.ValueOf(binary.Write)
var Read = reflect.ValueOf(binary.Read)
var Size = reflect.ValueOf(binary.Size)
var Varint = reflect.ValueOf(binary.Varint)
var ReadVarint = reflect.ValueOf(binary.ReadVarint)
var PutUvarint = reflect.ValueOf(binary.PutUvarint)
var Uvarint = reflect.ValueOf(binary.Uvarint)
var PutVarint = reflect.ValueOf(binary.PutVarint)
var ReadUvarint = reflect.ValueOf(binary.ReadUvarint)
