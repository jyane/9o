package main

import "fmt"

func genLvalue(node *Node) {
	if node.typ != NodeLvalue {
		panic("lvalue is not a variable.")
	}
	fmt.Println("  mov rax, rbp")
	fmt.Printf("  sub rax, %d\n", node.offset)
	fmt.Println("  push rax")
}

func gen(node *Node) {
	if node.typ == NodeNumber {
		fmt.Printf("  push %d\n", node.val)
		return
	}

	if node.typ == NodeReturn {
		gen(node.lhs)
		fmt.Println("  pop rax")
		fmt.Println("  mov rsp, rbp")
		fmt.Println("  pop rbp")
		fmt.Println("  ret")
		return
	}

	if node.typ == NodeLvalue {
		genLvalue(node)
		fmt.Println("  pop rax")
		fmt.Println("  mov rax, [rax]")
		fmt.Println("  push rax")
		return
	}

	if node.typ == NodeAssign {
		genLvalue(node.lhs)
		gen(node.rhs)
		fmt.Println("  pop rdi")
		fmt.Println("  pop rax")
		fmt.Println("  mov [rax], rdi")
		fmt.Println("  push rdi")
		return
	}

	// TODO: arbitrary if
	if node.typ == NodeIf {
		gen(node.cond)
		fmt.Println("  pop rax")
		fmt.Println("  cmp rax, 0")
		if node.elseb != nil {
			fmt.Printf("  je .Lelse00%d\n", 1)
			gen(node.thenb)
			fmt.Printf(".Lelse00%d:\n", 1)
			gen(node.elseb)
			fmt.Printf(".Lend00%d:\n", 1)
		} else {
			fmt.Printf("  je .Lend00%d\n", 1)
			gen(node.thenb)
			fmt.Printf(".Lend00%d:\n", 1)
		}
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

func Gen(roots []*Node) {
	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".global _main")
	fmt.Println("_main:")
	fmt.Println("  push rbp")
	fmt.Println("  mov rbp, rsp")
	fmt.Println("  sub rsp, 208")
	for _, v := range roots {
		gen(v)
		fmt.Println("  pop rax")
	}
	fmt.Println("  mov rsp, rbp")
	fmt.Println("  pop rbp")
	fmt.Println("  ret")
}
