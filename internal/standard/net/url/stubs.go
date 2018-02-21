package url

import (
	"net/url"
	"reflect"
)

var Parse = reflect.ValueOf(url.Parse)
var User = reflect.ValueOf(url.User)
var ParseRequestURI = reflect.ValueOf(url.ParseRequestURI)
var ParseQuery = reflect.ValueOf(url.ParseQuery)
var UserPassword = reflect.ValueOf(url.UserPassword)
var PathEscape = reflect.ValueOf(url.PathEscape)
var QueryUnescape = reflect.ValueOf(url.QueryUnescape)
var PathUnescape = reflect.ValueOf(url.PathUnescape)
var QueryEscape = reflect.ValueOf(url.QueryEscape)
