package smtp

import (
	"net/smtp"
	"reflect"
)

var Dial = reflect.ValueOf(smtp.Dial)
var NewClient = reflect.ValueOf(smtp.NewClient)
var SendMail = reflect.ValueOf(smtp.SendMail)
var PlainAuth = reflect.ValueOf(smtp.PlainAuth)
var CRAMMD5Auth = reflect.ValueOf(smtp.CRAMMD5Auth)
