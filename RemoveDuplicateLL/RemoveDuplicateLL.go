package main

import "fmt"

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

func main() {
	ll := &LinkedList{}
	ll.Append(&Node{value: 1})
	ll.Append(&Node{value: 3})
	ll.Append(&Node{value: 3})
	ll.Append(&Node{value: 3})
	ll.Append(&Node{value: 5})
	ll.Append(&Node{value: 5})
	ll.Append(&Node{value: 7})

	//removing duplicates
	ll.DeleteDuplicateSortedLL()
	//printing ll
	ll.PrintLL()

	ll1 := &LinkedList{}
	ll1.Append(&Node{value: 4})
	ll1.Append(&Node{value: 3})
	ll1.Append(&Node{value: 2})
	ll1.Append(&Node{value: 3})
	ll1.Append(&Node{value: 4})
	ll1.Append(&Node{value: 5})
	ll1.Append(&Node{value: 4})

	//removing duplicates
	ll1.DeleteDuplicateUnsortedLL()
	//printing ll
	ll1.PrintLL()
}

//PrintLL print linked list
func (ll *LinkedList) PrintLL() {
	fmt.Println("Linked List: ")
	fmt.Printf("Length: %v\n", ll.length)
	currentPosition := *ll.first
	for currentPosition.next != nil {
		fmt.Println(currentPosition)
		currentPosition = *currentPosition.next
	}
	fmt.Println(currentPosition)
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

// DeleteDuplicateSortedLL func
func (ll *LinkedList) DeleteDuplicateSortedLL() {
	currentNode := ll.first
	if currentNode == nil {
		return
	}
	for currentNode.next != nil {
		if currentNode.value == currentNode.next.value {
			nextNode := currentNode.next.next
			currentNode.next = nextNode
			ll.length--
		} else {
			currentNode = currentNode.next
		}
	}
}

// DeleteDuplicateUnsortedLL func
func (ll *LinkedList) DeleteDuplicateUnsortedLL() {
	s := make([]int, 1)
	currentNode := ll.first
	var previousNode *Node
	previousNode = nil
	for currentNode != nil {
		if findEle(s, currentNode.value) == true {
			previousNode.next = currentNode.next
			ll.length--
		} else {
			s = append(s, currentNode.value)
			previousNode = currentNode
		}
		currentNode = previousNode.next
	}
}

func findEle(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
