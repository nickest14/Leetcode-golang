// https://leetcode.com/problems/find-largest-value-in-each-tree-row/
// Level: Medium

package leetcode

import (
	structures "leetcode/go/structures"
)

// TreeNode is to define the struct
type TreeNode = structures.TreeNode

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var ans = []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		maxValue := queue[0].Val
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Val > maxValue {
				maxValue = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, maxValue)
	}

	return ans
}
