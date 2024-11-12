package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	items   [][]int
	queries []int
}

type question2070 struct {
	params
	ans []int
}

func Test_Problem2070(t *testing.T) {
	qs := []question2070{
		{
			params{items: [][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}, queries: []int{1, 2, 3, 4, 5, 6}},
			[]int{2, 4, 5, 5, 6, 6},
		},
		{
			params{items: [][]int{{1, 2}, {1, 2}, {1, 3}, {1, 4}}, queries: []int{1}},
			[]int{4},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximumBeauty(p.items, p.queries), ans)
	}
}
