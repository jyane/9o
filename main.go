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

	ts := Tokenize(args[0])
	// ts.print()
	node := Parse(ts)
	// node.print()
	Gen(node)
}
