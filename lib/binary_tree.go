package lib

import "fmt"

type Node struct {
	data  float32
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (bt *BinaryTree) Insert(value float32) {
	node := &Node{data: value}
	if bt.root == nil {
		bt.root = node
	} else {
		bt.InsertNode(bt.root, node)
	}
}

func (bt *BinaryTree) InsertNode(root, node *Node) {
	if node.data < root.data {
		if root.left == nil {
			root.left = node
		} else {
			bt.InsertNode(root.left, node)
		}
	} else if node.data > root.data {
		if root.right == nil {
			root.right = node
		} else {
			bt.InsertNode(root.right, node)
		}
	}
}

func (bt *BinaryTree) GetTree() {
	PrintNode(bt.root, "", true)
}

func PrintNode(node *Node, prefix string, isTail bool) {
	if node == nil {
		return
	}
	fmt.Printf("%s%s %d\n", prefix, ifThenElse(isTail, "└──", "├──"), int(node.data))
	newPrefix := prefix + ifThenElse(isTail, "    ", "│   ")
	PrintNode(node.left, newPrefix, node.right == nil)
	PrintNode(node.right, newPrefix, true)
}

func ifThenElse(condition bool, a, b string) string {
	if condition {
		return a
	}
	return b
}
