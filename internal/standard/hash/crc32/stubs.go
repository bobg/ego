package crc32

import (
	"hash/crc32"
	"reflect"
)

var NewIEEE = reflect.ValueOf(crc32.NewIEEE)
var New = reflect.ValueOf(crc32.New)
var Update = reflect.ValueOf(crc32.Update)
var MakeTable = reflect.ValueOf(crc32.MakeTable)
var Checksum = reflect.ValueOf(crc32.Checksum)
var ChecksumIEEE = reflect.ValueOf(crc32.ChecksumIEEE)
