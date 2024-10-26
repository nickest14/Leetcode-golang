package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	values  []int
	queries []int
}

type question2458 struct {
	params
	ans []int
}

func Test_Problem2458(t *testing.T) {
	qs := []question2458{
		{
			params{values: []int{5, 8, 9, 2, 1, 3, 7, 4, 6}, queries: []int{3, 2, 4, 8}},
			[]int{3, 2, 3, 2},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		root := structures.IntsToTreeNode(p.values)
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, treeQueries(root, p.queries), ans)
	}
}
