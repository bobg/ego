package ego

import (
	"go/ast"
	"go/token"

	"github.com/bobg/ego/refl"
)

func (s *Scope) evalStar(expr *ast.StarExpr) (*refl.Value, error) {
	sub, err := s.Eval1(expr.X)
	if err != nil {
		return nil, err
	}
	if sub.Kind() != refl.Ptr {
		// xxx err
	}
	return []*refl.Value{sub.Elem()}, nil
}

func (s *Scope) evalUnary(expr *ast.UnaryExpr) (*refl.Value, error) {
	sub, err := s.Eval1(expr.X)
	if err != nil {
		return nil, err
	}
	switch expr.Op {
	case token.ADD:
	case token.SUB:
	case token.XOR:
	case token.NOT:
	case token.AND:
	case token.XOR:
	case token.ARROW: // <-
	default:
		// xxx err
	}
}

func (s *Scope) evalBinary(expr *ast.BinaryExpr) (*refl.Value, error) {
	x, err := s.Eval1(expr.X)
	if err != nil {
		return nil, err
	}
	// boolean short-circuiting
	if expr.Op == token.LAND {
		b, err := x.Bool()
		if err != nil {
			return nil, err
		}
		if !b {
			return []*refl.Value{x}, nil
		}
	}
	if expr.Op == token.LOR {
		b, err := x.Bool()
		if err != nil {
			return nil, err
		}
		if b {
			return []*refl.Value{x}, nil
		}
	}
	y, err := s.Eval2(expr.Y)
	if err != nil {
		return nil, err
	}
	switch expr.Op {
	case token.ADD: // +
		if x.Kind() != y.Kind() {
			// xxx err
		}
		switch x.Kind() {
		case refl.Int:
			xv, _ := x.Int()
			yv, _ := y.Int()
			return []*refl.Value{refl.ValueOf(xv + yv)}, nil

		case refl.Int8:
			xv, _ := x.Int8()
			yv, _ := y.Int8()
			return []*refl.Value{refl.ValueOf(xv + yv)}, nil

		case refl.Int16:
			xv, _ := x.Int16()
			yv, _ := y.Int16()
			return []*refl.Value{refl.ValueOf(xv + yv)}, nil

		case refl.Int32:
			xv, _ := x.Int32()
			yv, _ := y.Int32()
			return []*refl.Value{refl.ValueOf(xv + yv)}, nil

		case refl.Int64:
			xv, _ := x.Int64()
			yv, _ := y.Int64()
			return []*refl.Value{refl.ValueOf(xv + yv)}, nil

		case refl.Uint:
			xv, _ := x.Uint()
			yv, _ := y.Uint()
			return []*refl.Value{refl.ValueOf(xv + yv)}, nil

		case refl.Uint8:
			xv, _ := x.Uint8()
			yv, _ := y.Uint8()
			return []*refl.Value{refl.ValueOf(xv + yv)}, nil

		case refl.Uint16:
			xv, _ := x.Uint16()
			yv, _ := y.Uint16()
			return []*refl.Value{refl.ValueOf(xv + yv)}, nil

		case refl.Uint32:
			xv, _ := x.Uint32()
			yv, _ := y.Uint32()
			return []*refl.Value{refl.ValueOf(xv + yv)}, nil

		case refl.Uint64:
			xv, _ := x.Uint64()
			yv, _ := y.Uint64()
			return []*refl.Value{refl.ValueOf(xv + yv)}, nil

		case refl.Uintptr:
			xv, _ := x.Uintptr()
			yv, _ := y.Uintptr()
			return []*refl.Value{refl.ValueOf(xv + yv)}, nil

		case refl.Float32:
			xv, _ := x.Float32()
			yv, _ := y.Float32()
			return []*refl.Value{refl.ValueOf(xv + yv)}, nil

		case refl.Float64:
			xv, _ := x.Float64()
			yv, _ := y.Float64()
			return []*refl.Value{refl.ValueOf(xv + yv)}, nil

		case refl.Complex64:
			xv, _ := x.Complex64()
			yv, _ := y.Complex64()
			return []*refl.Value{refl.ValueOf(xv + yv)}, nil

		case refl.Complex128:
			xv, _ := x.Complex128()
			yv, _ := y.Complex128()
			return []*refl.Value{refl.ValueOf(xv + yv)}, nil

		}

	case token.SUB: // -
		if x.Kind() != y.Kind() {
			// xxx err
		}
		switch x.Kind() {
		case refl.Int:
			xv, _ := x.Int()
			yv, _ := y.Int()
			return []*refl.Value{refl.ValueOf(xv - yv)}, nil

		case refl.Int8:
			xv, _ := x.Int8()
			yv, _ := y.Int8()
			return []*refl.Value{refl.ValueOf(xv - yv)}, nil

		case refl.Int16:
			xv, _ := x.Int16()
			yv, _ := y.Int16()
			return []*refl.Value{refl.ValueOf(xv - yv)}, nil

		case refl.Int32:
			xv, _ := x.Int32()
			yv, _ := y.Int32()
			return []*refl.Value{refl.ValueOf(xv - yv)}, nil

		case refl.Int64:
			xv, _ := x.Int64()
			yv, _ := y.Int64()
			return []*refl.Value{refl.ValueOf(xv - yv)}, nil

		case refl.Uint:
			xv, _ := x.Uint()
			yv, _ := y.Uint()
			return []*refl.Value{refl.ValueOf(xv - yv)}, nil

		case refl.Uint8:
			xv, _ := x.Uint8()
			yv, _ := y.Uint8()
			return []*refl.Value{refl.ValueOf(xv - yv)}, nil

		case refl.Uint16:
			xv, _ := x.Uint16()
			yv, _ := y.Uint16()
			return []*refl.Value{refl.ValueOf(xv - yv)}, nil

		case refl.Uint32:
			xv, _ := x.Uint32()
			yv, _ := y.Uint32()
			return []*refl.Value{refl.ValueOf(xv - yv)}, nil

		case refl.Uint64:
			xv, _ := x.Uint64()
			yv, _ := y.Uint64()
			return []*refl.Value{refl.ValueOf(xv - yv)}, nil

		case refl.Uintptr:
			xv, _ := x.Uintptr()
			yv, _ := y.Uintptr()
			return []*refl.Value{refl.ValueOf(xv - yv)}, nil

		case refl.Float32:
			xv, _ := x.Float32()
			yv, _ := y.Float32()
			return []*refl.Value{refl.ValueOf(xv - yv)}, nil

		case refl.Float64:
			xv, _ := x.Float64()
			yv, _ := y.Float64()
			return []*refl.Value{refl.ValueOf(xv - yv)}, nil

		case refl.Complex64:
			xv, _ := x.Complex64()
			yv, _ := y.Complex64()
			return []*refl.Value{refl.ValueOf(xv - yv)}, nil

		case refl.Complex128:
			xv, _ := x.Complex128()
			yv, _ := y.Complex128()
			return []*refl.Value{refl.ValueOf(xv - yv)}, nil

		}

	case token.MUL: // *
		if x.Kind() != y.Kind() {
			// xxx err
		}
		switch x.Kind() {
		case refl.Int:
			xv, _ := x.Int()
			yv, _ := y.Int()
			return []*refl.Value{refl.ValueOf(xv * yv)}, nil

		case refl.Int8:
			xv, _ := x.Int8()
			yv, _ := y.Int8()
			return []*refl.Value{refl.ValueOf(xv * yv)}, nil

		case refl.Int16:
			xv, _ := x.Int16()
			yv, _ := y.Int16()
			return []*refl.Value{refl.ValueOf(xv * yv)}, nil

		case refl.Int32:
			xv, _ := x.Int32()
			yv, _ := y.Int32()
			return []*refl.Value{refl.ValueOf(xv * yv)}, nil

		case refl.Int64:
			xv, _ := x.Int64()
			yv, _ := y.Int64()
			return []*refl.Value{refl.ValueOf(xv * yv)}, nil

		case refl.Uint:
			xv, _ := x.Uint()
			yv, _ := y.Uint()
			return []*refl.Value{refl.ValueOf(xv * yv)}, nil

		case refl.Uint8:
			xv, _ := x.Uint8()
			yv, _ := y.Uint8()
			return []*refl.Value{refl.ValueOf(xv * yv)}, nil

		case refl.Uint16:
			xv, _ := x.Uint16()
			yv, _ := y.Uint16()
			return []*refl.Value{refl.ValueOf(xv * yv)}, nil

		case refl.Uint32:
			xv, _ := x.Uint32()
			yv, _ := y.Uint32()
			return []*refl.Value{refl.ValueOf(xv * yv)}, nil

		case refl.Uint64:
			xv, _ := x.Uint64()
			yv, _ := y.Uint64()
			return []*refl.Value{refl.ValueOf(xv * yv)}, nil

		case refl.Uintptr:
			xv, _ := x.Uintptr()
			yv, _ := y.Uintptr()
			return []*refl.Value{refl.ValueOf(xv * yv)}, nil

		case refl.Float32:
			xv, _ := x.Float32()
			yv, _ := y.Float32()
			return []*refl.Value{refl.ValueOf(xv * yv)}, nil

		case refl.Float64:
			xv, _ := x.Float64()
			yv, _ := y.Float64()
			return []*refl.Value{refl.ValueOf(xv * yv)}, nil

		case refl.Complex64:
			xv, _ := x.Complex64()
			yv, _ := y.Complex64()
			return []*refl.Value{refl.ValueOf(xv * yv)}, nil

		case refl.Complex128:
			xv, _ := x.Complex128()
			yv, _ := y.Complex128()
			return []*refl.Value{refl.ValueOf(xv * yv)}, nil
		}

	case token.QUO: // /
		if x.Kind() != y.Kind() {
			// xxx err
		}
		// xxx check y != 0 ?
		switch x.Kind() {
		case refl.Int:
			xv, _ := x.Int()
			yv, _ := y.Int()
			return []*refl.Value{refl.ValueOf(xv / yv)}, nil

		case refl.Int8:
			xv, _ := x.Int8()
			yv, _ := y.Int8()
			return []*refl.Value{refl.ValueOf(xv / yv)}, nil

		case refl.Int16:
			xv, _ := x.Int16()
			yv, _ := y.Int16()
			return []*refl.Value{refl.ValueOf(xv / yv)}, nil

		case refl.Int32:
			xv, _ := x.Int32()
			yv, _ := y.Int32()
			return []*refl.Value{refl.ValueOf(xv / yv)}, nil

		case refl.Int64:
			xv, _ := x.Int64()
			yv, _ := y.Int64()
			return []*refl.Value{refl.ValueOf(xv / yv)}, nil

		case refl.Uint:
			xv, _ := x.Uint()
			yv, _ := y.Uint()
			return []*refl.Value{refl.ValueOf(xv / yv)}, nil

		case refl.Uint8:
			xv, _ := x.Uint8()
			yv, _ := y.Uint8()
			return []*refl.Value{refl.ValueOf(xv / yv)}, nil

		case refl.Uint16:
			xv, _ := x.Uint16()
			yv, _ := y.Uint16()
			return []*refl.Value{refl.ValueOf(xv / yv)}, nil

		case refl.Uint32:
			xv, _ := x.Uint32()
			yv, _ := y.Uint32()
			return []*refl.Value{refl.ValueOf(xv / yv)}, nil

		case refl.Uint64:
			xv, _ := x.Uint64()
			yv, _ := y.Uint64()
			return []*refl.Value{refl.ValueOf(xv / yv)}, nil

		case refl.Uintptr:
			xv, _ := x.Uintptr()
			yv, _ := y.Uintptr()
			return []*refl.Value{refl.ValueOf(xv / yv)}, nil

		case refl.Float32:
			xv, _ := x.Float32()
			yv, _ := y.Float32()
			return []*refl.Value{refl.ValueOf(xv / yv)}, nil

		case refl.Float64:
			xv, _ := x.Float64()
			yv, _ := y.Float64()
			return []*refl.Value{refl.ValueOf(xv / yv)}, nil

		case refl.Complex64:
			xv, _ := x.Complex64()
			yv, _ := y.Complex64()
			return []*refl.Value{refl.ValueOf(xv / yv)}, nil

		case refl.Complex128:
			xv, _ := x.Complex128()
			yv, _ := y.Complex128()
			return []*refl.Value{refl.ValueOf(xv / yv)}, nil

		}

	case token.REM: // %
		if x.Kind() != y.Kind() {
			// xxx err
		}
		// xxx check y != 0 ?
		switch x.Kind() {
		case refl.Int:
			xv, _ := x.Int()
			yv, _ := y.Int()
			return []*refl.Value{refl.ValueOf(xv % yv)}, nil

		case refl.Int8:
			xv, _ := x.Int8()
			yv, _ := y.Int8()
			return []*refl.Value{refl.ValueOf(xv % yv)}, nil

		case refl.Int16:
			xv, _ := x.Int16()
			yv, _ := y.Int16()
			return []*refl.Value{refl.ValueOf(xv % yv)}, nil

		case refl.Int32:
			xv, _ := x.Int32()
			yv, _ := y.Int32()
			return []*refl.Value{refl.ValueOf(xv % yv)}, nil

		case refl.Int64:
			xv, _ := x.Int64()
			yv, _ := y.Int64()
			return []*refl.Value{refl.ValueOf(xv % yv)}, nil

		case refl.Uint:
			xv, _ := x.Uint()
			yv, _ := y.Uint()
			return []*refl.Value{refl.ValueOf(xv % yv)}, nil

		case refl.Uint8:
			xv, _ := x.Uint8()
			yv, _ := y.Uint8()
			return []*refl.Value{refl.ValueOf(xv % yv)}, nil

		case refl.Uint16:
			xv, _ := x.Uint16()
			yv, _ := y.Uint16()
			return []*refl.Value{refl.ValueOf(xv % yv)}, nil

		case refl.Uint32:
			xv, _ := x.Uint32()
			yv, _ := y.Uint32()
			return []*refl.Value{refl.ValueOf(xv % yv)}, nil

		case refl.Uint64:
			xv, _ := x.Uint64()
			yv, _ := y.Uint64()
			return []*refl.Value{refl.ValueOf(xv % yv)}, nil

		case refl.Uintptr:
			xv, _ := x.Uintptr()
			yv, _ := y.Uintptr()
			return []*refl.Value{refl.ValueOf(xv % yv)}, nil
		}

	case token.AND: // &
	case token.OR: // |
	case token.XOR: // ^
	case token.SHL: // <<
	case token.SHR: // >>
	case token.AND_NOT: // &^
	case token.LAND, token.LOR:
		// Already know x is true (for and) or false (for or) thanks to
		// short-circuiting logic above.
		if y.Kind() != refl.Bool {
			// xxx err
		}
		return []*refl.Value{y}, nil

	case token.ARROW: // <-
	case token.EQL: // ==
	case token.LSS: // <
	case token.GTR: // >
	case token.NEQ: // !=
	case token.LEQ: // <=
	case token.GEQ: // >=
	}
}
