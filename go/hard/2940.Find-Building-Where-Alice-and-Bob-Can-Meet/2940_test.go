package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	heights []int
	queries [][]int
}

type question2940 struct {
	params
	ans []int
}

func Test_Problem2940(t *testing.T) {
	qs := []question2940{
		{
			params{heights: []int{6, 4, 8, 5, 2, 7}, queries: [][]int{{0, 1}, {0, 3}, {2, 4}, {3, 4}, {2, 2}}},
			[]int{2, 5, -1, 5, 2},
		},
		{
			params{heights: []int{5, 3, 8, 2, 6, 1, 4, 6}, queries: [][]int{{0, 7}, {3, 5}, {5, 2}, {3, 0}, {1, 6}}},
			[]int{7, 6, -1, 4, 6},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, leftmostBuildingQueries(p.heights, p.queries), ans)
	}
}
