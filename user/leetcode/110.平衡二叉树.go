/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 自顶向下
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := depth(root.Left)
	r := depth(root.Right)
	if l-r > 1 || l-r < -1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}
func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(depth(node.Left), depth(node.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
//自底向上 优化
func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := height(root.Right)
	if rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}