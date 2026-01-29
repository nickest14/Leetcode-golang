package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n     int
	edges [][]int
}

type question3650 struct {
	params
	ans int
}

func Test_Problem3650(t *testing.T) {
	qs := []question3650{
		{
			params{n: 4, edges: [][]int{{0, 1, 3}, {3, 1, 1}, {2, 3, 4}, {0, 2, 2}}},
			5,
		},
		{
			params{n: 4, edges: [][]int{{0, 2, 1}, {2, 1, 1}, {1, 3, 1}, {2, 3, 3}}},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minCost(p.n, p.edges), ans)
	}
}
