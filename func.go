package ego

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"

	"github.com/bobg/ego/refl"
	"github.com/pkg/errors"
)

func (s *Scope) evalFuncLit(expr *ast.FuncLit) (*refl.Value, error) {
	return s.makeFunc(expr.Type, expr.Body.List)
}

func (s *Scope) makeFunc(typ *ast.FuncType, stmts []ast.Stmt) (*refl.Value, error) {
	var (
		paramTypes, resultTypes     []reflect.Type // n.b. not refl.Type
		isVariadic, hasNamedReturns bool
	)

	for i, p := range typ.Params.List {
		if isVariadic {
			return nil, fmt.Errorf("no parameters may follow '...'")
		}
		if _, ok := p.Type.(*ast.Ellipsis); ok {
			isVariadic = true
		}
		t, err := s.EvalType(p.Type)
		if err != nil {
			return nil, errors.Wrapf(err, "eval type of param %d", i)
		}
		if t.R == nil {
			return nil, fmt.Errorf("unsupported parameter type")
		}
		for range p.Names {
			paramTypes = append(paramTypes, t.R)
		}
	}
	if typ.Results != nil {
		for i, r := range typ.Results.List {
			t, err := s.EvalType(r.Type)
			if err != nil {
				return nil, errors.Wrapf(err, "eval type of result %d", i)
			}
			if t.R == nil {
				return nil, fmt.Errorf("unsupported result type")
			}
			for i, n := range r.Names {
				if i == 0 {
					hasNamedReturns = (n != nil)
				} else if hasNamedReturns == (n == nil) {
					return nil, fmt.Errorf("either zero results may be named or all must be")
				}
				resultTypes = append(resultTypes, t.R)
			}
		}
	}

	ftype, err := s.EvalType(typ)
	if err != nil {
		return nil, errors.Wrap(err, "eval type of func lit")
	}
	if ftype.R == nil {
		return nil, fmt.Errorf("unsupported function type")
	}

	impl := func(args []reflect.Value) []reflect.Value {
		// check len(args) against expected number of params
		fscope := NewScope(s)
		fscope.deferrals = new([]deferral)

		defer func() {
			deferrals := *fscope.deferrals
			for i := len(deferrals) - 1; i >= 0; i-- {
				d := deferrals[i]
				if d.fun == nil {
					// xxx err - panic?
				}
				_, err := d.fun.Call(d.args)
				if err != nil {
					// xxx communicate err out - panic?
				}
			}
		}()

		var i int
		for _, p := range typ.Params.List {
			for _, n := range p.Names {
				arg := args[i]
				if !arg.Type().AssignableTo(paramTypes[i]) {
					// xxx err
				}
				if n != nil {
					fscope.Add(n.Name, refl.Value{R: arg})
				}
				i++
			}
			// xxx handle variadic final param
		}
		i = 0
		if typ.Results != nil {
			for _, r := range typ.Results.List {
				for _, n := range p.Names {
					if n != nil {
						fscope.Add(n.Name, NewValue(&Type{R: resultTypes[i]}))
					}
					i++
				}
			}
		}

		_, branch, err := fscope.ExecStmts(stmts)
		if err != nil {
			// xxx how to communicate this? panic?
		}
		var result []reflect.Value
		if branch != nil {
			if branch.typ != branchReturn {
				// xxx err
			}

			if hasNamedReturns && len(branch.vals) == 0 {
				for _, r := range typ.Results.List {
					for _, n := range r.Names {
						val := fscope.LookupValue(n.Name)
						if val == nil {
							// xxx err, but should be impossible
						}
						result = append(result, val)
					}
				}
			} else if len(branch.vals) == len(resultTypes) {
				for i, val := range branch.vals {
					if !val.Type().AssignableTo(refl.Type{R: resultTypes[i]}) {
						// xxx err
					}
					result = append(result, val)
				}
			}
		}
		return result
	}

	return refl.Value{R: reflect.MakeFunc(ftype.R, impl)}, nil
}

func (s *Scope) evalCall(expr *ast.CallExpr) ([]*refl.Value, error) {
	if ident, ok := expr.Fun.(*ast.Ident); ok {
		b := s.Lookup(ident.Name)
		if b, ok := b.(builtin); ok {
			return s.evalBuiltinCall(b, expr.Args, expr.Ellipsis != token.NoPos)
		}
	}

	ft, err := s.EvalType(expr.Fun)
	if err == nil {
		if len(expr.Args) != 1 {
			return nil, fmt.Errorf("got %d args for type conversion, want 1", len(expr.Args))
		}
		val, err := s.Eval1(expr.Args[0])
		if err != nil {
			return nil, err
		}
		return val.Convert(ft)
	}

	f, err := s.Eval1(expr.Fun)
	if err != nil {
		return nil, errors.Wrap(err, "eval func of call expr")
	}

	var vals []*refl.Value
	if len(expr.Args) == 1 {
		// "As a special case, if the return values of a function or
		// method g are equal in number and individually assignable to the
		// parameters of another function or method f, then the call
		// f(g(parameters_of_g)) will invoke f after binding the return
		// values of g to the parameters of f in order."
		// https://golang.org/ref/spec#Calls
		vals, err = s.Eval(expr.Args[0])
		if err != nil {
			return nil, err
		}
	} else {
		for _, a := range args {
			val, err := s.Eval1(a)
			if err != nil {
				return nil, err
			}
			if xxx /* val is not a reflect.Value */ {
				// xxx err
			}
			vals = append(vals, val.R)
		}
	}
	return f.Call(vals)
}

func (s *Scope) evalBuiltinCall(b builtin, args []ast.Expr, ellipsis bool) ([]*refl.Value, error) {
	switch b {
	case builtinBool:
		return s.doConvert(args, reflect.TypeOf(true))

	case builtinByte:
		return s.doConvert(args, reflect.TypeOf(byte(0)))

	case builtinComplex64:
		return s.doConvert(args, reflect.TypeOf(complex64(0)))

	case builtinComplex128:
		return s.doConvert(args, reflect.TypeOf(complex128(0)))

	case builtinError:
		return s.doConvert(args, reflect.TypeOf(error(nil)))

	case builtinFloat32:
		return s.doConvert(args, reflect.TypeOf(float32(0)))

	case builtinFloat64:
		return s.doConvert(args, reflect.TypeOf(float64(0)))

	case builtinInt:
		return s.doConvert(args, reflect.TypeOf(int(0)))

	case builtinInt8:
		return s.doConvert(args, reflect.TypeOf(int8(0)))

	case builtinInt16:
		return s.doConvert(args, reflect.TypeOf(int16(0)))

	case builtinInt32:
		return s.doConvert(args, reflect.TypeOf(int32(0)))

	case builtinInt64:
		return s.doConvert(args, reflect.TypeOf(int64(0)))

	case builtinRune:
		return s.doConvert(args, reflect.TypeOf(rune(0)))

	case builtinString:
		return s.doConvert(args, reflect.TypeOf(""))

	case builtinUint:
		return s.doConvert(args, reflect.TypeOf(uint(0)))

	case builtinUint8:
		return s.doConvert(args, reflect.TypeOf(uint8(0)))

	case builtinUint16:
		return s.doConvert(args, reflect.TypeOf(uint16(0)))

	case builtinUint32:
		return s.doConvert(args, reflect.TypeOf(uint32(0)))

	case builtinUint64:
		return s.doConvert(args, reflect.TypeOf(uint64(0)))

	case builtinUintptr:
		return s.doConvert(args, reflect.TypeOf(uintptr(0)))

	case builtinImag:
		return s.doConvert(args, reflect.TypeOf(uintptr(0)))

	case builtinReal:
		return s.doConvert(args, reflect.TypeOf(uintptr(0)))

	case builtinTrue, builtinFalse, builtinIota, builtinNil:
		return nil, fmt.Errorf("not a function")

	case builtinAppend:
		if len(args) < 2 {
			return nil, fmt.Errorf("got %d arg(s), want >=2", len(args))
		}
		var vals []*refl.Value
		for _, arg := range args {
			val, err := s.Eval1(arg)
			if err != nil {
				return nil, err
			}
			vals = append(vals, val)
		}
		slice := vals[0] // xxx check it's a slice
		if slice.Kind() != reflect.Slice {
			return nil, fmt.Errorf("first arg of append is not a slice")
		}
		elemType := slice.Type().Elem()
		for i := 1; i < len(vals)-1; i++ {
			if !vals[i].Type().AssignableTo(elemType) {
				return nil, fmt.Errorf("type mismatch in append")
			}
			slice = slice.Append(vals[i])
		}
		last := vals[len(vals)-1]
		if ellipsis {
			if last.Type() != slice.Type() {
				return nil, fmt.Errorf("'...' argument to append is not a slice")
			}
			if !last.Type().Elem().AssignableTo(elemType) {
				return nil, fmt.Errorf("'...' argument to append has wrong element type")
			}
			slice = slice.AppendSlice(last)
		} else {
			if !last.Type().AssignableTo(elemType) {
				return nil, fmt.Errorf("type mismatch in append")
			}
			slice = slice.Append(last)
		}
		return []*refl.Value{slice}, nil

	case builtinCap:
		if len(args) != 1 {
			return nil, fmt.Errorf("got %d args to cap, want 1", len(args))
		}
		val, err := s.Eval1(args[0])
		if err != nil {
			return nil, err
		}
		c, err := val.Cap()
		if err != nil {
			return nil, err
		}
		return []*refl.Value{{R: reflect.ValueOf(c)}}, nil

	case builtinClose:
		if len(args) != 1 {
			return nil, fmt.Errorf("got %d args to close, want 1", len(args))
		}
		val, err := s.Eval1(args[0])
		if err != nil {
			return nil, err
		}
		err = val.Close()
		return nil, err

	case builtinComplex:
		if len(args) != 2 {
			return nil, fmt.Errorf("got %d arg(s) to complex, want 2", len(args))
		}
		re, err := s.Eval1(args[0])
		if err != nil {
			return nil, err
		}
		im, err := s.Eval1(args[1])
		if err != nil {
			return nil, err
		}
		// xxx check types of re and im
		return []*refl.Value{{R: reflect.ValueOf(complex(re, im))}}, nil

	case builtinCopy:
		if len(args) != 2 {
			// xxx err
		}
		dst, err := s.Eval1(args[0])
		if err != nil {
			return nil, err
		}
		src, err := s.Eval1(args[1])
		if err != nil {
			return nil, err
		}
		switch dst.Kind() {
		case ArrayKind, SliceKind:
			// ok
		default:
			// xxx err
		}
		switch src.Kind() {
		case ArrayKind, SliceKind:
			// ok
		default:
			// xxx err
		}
		if dst.Type().Elem() != src.Type().Elem() {
			// xxx err
		}
		n := Copy(dst, src)
		return []*refl.Value{{R: reflect.ValueOf(n)}}, nil

	case builtinDelete:
		if len(args) != 2 {
			// xxx err
		}
		m, err := s.Eval1(args[0])
		if err != nil {
			return nil, err
		}
		if m.Kind() != MapKind {
			// xxx err
		}
		key, err := s.Eval1(args[1])
		if err != nil {
			return nil, err
		}
		if !key.Type().AssignableTo(m.Type().Key()) {
			// xxx err
		}
		m.Delete(key)
		return nil, nil

	case builtinLen:
		if len(args) != 1 {
			// xxx err
		}
		val, err := s.Eval1(args[0])
		if err != nil {
			return nil, err
		}
		l, err := val.Len()
		if err != nil {
			return nil, err
		}
		return []*refl.Value{{R: reflect.ValueOf(l)}}, nil

	case builtinMake:

	case builtinNew:
		if len(args) != 1 {
			// xxx err
		}
		typ, err := s.EvalType(args[0])
		if err != nil {
			return nil, err
		}
		return []*refl.Value{NewValue(typ)}, nil

	case builtinPanic:
	case builtinPrint:
	case builtinPrintln:
	case builtinRecover:
	}
}

func (s *Scope) doConvert(args []ast.Expr, rt reflect.Type) ([]*refl.Value, error) {
	if len(args) != 1 {
		// xxx err
	}
	val, err := s.Eval1(args[0])
	if err != nil {
		return nil, err
	}
	conv, err := val.Convert(&Type{R: rt})
	return []*refl.Value{conv}, err
}
