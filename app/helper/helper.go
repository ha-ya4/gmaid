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

// StTag structのtagの値を保持した構造体
type StTag struct {
	name, value string
}

// ParseTag structのtagをparseしてStTag構造体の配列で返す
func ParseTag(tag string) []StTag {
	sttag := []StTag{}
	bothEdges := string(tag[0])
	if bothEdges != "\"" && bothEdges != "`" {
		return sttag
	}

	tag = strings.Trim(tag, bothEdges)
	tags := strings.Split(tag, " ")
	for _, t := range tags {
		tvalues := strings.Split(t, ":")
		if len(tvalues) != 2 {
			continue
		}

		be := string(tvalues[1][0])
		if be != "\"" && be != "`" {
			continue
		}

		sttag = append(sttag, StTag{
			name:  tvalues[0],
			value: strings.Trim(tvalues[1], be),
		})
	}

	return sttag
}

// CreateTag structのタグをつ作る
func CreateTag(tag, value string) string {
	return fmt.Sprintf("`%s:\"%s\"`", tag, value)
}

// CombineTag 元々存在したtagの最後の文字[`]を取り除き、新たに作ったtagの最初の文字[`]を取り除いたものと結合する
func CombineTag(tags []StTag) string {
	tag := "`"
	for _, t := range tags {
		tag += fmt.Sprintf("%v:\"%v\" ", t.name, t.value)
	}
	tag = strings.TrimSuffix(tag, " ")
	tag += "`"
	return tag
}

func tagExist(tags []StTag, name string) bool {
	for _, t := range tags {
		if t.name == name {
			return true
		}
	}
	return false
}
