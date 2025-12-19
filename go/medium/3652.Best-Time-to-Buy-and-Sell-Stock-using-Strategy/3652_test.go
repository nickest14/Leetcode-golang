package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	prices   []int
	strategy []int
	k        int
}

type question3652 struct {
	params
	ans int64
}

func Test_Problem3652(t *testing.T) {
	qs := []question3652{
		{
			params{prices: []int{4, 2, 8}, strategy: []int{-1, 0, 1}, k: 2},
			10,
		},
		{
			params{prices: []int{5, 4, 3}, strategy: []int{1, 1, 0}, k: 2},
			9,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxProfit(p.prices, p.strategy, p.k), ans)
	}
}
