package ego

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"path"
	"unicode"
)

type Interp struct {
	scopes map[string]*Scope
}

func NewInterp() *Interp {
	return &Interp{scopes: make(map[string]*Scope)}
}

func (in *Interp) Load(importPath string, r io.Reader) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, xxxfilename, r, parser.DeclarationErrors)
	if err != nil {
		return err
	}

	// xxx check f.Name matches importPath

	pkgScope, ok := in.scopes[importPath]
	if !ok {
		pkgScope = NewScope(nil)
		in.scopes[importPath] = pkgScope
	}

	return in.ImportFile(pkgScope, f)
}

// ImportPath imports the package located at importPath. It creates a
// new top-level scope for the package, and subordinate scopes for
// each of its files. It also runs the init() function in each file, if any.
//
// If the path has already been imported by this interpreter, it is
// not re-imported; the previous resulting scope is returned instead.
func (in *Interp) ImportPath(importPath string) (*Scope, error) {
	if s, ok := in.scopes[importPath]; ok {
		return s, nil
	}

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, importPath, pkgFilesOnly, parser.DeclarationErrors)
	if err != nil {
		return nil, err
	}

	pkgName := path.Base(importPath)
	pkg, ok := pkgs[pkgName]
	if !ok {
		// xxx err
	}

	pkgScope := NewScope(nil)
	in.scopes[importPath] = pkgScope

	for _, f := range pkg.Files {
		err = in.doImportFile(pkgScope, f)
		if err != nil {
			return nil, err
		}
	}

	err = pkgScope.initializeVars()
	return pkgScope, err
}

func (in *Interp) ImportFile(pkgScope *Scope, f *ast.File) error {
	err := doImportFile(pkgScope, f)
	if err != nil {
		return err
	}
	err = pkgScope.initializeVars()
	return err
}

func (in *Interp) doImportFile(pkgScope *Scope, f *ast.File) error {
	fileScope := NewScope(pkgScope)
	for _, spec := range f.Imports {
		p, err := literalStr(spec.Path.Value)
		if err != nil {
			return nil, err
		}
		impScope, err := in.ImportPath(p)
		if err != nil {
			return nil, err
		}
		var importName string
		if spec.Name == nil {
			importName = path.Base(p)
		} else {
			importName = spec.Name.Name
		}
		if importName == "." {
			for k, v := range impScope.objs {
				if !isExported(k) {
					continue
				}
				// xxx should imported symbols overwrite symbols already in scope?
				fileScope.objs[k] = v
			}
		} else {
			// xxx SetScope (and other Scope.Set* functions) should ignore a
			// name of "_"
			fileScope.SetScope(importName, impScope)
		}
	}

	for _, decl := range f.Decls {
		// Each top-level decl gets a slot in pkgScope, but is evaluated
		// in fileScope (to allow imported symbols to resolve).
		switch decl := decl.(type) {
		case *ast.GenDecl:
			switch decl.Tok {
			case token.IMPORT:
				// ignore - already handled in the f.Imports clause above
				continue

			case token.CONST:
				var (
					lastExprList []ast.Expr
					lastType     ast.Expr
				)
				constScope := NewScope(fileScope)
				for i, spec := range decl.Specs {
					constScope.SetInt("iota", i)
					spec, ok := spec.(*ast.ValueSpec)
					if !ok {
						// xxx err
					}
					var (
						exprList  = spec.Values
						constType = spec.Type
					)
					if len(exprList) == 0 {
						if i == 0 {
							// xxx err
						}
						exprList = lastExprList
						constType = lastType // xxx can type be present when exprlist is absent?
					} else {
						lastExprList = exprList
						lastType = spec.Type
					}
					if len(spec.Names) != len(exprList) {
						// xxx err
					}
					var consts []Const
					for _, expr := range exprList {
						c, err := constScope.EvalConst(expr, constType)
						if err != nil {
							return nil, err
						}
						consts = append(consts, c)
					}
					for j, name := range spec.Names {
						pkgScope.SetConst(name.Name, consts[j])
					}
				}

			case token.TYPE:
				for i, spec := range decl.Specs {
					spec, ok := spec.(*ast.TypeSpec)
					if !ok {
						// xxx err
					}
					pkgScope.SetType(spec.Name.Name, spec.Type)
				}

			case token.VAR:
				for _, spec := range decl.Specs {
					spec, ok := spec.(*ast.ValueSpec)
					if !ok {
						// xxx err
					}
					if len(spec.Values) > 0 {
						// xxx let spec.Values be a single expr producing len(spec.Names) values
						if len(spec.Values) != len(spec.Names) {
							// xxx err
						}
						for i, name := range spec.Names {
							pkgScope.Add(name.Name, uninitialized(spec.Values[i]))
						}
					} else {
						for _, name := range spec.Names {
							pkgScope.Add(name.Name, NewValue(typ))
						}
					}
				}
			}

		case *ast.FuncDecl:
			// xxx check for init()
			if decl.Recv == nil {
				// regular function
				scope.SetFunc(decl.Name.Name, makeFunc(decl.Type, decl.Body.List))
			} else {
				// method
				if len(decl.Recv.List) != 1 {
					// xxx err
				}
				field := decl.Recv.List[0]
				if len(field.Names) != 1 {
					// xxx err
				}
				recvName := field.Names[0].Name
				// xxx extract receiver type, which must be an identifier or *identifier
				// (with parens possible)
				scope.AddMethod(decl.Name.Name, recvType, recvName, decl.Type, decl.Body.List)
			}

		default:
			// xxx err
		}
	}
	return nil
}

func isExported(name string) bool {
	return unicode.IsUpper(rune(name[0]))
}
