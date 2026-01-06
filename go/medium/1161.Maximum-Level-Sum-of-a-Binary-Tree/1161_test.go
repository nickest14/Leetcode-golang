package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	one []int
}

type question1161 struct {
	params
	ans int
}

func Test_Problem1161(t *testing.T) {
	qs := []question1161{
		{
			params{[]int{1, 7, 0, 7, -8, structures.NULL, structures.NULL}},
			2,
		},
		{
			params{[]int{989, structures.NULL, 10250, 98693, -89388, structures.NULL, structures.NULL, structures.NULL, -32127}},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		root := structures.IntsToTreeNode(p.one)
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxLevelSum(root), ans)
	}
}
