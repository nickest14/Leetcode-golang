package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	rating []int
}

type question1395 struct {
	params
	ans int
}

func Test_Problem1395(t *testing.T) {
	qs := []question1395{
		{
			params{rating: []int{2, 5, 3, 4, 1}},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numTeams(p.rating), ans)
	}
}
