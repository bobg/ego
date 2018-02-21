package bufio

import (
	"bufio"
	"reflect"
)

var ScanBytes = reflect.ValueOf(bufio.ScanBytes)
var ScanWords = reflect.ValueOf(bufio.ScanWords)
var NewScanner = reflect.ValueOf(bufio.NewScanner)
var ScanLines = reflect.ValueOf(bufio.ScanLines)
var ScanRunes = reflect.ValueOf(bufio.ScanRunes)
var NewReadWriter = reflect.ValueOf(bufio.NewReadWriter)
var NewReader = reflect.ValueOf(bufio.NewReader)
var NewWriterSize = reflect.ValueOf(bufio.NewWriterSize)
var NewWriter = reflect.ValueOf(bufio.NewWriter)
var NewReaderSize = reflect.ValueOf(bufio.NewReaderSize)
