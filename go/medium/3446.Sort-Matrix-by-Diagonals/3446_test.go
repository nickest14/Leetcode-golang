package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]int
}

type question3446 struct {
	params
	ans [][]int
}

func Test_Problem3446(t *testing.T) {
	qs := []question3446{
		{
			params{[][]int{{1, 7, 3}, {9, 8, 2}, {4, 5, 6}}},
			[][]int{{8, 2, 3}, {9, 6, 7}, {4, 5, 1}},
		},
		{
			params{[][]int{{0, 1}, {1, 2}}},
			[][]int{{2, 1}, {1, 0}},
		},
		{
			params{[][]int{{1}}},
			[][]int{{1}},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, sortMatrix(p.grid), ans)
	}
}
