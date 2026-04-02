package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	coins [][]int
}

type question3418 struct {
	params
	ans int
}

func Test_Problem3418(t *testing.T) {
	qs := []question3418{
		{
			params{coins: [][]int{{0, 1, -1}, {1, -2, 3}, {2, -3, 4}}},
			8,
		},
		{
			params{coins: [][]int{{10, 10, 10}, {10, 10, 10}}},
			40,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximumAmount(p.coins), ans)
	}
}
