package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]int
}

type question2290 struct {
	params
	ans int
}

func Test_Problem2290(t *testing.T) {
	qs := []question2290{
		{
			params{grid: [][]int{{0, 1, 1}, {1, 1, 0}, {1, 1, 0}}},
			2,
		},
		{
			params{grid: [][]int{{0, 1, 0, 0, 0}, {0, 1, 0, 1, 0}, {0, 0, 0, 1, 0}}},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumObstacles(p.grid), ans)
	}
}
