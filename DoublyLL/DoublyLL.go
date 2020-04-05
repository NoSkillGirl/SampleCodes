package main

import "fmt"

//Node struct
type Node struct {
	value    string
	next     *Node
	previous *Node
}

// LinkedList struct
type LinkedList struct {
	length int
	first  *Node
	last   *Node
}

func main() {
	ll := &LinkedList{}

	ll.Append(&Node{value: "n1"})
	ll.Append(&Node{value: "n2"})
	ll.Append(&Node{value: "n3"})
	ll.Append(&Node{value: "n4"})
	ll.Append(&Node{value: "n5"})
	ll.Append(&Node{value: "n6"})

	ll.PrintFromHead()
	ll.DeleteNode(3)
	ll.PrintFromHead()
	ll.DeleteNode(1)
	ll.PrintFromHead()
	ll.DeleteNode(4)
	ll.PrintFromHead()
	ll.ReverseLL()
	ll.PrintFromHead()

	ll.PrintFromTail()

	ll1 := &LinkedList{}
	ll1.InsertAtHead(&Node{value: "n1"})
	ll1.InsertAtHead(&Node{value: "n2"})
	ll1.InsertAtHead(&Node{value: "n3"})
	ll1.InsertAtHead(&Node{value: "n4"})
	ll1.InsertAtHead(&Node{value: "n5"})

	ll1.PrintFromHead()
}

//Append node - insert at end
func (ll *LinkedList) Append(newNode *Node) {
	if ll.length == 0 {
		ll.first = newNode
		ll.last = newNode
	} else {
		ll.last.next = newNode
		newNode.previous = ll.last
		ll.last = newNode
	}
	ll.length++
}

//InsertAtHead - insert at starting/head
func (ll *LinkedList) InsertAtHead(newNode *Node) {
	if ll.length == 0 {
		ll.first = newNode
		ll.last = newNode
	} else {
		newNode.next = ll.first
		ll.first.previous = newNode
		ll.first = newNode
	}
	ll.length++
}

//PrintFromHead to print LL
func (ll *LinkedList) PrintFromHead() {
	currentPosition := *ll.first
	fmt.Printf("Length: %v\n", ll.length)
	for currentPosition.next != nil {
		fmt.Println(currentPosition)
		currentPosition = *currentPosition.next
	}
	fmt.Println(currentPosition)
}

//PrintFromTail to print LL
func (ll *LinkedList) PrintFromTail() {
	fmt.Printf("Length: %v\n", ll.length)
	currentPosition := *ll.last
	for currentPosition.previous != nil {
		fmt.Println(currentPosition)
		currentPosition = *currentPosition.previous
	}
	fmt.Println(currentPosition)
}

//DeleteNode at provided location
func (ll *LinkedList) DeleteNode(pos int) {
	if pos > ll.length || pos < 1 {
		fmt.Println("Given position is not valid")
	}
	if pos == 1 {
		reqNode := ll.first.next
		reqNode.previous = nil
		ll.first = reqNode
	} else if pos == ll.length {
		reqNode := ll.last.previous
		reqNode.next = nil
		ll.last = reqNode
	} else {
		currentNode := ll.first
		for i := 1; i < pos-1; i++ {
			currentNode = currentNode.next
		}
		deleteNode := currentNode.next
		currentNode.next = deleteNode.next
		currentNode.next.previous = currentNode
	}
	ll.length--
}

//ReverseLL to reverse LL
func (ll *LinkedList) ReverseLL() {
	currentPosition := ll.first
	var tempNode *Node
	for currentPosition != nil {
		tempNode = currentPosition.previous
		currentPosition.previous = currentPosition.next
		currentPosition.next = tempNode
		currentPosition = currentPosition.previous
	}
	if tempNode != nil {
		ll.first = tempNode.previous
	}
}
