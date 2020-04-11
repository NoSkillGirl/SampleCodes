// Infix expression:The expression of the form a op b. When an operator is in-between
// every pair of operands.
// Prefix : An expression is called the prefix expression if the operator appears in the
// expression before the operands. Simply of the form (operator operand1 operand2).

// Infix: (a-b/c)*(a/k-l)
// Prefix: *-a/bc-/akl

package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func main() {
	infixExp := "(a-b/c)*(a/k-l)"
	fmt.Println("Infix expression: ", infixExp)
	rs := reverseString(infixExp)
	postfixExp := infixToPostfix(rs)
	prefixExp := reverseString(postfixExp)
	fmt.Println("Prefix expression: ", prefixExp)
}

func reverseString(s string) string {
	l := len(s)
	var rs string
	for i := l - 1; i >= 0; i-- {
		if string(s[i]) == "(" {
			rs = rs + ")"
		} else if string(s[i]) == ")" {
			rs = rs + "("
		} else {
			rs = rs + string(s[i])
		}
	}
	return rs
}

func prec(c string) int {
	if c == "^" {
		return 3
	} else if c == "*" || c == "/" {
		return 2
	} else if c == "+" || c == "-" {
		return 1
	} else {
		return 0
	}
}

func infixToPostfix(s string) string {
	st := stack.New()
	st.Push("N")
	l := len(s)
	var ns string
	for i := 0; i < l; i++ {
		if (string(s[i]) >= "a" && string(s[i]) <= "z") || (string(s[i]) >= "A" && string(s[i]) <= "Z") {
			ns = ns + string(s[i])
		} else if string(s[i]) == "(" {
			st.Push("(")
		} else if string(s[i]) == ")" {
			for st.Peek() != "N" && st.Peek() != "(" {
				c := st.Peek()
				cstr := c.(string)
				st.Pop()
				ns = ns + cstr
			}
			if st.Peek() == "(" {
				st.Pop()
			}
		} else {
			data := st.Peek()
			peekele := data.(string)
			for st.Peek() != "N" && (prec(string(s[i])) <= prec(peekele)) && st.Peek() != "(" {
				c := st.Peek()
				cstr := c.(string)
				st.Pop()
				ns = ns + cstr
			}
			st.Push(string(s[i]))
		}
	}
	for st.Peek() != "N" {
		c := st.Peek()
		cstr := c.(string)
		st.Pop()
		ns = ns + cstr
	}
	return ns
}
