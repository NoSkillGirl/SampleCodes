// Detect loop in a linked list
// Find length of loop in linked list

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
	// ll.Append(&Node{value: "n6"})
	// ll.Append(&Node{value: "n7"})
	ll.Append(&Node{value: "n5"})

	loopdetected, length := ll.LoopDetection()

	if loopdetected == true {
		fmt.Println("Loop is present in LL and Length of loop : ", length)
	} else {
		fmt.Println("Loop is not present in LL")
	}

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
	if newNode.value == "n5" {
		newNode.next = ll.first.next
	}
	ll.length++
}

//LoopDetection in ll
func (ll *LinkedList) LoopDetection() (bool, int) {
	endPosition := *ll.first
	reqPostion := *ll.first
	count := 0
	for endPosition.next != nil {
		endPosition = *endPosition.next
		count = count + 1
		if count%2 == 0 {
			reqPostion = *reqPostion.next
		}

		if endPosition == reqPostion {
			elementCount := 1
			endPosition = *endPosition.next
			for endPosition.next != reqPostion.next {
				endPosition = *endPosition.next
				elementCount = elementCount + 1
			}
			return true, elementCount
		}
	}
	return false, 0
}
