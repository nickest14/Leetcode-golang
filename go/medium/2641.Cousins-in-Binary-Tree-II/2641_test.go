package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	values []int
}

type question2583 struct {
	params
	ans []int
}

func Test_Problem2583(t *testing.T) {
	qs := []question2583{
		{
			params{values: []int{5, 4, 9, 1, 10, structures.NULL, 7}},
			[]int{0, 0, 0, 7, 7, structures.NULL, 11},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		root := structures.IntsToTreeNode(p.values)
		result := structures.TreeToints(replaceValueInTree(root))
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, result, ans)
	}
}
