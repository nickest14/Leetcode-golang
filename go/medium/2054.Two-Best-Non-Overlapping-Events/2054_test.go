package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	events [][]int
}

type question2054 struct {
	params
	ans int
}

func Test_Problem2054(t *testing.T) {
	qs := []question2054{
		{
			params{[][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}},
			4,
		},
		{
			params{[][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}}},
			5,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxTwoEvents(p.events), ans)
	}
}
