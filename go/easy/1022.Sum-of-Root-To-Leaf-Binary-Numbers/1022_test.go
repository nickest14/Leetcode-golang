package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	one []int
}

type question100 struct {
	params
	ans int
}

func Test_Problem100(t *testing.T) {
	qs := []question100{
		{
			params{[]int{1, 0, 1, 0, 1, 0, 1}},
			22,
		},
		{
			params{[]int{0}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		root := structures.IntsToTreeNode(p.one)
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, sumRootToLeaf(root), ans)
	}
}
