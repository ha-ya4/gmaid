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

// 文字列の書き方に違いがあっても期待される形にparseできているか
func TestParseTag(t *testing.T) {
	expect := []StTag{
		{name: "db", value: "users"},
		{name: "json", value: "users"},
	}

	tag := "`db:\"users\" json:\"users\"`"
	sttag := ParseTag(tag)
	assert.Exactly(t, expect, sttag)

	tag = "\"db:\"users\" json:\"users\"\""
	sttag = ParseTag(tag)
	assert.Exactly(t, expect, sttag)

	tag = "\"db:`users` json:`users`\""
	sttag = ParseTag(tag)
	assert.Exactly(t, expect, sttag)
}

func TestCreateTag(t *testing.T) {
	expect := "`json:\"test\"`"
	assert.True(t, expect == CreateTag("json", "test"))
}

func TestCombineTag(t *testing.T) {
	expect := "`db:\"users\" json:\"users\" sql:\"users\"`"
	tags := []StTag{
		{name: "db", value: "users"},
		{name: "json", value: "users"},
		{name: "sql", value: "users"},
	}
	tag := CombineTag(tags)
	assert.Equal(t, expect, tag)
}
