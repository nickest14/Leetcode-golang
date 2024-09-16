package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question1371 struct {
	params
	ans int
}

func Test_Problem1371(t *testing.T) {
	qs := []question1371{
		{
			params{s: "eleetminicoworoep"},
			13,
		},
		{
			params{s: "bcbcbc"},
			6,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findTheLongestSubstring(p.s), ans)
	}
}
