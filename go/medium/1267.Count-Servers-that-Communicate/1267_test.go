package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]int
}

type question1267 struct {
	params
	ans int
}

func Test_Problem1267(t *testing.T) {
	qs := []question1267{
		{
			params{[][]int{{1, 0}, {0, 1}}},
			0,
		},
		{
			params{[][]int{{1, 0}, {1, 1}}},
			3,
		},
		{
			params{[][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countServers(p.grid), ans)
	}
}
