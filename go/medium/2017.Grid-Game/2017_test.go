package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]int
}

type question2017 struct {
	params
	ans int64
}

func Test_Problem2017(t *testing.T) {
	qs := []question2017{
		{
			params{[][]int{{2, 5, 4}, {1, 5, 1}}},
			4,
		},
		{
			params{[][]int{{3, 3, 1}, {8, 5, 2}}},
			4,
		},
		{
			params{[][]int{{1, 3, 1, 15}, {1, 3, 3, 1}}},
			7,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, gridGame(p.grid), ans)
	}
}
