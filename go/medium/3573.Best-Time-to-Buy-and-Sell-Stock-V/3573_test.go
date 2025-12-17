package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	maximumProfit []int
	k             int
}

type question3573 struct {
	params
	ans int
}

func Test_Problem3573(t *testing.T) {
	qs := []question3573{
		{
			params{maximumProfit: []int{1, 7, 9, 8, 2}, k: 2},
			14,
		},
		{
			params{maximumProfit: []int{12, 16, 19, 19, 8, 1, 19, 13, 9}, k: 3},
			36,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximumProfit(p.maximumProfit, p.k), ans)
	}
}
