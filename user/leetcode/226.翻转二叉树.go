/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
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
// 一步一交换
// func invertTree(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return root
// 	}
// 	var temp *TreeNode
// 	temp = root.Left
// 	root.Left = root.Right
// 	root.Right = temp
// 	root.Left = invertTree(root.Left)
// 	root.Right = invertTree(root.Right)
// 	return root
// }
// 先交换后回溯
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

// @lc code=end
