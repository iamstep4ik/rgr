package lib

import "fmt"

type Node struct {
	Value float32
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value float32) bool {
	if n == nil {
		return false
	} else if value == n.Value {
		return false
	} else if value < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
			return true
		} else {
			return n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
			return true
		} else {
			return n.Right.Insert(value)
		}
	}
}

func PrintTree(node *Node, level int) {
	if node == nil {
		return
	}
	PrintTree(node.Right, level+1)
	for i := 0; i < level; i++ {
		fmt.Print("    ")
	}
	fmt.Println(node.Value)
	PrintTree(node.Left, level+1)
}
