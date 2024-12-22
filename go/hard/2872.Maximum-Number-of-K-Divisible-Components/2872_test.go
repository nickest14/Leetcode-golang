package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n      int
	edges  [][]int
	values []int
	k      int
}

type question2872 struct {
	params
	ans int
}

func Test_Problem2872(t *testing.T) {
	qs := []question2872{
		{
			params{n: 5, edges: [][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}}, values: []int{1, 8, 1, 4, 4}, k: 6},
			2,
		},
		{
			params{n: 7, edges: [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}}, values: []int{3, 0, 6, 1, 5, 2, 1}, k: 3},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxKDivisibleComponents(p.n, p.edges, p.values, p.k), ans)
	}
}
