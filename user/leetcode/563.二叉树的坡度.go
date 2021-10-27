/*
 * @lc app=leetcode.cn id=563 lang=golang
 *
 * [563] 二叉树的坡度
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
var total int

func findTilt(root *TreeNode) int {
	total = 0
	sum(root)
	return total
}

func sum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sl := sum(root.Left)
	sr := sum(root.Right)
	if sl > sr {
		total += sl - sr
	} else {
		total += sr - sl
	}
	return sl + sr + root.Val
}

// @lc code=end
