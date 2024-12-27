package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	values []int
}

type question102 struct {
	params
	ans []int
}

func Test_Problem102(t *testing.T) {
	qs := []question102{
		{
			params{[]int{1, 3, 2, 5, 3, structures.NULL, 9}},
			[]int{1, 3, 9},
		},
		{
			params{[]int{1, 2, 3}},
			[]int{1, 3},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		root := structures.IntsToTreeNode(p.values)
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, largestValues(root), ans)
	}
}
