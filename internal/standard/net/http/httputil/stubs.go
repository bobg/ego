package httputil

import (
	"net/http/httputil"
	"reflect"
)

var NewServerConn = reflect.ValueOf(httputil.NewServerConn)
var NewClientConn = reflect.ValueOf(httputil.NewClientConn)
var NewProxyClientConn = reflect.ValueOf(httputil.NewProxyClientConn)
var NewChunkedReader = reflect.ValueOf(httputil.NewChunkedReader)
var NewChunkedWriter = reflect.ValueOf(httputil.NewChunkedWriter)
var NewSingleHostReverseProxy = reflect.ValueOf(httputil.NewSingleHostReverseProxy)
var DumpRequestOut = reflect.ValueOf(httputil.DumpRequestOut)
var DumpResponse = reflect.ValueOf(httputil.DumpResponse)
var DumpRequest = reflect.ValueOf(httputil.DumpRequest)
