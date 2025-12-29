package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]int
}

type question1351 struct {
	params
	ans int
}

func Test_Problem1351(t *testing.T) {
	qs := []question1351{
		{
			params{grid: [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}},
			8,
		},
		{
			params{grid: [][]int{{3, 2}, {1, 0}}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countNegatives(p.grid), ans)
	}
}
