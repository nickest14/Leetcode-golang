// https://leetcode.com/problems/binary-tree-level-order-traversal/
// Level: Medium

package leetcode

import (
	structures "leetcode/go/structures"
)

// TreeNode is to define the struct
type TreeNode = structures.TreeNode

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	curNum, nextLevelNum, tempdata := 1, 0, []int{}
	for len(queue) != 0 {
		node := queue[0]
		if curNum > 0 {
			if node.Left != nil {
				nextLevelNum++
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				nextLevelNum++
				queue = append(queue, node.Right)
			}
			curNum--
			tempdata = append(tempdata, node.Val)
			queue = queue[1:]
		}

		if curNum == 0 {
			curNum = nextLevelNum
			ans = append(ans, tempdata)
			tempdata, nextLevelNum = []int{}, 0
		}
	}
	return ans
}
