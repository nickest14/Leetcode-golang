package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	edges1 [][]int
	edges2 [][]int
}

type question3373 struct {
	params
	ans []int
}

func Test_Problem3373(t *testing.T) {
	qs := []question3373{
		{
			params{edges1: [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}, edges2: [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 7}, {1, 4}, {4, 5}, {4, 6}}},
			[]int{8, 7, 7, 8, 8},
		},
		{
			params{edges1: [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}}, edges2: [][]int{{0, 1}, {1, 2}, {2, 3}}},
			[]int{3, 6, 6, 6, 6},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxTargetNodes(p.edges1, p.edges2), ans)
	}
}
