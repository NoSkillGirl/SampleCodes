package main

import (
	"fmt"
	"io"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(data int64) *BinaryTree {
	if t.root == nil {
		fmt.Println("1............")
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		fmt.Println("2............")
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data int64) {
	if n == nil {
		fmt.Println("n == nil return")
		return
	} else if data <= n.data {
		fmt.Println("else if .........data: ", data, " n.data: ", n.data)
		if n.left == nil {
			fmt.Println("n.left is nil")
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			fmt.Println("else..........n.left: ", n.left)
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			fmt.Println("n.right is nil")
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			fmt.Println("else..........n.right: ", n.right)
			n.right.insert(data)
		}
	}
}

func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func main() {
	tree := &BinaryTree{}
	// tree.insert(100).
	// 	insert(-20).
	// 	insert(-50).
	// 	insert(-15).
	// 	insert(-60).
	// 	insert(50).
	// 	insert(60).
	// 	insert(55).
	// 	insert(85).
	// 	insert(15).
	// 	insert(5).
	// 	insert(-10)
	// print(os.Stdout, tree.root, 0, 'M')

	tree.insert(4).
		insert(2).
		insert(6).
		insert(1).
		insert(3).
		insert(5).
		insert(7)

	fmt.Println("Preorder: ")
	tree.root.preorder()

	fmt.Println("Inorder: ")
	tree.root.inorder()

	fmt.Println("Postorder: ")
	tree.root.postorder()
}

func (n *BinaryNode) preorder() {
	if n == nil {
		return
	}

	fmt.Println(n.data)
	n.left.preorder()
	n.right.preorder()
}

func (n *BinaryNode) inorder() {
	if n == nil {
		return
	}
	n.left.inorder()
	fmt.Println(n.data)
	n.right.inorder()
}

func (n *BinaryNode) postorder() {
	if n == nil {
		return
	}
	n.left.postorder()
	n.right.postorder()
	fmt.Println(n.data)
}
