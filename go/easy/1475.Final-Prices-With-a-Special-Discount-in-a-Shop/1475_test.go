package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	prices []int
}

type question1475 struct {
	params
	ans []int
}

func Test_Problem1475(t *testing.T) {
	qs := []question1475{
		{
			params{prices: []int{8, 4, 6, 2, 3}},
			[]int{4, 2, 4, 2, 3},
		},
		{
			params{prices: []int{1, 2, 3, 4, 5}},
			[]int{1, 2, 3, 4, 5},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, finalPrices(p.prices), ans)
	}
}
