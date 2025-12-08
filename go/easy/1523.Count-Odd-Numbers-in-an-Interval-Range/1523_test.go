package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	low  int
	high int
}

type question1523 struct {
	params
	ans int
}

func Test_Problem1523(t *testing.T) {
	qs := []question1523{
		{
			params{low: 3, high: 7},
			3,
		},
		{
			params{low: 8, high: 10},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countOdds(p.low, p.high), ans)
	}
}
