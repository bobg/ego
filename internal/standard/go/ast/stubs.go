package ast

import (
	"go/ast"
	"reflect"
)

var NewObj = reflect.ValueOf(ast.NewObj)
var NewScope = reflect.ValueOf(ast.NewScope)
var Fprint = reflect.ValueOf(ast.Fprint)
var Print = reflect.ValueOf(ast.Print)
var NotNilFilter = reflect.ValueOf(ast.NotNilFilter)
var NewIdent = reflect.ValueOf(ast.NewIdent)
var IsExported = reflect.ValueOf(ast.IsExported)
var Walk = reflect.ValueOf(ast.Walk)
var Inspect = reflect.ValueOf(ast.Inspect)
var FilterFile = reflect.ValueOf(ast.FilterFile)
var PackageExports = reflect.ValueOf(ast.PackageExports)
var MergePackageFiles = reflect.ValueOf(ast.MergePackageFiles)
var FilterPackage = reflect.ValueOf(ast.FilterPackage)
var FileExports = reflect.ValueOf(ast.FileExports)
var FilterDecl = reflect.ValueOf(ast.FilterDecl)
var SortImports = reflect.ValueOf(ast.SortImports)
var NewPackage = reflect.ValueOf(ast.NewPackage)
var NewCommentMap = reflect.ValueOf(ast.NewCommentMap)
