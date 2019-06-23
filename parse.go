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
	NodeAssign         NodeType = "="
	NodeSemi           NodeType = ";"
	NodeLvalue         NodeType = "lvalue"
	NodeReturn         NodeType = "return"
)

type Node struct {
	typ    NodeType
	rhs    *Node
	lhs    *Node
	val    int
	offset int
}

func newNode(typ NodeType, rhs *Node, lhs *Node) *Node {
	return &Node{typ, rhs, lhs, 0, 0}
}

func newNumberNode(rhs *Node, lhs *Node, val int) *Node {
	node := newNode(NodeNumber, rhs, lhs)
	node.val = val
	return node
}

func newLvalueNode(rhs *Node, lhs *Node, offset int) *Node {
	node := newNode(NodeLvalue, rhs, lhs)
	node.offset = offset
	return node
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

var m = make(map[string]int)
var maxOffset = 0

func term(ts *TokenStream) *Node {
	t := ts.now()
	if ts.consume(TokenOpenParenthes) {
		node := expr(ts)
		if !ts.consume(TokenCloseParenthes) {
			panic("parse error: couldn't find a close parenthes")
		}
		return node
	} else if ts.consume(TokenNumber) {
		return newNumberNode(nil, nil, t.value)
	} else if ts.consume(TokenIdentifier) {
		offset := -1
		v, ok := m[t.name]
		if ok {
			offset = v
		} else {
			maxOffset = maxOffset + 8
			m[t.name] = maxOffset
			offset = maxOffset
		}
		return newLvalueNode(nil, nil, offset)
	}
	panic("parse error: unknown")
}

func unary(ts *TokenStream) *Node {
	if ts.consume(TokenPlus) {
		return term(ts)
	}
	if ts.consume(TokenMinus) {
		return newNode(NodeMinus, newNumberNode(nil, nil, 0), term(ts))
	}
	return term(ts)
}

func mul(ts *TokenStream) *Node {
	node := unary(ts)
	for {
		if ts.consume(TokenMultiply) {
			node = newNode(NodeMultiply, node, unary(ts))
		} else if ts.consume(TokenDivide) {
			node = newNode(NodeDivide, node, unary(ts))
		} else {
			return node
		}
	}
}

func add(ts *TokenStream) *Node {
	node := mul(ts)
	for {
		if ts.consume(TokenPlus) {
			node = newNode(NodePlus, node, mul(ts))
		} else if ts.consume(TokenMinus) {
			node = newNode(NodeMinus, node, mul(ts))
		} else {
			return node
		}
	}
}

func rational(ts *TokenStream) *Node {
	node := add(ts)
	for {
		if ts.consume(TokenLessEqual) {
			node = newNode(NodeLessEqual, node, add(ts))
		} else if ts.consume(TokenGreaterEqual) {
			node = newNode(NodeLessEqual, add(ts), node)
		} else if ts.consume(TokenLess) {
			node = newNode(NodeLess, node, add(ts))
		} else if ts.consume(TokenGreater) {
			node = newNode(NodeLess, add(ts), node)
		} else {
			return node
		}
	}
}

func equality(ts *TokenStream) *Node {
	node := rational(ts)
	for {
		if ts.consume(TokenEqual) {
			node = newNode(NodeEqual, node, rational(ts))
		} else if ts.consume(TokenNotEqual) {
			node = newNode(NodeNotEqual, node, rational(ts))
		} else {
			return node
		}
	}
}

func assign(ts *TokenStream) *Node {
	node := equality(ts)
	if ts.consume(TokenAssign) {
		node = newNode(NodeAssign, assign(ts), node)
	}
	return node
}

func expr(ts *TokenStream) *Node {
	return assign(ts)
}

func stmt(ts *TokenStream) *Node {
	node := &Node{}
	if ts.consume(TokenReturn) {
		node = newNode(NodeReturn, nil, expr(ts))
	} else {
		node = expr(ts)
	}
	if !ts.consume(TokenSemi) {
		panic("token is not ';'")
	}
	return node
}

func program(ts *TokenStream) []*Node {
	var nodes []*Node
	for !ts.isEnd() {
		nodes = append(nodes, stmt(ts))
	}
	return nodes
}

func Parse(ts *TokenStream) []*Node {
	return program(ts)
}
