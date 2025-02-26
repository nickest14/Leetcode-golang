// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/
// Level: Medium

package leetcode

import (
	structures "leetcode/go/structures"
)

// TreeNode is to define the struct
type TreeNode = structures.TreeNode

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	preIndex := len(preorder) - 1
	postIndex := len(postorder) - 1

	var makeTree func() *TreeNode
	makeTree = func() *TreeNode {
		node := &TreeNode{Val: postorder[postIndex]}
		postIndex--
		if node.Val != preorder[preIndex] { // post = [left, right], pre = [root, left, right]
			node.Right = makeTree()
		}

		if node.Val != preorder[preIndex] { // post = [left], pre = [root, left]
			node.Left = makeTree()
		}

		preIndex-- // post = [], pre = [root], root already used for node.Val

		return node
	}
	return makeTree()
}
