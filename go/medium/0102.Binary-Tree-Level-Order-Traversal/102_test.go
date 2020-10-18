package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	one []int
}

type question102 struct {
	params
	ans [][]int
}

func Test_Problem100(t *testing.T) {
	qs := []question102{
		// {
		// 	params{[]int{}},
		// 	[][]int{},
		// },
		// {
		// 	params{[]int{1}},
		// 	[][]int{{1}},
		// },
		{
			params{[]int{1, structures.NULL, 2}},
			[][]int{{1}, {2}},
		},
		{
			params{[]int{1, 2, 2, structures.NULL, 3, 3, structures.NULL}},
			[][]int{{1}, {2, 2}, {3, 3}},
		},
		{
			params{[]int{1, 3, 3, structures.NULL, 2}},
			[][]int{{1}, {3, 3}, {2}},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		root := structures.IntsToTreeNode(p.one)
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, levelOrder(root), ans)
	}
}
