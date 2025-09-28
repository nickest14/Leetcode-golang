package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	triangle [][]int
}

type question120 struct {
	params
	ans int
}

func Test_Problem120(t *testing.T) {
	qs := []question120{
		{
			params{[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}},
			11,
		},
		{
			params{[][]int{{-10}}},
			-10,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumTotal(p.triangle), ans)
	}
}
