package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	values []int
	k      int
}

type question2583 struct {
	params
	ans int64
}

func Test_Problem2583(t *testing.T) {
	qs := []question2583{
		{
			params{values: []int{5, 8, 9, 2, 1, 3, 7, 4, 6}, k: 2},
			13,
		},
		{
			params{values: []int{5, 8, 9, 2, 1, 3, 7}, k: 4},
			-1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		root := structures.IntsToTreeNode(p.values)
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, kthLargestLevelSum(root, p.k), ans)
	}
}
