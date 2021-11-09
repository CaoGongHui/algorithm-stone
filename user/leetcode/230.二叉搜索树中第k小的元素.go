/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	// 统计左子树的节点数
	leftCount := count(root.Left)
	if leftCount+1 == k {
		//左子树的节点数加1 等于k 那么root就是第k小的元素
		return root.Val
	} else if leftCount >= k {
		// 左子树的节点数大于等于k 那么在左子树中找第k小的元素
		return kthSmallest(root.Left, k)
	} else {
		// 左子树的节点数小于k 那么在右子树中找第k小的元素
		// 需要找的个数就是k减去左子树的节点数加上1 这个1 就是root节点
		return kthSmallest(root.Right, k-leftCount-1)
	}
}
func count(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return count(node.Left) + count(node.Right) + 1
}

// @lc code=end

// func kthSmallest(root *TreeNode, k int) int {
// 	if root == nil {
// 		return 0
// 	}
// 	res := 0
// 	var dfs func(*TreeNode)
// 	dfs = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		dfs(node.Left)
// 		k--
// 		if k == 0 {
// 			res = node.Val
// 			return
// 		}
// 		dfs(node.Right)
// 	}
// 	dfs(root)
// 	return res
// }
