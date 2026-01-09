package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	nodes []int
}

type question865 struct {
	params
	ans TreeNode
}

func Test_Problem865(t *testing.T) {
	qs := []question865{
		{
			params{[]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4}},
			*structures.IntsToTreeNode([]int{2, 7, 4}),
		},
		{
			params{[]int{1}},
			*structures.IntsToTreeNode([]int{1}),
		},
		{
			params{[]int{0, 1, 3, structures.NULL, 2}},
			*structures.IntsToTreeNode([]int{2}),
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		root := structures.IntsToTreeNode(p.nodes)

		fmt.Println("[input]:")
		fmt.Println(p)

		fmt.Println("[output]:")
		output := subtreeWithAllDeepest(root)
		fmt.Println(structures.TreeToints(output))

		fmt.Println("[answer]:")
		fmt.Println(structures.TreeToints(&ans))
	}
}
