package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	one []int
}

type question101 struct {
	params
	ans bool
}

func Test_Problem101(t *testing.T) {
	qs := []question101{
		{
			params{[]int{}},
			true,
		},
		{
			params{[]int{1}},
			true,
		},
		{
			params{[]int{1, structures.NULL, 2}},
			false,
		},
		{
			params{[]int{1, 2, 2, structures.NULL, 3, 3, structures.NULL}},
			true,
		},
		{
			params{[]int{1, 3, 3, structures.NULL, 2}},
			false,
		},
		{
			params{[]int{1, 2, 3}},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		root := structures.IntsToTreeNode(p.one)
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, isSymmetric(root), ans)
	}
}
