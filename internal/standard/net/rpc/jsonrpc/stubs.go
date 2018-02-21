package jsonrpc

import (
	"net/rpc/jsonrpc"
	"reflect"
)

var NewServerCodec = reflect.ValueOf(jsonrpc.NewServerCodec)
var ServeConn = reflect.ValueOf(jsonrpc.ServeConn)
var NewClient = reflect.ValueOf(jsonrpc.NewClient)
var Dial = reflect.ValueOf(jsonrpc.Dial)
var NewClientCodec = reflect.ValueOf(jsonrpc.NewClientCodec)
