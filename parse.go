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
	NodeEqual          NodeType = "=="
	NodeNotEqual       NodeType = "!="
	NodeLessEqual      NodeType = "<="
	NodeLess           NodeType = "<"
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

func unary(ts *TokenStream) *Node {
	if ts.consume(TokenPlus) {
		return term(ts)
	}
	if ts.consume(TokenMinus) {
		return &Node{NodeMinus, &Node{NodeNumber, nil, nil, 0}, term(ts), 0}
	}
	return term(ts)
}

func mul(ts *TokenStream) *Node {
	node := unary(ts)
	for {
		if ts.consume(TokenMultiply) {
			node = &Node{NodeMultiply, node, unary(ts), 0}
		} else if ts.consume(TokenDivide) {
			node = &Node{NodeDivide, node, unary(ts), 0}
		} else {
			return node
		}
	}
}

func add(ts *TokenStream) *Node {
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

func rational(ts *TokenStream) *Node {
	node := add(ts)
	for {
		if ts.consume(TokenLessEqual) {
			node = &Node{NodeLessEqual, node, add(ts), 0}
		} else if ts.consume(TokenGreaterEqual) {
			node = &Node{NodeLessEqual, add(ts), node, 0}
		} else if ts.consume(TokenLess) {
			node = &Node{NodeLess, node, add(ts), 0}
		} else if ts.consume(TokenGreater) {
			node = &Node{NodeLess, add(ts), node, 0}
		} else {
			return node
		}
	}
}

func equality(ts *TokenStream) *Node {
	node := rational(ts)
	for {
		if ts.consume(TokenEqual) {
			node = &Node{NodeEqual, node, rational(ts), 0}
		} else if ts.consume(TokenNotEqual) {
			node = &Node{NodeNotEqual, node, rational(ts), 0}
		} else {
			return node
		}
	}
}

func expr(ts *TokenStream) *Node {
	return equality(ts)
}

func Parse(ts *TokenStream) *Node {
	return expr(ts)
}
