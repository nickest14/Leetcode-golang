package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
	k int
}

type question2311 struct {
	params
	ans int
}

func Test_Problem2311(t *testing.T) {
	qs := []question2311{
		{
			params{s: "1001010", k: 5},
			5,
		},
		{
			params{s: "00101001", k: 1},
			6,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, longestSubsequence(p.s, p.k), ans)
	}
}
