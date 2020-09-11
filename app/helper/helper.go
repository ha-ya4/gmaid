package helper

import (
	"fmt"
	"strings"
)

// ToLowerCamel 新たなstringのスライスを作り、引数strの１文字目をlowerして入れる。２文字目移行もスライスにいれて最後にjoinする
// 他にいい方法があったら知りたい
func ToLowerCamel(str string) string {
	s := []string{}
	s = append(s, strings.ToLower(string(str[0])))
	s = append(s, string(str[1:]))
	return strings.Join(s, "")
}

// CreateTag structのタグをつ作る
func CreateTag(tag, value string) string {
	return fmt.Sprintf("`%s:\"%s\"`", tag, value)
}

//  CombineTag 元々存在したtagの最後の文字[`]を取り除き、新たに作ったtagの最初の文字[`]を取り除いたものと結合する
func CombineTag(old, tag, value string) string {
	s := []string{old[:len(old)-1], " "}
	t := CreateTag(tag, value)
	s = append(s, t[1:])
	return strings.Join(s, "")
}
