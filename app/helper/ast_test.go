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

func TestGetField(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "../../testdata/testgo.go", nil, parser.Mode(0))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	// structのフィールド情報のスライスを取得できているか
	fields := []string{}
	err = TakeOutStruct(f, testdata.StName, func(spec *ast.TypeSpec) error {
		field, err := GetStField(spec)
		assert.NoError(t, err)
		for _, v := range field {
			fields = append(fields, v.Name)
		}
		return err
	})
	assert.ElementsMatch(t, testdata.StField, fields)

	// structのフィールド情報のスライスを取得できているか
	err = TakeOutStruct(f, testdata.TypeAliasName, func(spec *ast.TypeSpec) error {
		_, err := GetStField(spec)
		return err
	})
	assert.EqualError(t, err, errNoField)
	t.Log(err)
}
