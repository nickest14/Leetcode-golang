package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	one []int
}

type question1382 struct {
	params
	ans []int
}

func Test_Problem1382(t *testing.T) {
	qs := []question1382{
		{
			params{[]int{1, structures.NULL, 2, structures.NULL, 3, structures.NULL, 4, structures.NULL, structures.NULL}},
			[]int{2, 1, 3, structures.NULL, structures.NULL, structures.NULL, 4},
		},
		{
			params{[]int{2, 1, 3}},
			[]int{2, 1, 3},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		root := structures.IntsToTreeNode(p.one)
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, structures.TreeToints(balanceBST(root)), ans)
	}
}
