package ego

import (
	"go/ast"
	"reflect"
)

func (s *Scope) execSelect(stmt *ast.SelectStmt, label string) (*Scope, *branch, error) {
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

		switch comm := clause.Comm.(type) {
		case nil:
			rcase.Dir = reflect.SelectDefault

		case *ast.SendStmt:
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

		case *ast.AssignStmt:
			// xxx

		case *ast.ExprStmt:
			// xxx

		default:
			// xxx err
		}

		rcases = append(rcases, rcase)
	}
	chosen, received, ok := reflect.Select(rcases)
	clause := commClauses[chosen]
	// xxx maybe assign to vars for received and/or ok, maybe in a new scope
	_, b, err := s.ExecStmts(clause.Body)
	if err != nil {
		return nil, nil, err
	}
	if b != nil {
		switch b.typ {
		case branchBreak:
		case branchContinue:
		case branchGoto:
		case branchFallthrough:
		case branchReturn:
		}
	}
	return origScope, nil, nil
}
