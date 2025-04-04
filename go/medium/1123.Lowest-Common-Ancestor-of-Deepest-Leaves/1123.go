// https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves/
// Level: Medium

package leetcode

import (
	structures "leetcode/go/structures"
)

// TreeNode is to define the struct
type TreeNode = structures.TreeNode

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode, depth int) (*TreeNode, int)
	dfs = func(node *TreeNode, depth int) (*TreeNode, int) {
		if node == nil {
			return nil, depth
		}
		lcaLeft, depthLeft := dfs(node.Left, depth+1)
		lcaRight, depthRight := dfs(node.Right, depth+1)

		if depthLeft == depthRight {
			return node, depthLeft
		} else if depthLeft > depthRight {
			return lcaLeft, depthLeft
		} else {
			return lcaRight, depthRight
		}
	}

	lcaNode, _ := dfs(root, 0)
	return lcaNode
}
