package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n          int
	rectangles [][]int
}

type question3394 struct {
	params
	ans bool
}

func Test_Problem3394(t *testing.T) {
	qs := []question3394{
		{
			params{n: 5, rectangles: [][]int{{1, 0, 5, 2}, {0, 2, 2, 4}, {3, 2, 5, 3}, {0, 4, 4, 5}}},
			true,
		},
		{
			params{n: 4, rectangles: [][]int{{0, 0, 1, 1}, {2, 0, 3, 4}, {0, 2, 2, 3}, {3, 0, 4, 3}}},
			true,
		},
		{
			params{n: 4, rectangles: [][]int{{0, 2, 2, 4}, {1, 0, 3, 2}, {2, 2, 3, 4}, {3, 0, 4, 2}, {3, 2, 4, 4}}},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, checkValidCuts(p.n, p.rectangles), ans)
	}
}
