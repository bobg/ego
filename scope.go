package ego

import (
	"go/ast"
	"reflect"
)

type lazytype struct {
	source *ast.TypeSpec
	typ    reflect.Type
}

type Scope struct {
	parent *Scope

	// values are reflect.Value or lazytype
	objs map[string]interface{}
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		parent: parent,
		objs:   make(map[string]interface{}),
	}
}

func (s *Scope) Lookup(name string) (reflect.Value, error) {
	if val, ok := s.objs[name]; ok {
		if val, ok := val.(reflect.Value); ok {
			return val
		}
		// xxx err
	}
	if s.parent != nil {
		return s.parent.Lookup(name)
	}
	// xxx err
}

func (s *Scope) LookupType(name string) (reflect.Type, error) {
	if val, ok := s.objs[name]; ok {
		if val, ok := val.(lazytype); ok {
			if val.typ == nil {
				// xxx compute val.typ
			}
			return val.typ, nil
		}
		// xxx err
	}
	if s.parent != nil {
		return s.parent.LookupType(name)
	}
	// xxx err
}

// Assign locates name and sets its value to val. It is an error for
// name not to exist.
func (s *Scope) Assign(name string, val reflect.Value) error {
	if _, ok := s.objs[name]; ok {
		s.objs[name] = val
		return nil
	}
	if s.parent != nil {
		return s.parent.Assign(name, val)
	}
	// xxx err
}

func (s *Scope) Add(name string, val reflect.Value) {
	s.objs[name] = val
}

func (s *Scope) AddType(name string, source *ast.TypeSpec) {
	s.objs[name] = lazytype{source: source}
}
