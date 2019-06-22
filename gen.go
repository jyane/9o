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
	case NodeEqual:
		fmt.Println("  cmp rax, rdi")
		fmt.Println("  sete al")
		fmt.Println("  movzx rax, al")
	case NodeLess:
		fmt.Println("  cmp rax, rdi")
		fmt.Println("  setl al")
		fmt.Println("  movzx rax, al")
	case NodeLessEqual:
		fmt.Println("  cmp rax, rdi")
		fmt.Println("  setle al")
		fmt.Println("  movzx rax, al")
	case NodeNotEqual:
		fmt.Println("  cmp rax, rdi")
		fmt.Println("  setne al")
		fmt.Println("  movzx rax, al")
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
