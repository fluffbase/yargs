package yargs

import (
	"strings"
	"regexp"

)

type Tokens []string

func FromSnakeCase(str string) Tokens {
	return strings.Split(str, "-")
}

func FromCamelCase(str string) Tokens {
	var matchAllCap   = regexp.MustCompile("[a-z0-9][A-Z]")
	found := matchAllCap.FindAllIndex([]byte(str), -1)
	previndex := 0
	i := 0 
	var tok Tokens
	for _, index := range(found) {
		i = index[0]+1
		tok = append(tok, str[previndex:i])
		previndex = i
	}
	tok = append(tok, str[previndex:])
	
	return tok
}

func (tok Tokens) ToSnakeCase() string {
	var strs []string
	for _,v := range(tok) {
		strs=append(strs, strings.ToLower(v))
	}
	return strings.Join(strs, "-")
}
func (tok Tokens) ToCamelCase() string {
	var strs []string
	for _,v := range(tok) {
		strs=append(strs, strings.Title(v))
	}
	return strings.Join(strs, "")
}
