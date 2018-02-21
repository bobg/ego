package io

import (
	"io"
	"reflect"
)

var MultiReader = reflect.ValueOf(io.MultiReader)
var MultiWriter = reflect.ValueOf(io.MultiWriter)
var Copy = reflect.ValueOf(io.Copy)
var CopyN = reflect.ValueOf(io.CopyN)
var NewSectionReader = reflect.ValueOf(io.NewSectionReader)
var ReadAtLeast = reflect.ValueOf(io.ReadAtLeast)
var WriteString = reflect.ValueOf(io.WriteString)
var CopyBuffer = reflect.ValueOf(io.CopyBuffer)
var TeeReader = reflect.ValueOf(io.TeeReader)
var ReadFull = reflect.ValueOf(io.ReadFull)
var LimitReader = reflect.ValueOf(io.LimitReader)
var Pipe = reflect.ValueOf(io.Pipe)
