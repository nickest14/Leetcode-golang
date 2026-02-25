// https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/
// Level: Easy

package leetcode

import (
	structures "leetcode/go/structures"
)

// TreeNode is to define the struct
type TreeNode = structures.TreeNode

func sumRootToLeaf(root *TreeNode) int {
	var dfs func(node *TreeNode, currentSum int) int
	dfs = func(node *TreeNode, currentSum int) int {
		if node == nil {
			return 0
		}
		currentSum = currentSum*2 + node.Val
		if node.Left == nil && node.Right == nil {
			return currentSum
		}
		return dfs(node.Left, currentSum) + dfs(node.Right, currentSum)
	}

	return dfs(root, 0)
}
