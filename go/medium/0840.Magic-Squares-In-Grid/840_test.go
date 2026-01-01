package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]int
}

type question840 struct {
	params
	ans int
}

func Test_Problem840(t *testing.T) {
	qs := []question840{
		{
			params{[][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}},
			1,
		},
		{
			params{[][]int{{8}}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numMagicSquaresInside(p.grid), ans)
	}
}
