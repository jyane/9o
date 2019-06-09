package main

import "fmt"

func gen(node *Node) {
	if node.typ == NodeNumber {
		fmt.Printf("  push %d\n", node.val)
		return
	}

	gen(node.lhs)
	gen(node.rhs)

	fmt.Println("  pop rax")
	fmt.Println("  pop rdi")

	switch node.typ {
	case NodePlus:
		fmt.Println("  add rax, rdi")
	case NodeMinus:
		fmt.Println("  sub rax, rdi")
	case NodeMultiply:
		fmt.Println("  imul rdi")
	case NodeDivide:
		fmt.Println("  cqo")
		fmt.Println("  idiv rdi")
	}

	fmt.Println("  push rax")
}

func Gen(root *Node) {
	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".global _main")
	fmt.Println("_main:")
	gen(root)
	fmt.Println("  pop rax")
	fmt.Println("  ret")
}
