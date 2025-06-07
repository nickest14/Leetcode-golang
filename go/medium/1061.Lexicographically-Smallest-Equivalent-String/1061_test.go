package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s1      string
	s2      string
	baseStr string
}

type question1061 struct {
	params
	ans string
}

func Test_Problem1061(t *testing.T) {
	qs := []question1061{
		{
			params{s1: "parker", s2: "morris", baseStr: "parser"},
			"makkek",
		},
		{
			params{s1: "hello", s2: "world", baseStr: "hold"},
			"hdld",
		},
		{
			params{s1: "leetcode", s2: "programs", baseStr: "sourcecode"},
			"aauaaaaada",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, smallestEquivalentString(p.s1, p.s2, p.baseStr), ans)
	}
}
