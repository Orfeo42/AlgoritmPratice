package leetcode

//https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	return ric(root, 0).node
}

func ric(node *TreeNode, depth int) Result {
	if node == nil {
		return Result{node: nil, depth: depth}
	}
	left := ric(node.Left, depth+1)
	right := ric(node.Right, depth+1)
	if left.depth == right.depth {
		return Result{node: node, depth: left.depth}
	}
	if left.depth > right.depth {
		return left
	}
	return right
}

type Result struct {
	node  *TreeNode
	depth int
}
