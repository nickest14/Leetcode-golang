package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums    []int
	queries [][]int
}

type question3653 struct {
	params
	ans int
}

func Test_Problem3653(t *testing.T) {
	qs := []question3653{
		{
			params{nums: []int{1, 1, 1}, queries: [][]int{{0, 2, 1, 4}}},
			4,
		},
		{
			params{nums: []int{2, 3, 1, 5, 4}, queries: [][]int{{1, 4, 2, 3}, {0, 2, 1, 2}}},
			31,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, xorAfterQueries(p.nums, p.queries), ans)
	}
}
