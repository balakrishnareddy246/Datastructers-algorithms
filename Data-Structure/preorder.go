package main

import (
	"fmt"

	"github.com/emirpasic/gods/stacks/arraystack"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func PreorderIterative(root *Node) {
	if root == nil {
		return
	}

	current := root
	stack := arraystack.New()
	stack.Push(current)
	for stack.Size() > 0 {
		temp, _ := stack.Pop()
		current = temp.(*Node)
		fmt.Printf("%d ", current.val)
		if current.right != nil {
			stack.Push(current.right)
		}
		if current.left != nil {
			stack.Push(current.left)
		}
	}
}

func main() {

	root := &Node{10, nil, nil}
	root.left = &Node{20, nil, nil}
	root.right = &Node{30, nil, nil}
	root.left.left = &Node{40, nil, nil}
	root.left.right = &Node{50, nil, nil}
	root.right.right = &Node{60, nil, nil}
	root.left.left.left = &Node{70, nil, nil}

	fmt.Println("Preorder Traversal - Iterative Solution : ")
	PreorderIterative(root)
}
