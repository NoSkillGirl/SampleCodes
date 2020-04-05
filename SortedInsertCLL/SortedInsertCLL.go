// Sorted insert for circular linked list
// Write a function to insert a new value in a sorted Circular Linked List (CLL)

package main

import (
	"fmt"
)

//Node struct
type Node struct {
	value int
	next  *Node
}

// LinkedList struct
type LinkedList struct {
	length int
	first  *Node
}

//insert, search, delete, update
func main() {
	ll := &LinkedList{}
	ll.Append(&Node{value: 1})
	ll.Append(&Node{value: 2})
	ll.Append(&Node{value: 3})
	ll.Append(&Node{value: 4})
	ll.PrintLL()

	ll1 := &LinkedList{}
	ll1.SortedAppend(&Node{value: 5})
	ll1.SortedAppend(&Node{value: 4})
	ll1.SortedAppend(&Node{value: 7})
	ll1.SortedAppend(&Node{value: 2})
	ll1.SortedAppend(&Node{value: 6})
	ll1.PrintLL()

}

//PrintLL print linked list
func (ll *LinkedList) PrintLL() {
	fmt.Println("Linked List: ")
	currentPosition := *ll.first
	fmt.Println(currentPosition)
	currentPosition = *currentPosition.next
	for currentPosition.next != ll.first {
		fmt.Println(currentPosition)
		currentPosition = *currentPosition.next
	}
	fmt.Println(currentPosition)
}

//Append node
func (ll *LinkedList) Append(newNode *Node) {
	if ll.length == 0 {
		ll.first = newNode
	} else if ll.length == 1 {
		ll.first.next = newNode
		newNode.next = ll.first
	} else {
		currentNode := ll.first.next
		for currentNode.next != ll.first {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
		newNode.next = ll.first
	}
	ll.length++
}

//SortedAppend node
func (ll *LinkedList) SortedAppend(newNode *Node) {
	if ll.length == 0 {
		newNode.next = newNode
		ll.first = newNode
	} else {
		if newNode.value < ll.first.value {
			currentNode := ll.first
			for currentNode.next != ll.first {
				currentNode = currentNode.next
			}
			currentNode.next = newNode
			newNode.next = ll.first
			ll.first = newNode
		} else {
			currentNode := ll.first
			for currentNode.next != ll.first && currentNode.next.value < newNode.value {
				currentNode = currentNode.next
			}
			newNode.next = currentNode.next
			currentNode.next = newNode
		}
	}
	ll.length++
}
