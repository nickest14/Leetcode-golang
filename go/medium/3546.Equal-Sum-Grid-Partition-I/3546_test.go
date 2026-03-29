package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]int
}

type question3546 struct {
	params
	ans bool
}

func Test_Problem3546(t *testing.T) {
	qs := []question3546{
		{
			params{grid: [][]int{{1, 4}, {2, 3}}},
			true,
		},
		{
			params{grid: [][]int{{1, 3}, {2, 4}}},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, canPartitionGrid(p.grid), ans)
	}
}
