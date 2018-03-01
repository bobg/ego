package ego

import (
	"go/ast"
	"go/token"
	"math"
	"reflect"
	"strconv"
)

func (scope *Scope) evalType(expr ast.Expr) (reflect.Type, error) {
	switch expr := expr.(type) {
	case *ast.Ident:
	case *ast.ParenExpr:
	case *ast.SelectorExpr:
	case *ast.StarExpr:
	case *ast.ArrayType:
		eltType, err := scope.evalType(expr.Elt)
		if err != nil {
			return nil, err
		}
		if expr.Len == nil {
			// slice
			return reflect.SliceOf(eltType), nil
		}
		if _, ok := expr.Len.(*ast.Ellipsis); ok {
			// xxx err
		}
		if lit, ok := expr.Len.(*ast.BasicLit); ok {
			if lit.Kind != token.INT {
				// xxx err
			}
			l, err := strconv.ParseInt(lit.Value, 0, 64)
			if err != nil {
				return nil, err
			}
			if l < 0 || l > math.MaxInt32 {
				// xxx err
			}
			return reflect.ArrayOf(l, eltType), nil
		}
		// xxx err

	case *ast.StructType:
		var fields []reflect.StructField
		for _, f := range expr.Fields.List {
			ftype, err := scope.evalType(f.Type)
			if err != nil {
				return nil, err
			}
			for _, ident := range f.Names {
				sf := reflect.StructField{
					Type: ftype,
				}
				if ident == nil {
					sf.Anonymous = true
				} else {
					sf.Name = ident.Name
				}
				if f.Tag != nil {
					sf.Tag = f.Tag.Value
				}
				fields = append(fields, sf)
			}
		}
		return reflect.StructOf(fields), nil

	case *ast.FuncType:
	case *ast.InterfaceType:
		// xxx shazbot, there's no reflect.InterfaceOf! See
		// https://github.com/golang/go/issues/4146

	case *ast.MapType:
		keyType, err := scope.evalType(expr.Key)
		if err != nil {
			return nil, err
		}
		valType, err := scope.evalType(expr.Value)
		if err != nil {
			return nil, err
		}
		return reflect.MapOf(keyType, valType), nil

	case *ast.ChanType:
		eltType, err := scope.evalType(expr.Value)
		if err != nil {
			return nil, err
		}
		var d reflect.ChanDir
		if (expr.Dir & ast.SEND) != 0 {
			d |= reflect.SendDir
		}
		if (expr.Dir & ast.RECV) != 0 {
			d |= reflect.RecvDir
		}
		return reflect.ChanOf(d, eltType), nil
	}
	// xxx err
}
