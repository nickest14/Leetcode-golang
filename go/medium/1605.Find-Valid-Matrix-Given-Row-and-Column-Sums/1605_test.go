package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	rowSum []int
	colSum []int
}

type question1605 struct {
	params
	ans [][]int
}

func Test_Problem1605(t *testing.T) {
	qs := []question1605{
		{
			params{rowSum: []int{3, 8}, colSum: []int{4, 7}},
			[][]int{{3, 0}, {1, 7}},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, restoreMatrix(p.rowSum, p.colSum), ans)
	}
}
