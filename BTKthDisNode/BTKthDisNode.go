// Print all nodes that are at distance k from a leaf node
// Given a Binary Tree and a positive integer k, print all nodes that are
// distance k from a leaf node.

// Here the meaning of distance is different from previous post. Here k distance
// from a leaf means k levels higher than a leaf node. For example if k is more
// than height of Binary Tree, then nothing should be printed. Expected time
// complexity is O(n) where n is the number nodes in the given Binary Tree.

//			1
//		/	|	\
//	   2	|	  3
//	 /  \	|	 /  \
//	4	 5	|   6    7
//		 	|    \
//			|     8

// Nodes at distance 1 from leaf node are 2, 6 and 3
// Nodes at distance 2 from leaf node are 1 and 3

package main

import "fmt"

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
	tree := &BinaryTree{}

	tree.root = &BinaryNode{data: 1, left: nil, right: nil}
	tree.root.left = &BinaryNode{data: 2, left: nil, right: nil}
	tree.root.right = &BinaryNode{data: 3, left: nil, right: nil}
	tree.root.left.left = &BinaryNode{data: 4, left: nil, right: nil}
	tree.root.left.right = &BinaryNode{data: 5, left: nil, right: nil}
	tree.root.right.left = &BinaryNode{data: 6, left: nil, right: nil}
	tree.root.right.right = &BinaryNode{data: 7, left: nil, right: nil}
	tree.root.right.left.right = &BinaryNode{data: 8, left: nil, right: nil}

	fmt.Println("Nodes at distance 2 are :")
	tree.root.printKDistantfromLeaf(2)

	fmt.Println("\nNodes at distance 1 are :")
	tree.root.printKDistantfromLeaf(1)
}

func (n *BinaryNode) printKDistantfromLeaf(k int) int {
	if n == nil {
		return -1
	}
	lk := n.left.printKDistantfromLeaf(k)
	rk := n.right.printKDistantfromLeaf(k)
	// boolean isLeaf = lk == -1 && lk == rk;
	var isLeaf bool
	if lk == -1 && lk == rk {
		isLeaf = true
	}

	if lk == 0 || rk == 0 || (isLeaf && k == 0) {
		fmt.Print(" ", n.data)
	}
	if isLeaf && k > 0 {
		return k - 1
	} // leaf node
	if lk > 0 && lk < k {
		return lk - 1 // parent of left leaf
	}
	if rk > 0 && rk < k {
		return rk - 1
	}

	return -2
}
