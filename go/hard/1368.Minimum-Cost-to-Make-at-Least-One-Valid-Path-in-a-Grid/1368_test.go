package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]int
}

type question1368 struct {
	params
	ans int
}

func Test_Problem1368(t *testing.T) {
	qs := []question1368{
		{
			params{[][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2}}},
			3,
		},
		{
			params{[][]int{{1, 1, 3}, {3, 2, 2}, {1, 1, 4}}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minCost(p.grid), ans)
	}
}
