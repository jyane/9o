package main

import "strconv"

type TokenType string

const (
	EOF    TokenType = "EOF"
	Plus   TokenType = "+"
	Minus  TokenType = "-"
	Number TokenType = "number"
)

type Token struct {
	typ   TokenType
	value int
}

func isDigit(r rune) bool {
	c := int(r) - '0'
	return 0 <= c && c <= 9
}

func tokenize(s string, index int) []Token {
	var res = []Token{}
	r := []rune(s)[index]
	N := len(s)
	if r == '+' {
		res = append(res, Token{Plus, 0})
	} else if r == '-' {
		res = append(res, Token{Minus, 0})
	} else if isDigit(r) {
		var ns string
		for i := index; i < N; i++ {
			if isDigit(rune(s[i])) {
				ns += string(s[i])
			} else {
				break
			}
		}
		num, err := strconv.Atoi(ns)
		if err != nil {
			panic("tokenize error: [" + ns + "]")
		}
		res = append(res, Token{Number, num})
	}
	if index < N-1 {
		res = append(res, tokenize(s, index+1)...)
	}
	return res
}

func Tokenize(s string) []Token {
	res := tokenize(s, 0)
	return append(res, Token{EOF, 0})
}
