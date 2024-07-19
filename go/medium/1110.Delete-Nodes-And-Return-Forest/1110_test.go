package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	one      []int
	toDelete []int
}

type question1110 struct {
	params
	ans []TreeNode
}

func Test_Problem1110(t *testing.T) {
	qs := []question1110{
		{
			params{[]int{1, 2, 3, 4, 5, 6, 7}, []int{3, 5}},
			[]TreeNode{*structures.IntsToTreeNode([]int{1, 2, structures.NULL, 4}), *structures.IntsToTreeNode([]int{6}), *structures.IntsToTreeNode([]int{7})},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		root := structures.IntsToTreeNode(p.one)

		fmt.Println("[input]:")
		fmt.Println(p)

		fmt.Println("[output]:")
		for _, tree := range delNodes(root, p.toDelete) {
			fmt.Println(structures.TreeToints(tree))
		}
		fmt.Println("[answer]:")
		for _, tree := range ans {
			fmt.Println(structures.TreeToints(&tree))
		}
	}
}
