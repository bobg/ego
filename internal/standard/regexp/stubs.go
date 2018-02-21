package regexp

import (
	"reflect"
	"regexp"
)

var MustCompilePOSIX = reflect.ValueOf(regexp.MustCompilePOSIX)
var QuoteMeta = reflect.ValueOf(regexp.QuoteMeta)
var Compile = reflect.ValueOf(regexp.Compile)
var CompilePOSIX = reflect.ValueOf(regexp.CompilePOSIX)
var Match = reflect.ValueOf(regexp.Match)
var MatchReader = reflect.ValueOf(regexp.MatchReader)
var MustCompile = reflect.ValueOf(regexp.MustCompile)
var MatchString = reflect.ValueOf(regexp.MatchString)
