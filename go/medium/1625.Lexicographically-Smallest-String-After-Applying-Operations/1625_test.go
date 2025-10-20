package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
	a int
	b int
}

type question1625 struct {
	params
	ans string
}

func Test_Problem1625(t *testing.T) {
	qs := []question1625{
		{
			params{s: "5525", a: 9, b: 2},
			"2050",
		},
		{
			params{s: "74", a: 5, b: 1},
			"24",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findLexSmallestString(p.s, p.a, p.b), ans)
	}
}
