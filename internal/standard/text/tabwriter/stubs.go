package tabwriter

import (
	"reflect"
	"text/tabwriter"
)

var NewWriter = reflect.ValueOf(tabwriter.NewWriter)
