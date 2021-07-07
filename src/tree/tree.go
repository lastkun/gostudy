package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Nil node")
		return
	}

	node.Value = value
}

//
//func (node *Node) Traverse() {
//	if node == nil {
//		return
//	}
//	node.Left.Traverse()
//	node.Print()
//	node.Right.Traverse()
//}

func (node *Node) Traverse() {
	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
	fmt.Println()
}

// 函数式编程
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}
