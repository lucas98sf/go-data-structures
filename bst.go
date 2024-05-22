package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
	depth int
}

func NewTree(rootValue int) *Node {
	return &Node{
		rootValue,
		nil,
		nil,
		0,
	}
}

// no default parameters huh
func (root *Node) AddNode(value int, _depth ...int) *Node {
	depth := 0
	if len(_depth) > 0 {
		depth = _depth[0]
	}
	if value < root.value {
		if root.left != nil {
			return root.left.AddNode(value, depth+1)
		}
		root.left = &Node{
			value,
			nil,
			nil,
			depth + 1,
		}
	} else {
		if root.right != nil {
			return root.right.AddNode(value, depth+1)
		}
		root.right = &Node{
			value,
			nil,
			nil,
			depth + 1,
		}
	}
	return root
}
func (root *Node) MaxDepth() int {
	lDepth := 0
	rDepth := 0
	if root.right != nil {
		lDepth = lDepth + 1 + root.right.MaxDepth()
	}
	if root.left != nil {
		rDepth = rDepth + 1 + root.left.MaxDepth()
	}
	return max(lDepth, rDepth)
}
func (root *Node) Depth(value int) int {
	if root.value == value {
		return root.depth
	}
	if root.left != nil {
		return root.left.Depth(value)
	}
	if root.right != nil {
		return root.right.Depth(value)
	}
	return 0
}
func (root *Node) Print(_space ...int) {
	space := 0
	if len(_space) > 0 {
		space = _space[0]
	}
	const DEFAULT_SPACE = 10
	space = space + DEFAULT_SPACE

	if root.right != nil {
		root.right.Print(space)
	}

	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
	// fmt.Printf("(depth: %d) ", root.depth)
	fmt.Println(root.value)

	if root.left != nil {
		root.left.Print(space)
	}
}

func main() {
	tree := NewTree(3)
	tree.AddNode(2)
	tree.AddNode(6)
	tree.AddNode(4)
	tree.AddNode(7)
	tree.AddNode(8)
	tree.AddNode(9)
	tree.AddNode(1)
	tree.Print()
}
