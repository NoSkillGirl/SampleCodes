// Given an expression string exp , write a program to examine whether the pairs
// and the orders of “{“,”}”,”(“,”)”,”[“,”]” are correct in exp.

// Input: exp = “[()]{}{[()()]()}”
// Output: Balanced
// Input: exp = “[(])”
// Output: Not Balanced

package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func main() {
	exp := "[()]{}{[()()]()}"
	l := len(exp)
	st := stack.New()
	balanced := true
	for i := 0; i < l; i++ {
		if string(exp[i]) == "(" || string(exp[i]) == "{" || string(exp[i]) == "[" {
			st.Push(string(exp[i]))
		} else if string(exp[i]) == ")" {
			if st.Peek() != "(" {
				balanced = false
				break
			}
			st.Pop()
		} else if string(exp[i]) == "}" {
			if st.Peek() != "{" {
				balanced = false
				break
			}
			st.Pop()
		} else if string(exp[i]) == "]" {
			if st.Peek() != "[" {
				balanced = false
				break
			}
			st.Pop()
		}
	}
	if balanced == true {
		fmt.Println("Balanced")
	} else {
		fmt.Println("Not Balanced")
	}
}
