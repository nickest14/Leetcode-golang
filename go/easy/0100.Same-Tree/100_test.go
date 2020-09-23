package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	one     []int
	another []int
}

type question100 struct {
	params
	ans bool
}

func Test_Problem100(t *testing.T) {
	qs := []question100{
		{
			params{[]int{}, []int{}},
			true,
		},
		{
			params{[]int{}, []int{1}},
			false,
		},
		{
			params{[]int{2, 1}, []int{2, 1}},
			false,
		},
		{
			params{[]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
			true,
		},
		{
			params{[]int{1, structures.NULL, 2}, []int{1, 2, 2}},
			false,
		},

		{
			params{[]int{3, 2, 1}, []int{1, 2, 3}},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		rootOne := structures.IntsToTreeNode(p.one)
		rootTwo := structures.IntsToTreeNode(p.another)
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, isSameTree(rootOne, rootTwo), ans)
	}
}
