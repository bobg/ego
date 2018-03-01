package ego

import (
	"go/ast"
	"reflect"
)

func (s *Scope) execSelect(stmt *ast.SelectStmt) (*Scope, *branch, error) {
	origScope := s
	var (
		rcases      []reflect.SelectCase
		commClauses []*ast.CommClause
	)

	for _, sub := range stmt.Body.List {
		clause, ok := sub.(*ast.CommClause)
		if !ok {
			// xxx err
		}
		commClauses = append(commClauses, clause)
		var rcase reflect.SelectCase
		if clause.Comm == nil {
			rcase.Dir = reflect.SelectDefault
		} else if clause, ok := clause.Comm.(*ast.SendStmt); ok {
			ch, err := s.Eval1(clause.Chan)
			if err != nil {
				return nil, err
			}
			if ch.Kind() != refl.Chan {
				// xxx err
			}
			sendval, err := s.Eval1(clause.Value)
			if err != nil {
				return nil, err
			}
			if sendval.Type() != ch.Type().Elem() {
				// xxx err
			}
			rcase.Dir = reflect.SelectSend
			rcase.Chan = ch
			rcase.Send = sendval
		} else {
			// xxx check it's <-ch or x := <-ch or x, ok := ch
		}
		rcases = append(rcases, rcase)
	}
	chosen, received, ok := reflect.Select(rcases)
	clause := commClauses[chosen]
	// xxx maybe assign to vars for received and/or ok, maybe in a new scope
	s, b, err := s.ExecStmts(clause.Body)
	if err != nil {
		return nil, nil, err
	}
	if b != nil {
		// xxx handle branching
	}
	return origScope, nil, nil
}
