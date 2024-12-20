package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	one []int
}

type question2415 struct {
	params
	ans TreeNode
}

func Test_Problem2415(t *testing.T) {
	qs := []question2415{
		{
			params{[]int{2, 3, 5, 8, 13, 21, 34}},
			*structures.IntsToTreeNode([]int{2, 5, 3, 8, 13, 21, 34}),
		},
		{
			params{[]int{0, 1, 2, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2}},
			*structures.IntsToTreeNode([]int{0, 2, 1, 0, 0, 0, 0, 2, 2, 2, 2, 1, 1, 1, 1}),
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		root := structures.IntsToTreeNode(p.one)

		fmt.Println("[input]:")
		fmt.Println(p)

		fmt.Println("[output]:")
		output := reverseOddLevels(root)
		fmt.Println(structures.TreeToints(output))

		fmt.Println("[answer]:")
		fmt.Println(structures.TreeToints(&ans))
	}
}
