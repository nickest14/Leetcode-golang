package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n     int
	edges [][]int
}

type question2685 struct {
	params
	ans int
}

func Test_Problem2685(t *testing.T) {
	qs := []question2685{
		{
			params{n: 6, edges: [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}}},
			3,
		},
		{
			params{n: 6, edges: [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}}},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countCompleteComponents(p.n, p.edges), ans)
	}
}
