package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]int
}

type question1536 struct {
	params
	ans int
}

func Test_Problem1536(t *testing.T) {
	qs := []question1536{
		{
			params{[][]int{{0, 0, 1}, {1, 1, 0}, {1, 0, 0}}},
			3,
		},
		{
			params{[][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}}},
			-1,
		},
		{
			params{[][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 1}}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minSwaps(p.grid), ans)
	}
}
