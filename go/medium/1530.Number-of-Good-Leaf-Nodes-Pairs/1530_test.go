package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	nodes    []int
	distance int
}

type question1530 struct {
	params
	ans int
}

func Test_Problem1530(t *testing.T) {
	qs := []question1530{
		{
			params{[]int{1, 2, 3, structures.NULL, 4}, 3},
			1,
		},
		{
			params{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		root := structures.IntsToTreeNode(p.nodes)
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countPairs(root, p.distance), ans)
	}
}
