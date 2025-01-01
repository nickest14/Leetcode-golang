package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	low  int
	high int
	zero int
	one  int
}

type question2466 struct {
	params
	ans int
}

func Test_Problem2466(t *testing.T) {
	qs := []question2466{
		{
			params{low: 3, high: 3, zero: 1, one: 1},
			8,
		},
		{
			params{low: 2, high: 3, zero: 1, one: 2},
			5,
		},
		{
			params{low: 200, high: 200, zero: 10, one: 1},
			764262396,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countGoodStrings(p.low, p.high, p.zero, p.one), ans)
	}
}
