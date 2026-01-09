// https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/
// Level: Medium

package leetcode

import (
	structures "leetcode/go/structures"
)

// TreeNode is to define the struct
type TreeNode = structures.TreeNode

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode) (int, *TreeNode)
	dfs = func(node *TreeNode) (int, *TreeNode) {
		if node == nil {
			return 0, nil
		}

		ld, ln := dfs(node.Left)
		rd, rn := dfs(node.Right)
		if ld > rd {
			return ld + 1, ln
		}
		if rd > ld {
			return rd + 1, rn
		}
		return ld + 1, node
	}
	_, ans := dfs(root)
	return ans
}
