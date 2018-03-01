package ego

import (
	"go/ast"

	"github.com/bobg/ego/refl"
)

func (s *Scope) execDefer(stmt *ast.DeferStmt) (*Scope, *branch, error) {
	f, err := s.Eval1(stmt.Call.Fun)
	if err != nil {
		return nil, nil, err
	}
	var vals []refl.Value
	for _, a := range args {
		val, err := s.Eval1(a)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	// xxx walk up the scope hierarchy to the nearest enclosing function
	// scope and add (f,vals) to the defer list
	return s, nil, nil
}
