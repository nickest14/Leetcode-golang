package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question1930 struct {
	params
	ans int
}

func Test_Problem1930(t *testing.T) {
	qs := []question1930{
		{
			params{"aabca"},
			3,
		},
		{
			params{"bbcbaba"},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countPalindromicSubsequence(p.s), ans)
	}
}
