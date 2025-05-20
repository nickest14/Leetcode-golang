package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums    []int
	queries [][]int
}

type question3355 struct {
	params
	ans bool
}

func Test_Problem3355(t *testing.T) {
	qs := []question3355{
		{
			params{nums: []int{1, 0, 1}, queries: [][]int{{0, 2}}},
			true,
		},
		{
			params{nums: []int{4, 3, 2, 1}, queries: [][]int{{1, 3}, {0, 2}}},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, isZeroArray(p.nums, p.queries), ans)
	}
}
