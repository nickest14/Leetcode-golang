// https://leetcode.com/problems/reverse-odd-levels-of-binary-tree/
// Level: Medium

package leetcode

import (
	structures "leetcode/go/structures"
)

// TreeNode is to define the struct
type TreeNode = structures.TreeNode

func reverseOddLevels(root *TreeNode) *TreeNode {
	var reverse func(node1, node2 *TreeNode, isOdd bool)
	reverse = func(node1, node2 *TreeNode, isOdd bool) {
		if node1 == nil || node2 == nil {
			return
		}
		if isOdd {
			node1.Val, node2.Val = node2.Val, node1.Val
		}
		reverse(node1.Left, node2.Right, !isOdd)
		reverse(node1.Right, node2.Left, !isOdd)
	}

	reverse(root.Left, root.Right, true)
	return root
}
