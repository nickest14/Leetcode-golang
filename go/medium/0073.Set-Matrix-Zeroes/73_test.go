package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	matrix [][]int
}

type question73 struct {
	params
	ans [][]int
}

func Test_Problem73(t *testing.T) {
	qs := []question73{
		{
			params{matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}},
			[][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			params{matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}},
			[][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		setZeroes(p.matrix)
		fmt.Printf("[output]: %v    [answer]: %v\n", p.matrix, ans)
	}
}
