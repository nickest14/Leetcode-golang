package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]int
	k    int
}

type question3070 struct {
	params
	ans int
}

func Test_Problem3070(t *testing.T) {
	qs := []question3070{
		{
			params{grid: [][]int{{7, 6, 3}, {6, 6, 1}}, k: 18},
			4,
		},
		{
			params{grid: [][]int{{7, 2, 9}, {1, 5, 0}, {2, 6, 6}}, k: 20},
			6,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countSubmatrices(p.grid, p.k), ans)
	}
}
