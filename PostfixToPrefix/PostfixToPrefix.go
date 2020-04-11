// Postfix: An expression is called the postfix expression if the operator appears in the
// expression after the operands. Simply of the form (operand1 operand2 operator).

// Prefix : An expression is called the prefix expression if the operator appears in the
// expression before the operands. Simply of the form (operator operand1 operand2).

// Postfix : ABC/-AK/L-*
// Prefix :  *-A/BC-/AKL

package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

// function to check if character is operator or not
func isOperator(x string) bool {
	if x == "+" || x == "-" || x == "/" || x == "*" {
		return true
	}
	return false
}

func main() {
	postExp := "ABC/-AK/L-*"
	fmt.Println("Postfix expression: ", postExp)
	s := stack.New()
	l := len(postExp)
	for i := 0; i < l; i++ {
		if isOperator(string(postExp[i])) == true {
			op1 := s.Peek().(string)
			s.Pop()
			op2 := s.Peek().(string)
			s.Pop()
			temp := string(postExp[i]) + op2 + op1
			s.Push(temp)
		} else {
			s.Push(string(postExp[i]))
		}
	}
	fmt.Println("Prefix expression: ", s.Peek())
}
