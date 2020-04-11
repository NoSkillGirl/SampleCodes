// Prefix : An expression is called the prefix expression if the operator appears in the
// expression before the operands. Simply of the form (operator operand1 operand2).

// Postfix: An expression is called the postfix expression if the operator appears in the
// expression after the operands. Simply of the form (operand1 operand2 operator).

// Prefix :  *-A/BC-/AKL
// Postfix : ABC/-AK/L-*

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
	preExp := "*-A/BC-/AKL"
	fmt.Println("Prefix expression: ", preExp)
	s := stack.New()
	l := len(preExp)
	for i := l - 1; i >= 0; i-- {
		if isOperator(string(preExp[i])) == true {
			op1 := s.Peek().(string)
			s.Pop()
			op2 := s.Peek().(string)
			s.Pop()
			temp := op1 + op2 + string(preExp[i])
			s.Push(temp)
		} else {
			s.Push(string(preExp[i]))
		}
	}
	fmt.Println("Postfix expression: ", s.Peek())
}
