package main

import (
	"flag"
	"fmt"
	"strconv"
)

type Type int

const (
	Eof = iota
	Plus
	Minus
	Number
)

type Token struct {
	typ   Type
	value int
}

func isNumber(r rune) bool {
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
	} else if isNumber(r) {
		var ns string
		for i := index; i < N; i++ {
			if isNumber(rune(s[i])) {
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

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: ")
		panic("too few argments")
	}

	tokens := tokenize(args[0], 0)

	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".global _main\n")
	fmt.Printf("_main:\n")
	fmt.Printf("  mov rax, %d\n", tokens[0].value)
	i := 1
	N := len(tokens)
	for i < N {
		if tokens[i].typ == Plus {
			i++
			if tokens[i].typ == Number {
				fmt.Printf("  add rax, %d\n", tokens[i].value)
			}
		} else if tokens[i].typ == Minus {
			i++
			if tokens[i].typ == Number {
				fmt.Printf("  sub rax, %d\n", tokens[i].value)
			}
		}
		i++
	}
	fmt.Printf("  ret\n")
}
