package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	values []int
}

type question2471 struct {
	params
	ans int
}

func Test_Problem2471(t *testing.T) {
	qs := []question2471{
		{
			params{values: []int{1, 4, 3, 7, 6, 8, 5, structures.NULL, structures.NULL, structures.NULL, structures.NULL, 9, structures.NULL, 10}},
			3,
		},
		{
			params{values: []int{1, 3, 2, 7, 6, 5, 4}},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		root := structures.IntsToTreeNode(p.values)
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumOperations(root), ans)
	}
}
