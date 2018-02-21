package driver

import (
	"database/sql/driver"
	"reflect"
)

var IsValue = reflect.ValueOf(driver.IsValue)
var IsScanValue = reflect.ValueOf(driver.IsScanValue)
