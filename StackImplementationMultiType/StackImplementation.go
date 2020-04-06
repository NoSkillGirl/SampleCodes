package main

import (
	"fmt"
)

type item struct {
	//value as interface type to hold any data type
	value interface{}
	next  *item
}

//Stack struct
type Stack struct {
	top  *item
	size int
}

//Len - length of stack
func (stack *Stack) Len() int {
	return stack.size
}

//Push - add element to stack
func (stack *Stack) Push(value interface{}) {
	stack.top = &item{
		value: value,
		next:  stack.top,
	}
	stack.size++
}

//Pop - remove element from stack
func (stack *Stack) Pop() (value interface{}) {
	if stack.Len() > 0 {
		value = stack.top.value
		stack.top = stack.top.next
		stack.size--
		return
	}

	return nil
}

func main() {
	stack := new(Stack)
	// Push different data type to the stack
	stack.Push(1)
	stack.Push("Welcome")
	stack.Push(4.0)

	// Pop until stack is empty
	for stack.Len() > 0 {
		fmt.Println(stack.Pop())
	}
}
