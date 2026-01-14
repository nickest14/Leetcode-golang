package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	squares [][]int
}

type question3453 struct {
	params
	ans float64
}

func Test_Problem3453(t *testing.T) {
	qs := []question3453{
		{
			params{[][]int{{0, 0, 1}, {2, 2, 1}}},
			1.00000,
		},
		{
			params{[][]int{{0, 0, 2}, {1, 1, 1}}},
			1.16667,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, separateSquares(p.squares), ans)
	}
}
