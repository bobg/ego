package ego

import (
	"go/ast"

	"github.com/bobg/ego/refl"
)

type branchType int

const (
	branchBreak branchType = 1 + iota
	branchContinue
	branchGoto
	branchFallthrough
	branchReturn
)

type branch struct {
	typ   branchType
	vals  []refl.Value
	label string
}

func (s *Scope) ExecStmts(stmts []ast.Stmt) (*Scope, *branch, error) {
	origScope := s
	scopes := make([]*Scope, len(stmts))

	for i := 0; i < len(stmts); {
		stmt := stmts[i]
		scopes[i] = s
		s, b, err := s.ExecStmt(stmt)
		if err != nil {
			return nil, nil, err
		}
		if b != nil {
			switch b.typ {
			case branchBreak:
				return origScope, b, nil

			case branchContinue:
				return origScope, b, nil

			case branchGoto:
				var found bool
				for j := 0; j <= i; j++ {
					if l, ok := stmts[j].(*ast.LabeledStmt); ok && l.Label.Name == b.label {
						i = j
						s = scopes[j]
						found = true
						break
					}
				}
				if !found {
					var newScope bool
					for j := i + 1; j < len(stmts); j++ {
						if l, ok := stmts[j].(*ast.LabeledStmt); ok && l.Label.Name == b.label {
							if newScope {
								// xxx err - cannot goto into a new scope
							}
							i = j
							found = true
							break
						}
						if isNewScopeStmt(stmt) {
							newScope = true
						}
					}
				}
				if found {
					continue
				}
				return origScope, b, nil

			case branchFallthrough:
				return origScope, b, nil

			case branchReturn:
				return origScope, b, nil
			}
		}
	}
	return origScope, nil, nil
}

func (s *Scope) ExecStmt(stmt ast.Stmt) (*Scope, *branch, error) {
	switch stmt := stmt.(type) {
	case *ast.DeclStmt:
		// xxx

	case *ast.EmptyStmt:
		return s, nil, nil

	case *ast.LabeledStmt:
		switch inner := stmt.Stmt.(type) {
		case *ast.ForStmt:
			return s.execFor(inner, stmt.Label.Name)

		case *ast.SwitchStmt:
			return s.execSwitch(inner, stmt.Label.Name)

		case *ast.TypeSwitchStmt:
			return s.execTypeSwitch(inner, stmt.Label.Name)

		case *ast.SelectStmt:
			return s.execSelect(inner, stmt.Label.Name)
		}
		return s.ExecStmt(stmt.Stmt)

	case *ast.ExprStmt:
		_, err := s.Eval(stmt.X)
		return s, nil, err

	case *ast.SendStmt:
		return s.execSend(stmt)

	case *ast.IncDecStmt:
		return s.execIncDec(stmt)

	case *ast.AssignStmt:
		return s.execAssign(stmt)

	case *ast.GoStmt:
		go func() {
			// xxx need some way to communicate errors
			s.evalCall(stmt.Call)
		}()

	case *ast.DeferStmt:
		return s.execDefer(stmt)

	case *ast.ReturnStmt:
		return s.execReturn(stmt)

	case *ast.BranchStmt:
		return s.execBranch(stmt)

	case *ast.BlockStmt:
		return s.ExecStmts(stmt.List)

	case *ast.IfStmt:
		return s.execIf(stmt)

	case *ast.CaseClause:
		// xxx error - outside switch or typeswitch

	case *ast.SwitchStmt:
		return s.execSwitch(stmt, "")

	case *ast.TypeSwitchStmt:
		return s.execTypeSwitch(stmt, "")

	case *ast.CommClause:
		// xxx error - outside select

	case *ast.SelectStmt:
		return s.execSelect(stmt, "")

	case *ast.ForStmt:
		return s.execFor(stmt, "")

	case *ast.RangeStmt:
		return s.execRange(stmt)
	}
}
