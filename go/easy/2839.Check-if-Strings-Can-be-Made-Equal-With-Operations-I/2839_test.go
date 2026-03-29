package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s1 string
	s2 string
}

type question2839 struct {
	params
	ans bool
}

func Test_Problem2839(t *testing.T) {
	qs := []question2839{
		{
			params{s1: "abcd", s2: "cdab"},
			true,
		},
		{
			params{s1: "abcd", s2: "dacb"},
			false,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, canBeEqual(p.s1, p.s2), ans)
	}
}
