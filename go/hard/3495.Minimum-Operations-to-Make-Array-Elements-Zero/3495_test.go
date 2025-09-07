package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	queries [][]int
}

type question3495 struct {
	params
	ans int64
}

func Test_Problem3495(t *testing.T) {
	qs := []question3495{
		{
			params{queries: [][]int{{1, 2}, {2, 4}}},
			3,
		},
		{
			params{queries: [][]int{{2, 6}}},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minOperations(p.queries), ans)
	}
}
