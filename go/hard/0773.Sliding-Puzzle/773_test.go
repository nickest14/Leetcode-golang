package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	board [][]int
}

type question773 struct {
	params
	ans int
}

func Test_Problem773(t *testing.T) {
	qs := []question773{
		{
			params{[][]int{{1, 2, 3}, {4, 0, 5}}},
			1,
		},
		{
			params{[][]int{{4, 1, 2}, {5, 0, 3}}},
			5,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, slidingPuzzle(p.board), ans)
	}
}
