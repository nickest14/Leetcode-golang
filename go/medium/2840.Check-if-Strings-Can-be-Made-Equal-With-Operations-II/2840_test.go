package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s1 string
	s2 string
}

type question2840 struct {
	params
	ans bool
}

func Test_Problem2840(t *testing.T) {
	qs := []question2840{
		{
			params{s1: "abcdba", s2: "cabdab"},
			true,
		},
		{
			params{s1: "abe", s2: "bea"},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, checkStrings(p.s1, p.s2), ans)
	}
}
