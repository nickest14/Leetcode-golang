package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]int
}

type question1594 struct {
	params
	ans int
}

func Test_Problem1594(t *testing.T) {
	qs := []question1594{
		{
			params{grid: [][]int{{-1, -2, -3}, {-2, -3, -3}, {-3, -3, -2}}},
			-1,
		},
		{
			params{grid: [][]int{{1, -2, 1}, {1, -2, 1}, {3, -4, 1}}},
			8,
		},
		{
			params{grid: [][]int{{1, 3}, {0, -4}}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxProductPath(p.grid), ans)
	}
}
