// https://leetcode.com/problems/linked-list-in-binary-tree/
// Level: Medium

package leetcode

import (
	structures "leetcode/go/structures"
)

// ListNode is to define the struct
type ListNode = structures.ListNode

// TreeNode is to define the struct
type TreeNode = structures.TreeNode

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}

	var checkPath func(*ListNode, *TreeNode) bool
	checkPath = func(head *ListNode, root *TreeNode) bool {
		if head == nil {
			return true
		}
		if root == nil || head.Val != root.Val {
			return false
		}
		return checkPath(head.Next, root.Left) || checkPath(head.Next, root.Right)
	}

	if checkPath(head, root) {
		return true
	}
	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}
