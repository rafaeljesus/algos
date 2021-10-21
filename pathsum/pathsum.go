package pathsum

type Node struct {
	Value       int
	Left, Right *Node
}

func HasPathSum(n *Node, sum int) bool {
	if n == nil {
		return false
	}
	if n.Left == nil && n.Right == nil {
		return n.Value-sum == 0
	}
	return HasPathSum(n.Left, sum-n.Value) || HasPathSum(n.Right, sum-n.Value)
}
