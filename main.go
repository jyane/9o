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

	debug := false
	if len(args) == 2 && args[1] == "--debug" {
		debug = true
	}

	src := Read(args[0])
	ts := Tokenize(src)

	if debug {
		ts.print()
	}

	nodes := Parse(ts)

	if debug {
		for _, node := range nodes {
			node.print()
		}
	}

	asm := Gen(nodes)
	Write("out.s", asm)
}
