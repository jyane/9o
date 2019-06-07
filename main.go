package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	flag.Parse()
	args := flag.Args()
	num, err := strconv.Atoi(args[0])
	if err != nil {
		panic("error")
	}
	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".global _main\n")
	fmt.Printf("_main:\n")
	fmt.Printf("  mov rax, %d\n", num)
	fmt.Printf("  ret\n")
}
