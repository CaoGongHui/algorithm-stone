/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}
func check(p *TreeNode, q *TreeNode) bool {
	if q == nil || p == nil {
		return p == q
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

// @lc code=end
