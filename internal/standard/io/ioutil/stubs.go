package ioutil

import (
	"io/ioutil"
	"reflect"
)

var TempFile = reflect.ValueOf(ioutil.TempFile)
var TempDir = reflect.ValueOf(ioutil.TempDir)
var ReadDir = reflect.ValueOf(ioutil.ReadDir)
var ReadAll = reflect.ValueOf(ioutil.ReadAll)
var ReadFile = reflect.ValueOf(ioutil.ReadFile)
var WriteFile = reflect.ValueOf(ioutil.WriteFile)
var NopCloser = reflect.ValueOf(ioutil.NopCloser)
