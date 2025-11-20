package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	intervals [][]int
}

type question757 struct {
	params
	ans int
}

func Test_Problem757(t *testing.T) {
	qs := []question757{
		{
			params{intervals: [][]int{{1, 3}, {3, 7}, {8, 9}}},
			5,
		},
		{
			params{intervals: [][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}}},
			3,
		},
		{
			params{intervals: [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}}},
			5,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, intersectionSizeTwo(p.intervals), ans)
	}
}
