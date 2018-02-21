package time

import (
	"reflect"
	"time"
)

var NewTicker = reflect.ValueOf(time.NewTicker)
var Tick = reflect.ValueOf(time.Tick)
var Until = reflect.ValueOf(time.Until)
var Date = reflect.ValueOf(time.Date)
var Unix = reflect.ValueOf(time.Unix)
var Since = reflect.ValueOf(time.Since)
var Now = reflect.ValueOf(time.Now)
var NewTimer = reflect.ValueOf(time.NewTimer)
var AfterFunc = reflect.ValueOf(time.AfterFunc)
var Sleep = reflect.ValueOf(time.Sleep)
var After = reflect.ValueOf(time.After)
var LoadLocation = reflect.ValueOf(time.LoadLocation)
var FixedZone = reflect.ValueOf(time.FixedZone)
var Parse = reflect.ValueOf(time.Parse)
var ParseInLocation = reflect.ValueOf(time.ParseInLocation)
var ParseDuration = reflect.ValueOf(time.ParseDuration)
