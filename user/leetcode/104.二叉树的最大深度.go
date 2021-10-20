/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 *
 * https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (76.64%)
 * Likes:    1004
 * Dislikes: 0
 * Total Accepted:    527.8K
 * Total Submissions: 688.9K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最大深度。
 *
 * 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例：
 * 给定二叉树 [3,9,20,null,null,15,7]，
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回它的最大深度 3 。
 *
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 广度优先搜索
func maxDepthByWidth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	query := []*TreeNode{root}
	ans := 0
	for len(query) > 0 {
		sz := len(query)
		for sz > 0 {
			node := query[0]
			query = query[1:]
			if node.Left != nil {
				query = append(query, node.Left)
			}
			if node.Right != nil {
				query = append(query, node.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}

// @lc code=end
