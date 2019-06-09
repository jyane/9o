package main

import (
	"fmt"
	"strconv"
)

type TokenType string

const (
	TokenEOF      TokenType = "EOF"
	TokenPlus     TokenType = "+"
	TokenMinus    TokenType = "-"
	TokenMultiply TokenType = "*"
	TokenDivide   TokenType = "/"
	TokenNumber   TokenType = "number"
)

type Token struct {
	typ   TokenType
	value int
}

type TokenStream struct {
	tokens []*Token
	pos    int
}

func newTokenStream() *TokenStream {
	return &TokenStream{[]*Token{}, 0}
}

func (ts *TokenStream) consume(tt TokenType) bool {
	if ts.pos < len(ts.tokens) && ts.tokens[ts.pos].typ == tt {
		ts.pos++
		return true
	}
	return false
}

func (ts *TokenStream) now() *Token {
	return ts.tokens[ts.pos]
}

func (ts *TokenStream) add(token *Token) *TokenStream {
	ts.tokens = append(ts.tokens, token)
	return ts
}

func (ts *TokenStream) merge(that *TokenStream) *TokenStream {
	ts.tokens = append(ts.tokens, that.tokens...)
	return ts
}

func (ts *TokenStream) print() {
	fmt.Print("TokenStream{tokens=")
	for _, v := range ts.tokens {
		fmt.Print(v)
	}
	fmt.Printf("pos=%d}\n", ts.pos)
}

func isDigit(r rune) bool {
	c := int(r) - '0'
	return 0 <= c && c <= 9
}

func tokenize(s string, index int) *TokenStream {
	ts := newTokenStream()
	r := []rune(s)[index]
	N := len(s)
	if r == '+' {
		ts.add(&Token{TokenPlus, 0})
	} else if r == '-' {
		ts.add(&Token{TokenMinus, 0})
	} else if r == '*' {
		ts.add(&Token{TokenMultiply, 0})
	} else if r == '/' {
		ts.add(&Token{TokenDivide, 0})
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
		index = index + len(ns) - 1
		ts.add(&Token{TokenNumber, num})
	}
	if index < N-1 {
		ts.merge(tokenize(s, index+1))
	}
	return ts
}

func Tokenize(s string) *TokenStream {
	ts := tokenize(s, 0)
	ts.add(&Token{TokenEOF, 0})
	return ts
}
