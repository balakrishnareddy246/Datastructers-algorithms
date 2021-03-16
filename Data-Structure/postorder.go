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

func postorderIterative(root *Node) {
	if root == nil {
		return
	}

	temp := root
	stack := arraystack.New()
	for stack.Size() > 0 || temp != nil {
		if temp != nil {
			stack.Push(temp)
			temp = temp.left
		} else {
			obj, _ := stack.Pop()
			temp = temp.right
			fmt.Printf("%d ", temp.val)
			temp = obj.(*Node)
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

	fmt.Println("Inorder Traversal - iterative solution : ")
	postorderIterative(root)
}
