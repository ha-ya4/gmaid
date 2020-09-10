package helper

import (
	"errors"
	"go/ast"
)

const (
	errNoField = "field not found"
)

// TakeOutStruct 受け取ったast.Nodeの中から指定した名前のtypeを探し引数fnの関数に渡す
func TakeOutStruct(f *ast.File, stname string, fn func(spec *ast.TypeSpec) error) error {
	err := errors.New("specified struct not found")

	ast.Inspect(f, func(n ast.Node) bool {
		if gendecl, ok := n.(*ast.GenDecl); ok {
			for _, spec := range gendecl.Specs {
				if tspec, ok := spec.(*ast.TypeSpec); ok {
					if tspec.Name.Name == stname {
						err = fn(tspec)
						return false
					}
				}
			}
		}
		return true
	})

	return err
}

// StField structのフィールドの名前とtypeを保持した構造体
type StField struct {
	Name string
	Typ  string
}

// GetStField 引数で受け取ったnodeからフィールドの情報を抜き出した配列で返す
func GetStField(node *ast.TypeSpec) ([]StField, error) {
	var err error
	var fields []StField

	if sttype, ok := node.Type.(*ast.StructType); ok {
		for _, field := range sttype.Fields.List {
			var typ string
			if t, ok := field.Type.(*ast.Ident); ok {
				typ = t.Name
			} else {
				continue
			}

			for _, name := range field.Names {
				f := StField{Typ: typ}
				f.Name = name.Name
				fields = append(fields, f)
			}
		}
	}

	if len(fields) == 0 {
		err = errors.New(errNoField)
	}

	return fields, err
}
