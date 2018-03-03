package ego

import "go/ast"

func (s *Scope) execFor(stmt *ast.ForStmt, label string) (*Scope, *branch, error) {
	origScope := s
	if stmt.Init != nil {
		var err error
		s, _, err = s.Exec(stmt.Init) // No branching possible here, only "simple statments" allowed.
		if err != nil {
			return nil, nil, err
		}
	}
	for {
		if stmt.Cond != nil {
			val, err := s.Eval1(stmt.Cond)
			if err != nil {
				return nil, nil, err
			}
			b, err := val.Bool()
			if err != nil {
				return nil, nil, err
			}
			if !b {
				return origScope, nil, nil
			}
		}
		_, b, err := s.ExecStmts(stmt.Body.List)
		if err != nil {
			return nil, nil, err
		}
		if b != nil {
			switch b.typ {
			case branchBreak:
				if b.label == "" || b.label == label {
					return origScope, nil, nil
				}

			case branchContinue:
				if b.label != "" && b.label != label {
					return origScope, b, nil
				}

			case branchGoto:
				return origScope, b, nil

			case branchFallthrough:
				// xxx err

			case branchReturn:
				return origScope, b, nil
			}
		}
		if stmt.Post != nil {
			_, _, err = s.ExecStmt(stmt.Post)
			if err != nil {
				return nil, nil, err
			}
		}
	}
}

func (s *Scope) execRange(stmt *ast.RangeStmt, label string) (*Scope, *branch, error) {
	origScope := s

	if stmt.Tok == token.DEFINE {
		s = NewScope(s)
	}

	// xxx do not eval X if at most one iteration var is present and
	// len(X) is constant (https://golang.org/ref/spec#RangeClause)
	rexp, err := s.Eval1(stmt.X)
	if err != nil {
		return nil, nil, err
	}

	switch rexp.Kind() {
	case refl.Array, refl.Slice, refl.String:
		l, _ := rexp.Len()
		for i := 0; i < l; i++ {
			// xxx assign to key, value if present
			_, b, err := s.ExecStmts(stmt.Body.List)
			if err != nil {
				return nil, nil, err
			}
			if b != nil {
				// xxx
			}
		}
		return origScope, nil, nil

	case refl.Map:
	case refl.Chan:
	}
}
