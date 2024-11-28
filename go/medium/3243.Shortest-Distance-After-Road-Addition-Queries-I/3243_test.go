package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n       int
	queries [][]int
}

type question3243 struct {
	params
	ans []int
}

func Test_Problem3243(t *testing.T) {
	qs := []question3243{
		{
			params{n: 5, queries: [][]int{{2, 4}, {0, 2}, {0, 4}}},
			[]int{3, 2, 1},
		},
		{
			params{n: 4, queries: [][]int{{0, 3}, {0, 2}}},
			[]int{1, 1},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, shortestDistanceAfterQueries(p.n, p.queries), ans)
	}
}
