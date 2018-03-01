package ego

import "go/ast"

func (s *Scope) execIf(stmt *ast.IfStmt) (*Scope, *branch, error) {
	origScope := s
	if stmt.Init != nil {
		s, b, err := s.ExecStmt(stmt.Init)
		if err != nil {
			return nil, nil, err
		}
		if b != nil {
			// xxx impossible?
		}
	}
	cond, err := s.Eval1(stmt.Cond)
	if err != nil {
		return nil, nil, err
	}
	// xxx check cond is boolean
	if xxx /* cond is true */ {
		_, b, err := s.ExecStmts(stmt.Body.List)
		return origScope, b, err
	}
	if stmt.Else != nil {
		_, b, err := s.ExecStmt(stmt.Else)
		return origScope, b, err
	}
	return origScope, nil, nil
}

func (s *Scope) execSwitch(stmt *ast.SwitchStmt) (*Scope, *branch, error) {
}

func (s *Scope) execTypeSwitch(stmt *ast.TypeSwitchStmt) (*Scope, *branch, error) {
}
