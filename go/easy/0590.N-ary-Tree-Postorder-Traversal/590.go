// https://leetcode.com/problems/n-ary-tree-postorder-traversal/
// Level: Easy

package leetcode

import (
	"leetcode/go/structures"
)

type Node = structures.Node

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	stack := []*Node{root}
	ans := []int{}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			ans = append(ans, node.Val)
			stack = append(stack, node.Children...)
		}
	}

	// Reverse the result for postorder
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}
