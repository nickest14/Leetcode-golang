// https://leetcode.com/problems/binary-tree-level-order-traversal/
// Level: Medium

package leetcode

import (
	structures "leetcode/go/structures"
)

// TreeNode is to define the struct
type TreeNode = structures.TreeNode

func maxLevelSum(root *TreeNode) int {
	values := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelValue := 0
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			levelValue += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		values = append(values, levelValue)
	}

	maxValue := values[0]
	maxIdx := 0
	for i, v := range values {
		if v > maxValue {
			maxValue = v
			maxIdx = i
		}
	}
	return maxIdx + 1
}
