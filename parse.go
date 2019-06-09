package main

import "fmt"

type NodeType string

const (
	NodeNumber         NodeType = "number"
	NodePlus           NodeType = "+"
	NodeMinus          NodeType = "-"
	NodeMultiply       NodeType = "*"
	NodeDivide         NodeType = "/"
	NodeOpenParenthes  NodeType = "("
	NodeCloseParenthes NodeType = ")"
)

type Node struct {
	typ NodeType
	rhs *Node
	lhs *Node
	val int
}

func (node *Node) print() {
	if node.lhs != nil {
		node.lhs.print()
	}
	fmt.Println(node)
	if node.rhs != nil {
		node.rhs.print()
	}
}

func term(ts *TokenStream) *Node {
	t := ts.now()
	if ts.consume(TokenOpenParenthes) {
		node := expr(ts)
		if !ts.consume(TokenCloseParenthes) {
			panic("parse error: couldn't find a close parenthes")
		}
		return node
	} else if ts.consume(TokenNumber) {
		return &Node{NodeNumber, nil, nil, t.value}
	}
	panic("parse error: unknown")
}

func mul(ts *TokenStream) *Node {
	node := term(ts)
	for {
		if ts.consume(TokenMultiply) {
			node = &Node{NodeMultiply, node, term(ts), 0}
		} else if ts.consume(TokenDivide) {
			node = &Node{NodeDivide, node, term(ts), 0}
		} else {
			return node
		}
	}
}

func expr(ts *TokenStream) *Node {
	node := mul(ts)
	for {
		if ts.consume(TokenPlus) {
			node = &Node{NodePlus, node, mul(ts), 0}
		} else if ts.consume(TokenMinus) {
			node = &Node{NodeMinus, node, mul(ts), 0}
		} else {
			return node
		}
	}
}

func Parse(ts *TokenStream) *Node {
	return expr(ts)
}
