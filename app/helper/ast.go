package helper

import (
	"errors"
	"fmt"
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

// AddStTag filedに対してtagを追加する
// type point struct {
//  	x, y int
// }
// このように複数のfieldが並んでるタイプには未対応
// たぶんこの２つに対してtagをつけれないと思うから
func AddStTag(node *ast.StructType, tag string, nameConverter func(str string) string) error {
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
		sttag := ParseTag(f.Tag.Value)
		// すでに存在しているtagならエラー
		if tagExist(sttag, tag) {
			return fmt.Errorf("%v tag already exists", tag)
		}
		sttag = append(sttag, StTag{name: tag, value: nameConverter(f.Names[0].Name)})
		f.Tag.Value = CombineTag(sttag)
	}
	return nil
}
