package lib

import "fmt"

type Node struct {
	Value float32
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value float32) {
	if n == nil {
		return
	} else if value <= n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

func PrintTree(node *Node, level int) {
	if node == nil {
		return
	}
	PrintTree(node.Right, level+2)
	for i := 0; i < level; i++ {
		fmt.Print("   ")
	}
	fmt.Println(node.Value)
	PrintTree(node.Left, level+2)
}
