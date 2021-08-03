package main

import "fmt"

type Node struct {
	value       int
	left, right *Node
}

func (node Node) print() {
	fmt.Println(node.value)
}

func (node *Node) setValue(value int) {
	node.value = value
}

func createNode(value int) *Node {
	return &Node{value: value}
}
func (node *Node) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

/**
     3
   0    5
  9 2  6
*/

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.left.TraverseFunc(f)
	f(node)
	node.right.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()

	return out
}

func main() {
	var root Node
	//golang中都是值传递
	root = Node{value: 3}

	root.left = &Node{} // 默认值为0
	root.left.left = createNode(9)
	root.right = &Node{5, nil, nil}
	root.left.right = createNode(2)

	root.right.left = new(Node)
	root.right.left.setValue(6)
	//root.right.left.print()

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		node.print()
		if node.value > maxNode {
			maxNode = node.value
		}
	}
	fmt.Println("Max node :", maxNode)

}
