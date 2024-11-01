package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]int
}

type question2684 struct {
	params
	ans int
}

func Test_Problem2684(t *testing.T) {
	qs := []question2684{
		{
			params{[][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxMoves(p.grid), ans)
	}
}
