package user

import (
	"os/user"
	"reflect"
)

var LookupGroup = reflect.ValueOf(user.LookupGroup)
var LookupGroupId = reflect.ValueOf(user.LookupGroupId)
var Current = reflect.ValueOf(user.Current)
var Lookup = reflect.ValueOf(user.Lookup)
var LookupId = reflect.ValueOf(user.LookupId)
