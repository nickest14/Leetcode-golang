package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	questions [][]int
}

type question2140 struct {
	params
	ans int64
}

func Test_Problem2140(t *testing.T) {
	qs := []question2140{
		{
			params{[][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}},
			5,
		},
		{
			params{[][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}},
			7,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, mostPoints(p.questions), ans)
	}
}
