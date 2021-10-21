package pathsum

import "testing"

func TestPathSum(t *testing.T) {
	/* Constructed binary tree is
				10
			 /  \
		 8     2
		/ \   /
	 3   5 2
	*/
	root := &Node{Value: 10}
	root.Left = &Node{Value: 8}
	root.Right = &Node{Value: 2}
	root.Left.Left = &Node{Value: 3}
	root.Left.Right = &Node{Value: 5}
	root.Right.Left = &Node{Value: 2}
	if !HasPathSum(root, 21) {
		t.Error("expected HasPathSum to return true")
	}
}
