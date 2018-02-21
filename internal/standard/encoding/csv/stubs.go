package csv

import (
	"encoding/csv"
	"reflect"
)

var NewWriter = reflect.ValueOf(csv.NewWriter)
var NewReader = reflect.ValueOf(csv.NewReader)
