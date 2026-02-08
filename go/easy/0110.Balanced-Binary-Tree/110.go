// https://leetcode.com/problems/balanced-binary-tree/
// Level: Easy

package leetcode

import (
	structures "leetcode/go/structures"
)

// TreeNode is to define the struct
type TreeNode = structures.TreeNode

func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight < 0 || rightHeight < 0 {
		return -1
	}
	diff := leftHeight - rightHeight
	if diff < 0 {
		diff = -diff
	}
	if diff > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}
