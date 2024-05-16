package lib

import (
	"fmt"
	"strings"
)

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
	PrintNode(bt.root, 0, false, bt.root.right != nil)
}

func PrintNode(node *Node, count int, isParentLeft, isParentHaveRight bool) {
	postfix := ""
	if count == 0 {
		fmt.Println(node.data)
	} else {
		if isParentLeft && isParentHaveRight {
			postfix = "├──"
		} else {
			postfix = "└──"
		}
		fmt.Printf("%s%s %f\n", strings.Repeat(" ", count-4), postfix, node.data)
	}

	if node.left != nil {
		PrintNode(node.left, count+4, true, node.right != nil)
	}
	if node.right != nil {
		PrintNode(node.right, count+4, false, false)
	}
}
