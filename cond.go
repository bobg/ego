package ego

import (
	"fmt"
	"go/ast"
	"go/token"
)

func (s *Scope) execIf(stmt *ast.IfStmt) (*Scope, *branch, error) {
	origScope := s
	if stmt.Init != nil {
		var err error
		s, _, err = s.ExecStmt(stmt.Init)
		if err != nil {
			return nil, nil, err
		}
	}
	cond, err := s.Eval1(stmt.Cond)
	if err != nil {
		return nil, nil, err
	}
	boolval, err := cond.Bool()
	if err != nil {
		return nil, nil, err
	}
	if boolval {
		_, b, err := s.ExecStmts(stmt.Body.List)
		return origScope, b, err
	}
	if stmt.Else != nil {
		_, b, err := s.ExecStmt(stmt.Else)
		return origScope, b, err
	}
	return origScope, nil, nil
}

func (s *Scope) execSwitch(stmt *ast.SwitchStmt, label string) (*Scope, *branch, error) {
	var err error

	origScope := s
	if stmt.Init != nil {
		s, _, err = s.ExecStmt(stmt.Init)
		if err != nil {
			return nil, nil, err
		}
	}
	var tag *refl.Value
	if stmt.Tag != nil {
		tag, err = s.Eval1(stmt.Tag)
		if err != nil {
			return nil, nil, err
		}
	} else {
		tag = refl.ValueOf(true)
	}
	fallingThrough := false
	for _, item := range stmt.Body.List {
		clause, ok := item.(*ast.CaseClause)
		if !ok {
			return nil, nil, fmt.Errorf("got statement type %T in switch, want CaseClause", item)
		}
		if !fallingThrough && len(clause.List) > 0 {
			match := false
			for _, e := range clause.List {
				v, err := s.Eval1(e)
				if err != nil {
					return nil, nil, err
				}
				if xxx /* v == tag, with proper type jiggering */ {
					match = true
					break
				}
			}
			if !match {
				continue
			}
		}
		// default clause, falling through, or match
		fallingThrough = false
		_, b, err := s.ExecStmts(clause.Body)
		if err != nil {
			return nil, nil, err
		}
		if b != nil {
			switch b.typ {
			case branchBreak:
				if b.label == "" || b.label == label {
					return origScope, nil, nil
				}
				return origScope, b, nil

			case branchContinue:
				return origScope, b, nil

			case branchGoto:
				return origScope, b, nil

			case branchFallthrough:
				fallingThrough = true
			}
		}
		if !fallingThrough {
			return origScope, nil, nil
		}
	}
	return origScope, nil, nil
}

func (s *Scope) execTypeSwitch(stmt *ast.TypeSwitchStmt, label string) (*Scope, *branch, error) {
	var err error

	origScope := s
	if stmt.Init != nil {
		s, _, err = s.ExecStmt(stmt)
		if err != nil {
			return nil, nil, err
		}
	}

	var (
		varName    string // for "switch x := y.(type)"
		assignExpr ast.Expr
	)
	switch a := stmt.Assign.(type) {
	case *ast.AssignStmt:
		if a.Tok != token.DEFINE {
			return nil, nil, fmt.Errorf("assignment must use := in type switch")
		}
		if len(a.Lhs) != 1 {
			return nil, nil, fmt.Errorf("got %d items on lhs of assignment in type switch, may only have 1", len(a.Lhs))
		}
		id, ok := a.Lhs[0].(*ast.Ident)
		if !ok {
			return nil, nil, fmt.Errorf("got %T on lhs of assignment in type switch, want Ident", a.Lhs[0])
		}
		varName = id.Name
		if len(a.Rhs) != 1 {
			return nil, nil, fmt.Errorf("got %d items on rhs of assignment in type switch, may only have 1", len(a.Rhs))
		}
		assignExpr = a.Rhs[0]

	case *ast.ExprStmt:
		assignExpr = a.X

	default:
		return nil, nil, fmt.Errorf("got %T in type switch guard, want ExprStmt or AssignStmt", a)
	}

	val, err := s.Eval1(assignExpr)
	if err != nil {
		return nil, nil, err
	}

	for _, clause := range stmt.Body.List {
		clause, ok := clause.(*ast.CaseClause)
		if !ok {
			return nil, nil, fmt.Errorf("got statement type %T in type switch, want CaseClause", item)
		}
		var matchType *refl.Type
		if len(clause.List) > 0 {
			for _, e := range clause.List {
				t, err := s.EvalType(e)
				if err != nil {
					return nil, nil, err
				}
				if val.Type().AssignableTo(t) {
					matchType = t
					break
				}
			}
			if matchType == nil {
				continue
			}
			// xxx if len(clause.List) > 1, varName does not get the type
			// matchType, but the type of assignExpr.
			if varName != "" {
				s = NewScope(s)

				// This makes sure variable varName has the right type in
				// scope s, even if val is a different (but assignable) type.
				v := refl.NewValue(matchType)
				v.Set(val)
				s.Add(varName, v)
			}
		}
		// default clause or match
		_, b, err := s.ExecStmts(clause.Body)
		if err != nil {
			return nil, nil, err
		}
		if b != nil {
			switch b.typ {
			case branchBreak:
				if b.label == "" || b.label == label {
					return origScope, nil, nil
				}
				return origScope, b, nil

			case branchContinue:
				return origScope, b, nil

			case branchGoto:
				return origScope, b, nil

			case branchFallthrough:
				return nil, nil, fmt.Errorf("fallthrough not allowed in type switch")
			}
		}
		break
	}

	return origScope, nil, nil
}
