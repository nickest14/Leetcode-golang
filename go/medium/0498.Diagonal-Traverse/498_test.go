package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	mat [][]int
}

type question498 struct {
	params
	ans []int
}

func Test_Problem498(t *testing.T) {

	qs := []question498{
		{
			params{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			[]int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			params{[][]int{{1, 2}, {3, 4}}},
			[]int{1, 2, 3, 4},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v, [answer]: %v\n", p, findDiagonalOrder(p.mat), ans)
	}
}
