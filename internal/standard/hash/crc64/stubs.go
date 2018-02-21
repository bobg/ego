package crc64

import (
	"hash/crc64"
	"reflect"
)

var MakeTable = reflect.ValueOf(crc64.MakeTable)
var Checksum = reflect.ValueOf(crc64.Checksum)
var New = reflect.ValueOf(crc64.New)
var Update = reflect.ValueOf(crc64.Update)
