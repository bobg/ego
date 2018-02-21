package rand

import (
	"math/rand"
	"reflect"
)

var NewZipf = reflect.ValueOf(rand.NewZipf)
var Uint64 = reflect.ValueOf(rand.Uint64)
var Int = reflect.ValueOf(rand.Int)
var Int31n = reflect.ValueOf(rand.Int31n)
var Float32 = reflect.ValueOf(rand.Float32)
var Perm = reflect.ValueOf(rand.Perm)
var NormFloat64 = reflect.ValueOf(rand.NormFloat64)
var Seed = reflect.ValueOf(rand.Seed)
var ExpFloat64 = reflect.ValueOf(rand.ExpFloat64)
var New = reflect.ValueOf(rand.New)
var Int63 = reflect.ValueOf(rand.Int63)
var Uint32 = reflect.ValueOf(rand.Uint32)
var Intn = reflect.ValueOf(rand.Intn)
var Int63n = reflect.ValueOf(rand.Int63n)
var NewSource = reflect.ValueOf(rand.NewSource)
var Int31 = reflect.ValueOf(rand.Int31)
var Float64 = reflect.ValueOf(rand.Float64)
var Read = reflect.ValueOf(rand.Read)
