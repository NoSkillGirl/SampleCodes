// Slice Implementation - Uniform Type

package main

import (
	"fmt"
)

type customQueue struct {
	stack []string
}

func (c *customQueue) Push(name string) {
	c.stack = append(c.stack, name)
}

func (c *customQueue) Pop() error {
	len := len(c.stack)
	if len > 0 {
		c.stack = c.stack[:len-1]
		return nil
	}
	return fmt.Errorf("Pop Error: Queue is empty")
}

func (c *customQueue) Front() (string, error) {
	len := len(c.stack)
	if len > 0 {
		return c.stack[len-1], nil
	}
	return "", fmt.Errorf("Peep Error: Queue is empty")
}

func (c *customQueue) Size() int {
	return len(c.stack)
}

func (c *customQueue) Empty() bool {
	return len(c.stack) == 0
}

func main() {
	customQueue := &customQueue{
		stack: make([]string, 0),
	}
	fmt.Println("Push: A")
	customQueue.Push("A")
	fmt.Println("push: B")
	customQueue.Push("B")
	fmt.Printf("Size: %d\n", customQueue.Size())
	for customQueue.Size() > 0 {
		frontVal, _ := customQueue.Front()
		fmt.Printf("Front: %s\n", frontVal)
		fmt.Printf("Pop: %s\n", frontVal)
		customQueue.Pop()
	}
	fmt.Printf("Size: %d\n", customQueue.Size())
}
