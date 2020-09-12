package helper

import (
	"errors"
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
	err = TakeOutStruct(f, testdata.StName, func(spec *ast.TypeSpec) error {
		name = spec.Name.Name
		return nil
	})
	assert.NoError(t, err)
	assert.True(t, name == testdata.StName)

	// 指定されたstructが見つからずエラーになり変数nameがかわってないか
	name = ""
	err = TakeOutStruct(f, "testestest", func(spec *ast.TypeSpec) error {
		name = spec.Name.Name
		return nil
	})
	assert.Error(t, err)
	assert.True(t, name == "")

	// 指定したstructが見つかり変数nameが期待通りに書き換わっているが,引数に渡した関数内でエラーがでたときにエラーになるか
	es := "error"
	name = ""
	err = TakeOutStruct(f, testdata.StName, func(spec *ast.TypeSpec) error {
		name = spec.Name.Name
		return errors.New(es)
	})
	assert.EqualError(t, err, es)
	assert.True(t, name == testdata.StName)
}

func TestAddStTag(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "../../testdata/testgo.go", nil, parser.Mode(0))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	var result []*ast.Field
	err = TakeOutStruct(f, testdata.StName, func(spec *ast.TypeSpec) error {
		sttype, ok := spec.Type.(*ast.StructType)
		if ok {
			AddStTag(sttype, "json", ToLowerCamel)
		}
		result = sttype.Fields.List
		return nil
	})

	for _, r := range result {
		assert.NotEmpty(t, r.Tag)
	}
}
