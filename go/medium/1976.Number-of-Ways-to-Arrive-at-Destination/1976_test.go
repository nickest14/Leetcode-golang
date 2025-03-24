package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n     int
	roads [][]int
}

type question1976 struct {
	params
	ans int
}

func Test_Problem1976(t *testing.T) {
	qs := []question1976{
		{
			params{n: 7, roads: [][]int{{0, 6, 7}, {0, 1, 2}, {1, 2, 3}, {1, 3, 3}, {6, 3, 3}, {3, 5, 1}, {6, 5, 1}, {2, 5, 1}, {0, 4, 5}, {4, 6, 2}}},
			4,
		},
		{
			params{n: 2, roads: [][]int{{1, 0, 10}}},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countPaths(p.n, p.roads), ans)
	}
}
