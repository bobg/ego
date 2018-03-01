package ego

import (
	"go/ast"
	"go/token"

	"reflect"
)

func (scope *Scope) evalCall(expr *ast.CallExpr) ([]reflect.Value, error) {
	if name, ok := expr.Fun.(*ast.Ident); ok {
		switch name.Name {
		case "close":
			if len(expr.Args) != 1 {
				return nil, errArgCount
			}
			v, err := scope.eval1(expr.Args[0])
			if err != nil {
				return nil, err
			}
			if v.Kind() != reflect.Chan {
				return nil, errArgType
			}
			v.Close()
			return nil, nil

		case "len":
			if len(expr.Args) != 1 {
				return nil, errArgCount
			}
			v, err := scope.eval1(expr.Args[0])
			if err != nil {
				return nil, err
			}
			v = resolveArrayPtr(v)
			switch v.Kind() {
			case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
				return []reflect.Value{reflect.ValueOf(v.Len())}, nil
			}
			return nil, errArgType

		case "cap":
			if len(expr.Args) != 1 {
				return nil, errArgCount
			}
			v, err := scope.eval1(expr.Args[0])
			if err != nil {
				return nil, err
			}
			v = resolveArrayPtr(v)
			switch v.Kind() {
			case reflect.Array, reflect.Chan, reflect.Slice:
				return []reflect.Value{reflect.ValueOf(v.Cap())}, nil
			}
			return nil, errArgType

		case "new":
			if len(expr.Args) != 1 {
				return nil, errArgCount
			}
			typ, err := scope.evalType(expr.Args[0])
			if err != nil {
				return nil, err
			}
			return []reflect.Value{reflect.New(typ)}, nil

		case "make":
			if len(expr.Args) < 1 {
				return nil, errArgCount
			}
			typ, err := scope.evalType(expr.Args[0])
			if err != nil {
				return nil, err
			}
			switch typ.Kind() {
			// xxx
			}

		case "append":
		case "delete":
		case "complex":
		case "real":
		case "imag":
		case "panic":
		case "recover":
		case "print":
		case "println":
		}
	}

	fun, err := scope.eval1(expr.Fun)
	if err != nil {
		return nil, err
	}
	switch fun.Kind() {
	case reflect.Func:
		var (
			args     []reflect.Value
			variadic = expr.Ellipsis != token.NoPos
		)
		for _, arg := range expr.Args {
			v, err := scope.eval1(arg)
			if err != nil {
				return nil, err
			}
			append(args, v)
		}
		vals := fun.Call(args)
		return vals, nil

		// xxx other cases, e.g. a type conversion
	}
	return nil, errArgType
}

func resolveArrayPtr(v reflect.Value) reflect.Value {
	if v.Kind() != reflect.Ptr {
		return v
	}
	if v.IsNil() {
		return v
	}
	if reflect.TypeOf(v).Elem().Kind() != reflect.Array {
		return v
	}
	return v.Elem()
}

type wrappedFn func([]reflect.Value) []reflect.Value

// Function fnWrapper is for ordinary functions only.  Because reflect
// does not allow adding methods to types, we have to simulate methods
// ourselves. We do this with methodWrapper, below, which allows
// calling methods like:
//   wrappedMethod(receiver)(other args...)
func fnWrapper(scope *Scope, recv *ast.FieldList, typ *ast.FuncType, body *ast.BlockStmt) (wrappedFn, error) {
	var ftypes, rtypes []reflect.Type
	for _, f := range typ.Params.List {
		ftype, err := scope.evalType(f.Type)
		if err != nil {
			return nil, err
		}
		for range f.Names {
			ftypes = append(ftypes, ftype)
		}
	}
	if typ.Results != nil {
		for _, f := range typ.Results.List {
			rtype, err := scope.evalType(f.Type)
			if err != nil {
				return nil, err
			}
			for range f.Names {
				rtypes = append(rtypes, rtype)
			}
		}
	}

	return func(args []reflect.Value) []reflect.Value {
		scope = NewScope(scope)
		// xxx range-check len(args) against number of params
		// xxx handle variadic calls
		var i int
		for _, f := range typ.Params.List {
			for _, n := range f.Names {
				if args[i].Type() != ftypes[i] {
					// xxx err
				}
				if n != nil {
					scope.Add(n.Name, args[i])
				}
				i++
			}
		}
		i = 0
		if typ.Results != nil {
			for _, f := range typ.Results.List {
				for _, n := range f.Names {
					if n != nil {
						scope.Add(n.Name, reflect.New(rtypes[i]))
					}
					i++
				}
			}
		}

		branch, err = execStmts(scope, lit.Body.List)
		if err != nil {
			// xxx how to communicate errors?
		}
		if branch != nil {
			// xxx handle branching
		}
		// xxx unmarshal return val(s) from activation record and return them
	}
}

func methodWrapper(scope *Scope, recv *ast.FieldList, typ *ast.FuncType, body *ast.BlockStmt) (func(reflect.Value) wrappedFn, error) {
	// ...
}
