package fmt

import (
	"fmt"
	"reflect"
)

var Scan = reflect.ValueOf(fmt.Scan)
var Scanf = reflect.ValueOf(fmt.Scanf)
var Sscanln = reflect.ValueOf(fmt.Sscanln)
var Fscan = reflect.ValueOf(fmt.Fscan)
var Sscan = reflect.ValueOf(fmt.Sscan)
var Fscanln = reflect.ValueOf(fmt.Fscanln)
var Fscanf = reflect.ValueOf(fmt.Fscanf)
var Scanln = reflect.ValueOf(fmt.Scanln)
var Sscanf = reflect.ValueOf(fmt.Sscanf)
var Print = reflect.ValueOf(fmt.Print)
var Printf = reflect.ValueOf(fmt.Printf)
var Fprintf = reflect.ValueOf(fmt.Fprintf)
var Errorf = reflect.ValueOf(fmt.Errorf)
var Sprintln = reflect.ValueOf(fmt.Sprintln)
var Fprint = reflect.ValueOf(fmt.Fprint)
var Fprintln = reflect.ValueOf(fmt.Fprintln)
var Println = reflect.ValueOf(fmt.Println)
var Sprintf = reflect.ValueOf(fmt.Sprintf)
var Sprint = reflect.ValueOf(fmt.Sprint)
