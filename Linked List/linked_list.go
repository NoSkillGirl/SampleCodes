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

//insert, search, delete, update
func main() {
	ll := &LinkedList{}

	ll.Append(&Node{value: "n1"})
	ll.Append(&Node{value: "n2"})
	ll.Append(&Node{value: "n3"})
	ll.Insert(&Node{value: "n4"}, 1)
	ll.Insert(&Node{value: "n5"}, 8)
	ll.Insert(&Node{value: "n5"}, 3)
	ll.Delete(3)
	ll.Update("uyfghuif", 1)
	ll.Update("tyfdgui", 3)

	//Print all linked list value
	currentPosition := *ll.first
	fmt.Printf("Length: %v\n", ll.length)
	for currentPosition.next != nil {
		fmt.Println(currentPosition.value)
		currentPosition = *currentPosition.next
	}
	fmt.Println("Last: ", currentPosition.value)
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

//Insert function
func (ll *LinkedList) Insert(newNode *Node, pos int) {
	if ll.length < pos {
		fmt.Println("Linked in list is less than the number provider")
	} else if pos == 1 {
		newNode.next = ll.first
		ll.first = newNode
		ll.length++
	} else {
		currentNode := ll.first
		for i := 1; i < pos-1; i++ {
			currentNode = currentNode.next
		}
		(*newNode).next = currentNode.next
		currentNode.next = newNode
		ll.length++
	}
}

//Delete function
func (ll *LinkedList) Delete(pos int) {
	if ll.length < pos {
		fmt.Println("Linked in list is less than the number provider")
	} else if pos == 1 {
		ll.first = ll.first.next
		ll.length--
	} else {
		currentNode := ll.first
		for i := 1; i < pos-1; i++ {
			currentNode = currentNode.next
		}
		deleteNode := currentNode.next
		currentNode.next = deleteNode.next
		ll.length--
	}
}

//Update function
func (ll *LinkedList) Update(updateValue string, pos int) {
	if ll.length < pos {
		fmt.Println("Linked in list is less than the number provider")
	} else if pos == 1 {
		ll.first.value = updateValue
	} else {
		currentNode := ll.first
		for i := 1; i < pos; i++ {
			currentNode = currentNode.next
		}
		currentNode.value = updateValue
	}
}
