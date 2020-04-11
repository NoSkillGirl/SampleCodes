// Infix : An expression is called the Infix expression if the operator appears in between
// the operands in the expression. Simply of the form (operand1 operator operand2).
// Prefix : An expression is called the prefix expression if the operator appears in the
// expression before the operands. Simply of the form (operator operand1 operand2).

// Prefix :  *-A/BC-/AKL
// Infix : ((A-(B/C))*((A/K)-L))

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
			temp := "(" + op1 + string(preExp[i]) + op2 + ")"
			s.Push(temp)
		} else {
			s.Push(string(preExp[i]))
		}
	}
	fmt.Println("Infix expression: ", s.Peek())
}
