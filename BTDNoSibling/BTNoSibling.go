// Print all nodes that don’t have sibling
// Given a Binary Tree, print all nodes that don’t have a sibling (a sibling is a node
// that has same parent. In a Binary Tree, there can be at most one sibling).
// Root should not be printed as root cannot have a sibling.

// For example, the output should be “4 5 6” for the following tree.
//			1
//		/	|	\
//	   2	|	  3
//		\	|	 /
//		 4	| 	 5
//		 	|   /
//			|  6

package main

import (
	"fmt"
)

// BinaryNode struct
type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

//BinaryTree struct
type BinaryTree struct {
	root *BinaryNode
}

func main() {
	//create binary tree given in the above example
	tree := &BinaryTree{}
	tree.root = &BinaryNode{data: 1, left: nil, right: nil}
	tree.root.left = &BinaryNode{data: 2, left: nil, right: nil}
	tree.root.right = &BinaryNode{data: 3, left: nil, right: nil}
	tree.root.left.right = &BinaryNode{data: 4, left: nil, right: nil}
	tree.root.right.left = &BinaryNode{data: 5, left: nil, right: nil}
	tree.root.right.left.left = &BinaryNode{data: 6, left: nil, right: nil}
	tree.root.printSingles()
}

func (n *BinaryNode) printSingles() {
	// Base case
	if n == nil {
		return
	}

	// If this is an internal node, recur for left and right subtrees
	if n.left != nil && n.right != nil {
		n.left.printSingles()
		n.right.printSingles()
	} else if n.right != nil {
		// If left child is NULL and right is not, print right child
		// and recur for right child
		fmt.Print(n.right.data, " ")
		n.right.printSingles()
	} else if n.left != nil {
		// If right child is NULL and left is not, print left child
		// and recur for left child
		fmt.Print(n.left.data, " ")
		n.left.printSingles()
	}
}
