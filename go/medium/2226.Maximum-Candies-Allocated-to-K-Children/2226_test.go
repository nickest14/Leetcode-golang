package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	candies []int
	k       int64
}

type question2226 struct {
	params
	ans int
}

func Test_Problem2226(t *testing.T) {
	qs := []question2226{
		{
			params{candies: []int{1}, k: 1},
			0,
		},
		{
			params{candies: []int{5, 8, 6}, k: 3},
			5,
		},
		{
			params{candies: []int{2, 5}, k: 11},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximumCandies(p.candies, p.k), ans)
	}
}
