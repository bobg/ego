package sql

import (
	"database/sql"
	"reflect"
)

var Register = reflect.ValueOf(sql.Register)
var Open = reflect.ValueOf(sql.Open)
var Named = reflect.ValueOf(sql.Named)
var Drivers = reflect.ValueOf(sql.Drivers)
