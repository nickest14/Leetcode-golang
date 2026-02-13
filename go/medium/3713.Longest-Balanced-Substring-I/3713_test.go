package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question3713 struct {
	params
	ans int
}

func Test_Problem3713(t *testing.T) {
	qs := []question3713{
		{
			params{s: "abbac"},
			4,
		},
		{
			params{s: "zzabccy"},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, longestBalanced(p.s), ans)
	}
}
