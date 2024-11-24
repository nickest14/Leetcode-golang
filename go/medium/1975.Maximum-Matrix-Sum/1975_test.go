package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	matrix [][]int
}

type question1975 struct {
	params
	ans int64
}

func Test_Problem1975(t *testing.T) {
	qs := []question1975{
		{
			params{[][]int{{1, -1}, {-1, 1}}},
			4,
		},
		{
			params{[][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}}},
			16,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxMatrixSum(p.matrix), ans)
	}
}
