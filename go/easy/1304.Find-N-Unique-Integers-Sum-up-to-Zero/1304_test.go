package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question1323 struct {
	params
	ans []int
}

func Test_Problem1323(t *testing.T) {
	qs := []question1323{
		{
			params{5},
			[]int{-10, 1, 2, 3, 4},
		},
		{
			params{3},
			[]int{-3, 1, 2},
		},
		{
			params{1},
			[]int{0},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, sumZero(p.n), ans)
	}
}
