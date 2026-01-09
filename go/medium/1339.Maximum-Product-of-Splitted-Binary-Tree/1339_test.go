package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	nodes []int
}

type question1339 struct {
	params
	ans int
}

func Test_Problem1339(t *testing.T) {
	qs := []question1339{
		{
			params{[]int{1, 2, 3, 4, 5, 6}},
			110,
		},
		{
			params{[]int{1, structures.NULL, 2, 3, 4, structures.NULL, structures.NULL, 5, 6}},
			90,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		root := structures.IntsToTreeNode(p.nodes)
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxProduct(root), ans)
	}
}
