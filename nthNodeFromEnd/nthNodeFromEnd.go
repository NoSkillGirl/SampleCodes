// Program for n’th node from the end of a Linked List
// Given a Linked List and a number n, write a function that returns the value at
// the n’th node from the end of the Linked List.

// Time Complexity: O(n)

package main

import "fmt"

//Node struct
type Node struct {
	value string
	next  *Node
}

// LinkedList struct
type LinkedList struct {
	length int
	first  *Node
}

func main() {
	ll := &LinkedList{}

	ll.Append(&Node{value: "n1"})
	ll.Append(&Node{value: "n2"})
	ll.Append(&Node{value: "n3"})
	ll.Append(&Node{value: "n4"})
	ll.Append(&Node{value: "n5"})
	ll.Append(&Node{value: "n6"})

	nthValueFromEnd := ll.ValueFromEndNode(2)
	fmt.Println("Required Value: ", nthValueFromEnd)
}

//Append node
func (ll *LinkedList) Append(newNode *Node) {
	if ll.length == 0 {
		ll.first = newNode
	} else {
		currentNode := ll.first
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
	ll.length++
}

//ValueFromEndNode function
func (ll *LinkedList) ValueFromEndNode(pos int) string {
	endPosition := *ll.first
	reqPostion := *ll.first
	if pos > 1 {
		for i := 1; i < pos; i++ {
			endPosition = *endPosition.next
		}

		for endPosition.next != nil {
			endPosition = *endPosition.next
			reqPostion = *reqPostion.next
		}
	} else {
		for reqPostion.next != nil {
			reqPostion = *reqPostion.next
		}
	}
	return reqPostion.value
}
