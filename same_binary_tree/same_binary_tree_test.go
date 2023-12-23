package same_binary_tree

import "testing"

// TikTok SRE Interview Question

// Given the roots of two binary trees p and q, write a function to check if they are the same or not.
// Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
// Example 1:
// Input: p = [1,2,3], q = [1,2,3]
// Output: true
// Example 2:
// Input: p = [1,2], q = [1,null,2]
// Output: false
// Example 3:
// Input: p = [1,2,1], q = [1,1,2]
// Output: false

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func TestIsSameTree(t *testing.T) {
	p := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

	if !isSameTree(p, q) {
		t.Errorf("isSameTree(p, q) = %v; want true", isSameTree(p, q))
	}

	p = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	q = &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}

	if isSameTree(p, q) {
		t.Errorf("isSameTree(p, q) = %v; want false", isSameTree(p, q))
	}

	p = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}}
	q = &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}

	if isSameTree(p, q) {
		t.Errorf("isSameTree(p, q) = %v; want false", isSameTree(p, q))
	}
}
