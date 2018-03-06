package ego

import (
	"go/ast"
	"reflect"

	"github.com/bobg/ego/refl"
)

func (s *Scope) evalSelector(expr *ast.SelectorExpr) (*refl.Value, error) {
	// package.Identifier
	if id, ok := expr.(*ast.Ident); ok {
		if pkgScope := s.LookupScope(id.Name); pkgScope != nil {
			// xxx check that id.Name is capitalized
			if val := pkgScope.Lookup(expr.Sel.Name); val != nil {
				return val, nil
			}
			// xxx err
		}
	}

	base, err := s.Eval1(expr.X)
	if err != nil {
		return nil, err
	}

	baseType, err := base.Type()
	if err != nil {
		return nil, err
	}

	meth, field, isPtr, err := findShallowestMethodOrField(baseType, expr.Sel.Name)
	if err != nil {
		return nil, err
	}

	if field != nil {
		return base.FieldByIndex(field.Index)
	}

	if meth == nil {
		// xxx err
	}

	mscope := NewScope(meth.scope)
	mscope.Add(meth.recvName, base) // xxx adjust according to isPtr
	return mscope.makeFunc(meth.typ, meth.stmts)
}

// xxx should error if there are two items with the same name at a
// given depth (https://golang.org/ref/spec#Selectors)
func findShallowestMethodOrField(typ *refl.Type, sel string) (*refl.Method, *refl.StructField, bool, error) {
	queue := []*refl.Type{typ}
	for len(queue) > 0 {
		typ = queue[0]
		queue = queue[1:]

		var (
			baseType = typ
			isPtr    bool
		)
		if baseType.Kind() == reflect.Ptr {
			isPtr = true
			baseType = baseType.Elem()
		}

		m, ok := baseType.MethodByName(sel)
		if ok {
			return m, nil, isPtr, nil
		}
		if baseType.Kind() == reflect.Struct {
			f, ok := baseType.FieldByName(sel)
			if ok {
				return nil, f, isPtr, nil
			}
			for i := 0; i < baseType.NumField(); i++ {
				f := baseType.FieldByIndex(i)
				if f.Anonymous {
					queue = append(queue, f.Type)
				}
			}
		}
	}
	return nil, nil, false, nil // xxx error?
}
