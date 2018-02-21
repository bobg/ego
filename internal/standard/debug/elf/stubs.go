package elf

import (
	"debug/elf"
	"reflect"
)

var Open = reflect.ValueOf(elf.Open)
var NewFile = reflect.ValueOf(elf.NewFile)
var ST_BIND = reflect.ValueOf(elf.ST_BIND)
var ST_INFO = reflect.ValueOf(elf.ST_INFO)
var R_INFO32 = reflect.ValueOf(elf.R_INFO32)
var R_TYPE64 = reflect.ValueOf(elf.R_TYPE64)
var ST_VISIBILITY = reflect.ValueOf(elf.ST_VISIBILITY)
var R_TYPE32 = reflect.ValueOf(elf.R_TYPE32)
var R_SYM64 = reflect.ValueOf(elf.R_SYM64)
var R_INFO = reflect.ValueOf(elf.R_INFO)
var R_SYM32 = reflect.ValueOf(elf.R_SYM32)
var ST_TYPE = reflect.ValueOf(elf.ST_TYPE)
