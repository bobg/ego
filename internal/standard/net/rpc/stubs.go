package rpc

import (
	"net/rpc"
	"reflect"
)

var NewClientWithCodec = reflect.ValueOf(rpc.NewClientWithCodec)
var NewClient = reflect.ValueOf(rpc.NewClient)
var DialHTTP = reflect.ValueOf(rpc.DialHTTP)
var DialHTTPPath = reflect.ValueOf(rpc.DialHTTPPath)
var Dial = reflect.ValueOf(rpc.Dial)
var RegisterName = reflect.ValueOf(rpc.RegisterName)
var Register = reflect.ValueOf(rpc.Register)
var ServeConn = reflect.ValueOf(rpc.ServeConn)
var ServeRequest = reflect.ValueOf(rpc.ServeRequest)
var HandleHTTP = reflect.ValueOf(rpc.HandleHTTP)
var NewServer = reflect.ValueOf(rpc.NewServer)
var Accept = reflect.ValueOf(rpc.Accept)
var ServeCodec = reflect.ValueOf(rpc.ServeCodec)
