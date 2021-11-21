package invert_binary_tree

import (
	"fmt"
	"testing"
)

type BinaryTree struct {
	Value interface{}
	Left  *BinaryTree
	Right *BinaryTree
}

// time O(n) | space O(d) n is the number of nodes, and d is the depth (height) of the tree
func (tree *BinaryTree) Invert() {
	queue := []*BinaryTree{tree}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current != nil {
			current.Left, current.Right = current.Right, current.Right
			queue = append(queue, current.Left, current.Right)
		}
	}
}

/*
queue = [1]
len queue > 0? y
curr = 1
queue = []
curr != nil? y
curr.left = 3
curr.right = 2
queue = [3, 2]

len queue > 0? y
curr = 3
queue = [2]
curr != nil? y
curr.left = 7
curr.right = 6
len queue > 0? y
curr = 2
queue = []
curr != nil? y
curr.left = 5
curr.right = 4

queue = [5, 4]
len queue > 0? y
curr = 5
queue = [4]
curr != nil? y
curr.left = nil
curr.right = nil
queue = [4, nil, nil]
len queue > 0? y
curr = 4
queue = [nil, nil]
curr != nil? y
curr.left = 9
curr.right = 8
queue = [nil, nil, nil, nil]
...
curr = nil
queue[nil]
if curr != nil? n
curr = nil
queue = []
if curr != nil? n
len queue > 0? n
						1
				3				2
			7		6		5		4
								9		8
*/

func Test(t *testing.T) {
	tree := &BinaryTree{
		Value: 1,
		Left: &BinaryTree{
			Value: 2,
			Left: &BinaryTree{
				Value: 4,
				Left: &BinaryTree{
					Value: 8,
				},
				Right: &BinaryTree{
					Value: 9,
				},
			},
			Right: &BinaryTree{
				Value: 5,
			},
		},
		Right: &BinaryTree{
			Value: 3,
			Left: &BinaryTree{
				Value: 6,
			},
			Right: &BinaryTree{
				Value: 7,
			},
		},
	}
	tree.Invert()
	fmt.Println(tree)
}

/*
					1
			2				3
		4		5		6		7
	8		9
*/

// swap left to right on every iteration
// tree.left, tree.right = tree.right, tree.left
/*
				1
		3				2
	6		7		4		5
				8		9
*/

/*
				1
		3				2
	7		6		5		4
						8		9
*/

/*
				1
		3				2
	7		6		5		4
						9		8
*/
