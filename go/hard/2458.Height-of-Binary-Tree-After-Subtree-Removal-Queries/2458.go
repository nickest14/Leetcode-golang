// https://leetcode.com/problems/height-of-binary-tree-after-subtree-removal-queries/
// Level: Hard

package leetcode

import (
	structures "leetcode/go/structures"
	"sort"
)

// TreeNode is to define the struct
type TreeNode = structures.TreeNode

func treeQueries(root *TreeNode, queries []int) []int {
	depth := map[int]int{}
	height := map[int]int{}

	var dfs func(node *TreeNode, curDepth int) int
	dfs = func(node *TreeNode, curDepth int) int {
		if node == nil {
			return -1
		}
		depth[node.Val] = curDepth
		h := max(dfs(node.Left, curDepth+1), dfs(node.Right, curDepth+1)) + 1
		height[node.Val] = h
		return h
	}
	dfs(root, 0)

	cousins := map[int][][2]int{}
	for val, dep := range depth {
		cousins[dep] = append(cousins[dep], [2]int{-height[val], val})
		sort.Slice(cousins[dep], func(i, j int) bool {
			return cousins[dep][i][0] < cousins[dep][j][0]
		})
		if len(cousins[dep]) > 2 {
			cousins[dep] = cousins[dep][:2]
		}
	}

	ans := make([]int, len(queries))
	for i, query := range queries {
		dep := depth[query]
		if len(cousins[dep]) == 1 {
			ans[i] = dep - 1
		} else {
			ind := 0
			if cousins[dep][0][1] == query {
				ind = 1
			}
			ans[i] = -cousins[dep][ind][0] + dep
		}
	}
	return ans
}
