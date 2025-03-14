package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums    []int
	queries [][]int
}

type question3356 struct {
	params
	ans int
}

func Test_Problem3356(t *testing.T) {
	qs := []question3356{
		{
			params{nums: []int{2, 0, 2}, queries: [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}}},
			2,
		},
		{
			params{nums: []int{4, 3, 2, 1}, queries: [][]int{{1, 3, 2}, {0, 2, 1}}},
			-1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minZeroArray(p.nums, p.queries), ans)
	}
}
