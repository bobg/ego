package testing

import (
	"reflect"
	"testing"
)

var RunBenchmarks = reflect.ValueOf(testing.RunBenchmarks)
var Benchmark = reflect.ValueOf(testing.Benchmark)
var Short = reflect.ValueOf(testing.Short)
var CoverMode = reflect.ValueOf(testing.CoverMode)
var RunTests = reflect.ValueOf(testing.RunTests)
var MainStart = reflect.ValueOf(testing.MainStart)
var Main = reflect.ValueOf(testing.Main)
var Verbose = reflect.ValueOf(testing.Verbose)
var RunExamples = reflect.ValueOf(testing.RunExamples)
var AllocsPerRun = reflect.ValueOf(testing.AllocsPerRun)
var Coverage = reflect.ValueOf(testing.Coverage)
var RegisterCover = reflect.ValueOf(testing.RegisterCover)
