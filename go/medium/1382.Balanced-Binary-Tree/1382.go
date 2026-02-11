// https://leetcode.com/problems/balance-a-binary-search-tree/
// Level: Medium

package leetcode

import (
	structures "leetcode/go/structures"
)

// TreeNode is to define the struct
type TreeNode = structures.TreeNode

func balanceBST(root *TreeNode) *TreeNode {
	var nodes []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		nodes = append(nodes, node.Val)
		dfs(node.Right)
	}
	dfs(root)

	var build func(int, int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		m := (left + right) / 2
		cur := &TreeNode{Val: nodes[m]}
		cur.Left = build(left, m-1)
		cur.Right = build(m+1, right)
		return cur
	}
	return build(0, len(nodes)-1)
}
