package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n     int
	edges [][]int
}

type question2924 struct {
	params
	ans int
}

func Test_Problem2924(t *testing.T) {
	qs := []question2924{
		{
			params{n: 3, edges: [][]int{{0, 1}, {1, 2}}},
			0,
		},
		{
			params{n: 4, edges: [][]int{{0, 2}, {1, 3}, {1, 2}}},
			-1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findChampion(p.n, p.edges), ans)
	}
}
