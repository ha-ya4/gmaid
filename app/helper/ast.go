package helper

import (
	"errors"
	"go/ast"
	"go/token"
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
		fields, err = GetStFieldStruct(sttype)
		if err != nil {
			return fields, err
		}
	}

	if len(fields) == 0 {
		err = errors.New(errNoField)
	}

	return fields, err
}

// GetStFieldStruct 引数で受け取ったnodeからフィールドの情報を抜き出した配列で返す
func GetStFieldStruct(node *ast.StructType) ([]StField, error) {
	var fields []StField

	for _, field := range node.Fields.List {
			var typ string
			if t, ok := field.Type.(*ast.Ident); ok {
				typ = t.Name
			} else {
				continue
			}

		if len(field.Names) > 1 {
			return fields, errors.New("does not support multiple fields")
		}

		name := field.Names[0]
				f := StField{Typ: typ}
				f.Name = name.Name
				fields = append(fields, f)
			}

	return fields, nil
		}

// AddStTag filedに対してtagを追加する
// type point struct {
//  	x, y int
// }
// このように複数のfieldが並んでるタイプには未対応
// たぶんこの２つに対してtagをつけれないと思うから
func AddStTag(node *ast.StructType, tag string, nameConverter func(str string) string) {
	for _, f := range node.Fields.List {
		// すでにタグがあった場合となかった場合で分ける
		// ない場合は新規で作るだけ
		if f.Tag == nil {
			f.Tag = &ast.BasicLit{
				Kind:  token.STRING,
				Value: CreateTag(tag, nameConverter(f.Names[0].Name)),
	}
			continue
		}

		// ある場合は元々あるタグと結合する
		f.Tag.Value = CombineTag(f.Tag.Value, "json", nameConverter(f.Names[0].Name))
	}
}

	}

	return fields, err
}
