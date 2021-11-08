/*
/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}
func helper(node *TreeNode, p, q int) bool {
	if node == nil {
		return true
	}
	if node.Val <= p || node.Val >= q {
		return false
	}
	return helper(node.Left, p, node.Val) && helper(node.Right, node.Val, q)
}
