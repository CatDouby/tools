package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "sample.go", nil, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	for _, decl := range f.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
			ast.Print(nil, genDecl)
			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				if stype, ok := typeSpec.Type.(*ast.StructType); ok {
					for _, field := range stype.Fields.List {
						fmt.Println(field)
					}
				}
			}
		}
	}
}
