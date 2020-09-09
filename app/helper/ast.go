package helper

import (
	"go/ast"
	"errors"
)

// TakeOutStruct 受け取ったast.Nodeの中から指定した名前のtypeを探し引数fnの関数に渡す
func TakeOutStruct(f  *ast.File, stname string, fn func(spec *ast.TypeSpec)) error {
	err := errors.New("specified struct not found")

	ast.Inspect(f, func(n ast.Node) bool {
		if gendecl, ok := n.(*ast.GenDecl); ok {
			for _, spec := range gendecl.Specs {
				if tspec, ok := spec.(*ast.TypeSpec); ok {
					if tspec.Name.Name == stname {
						fn(tspec)
						err = nil
						return false
					}
				}
			}
		}
		return true
	})

	return err
}
