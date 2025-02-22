// https://leetcode.com/problems/find-elements-in-a-contaminated-binary-tree/
// Level: Medium

package leetcode

import "leetcode/go/structures"

type TreeNode = structures.TreeNode

type FindElements struct {
	nodes map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	fe := FindElements{nodes: make(map[int]bool)}
	if root != nil {
		root.Val = 0
		fe.recover(root)
	}
	return fe
}

func (this *FindElements) recover(root *TreeNode) {
	if root == nil {
		return
	}

	this.nodes[root.Val] = true

	if root.Left != nil {
		root.Left.Val = root.Val*2 + 1
		this.recover(root.Left)
	}
	if root.Right != nil {
		root.Right.Val = root.Val*2 + 2
		this.recover(root.Right)
	}
}

func (this *FindElements) Find(target int) bool {
	_, exists := this.nodes[target]
	return exists
}
