package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToLowerCamel(t *testing.T) {
	expect := "helloWorld"
	result := ToLowerCamel("HelloWorld")
	assert.True(t, expect == result)
}

func TestCreateTag(t *testing.T) {
	expect := "`json:\"test\"`"
	assert.True(t, expect == CreateTag("json", "test"))
}

func TestCombineTag(t *testing.T) {
	expect := "`db:\"users\" json:\"users\"`"
	tag := CombineTag("`db:\"users\"`", "json", "users")
	assert.True(t, expect == tag)
	t.Log(expect)
	t.Log(tag)
}
