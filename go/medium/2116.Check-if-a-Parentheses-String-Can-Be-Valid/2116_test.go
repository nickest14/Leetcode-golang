package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s      string
	locked string
}

type question2116 struct {
	params
	ans bool
}

func Test_Problem2116(t *testing.T) {
	qs := []question2116{
		{
			params{s: "))()))", locked: "010100"},
			true,
		},
		{
			params{s: "()()", locked: "0000"},
			true,
		},
		{
			params{s: ")", locked: "0"},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, canBeValid(p.s, p.locked), ans)
	}
}
