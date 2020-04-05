// Find the middle of a given linked list

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
	ll.Append(&Node{value: "n7"})

	midEle := ll.FindMidValue()
	fmt.Println("Required Value: ", midEle)
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

//FindMidValue func
func (ll *LinkedList) FindMidValue() string {
	endPosition := *ll.first
	reqPostion := *ll.first
	count := 0
	for endPosition.next != nil {
		endPosition = *endPosition.next
		count = count + 1
		if count%2 == 0 {
			reqPostion = *reqPostion.next
		}
	}

	return reqPostion.value
}
