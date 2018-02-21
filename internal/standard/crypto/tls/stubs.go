package tls

import (
	"crypto/tls"
	"reflect"
)

var Server = reflect.ValueOf(tls.Server)
var NewListener = reflect.ValueOf(tls.NewListener)
var DialWithDialer = reflect.ValueOf(tls.DialWithDialer)
var LoadX509KeyPair = reflect.ValueOf(tls.LoadX509KeyPair)
var Client = reflect.ValueOf(tls.Client)
var Listen = reflect.ValueOf(tls.Listen)
var Dial = reflect.ValueOf(tls.Dial)
var X509KeyPair = reflect.ValueOf(tls.X509KeyPair)
var NewLRUClientSessionCache = reflect.ValueOf(tls.NewLRUClientSessionCache)
