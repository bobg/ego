package log

import (
	"log"
	"reflect"
)

var SetOutput = reflect.ValueOf(log.SetOutput)
var Fatalf = reflect.ValueOf(log.Fatalf)
var Panicln = reflect.ValueOf(log.Panicln)
var Println = reflect.ValueOf(log.Println)
var Fatalln = reflect.ValueOf(log.Fatalln)
var Output = reflect.ValueOf(log.Output)
var New = reflect.ValueOf(log.New)
var Flags = reflect.ValueOf(log.Flags)
var SetFlags = reflect.ValueOf(log.SetFlags)
var Prefix = reflect.ValueOf(log.Prefix)
var SetPrefix = reflect.ValueOf(log.SetPrefix)
var Print = reflect.ValueOf(log.Print)
var Printf = reflect.ValueOf(log.Printf)
var Fatal = reflect.ValueOf(log.Fatal)
var Panic = reflect.ValueOf(log.Panic)
var Panicf = reflect.ValueOf(log.Panicf)
