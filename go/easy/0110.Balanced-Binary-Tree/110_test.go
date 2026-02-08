package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	one []int
}

type question110 struct {
	params
	ans bool
}

func Test_Problem110(t *testing.T) {
	qs := []question110{
		{
			params{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			true,
		},
		{
			params{[]int{1, 2, 2, 3, 3, structures.NULL, structures.NULL, 4, 4}},
			false,
		},
		{
			params{[]int{}},
			true,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		root := structures.IntsToTreeNode(p.one)
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, isBalanced(root), ans)
	}
}
