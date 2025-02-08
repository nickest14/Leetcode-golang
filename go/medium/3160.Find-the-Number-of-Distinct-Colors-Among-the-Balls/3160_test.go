package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	limit   int
	queries [][]int
}

type question3160 struct {
	params
	ans []int
}

func Test_Problem3160(t *testing.T) {
	qs := []question3160{
		{
			params{limit: 4, queries: [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}}},
			[]int{1, 2, 2, 3},
		},
		{
			params{limit: 4, queries: [][]int{{0, 1}, {1, 2}, {2, 2}, {3, 4}, {4, 5}}},
			[]int{1, 2, 2, 3, 4},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, queryResults(p.limit, p.queries), ans)
	}
}
