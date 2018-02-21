package path

import (
	"path"
	"reflect"
)

var Clean = reflect.ValueOf(path.Clean)
var Split = reflect.ValueOf(path.Split)
var Join = reflect.ValueOf(path.Join)
var Ext = reflect.ValueOf(path.Ext)
var Base = reflect.ValueOf(path.Base)
var IsAbs = reflect.ValueOf(path.IsAbs)
var Dir = reflect.ValueOf(path.Dir)
var Match = reflect.ValueOf(path.Match)
