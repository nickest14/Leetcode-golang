package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]int
}

type question3195 struct {
	params
	ans int
}

func Test_Problem3195(t *testing.T) {
	qs := []question3195{
		{
			params{[][]int{{0, 1, 0}, {1, 0, 1}}},
			6,
		},
		{
			params{[][]int{{1, 0}, {0, 0}}},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumArea(p.grid), ans)
	}
}
