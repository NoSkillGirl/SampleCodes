// Infix expression: The expression of the form a op b. When an operator is in-between
// every pair of operands.
// Postfix expression: The expression of the form a b op. When an operator is followed
// for every pair of operands.

// Postfix : ab*c+
// Infix : ((a*b)+c)

package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func main() {
	st := stack.New()
	s := "ab*c+"
	fmt.Println("Postfix expression: ", s)
	l := len(s)
	for i := 0; i < l; i++ {
		if (string(s[i]) >= "a" && string(s[i]) <= "z") || (string(s[i]) >= "A" && string(s[i]) <= "Z") {
			st.Push(string(s[i]))
		} else {
			op1 := st.Peek().(string)
			st.Pop()
			op2 := st.Peek().(string)
			st.Pop()
			temp := "(" + op2 + string(s[i]) + op1 + ")"
			st.Push(temp)
		}
	}
	fmt.Println("Postfix expression: ", st.Peek())
}
