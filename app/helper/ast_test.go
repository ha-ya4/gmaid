package helper

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ha-ya4/gmaid/testdata"
)

func TestTakeOutStruct(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "../../testdata/testgo.go", nil, parser.Mode(0))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	// 指定したstructが見つかり変数nameが期待通りに書き換わっているか
	name := ""
	err = TakeOutStruct(f, testdata.StName, func(spec *ast.TypeSpec) {
		name = spec.Name.Name
	})
	assert.NoError(t, err)
	assert.True(t, name == testdata.StName)

	// 指定されたstructが見つからずエラーになり変数nameがかわってないか
	name = ""
	err = TakeOutStruct(f, "testestest", func(spec *ast.TypeSpec) {
		name = spec.Name.Name
	})
	assert.Error(t, err)
	assert.True(t, name == "")
}