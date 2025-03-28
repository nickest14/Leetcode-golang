package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid    [][]int
	queries []int
}

type question2503 struct {
	params
	ans []int
}

func Test_Problem2503(t *testing.T) {
	qs := []question2503{
		{
			params{grid: [][]int{{1, 2, 3}, {2, 5, 7}, {3, 5, 1}}, queries: []int{5, 6, 2}},
			[]int{5, 8, 1},
		},
		{
			params{grid: [][]int{{5, 2, 1}, {1, 1, 2}}, queries: []int{3}},
			[]int{0},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxPoints(p.grid, p.queries), ans)
	}
}
