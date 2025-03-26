package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]int
	x    int
}

type question2033 struct {
	params
	ans int
}

func Test_Problem2033(t *testing.T) {
	qs := []question2033{
		{
			params{grid: [][]int{{2, 4}, {6, 8}}, x: 2},
			4,
		},
		{
			params{grid: [][]int{{1, 5}, {2, 3}}, x: 1},
			5,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minOperations(p.grid, p.x), ans)
	}
}
