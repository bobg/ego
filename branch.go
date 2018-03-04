package ego

import (
	"errors"
	"go/ast"
	"go/token"

	"github.com/bobg/ego/refl"
)

var errBadBranch = errors.New("bad branch")

func (s *Scope) execReturn(stmt *ast.ReturnStmt) (*Scope, *branch, error) {
	var (
		vals []*refl.Value
		err  error
	)
	if len(stmt.Results) == 1 {
		vals, err = s.Eval(stmt.Results[0])
		if err != nil {
			return nil, nil, err
		}
	} else {
		for _, r := range stmt.Results {
			val, err := s.Eval1(r)
			if err != nil {
				return nil, nil, err
			}
			vals = append(vals, val)
		}
	}
	return s, &branch{typ: branchReturn, vals: vals}, nil
}

func (s *Scope) execBranch(stmt *ast.BranchStmt) (*Scope, *branch, error) {
	var label string
	if stmt.Label != nil {
		label = stmt.Label.Name
	}
	switch stmt.Tok {
	case token.BREAK:
		return s, &branch{typ: branchBreak, label: label}, nil
	case token.CONTINUE:
		return s, &branch{typ: branchContinue, label: label}, nil
	case token.GOTO:
		return s, &branch{typ: branchGoto, label: label}, nil
	case token.FALLTHROUGH:
		return s, &branch{typ: branchFallthrough}, nil
	}
	return nil, nil, errBadBranch
}
