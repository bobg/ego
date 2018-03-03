package ego

import (
	"sync"

	"github.com/bobg/ego/refl"
)

type deferral struct {
	fun  *refl.Value
	args []*refl.Value
}

type Scope struct {
	mu        sync.RWMutex
	parent    *Scope
	objs      map[string]interface{}
	deferrals *[]deferral // nil pointer means not a function scope
}

type uninitialized ast.Expr

type builtin int

const (
	builtinBool builtin = 1 + iota
	builtinByte
	builtinComplex64
	builtinComplex128
	builtinError
	builtinFloat32
	builtinFloat64
	builtinInt
	builtinInt8
	builtinInt16
	builtinInt32
	builtinInt64
	builtinRune
	builtinString
	builtinUint
	builtinUint8
	builtinUint16
	builtinUint32
	builtinUint64
	builtinUintptr
	builtinTrue
	builtinFalse
	builtinIota
	builtinNil
	builtinAppend
	builtinCap
	builtinClose
	builtinComplex
	builtinCopy
	builtinDelete
	builtinImag
	builtinLen
	builtinMake
	builtinNew
	builtinPanic
	builtinPrint
	builtinPrintln
	builtinReal
	builtinRecover
)

var universe = &Scope{
	objs: map[string]interface{}{
		"bool":       builtinBool,
		"byte":       builtinByte,
		"complex64":  builtinComplex64,
		"complex128": builtinComplex128,
		"error":      builtinError,
		"float32":    builtinFloat32,
		"float64":    builtinFloat64,
		"int":        builtinInt,
		"int8":       builtinInt8,
		"int16":      builtinInt16,
		"int32":      builtinInt32,
		"int64":      builtinInt64,
		"rune":       builtinRune,
		"string":     builtinString,
		"uint":       builtinUint,
		"uint8":      builtinUint8,
		"uint16":     builtinUint16,
		"uint32":     builtinUint32,
		"uint64":     builtinUint64,
		"uintptr":    builtinUintptr,
		"true":       builtinTrue,
		"false":      builtinFalse,
		"iota":       builtinIota,
		"nil":        builtinNil,
		"append":     builtinAppend,
		"cap":        builtinCap,
		"close":      builtinClose,
		"complex":    builtinComplex,
		"copy":       builtinCopy,
		"delete":     builtinDelete,
		"imag":       builtinImag,
		"len":        builtinLen,
		"make":       builtinMake,
		"new":        builtinNew,
		"panic":      builtinPanic,
		"print":      builtinPrint,
		"println":    builtinPrintln,
		"real":       builtinReal,
		"recover":    builtinRecover,
	},
}

func NewScope(parent *Scope) *Scope {
	if parent == nil {
		parent = universe
	}
	return &Scope{
		parent: parent,
		objs:   make(map[string]interface{}),
	}
}

func (s *Scope) addDefer(f *refl.Value, args []*refl.Value) {
	if s.deferrals != nil {
		(*s.deferrals) = append((*s.deferrals), deferral{f, args})
	} else if s.parent != nil {
		s.parent.addDefer(f, args)
	}
}

// xxx check first that name is undefined?
func (s *Scope) Add(name string, val interface{}) {
	s[name] = val
}

func (s *Scope) Set(name string, val interface{}) error {
	if _, ok := s[name]; ok {
		s[name] = val
		return nil
	}
	if s.parent == nil || s.parent.isUniverse() {
		// xxx err
	}
	return s.parent.Set(name, val)
}

func (s *Scope) Lookup(name) interface{} {
	if val, ok := s[name]; ok {
		return val
	}
	if s.parent == nil {
		return nil
	}
	return s.parent.Lookup(name)
}

func (s *Scope) LookupValue(name) *refl.Value {
	l := s.Lookup(name)
	val, _ := l.(*refl.Value)
	return val
}
