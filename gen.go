package main

import (
	"fmt"
	"strings"
)

func genLvalue(node *Node) string {
	var sb strings.Builder
	if node.typ != NodeLvalue {
		panic("lvalue is not a variable.")
	}
	sb.WriteString("  mov rax, rbp\n")
	fmt.Fprintf(&sb, "  sub rax, %d\n", node.offset)
	sb.WriteString("  push rax\n")
	return sb.String()
}

func gen(node *Node) string {
	var sb strings.Builder
	if node.typ == NodeNumber {
		fmt.Fprintf(&sb, "  push %d\n", node.val)
		return sb.String()
	}

	if node.typ == NodeReturn {
		sb.WriteString(gen(node.lhs))
		sb.WriteString("  pop rax\n")
		sb.WriteString("  mov rsp, rbp\n")
		sb.WriteString("  pop rbp\n")
		sb.WriteString("  ret\n")
		return sb.String()
	}

	if node.typ == NodeLvalue {
		sb.WriteString(genLvalue(node))
		sb.WriteString("  pop rax\n")
		sb.WriteString("  mov rax, [rax]\n")
		sb.WriteString("  push rax\n")
		return sb.String()
	}

	if node.typ == NodeAssign {
		sb.WriteString(genLvalue(node.lhs))
		sb.WriteString(gen(node.rhs))
		sb.WriteString("  pop rdi\n")
		sb.WriteString("  pop rax\n")
		sb.WriteString("  mov [rax], rdi\n")
		sb.WriteString("  push rdi\n")
		return sb.String()
	}

	// TODO: arbitrary if
	if node.typ == NodeIf {
		sb.WriteString(gen(node.cond))
		sb.WriteString("  pop rax\n")
		sb.WriteString("  cmp rax, 0\n")
		if node.elseb != nil {
			fmt.Fprintf(&sb, "  je .Lelse00%d\n", 1)
			sb.WriteString(gen(node.thenb))
			fmt.Fprintf(&sb, ".Lelse00%d:\n", 1)
			sb.WriteString(gen(node.elseb))
			fmt.Fprintf(&sb, ".Lend00%d:\n", 1)
		} else {
			fmt.Fprintf(&sb, "  je .Lend00%d\n", 1)
			sb.WriteString(gen(node.thenb))
			fmt.Fprintf(&sb, ".Lend00%d:\n", 1)
		}
		return sb.String()
	}

	// TODO: arbitrary while
	if node.typ == NodeWhile {
		fmt.Fprintf(&sb, ".Lbegin00%d:\n", 1)
		sb.WriteString(gen(node.cond))
		sb.WriteString("  pop rax\n")
		sb.WriteString("  cmp rax, 0\n")
		fmt.Fprintf(&sb, "  je .Lend00%d\n", 1)
		sb.WriteString(gen(node.thenb))
		fmt.Fprintf(&sb, "jmp .Lbegin00%d\n", 1)
		fmt.Fprintf(&sb, ".Lend00%d:", 1)
		return sb.String()
	}

	if node.typ == NodeBlock {
		for _, stmt := range node.stmts {
			sb.WriteString(gen(stmt))
			sb.WriteString("  pop rax\n")
		}
		return sb.String()
	}

	sb.WriteString(gen(node.lhs))
	sb.WriteString(gen(node.rhs))

	sb.WriteString("  pop rax\n")
	sb.WriteString("  pop rdi\n")

	switch node.typ {
	case NodePlus:
		sb.WriteString("  add rax, rdi\n")
	case NodeMinus:
		sb.WriteString("  sub rax, rdi\n")
	case NodeMultiply:
		sb.WriteString("  imul rdi\n")
	case NodeDivide:
		sb.WriteString("  cqo\n")
		sb.WriteString("  idiv rdi\n")
	case NodeEqual:
		sb.WriteString("  cmp rax, rdi\n")
		sb.WriteString("  sete al\n")
		sb.WriteString("  movzx rax, al\n")
	case NodeLess:
		sb.WriteString("  cmp rax, rdi\n")
		sb.WriteString("  setl al\n")
		sb.WriteString("  movzx rax, al\n")
	case NodeLessEqual:
		sb.WriteString("  cmp rax, rdi\n")
		sb.WriteString("  setle al\n")
		sb.WriteString("  movzx rax, al\n")
	case NodeNotEqual:
		sb.WriteString("  cmp rax, rdi\n")
		sb.WriteString("  setne al\n")
		sb.WriteString("  movzx rax, al\n")
	}

	sb.WriteString("  push rax\n")
	return sb.String()
}

func Gen(roots []*Node) string {
	var sb strings.Builder
	sb.WriteString(".intel_syntax noprefix\n")
	sb.WriteString(".global _main\n")
	sb.WriteString("_main:\n")
	sb.WriteString("  push rbp\n")
	sb.WriteString("  mov rbp, rsp\n")
	sb.WriteString("  sub rsp, 208\n")
	for _, v := range roots {
		sb.WriteString(gen(v))
		sb.WriteString("  pop rax\n")
	}
	sb.WriteString("  mov rsp, rbp\n")
	sb.WriteString("  pop rbp\n")
	sb.WriteString("  ret\n")
	return sb.String()
}
