package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]int
}

type question1895Test_Problem1895 struct {
	params
	ans int
}

func Test_Problem1895(t *testing.T) {
	qs := []question1895Test_Problem1895{
		{
			params{grid: [][]int{{7, 1, 4, 5, 6}, {2, 5, 1, 6, 4}, {1, 5, 4, 3, 2}, {1, 2, 7, 3, 4}}},
			3,
		},
		{
			params{grid: [][]int{{5, 1, 3, 1}, {9, 3, 3, 1}, {1, 3, 3, 8}}},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, largestMagicSquare(p.grid), ans)
	}
}
