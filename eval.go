package ego

import (
	"go/ast"

	"github.com/bobg/ego/refl"
)

func (s *Scope) Eval1(expr ast.Expr) (refl.Value, error) {
	vals, err := s.Eval(expr)
	if err != nil {
		return nil, err
	}
	if len(vals) != 1 {
		// xxx err
	}
	return vals[0], nil
}

// Eval evaluates non-type expressions.
func (s *Scope) Eval(expr ast.Expr) ([]*refl.Value, error) {
	switch expr := expr.(type) {
	case *ast.Ident:
		return s.LookupValue(expr.Name)

	case *ast.Ellipsis:
		// xxx error - ellipsis out of context

	case *ast.BasicLit:
		val, err := s.evalBasicLit(expr)
		return []*refl.Value{val}, err

	case *ast.FuncLit:
		wrappedFn, err := s.evalFuncLit(expr)
		return []*refl.Value{wrappedFn}, err

	case *ast.CompositeLit:
		val, err := s.evalCompositeLit(expr)
		return []*refl.Value{val}, err

	case *ast.ParenExpr:
		return s.Eval(expr.X)

	case *ast.SelectorExpr:
		val, err := s.evalSelector(expr)
		return []*refl.Value{val}, err

	case *ast.IndexExpr:
		val, err := s.evalIndex(expr)
		return []*refl.Value{val}, err

	case *ast.SliceExpr:
		val, err := s.evalSlice(expr)
		return []*refl.Value{val}, err

	case *ast.TypeAssertExpr:
		val, err := s.evalTypeAssert(expr)
		return []*refl.Value{val}, err

	case *ast.CallExpr:
		return s.evalCall(expr)

	case *ast.StarExpr:
		val, err := s.evalStar(expr)
		return []*refl.Value{val}, err

	case *ast.UnaryExpr:
		val, err := s.evalUnary(expr)
		return []*refl.Value{val}, err

	case *ast.BinaryExpr:
		val, err := s.evalBinary(expr)
		return []*refl.Value{val}, err

	case *ast.KeyValueExpr:
		// xxx error - out of context
	}
}
