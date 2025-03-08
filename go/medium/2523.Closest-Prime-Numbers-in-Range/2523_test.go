package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	left  int
	right int
}

type question2523 struct {
	params
	ans []int
}

func Test_Problem2523(t *testing.T) {
	qs := []question2523{
		{
			params{left: 10, right: 19},
			[]int{11, 13},
		},
		{
			params{left: 4, right: 6},
			[]int{-1, -1},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, closestPrimes(p.left, p.right), ans)
	}
}
