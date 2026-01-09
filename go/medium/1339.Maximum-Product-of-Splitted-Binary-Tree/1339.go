// https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/
// Level: Medium

package leetcode

import (
	structures "leetcode/go/structures"
)

// TreeNode is to define the struct
type TreeNode = structures.TreeNode

func maxProduct(root *TreeNode) int {
	const MOD int = 1000000007
	var total, best int

	var sum func(*TreeNode) int
	sum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return node.Val + sum(node.Left) + sum(node.Right)
	}

	total = sum(root)

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		s := node.Val + dfs(node.Left) + dfs(node.Right)
		prod := s * (total - s)
		if prod > best {
			best = prod
		}
		return s
	}

	dfs(root)
	return best % MOD
}
