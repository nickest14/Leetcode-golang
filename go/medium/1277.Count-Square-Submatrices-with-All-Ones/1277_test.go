package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	matrix [][]int
}

type question1277 struct {
	params
	ans int
}

func Test_Problem1277(t *testing.T) {
	qs := []question1277{
		// {
		// 	params{[][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}},
		// 	15,
		// },
		{
			params{[][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}},
			7,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countSquares(p.matrix), ans)
	}
}
