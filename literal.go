package ego

import (
	"go/ast"
	"reflect"
	"strconv"

	"github.com/bobg/ego/refl"
)

func (s *Scope) evalBasicLit(expr *ast.BasicLit) (*refl.Value, error) {
	if x, err := strconv.ParseUint(expr.Value, 0, 64); err == nil {
		return &Value{R: reflect.ValueOf(x)}, nil // xxx really we want an untyped int constant
	}
	if x, err := strconv.ParseInt(expr.Value, 0, 64); err == nil {
		return &Value{R: reflect.ValueOf(x)}, nil
	}
	if x, err := strconv.ParseFloat(expr.Value, 64); err == nil {
		return &Value{R: reflect.ValueOf(x)}, nil
	}
	// xxx imaginary literals
	// xxx rune literals
	if x, err := strconv.Unquote(expr.Value); err == nil {
		return &Value{R: reflect.ValueOf(x)}, nil
	}
	// xxx err
}

func (s *Scope) evalCompositeLit(expr *ast.CompositeLit) (*Value, error) {
	typ, err := s.EvalType(expr.Type)
	if err != nil {
		return nil, err
	}
	switch typ.Kind() {
	case StructKind:
		return s.evalStructLit(typ, expr)
	case ArrayKind:
		return s.evalArrayLit(typ, expr)
	case SliceKind:
		return s.evalSliceLit(typ, expr)
	case MapKind:
		return s.evalMapLit(typ, expr)
	}
	// xxx err
}

func (s *Scope) evalStructLit(typ Type, expr *ast.CompositeLit) (*Value, error) {
	var (
		expectKV bool
		result   = NewValue(typ)
	)
	for i, elt := range expr.Elts {
		if elt, ok := elt.(*ast.KeyValueExpr); ok {
			if i == 0 {
				expectKV = true
			} else if !expectKV {
				// xxx err
			}
			key, ok := elt.Key.(*ast.Ident)
			if !ok {
				// xxx err
			}
			val, err := s.Eval1(elt.Value)
			if err != nil {
				return nil, err
			}
			tField, ok := typ.FieldByName(key.Name)
			if !ok {
				// xxx err
			}
			vField := result.FieldByIndex(tField.Index)
			// xxx check val is assignable to vField's type
			vField.Set(val)
		} else {
			if expectKV {
				// xxx err
			}
			if i >= typ.NumField() {
				// xxx err
			}
			val, err := s.Eval1(elt.Value)
			if err != nil {
				return nil, err
			}
			vField := result.FieldByIndex(i)
			// xxx check val is assignable to vField's type
			vField.Set(val)
		}
	}
	return result, nil
}

func (s *Scope) evalArrayLit(typ Type, expr *ast.CompositeLit) (*Value, error) {
	// xxx
}

func (s *Scope) evalSliceLit(typ Type, expr *ast.CompositeLit) (*Value, error) {
	result := NewValue(typ)
	for i, elt := range expr.Elts {
		val, err := s.Eval1(elt)
		if err != nil {
			return nil, err
		}
		if !val.Type().AssignableTo(typ.Elem()) {
			// xxx err
		}
		result = result.Append(val)
	}
	return result, nil
}

func (s *Scope) evalMapLit(typ Type, expr *ast.CompositeLit) (*Value, error) {
	result := NewValue(typ)
	for i, elt := range expr.Elts {
		elt, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			// xxx err
		}
		kv, err := s.Eval1(elt.Key)
		if err != nil {
			return nil, err
		}
		if !kv.Type().AssignableTo(typ.Key()) {
			// xxx err
		}
		vv, err := s.Eval1(elt.Value)
		if err != nil {
			return nil, err
		}
		if !vv.Type().AssignableTo(typ.Elem()) {
			// xxx err
		}
		result.SetMapIndex(kv, vv)
	}
	return result, nil
}
