package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n       int
	queries [][]int
}

type question2438 struct {
	params
	ans []int
}

func Test_Problem2438(t *testing.T) {
	qs := []question2438{
		{
			params{n: 15, queries: [][]int{{0, 1}, {2, 2}, {0, 3}}},
			[]int{2, 4, 64},
		},
		{
			params{n: 2, queries: [][]int{{0, 0}}},
			[]int{2},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, productQueries(p.n, p.queries), ans)
	}
}
