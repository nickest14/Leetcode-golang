package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	graph [][]int
}

type question802 struct {
	params
	ans []int
}

func Test_Problem802(t *testing.T) {
	qs := []question802{
		{
			params{[][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}},
			[]int{2, 4, 5, 6},
		},
		{
			params{[][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}},
			[]int{4},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, eventualSafeNodes(p.graph), ans)
	}
}
