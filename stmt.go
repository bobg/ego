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
	for _, stmt := range stmts {
		s, b, err := s.ExecStmt(stmt)
		if err != nil {
			return nil, nil, err
		}
		if b != nil {
			switch b.typ {
			case branchBreak:
				// xxx
			case branchContinue:
				// xxx
			case branchGoto:
				// xxx
			case branchFallthrough:
				// xxx
			case branchReturn:
				return s, b, nil
			}
		}
	}
	return s, nil, nil
}

func (s *Scope) ExecStmt(stmt ast.Stmt) (*Scope, *branch, error) {
	switch stmt := stmt.(type) {
	case *ast.DeclStmt:
		// xxx

	case *ast.EmptyStmt:
		return s, nil, nil

	case *ast.LabeledStmt:
		s, b, err := s.ExecStmt(stmt.Stmt)
		if err != nil {
			return nil, nil, err
		}
		if b != nil {
			switch branch.typ {
			case branchBreak:
			case branchContinue:
			case branchGoto:
			case branchFallthrough:
			case branchReturn:
				return s, b, nil
			}
		}
		return s, nil, nil

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
		return s.execSwitch(stmt)

	case *ast.TypeSwitchStmt:
		return s.execTypeSwitch(stmt)

	case *ast.CommClause:
		// xxx error - outside select

	case *ast.SelectStmt:
		return s.execSelect(stmt)

	case *ast.ForStmt:
		return s.execFor(stmt)

	case *ast.RangeStmt:
		return s.execRange(stmt)
	}
}
