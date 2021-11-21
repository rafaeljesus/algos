package main

type BinaryTree struct {
	Value interface{}
	Left  *BinaryTree
	Right *BinaryTree
}

func DFS(tree *BinaryTree) {
	if tree == nil {
		return
	}
	DFS(tree.Left)
	DFS(tree.Right)
}
