package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: ")
		panic("too few argments")
	}

	tokens := Tokenize(args[0])

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
