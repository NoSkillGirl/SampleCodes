// Infix expression:The expression of the form a op b. When an operator is in-between
// every pair of operands.
// Postfix expression:The expression of the form a b op. When an operator is followed
// for every pair of operands.

// Infix: a+b*(c^d-e)^(f+g*h)-i
// Postfix: abcd^e-fgh*+^*+i-

package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

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

func main() {
	st := stack.New()
	st.Push("N")
	s := "a+b*(c^d-e)^(f+g*h)-i"
	fmt.Println("Infix expression : ", s)
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
	fmt.Println("Postfix expression : ", ns)
}
