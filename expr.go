package ego

import (
	"go/ast"
	"go/token"
	"math"
	"reflect"
	"strconv"
)

func (scope *Scope) eval1(expr ast.Expr) (val reflect.Value, err error) {
	vals, err := scope.eval(expr)
	if err != nil {
		return val, err
	}
	if len(vals) != 1 {
		return val, xxxerr
	}
	return vals[0], nil
}

func (scope *Scope) eval(expr ast.Expr) ([]reflect.Value, error) {
	switch expr := expr.(type) {
	case *ast.Ident:
		return scope.Lookup(expr.Name)

	case *ast.Ellipsis:
		// xxx

	case *ast.BasicLit:
		switch expr.Kind {
		case token.INT:
			for _, bitsize := range []int{8, 16, 32, 64} {
				val, err := strconv.ParseInt(expr.Value, 0, bitsize)
				if err == nil {
					switch bitsize {
					case 8:
						x := int8(val)
						return []reflect.Value{reflect.ValueOf(x)}, nil
					case 16:
						x := int16(val)
						return []reflect.Value{reflect.ValueOf(x)}, nil
					case 32:
						x := int32(val)
						return []reflect.Value{reflect.ValueOf(x)}, nil
					case 64:
						return []reflect.Value{reflect.ValueOf(val)}, nil
					}
				}
				uval, err := strconv.ParseUint(expr.Value, 0, bitsize)
				if err == nil {
					switch bitsize {
					case 8:
						x := uint8(uval)
						return []reflect.Value{reflect.ValueOf(x)}, nil
					case 16:
						x := uint16(uval)
						return []reflect.Value{reflect.ValueOf(x)}, nil
					case 32:
						x := uint32(uval)
						return []reflect.Value{reflect.ValueOf(x)}, nil
					case 64:
						return []reflect.Value{reflect.ValueOf(uval)}, nil
					}
				}
			}
			// xxx err

		case token.FLOAT:
			for _, bitsize := range []int{32, 64} {
				val, err := strconv.ParseFloat(expr.Value, bitsize)
				if err == nil {
					switch bitsize {
					case 32:
						x := float32(val)
						return []reflect.Value{reflect.ValueOf(x)}, nil
					case 64:
						return []reflect.Value{reflect.ValueOf(val)}, nil
					}
				}
			}

		case token.IMAG:
			// xxx

		case token.CHAR:
			// xxx double-check expr.Value does not need parsing/escaping
			return []reflect.Value{reflect.ValueOf(expr.Value[0])}, nil

		case token.STRING:
			// xxx double-check expr.Value does not need parsing/escaping
			return []reflect.Value{reflect.ValueOf(expr.Value)}, nil
		}

	case *ast.FuncLit:
		typ := convFuncType(expr.Type)
		return []reflect.Value{reflect.MakeFunc(typ, fnWrapper(scope, nil, expr.Type, expr.Body))}, nil

	case *ast.CompositeLit:
		typ, err := scope.evalType(expr.Type)
		if err != nil {
			return nil, err
		}
		switch typ.Kind() {
		case reflect.Struct:
			var (
				result   = reflect.New(typ)
				elemType = typ.Elem()
				expectKV bool
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
					val, err := scope.eval1(elt.Value)
					if err != nil {
						return nil, err
					}
					tField, ok := typ.FieldByName(key.Name)
					if !ok {
						// xxx err
					}
					vField := result.FieldByIndex(tField.Index)
					vField.Set(val)
				} else {
					if expectKV {
						// xxx err
					}
					if i >= typ.NumField() {
						// xxx err
					}
					val, err := scope.eval1(elt.Value)
					if err != nil {
						return nil, err
					}
					vField := result.FieldByIndex(i)
					vField.Set(val)
				}
			}
			return result, nil

		case reflect.Array:
		case reflect.Slice:
			result := reflect.New(typ)
			for _, elt := range expr.Elts {
				val, err := scope.eval1(elt)
				if err != nil {
					return nil, err
				}
				if !reflect.TypeOf(val).AssignableTo(elemType) {
					// xxx err
				}
				result = reflect.Append(result, val)
			}
			return result, nil

		case reflect.Map:
			// xxx
		}
		// xxx err

	case *ast.ParenExpr:
		return scope.eval(expr.X)

	case *ast.SelectorExpr:
		baseVal, err := scope.eval1(expr.X)
		if err != nil {
			return nil, err
		}
		// xxx resolve expr.Sel in the context of baseVal

	case *ast.IndexExpr:
		baseVal, err := scope.eval1(expr.X)
		if err != nil {
			return nil, err
		}
		indexVal, err := scope.eval1(expr.Index)
		if err != nil {
			return nil, err
		}
		baseValType := reflect.TypeOf(baseVal)
		switch baseValType.Kind() {
		case reflect.Array, reflect.Slice, reflect.String:
			indexValType := reflect.TypeOf(indexVal)
			switch indexValType {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				index := int(indexVal.Int()) // xxx do range checking
				if index < 0 || index >= baseVal.Len() {
					// xxx err
				}
				return []reflect.Value{baseVal.Index(index)}, nil
			default:
				// xxx err
			}

		case reflect.Map:
			if reflect.TypeOf(baseVal[0]).Key() != reflect.TypeOf(indexVal[0]) {
				// xxx err
			}
			return []reflect.Value{baseVal[0].MapIndex(indexVal[0])}, nil

		default:
			// xxx err
		}

	case *ast.SliceExpr:
		sub, err := scope.eval1(expr.X)
		if err != nil {
			return nil, err
		}
		sub = resolveArrayPtr(sub)
		switch sub.Kind() {
		case reflect.Array, reflect.Slice, reflect.String:
			// ok
		default:
			// xxx error
		}
		var (
			res            reflect.Value
			low, high, max int
		)
		if expr.Low != nil {
			lowVal, err := scope.eval1(expr.Low)
			if err != nil {
				return nil, err
			}
			low, err = toInt(lowVal)
			if err != nil {
				return nil, err
			}
			if low < 0 {
				// xxx err
			}
		}
		if expr.High == nil {
			high = sub.Len()
		} else {
			highVal, err := scope.eval1(expr.High)
			if err != nil {
				return nil, err
			}
			high, err = toInt(highVal)
			if err != nil {
				return nil, err
			}
			if high < low {
				// xxx err
			}
		}

		if expr.Slice3 {
			if expr.Max == nil {
				// xxx err
			}
			maxVal, err := scope.eval1(expr.Max)
			if err != nil {
				return nil, err
			}
			max, err = toInt(maxVal)
			if err != nil {
				return nil, err
			}
			if max < high {
				// xxx err
			}
			res = sub.Slice3(low, high, max)
		} else {
			res = sub.Slice(low, high)
		}
		return []reflect.Value{res}, nil

	case *ast.TypeAssertExpr:
		

	case *ast.CallExpr:
		return scope.evalCall(expr)

	case *ast.StarExpr:
		sub, err := scope.eval1(expr.X)
		if err != nil {
			return nil, err
		}
		if sub.Kind() != reflect.Ptr {
			// xxx err
		}
		return []reflect.Value{sub.Elem()}, nil

	case *ast.UnaryExpr:
		sub, err := scope.eval1(expr.X)
		if err != nil {
			return nil, err
		}
		switch expr.Op {
		case token.ADD:
		case token.SUB:
		case token.XOR:
		case token.NOT:
		default:
			// xxx err
		}

	case *ast.BinaryExpr:
	case *ast.KeyValueExpr:
	}
}

func toInt(val reflect.Value) (int, error) {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		x := val.Int()
		if x > math.MaxInt32 {
			// xxx err
		}
		if x < math.MinInt32 {
			// xxx err
		}
		return int(x), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		x := val.Uint()
		if x > math.MaxInt32 {
			// xxx err
		}
		return int(x), nil
	}
	// xxx error
}
