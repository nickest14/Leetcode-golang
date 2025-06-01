package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	board [][]int
}

type question909 struct {
	params
	ans int
}

func Test_Problem909(t *testing.T) {
	qs := []question909{
		{
			params{[][]int{{-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 35, -1, -1, 13, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 15, -1, -1, -1, -1}}},
			4,
		},
		{
			params{[][]int{{-1, -1}, {-1, 3}}},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, snakesAndLadders(p.board), ans)
	}
}
