// https://leetcode.com/problems/delete-nodes-and-return-forest/
// Level: Medium

package leetcode

import (
	structures "leetcode/go/structures"
)

// TreeNode is to define the struct
type TreeNode = structures.TreeNode

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	toDeleteMap := make(map[int]bool)
	for _, val := range to_delete {
		toDeleteMap[val] = true
	}
	var remainNodes []*TreeNode

	var dfs func(*TreeNode, bool) *TreeNode
	dfs = func(node *TreeNode, isRoot bool) *TreeNode {
		if node == nil {
			return nil
		}

		deleted := toDeleteMap[node.Val]

		if isRoot && !deleted {
			remainNodes = append(remainNodes, node)
		}

		node.Left = dfs(node.Left, deleted)
		node.Right = dfs(node.Right, deleted)

		if deleted {
			return nil
		}
		return node
	}

	dfs(root, true)
	return remainNodes
}
