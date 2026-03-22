package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]int
	x    int
	y    int
	k    int
}

type question3264 struct {
	params
	ans [][]int
}

func Test_Problem3264(t *testing.T) {
	qs := []question3264{
		{
			params{grid: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, x: 1, y: 0, k: 3},
			[][]int{{1, 2, 3, 4}, {13, 14, 15, 8}, {9, 10, 11, 12}, {5, 6, 7, 16}},
		},
		{
			params{grid: [][]int{{3, 4, 2, 3}, {2, 3, 4, 2}}, x: 0, y: 2, k: 2},
			[][]int{{3, 4, 4, 2}, {2, 3, 2, 3}},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, reverseSubmatrix(p.grid, p.x, p.y, p.k), ans)
	}
}
