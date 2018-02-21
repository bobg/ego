package big

import (
	"math/big"
	"reflect"
)

var ParseFloat = reflect.ValueOf(big.ParseFloat)
var Jacobi = reflect.ValueOf(big.Jacobi)
var NewInt = reflect.ValueOf(big.NewInt)
var NewRat = reflect.ValueOf(big.NewRat)
var NewFloat = reflect.ValueOf(big.NewFloat)
