package ego

import (
	"go/ast"
	"go/token"
	"reflect"
)

func execStmts(scope *Scope, stmts []ast.Stmt) (interface{}, error) {
	for _, stmt := range stmts {
		branch, err := execStmt(scope, stmt)
		if err != nil {
			return nil, err
		}
		if branch != nil {
			switch branch := branch.(type) {
			case retVals:
				return branch, nil

			case *ast.BranchStmt:
				switch branch.Tok {
				case token.BREAK, token.CONTINUE, token.FALLTHROUGH:
					return branch, nil

				case token.GOTO:
					// xxx
				}
			}
			// xxx err
		}
	}
	return nil
}

func execStmt(scope *Scope, stmt ast.Stmt) (interface{}, error) {
	switch stmt := stmt.(type) {
	case *ast.DeclStmt:
		// xxx

	case *ast.EmptyStmt:
		return nil, nil

	case *ast.LabeledStmt:
		return execStmt(scope, stmt.Stmt)

	case *ast.ExprStmt:
		_, err := scope.eval(stmt.X)
		return nil, err

	case *ast.SendStmt:
		chans, err := scope.eval(stmt.Chan)
		if err != nil {
			return nil, err
		}
		if len(chans) != 1 {
			// xxx err
		}
		ch := chans[0]
		vals, err := scope.eval(stmt.Value)
		if err != nil {
			return err
		}
		if len(vals) != 1 {
			// xxx err
		}
		val := vals[0]
		if ch.Kind() != reflect.Chan {
			// xxx err
		}
		if reflect.TypeOf(ch).Elem() != reflect.TypeOf(val) {
			// xxx err
		}
		ch.Send(val)

	case *ast.IncDecStmt:
		subs, err := scope.eval(stmt.X)
		if err != nil {
			return err
		}
		if len(subs) != 1 {
			// xxx err
		}
		sub := subs[0]
		// xxx

	case *ast.AssignStmt:
		if stmt.Tok == token.DEFINE {
			// xxx
		} else {
			// xxx
		}

	case *ast.GoStmt:
		go func() {
			// xxx need some way to communicate errors
			evalCallExpr(scope, stmt.Call)
		}()

	case *ast.DeferStmt:
		// xxx add stmt.Call to the defer list of the innermost executing CallExpr

	case *ast.ReturnStmt:
		var (
			vals []reflect.Value
			err  error
		)
		if len(stmt.Results) == 1 {
			vals, err = scope.eval(stmt.Results[0])
			if err != nil {
				return nil, err
			}
		} else {
			for _, expr := range stmt.Results {
				val, err := scope.eval1(expr)
				if err != nil {
					return nil, err
				}
				vals = append(vals, val)
			}
		}
		return &retVals{vals}, nil

	case *ast.BranchStmt:
		return stmt, nil

	case *ast.BlockStmt:
		return execStmts(scope, stmt.List)

	case *ast.IfStmt:
		// xxx maybe new scope
		if stmt.Init != nil {
			err := execStmt(scope, stmt.Init)
			if err != nil {
				return err
			}
		}
		conds, err := scope.eval(stmt.Cond)
		if err != nil {
			return err
		}
		if len(conds) != 1 {
			// xxx err
		}
		cond := conds[0]
		if reflect.TypeOf(cond) != reflect.Bool {
			// xxx err
		}
		if cond.Bool() {
			return execStmts(scope, stmt.Body.List)
		}
		if stmt.Else != nil {
			return execStmt(scope, stmt.Else)
		}
		return nil

	case *ast.CaseClause:
		// xxx error - caseclause outside switch or typeswitch

	case *ast.SwitchStmt:
	case *ast.TypeSwitchStmt:
	case *ast.CommClause:
		// xxx error - commclause outside select

	case *ast.SelectStmt:
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
				ch, err := scope.eval1(clause.Chan)
				if err != nil {
					return nil, err
				}
				if ch.Kind() != reflect.Chan {
					// xxx err
				}
				sendval, err := scope.eval1(clause.Value)
				if err != nil {
					return nil, err
				}
				if reflect.TypeOf(sendval) != reflect.TypeOf(ch).Elem() {
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
		branch, err := execStmts(scope, clause.Body)
		if err != nil {
			return nil, err
		}
		if branch != nil {
			// xxx handle branching
		}
		return nil, nil

	case *ast.ForStmt:
		// xxx maybe new scope
		if stmt.Init != nil {
			err := execStmt(scope, stmt.Init)
			if err != nil {
				return err
			}
		}
		for {
			conds, err := scope.eval(stmt.Cond)
			if err != nil {
				return err
			}
			if len(conds) != 1 {
				// xxx err
			}
			cond := conds[0]
			if cond.Kind() != reflect.Bool {
				// xxx err
			}
			if !cond.Bool() {
				// exit the loop
				return nil
			}
			branch, err = execStmts(scope, stmt.Body.List)
			if err != nil {
				return err
			}
			if branch != nil {
				switch branch := branch.(type) {
				case retVals:
					return branch, nil
				case *ast.BranchStmt:
					switch branch.Tok {
					case token.BREAK:
						// xxx handle "break LABEL"
					case token.CONTINUE:
						// xxx handle "continue LABEL"
					case token.FALLTHROUGH:
						// xxx error
					case token.GOTO:
						// xxx
					}
				}
			}
			if stmt.Post != nil {
				err = execStmt(scope, stmt.Post)
				if err != nil {
					return err
				}
			}
		}

	case *ast.RangeStmt:
		// xxx maybe new scope
		if stmt.Key != nil || stmt.Value != nil {
			if stmt.Tok != token.DEFINE {
				// xxx define new key and/or value in scope
			}
		}
		rexps, err := scope.eval(stmt.X)
		if err != nil {
			return err
		}
		if len(rexps) != 1 {
			// xxx err
		}
		rexp := rexps[0]
		rexp = resolveArrayPtr(rexp)

		switch rexp.Kind() {
		case reflect.Array, reflect.Slice, reflect.String:
			l := rexp.Len()
			for i := 0; i < l; i++ {
				el := rexp.Index(i)
				// xxx assign to key if present, value if present
				err = execStmts(scope, stmt.Body.List)
				if err != nil {
					return err
				}
			}
			return nil

		case reflect.Map:
			keys := rexp.MapKeys()
			for _, k := range keys {
				v := rexp.MapIndex(k)
				// xxx assign to key if present, value if present
				err = execStmts(scope, stmt.Body.List)
				if err != nil {
					return err
				}
			}
			return nil

		case reflect.Chan:
			if (reflect.TypeOf(rexp).ChanDir() & reflect.RecvDir) == 0 {
				// xxx err
			}
			for {
				v, ok := rexp.Recv()
				if !ok {
					return nil
				}
				// xxx assign to key if present
				err = execStmts(scope, stmt.Body.List)
				if err != nil {
					return nil
				}
			}
		}
	}
}

type retVals struct {
	vals []reflect.Value
}
