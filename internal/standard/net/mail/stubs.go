package mail

import (
	"net/mail"
	"reflect"
)

var ParseAddress = reflect.ValueOf(mail.ParseAddress)
var ReadMessage = reflect.ValueOf(mail.ReadMessage)
var ParseDate = reflect.ValueOf(mail.ParseDate)
var ParseAddressList = reflect.ValueOf(mail.ParseAddressList)
