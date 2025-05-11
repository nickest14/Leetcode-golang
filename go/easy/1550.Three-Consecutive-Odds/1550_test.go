package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	arr []int
}

type question1550 struct {
	params
	ans bool
}

func Test_Problem1550(t *testing.T) {
	qs := []question1550{
		{
			params{arr: []int{2, 6, 4, 1}},
			false,
		},
		{
			params{arr: []int{1, 2, 34, 3, 4, 5, 7, 23, 12}},
			true,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, threeConsecutiveOdds(p.arr), ans)
	}
}
