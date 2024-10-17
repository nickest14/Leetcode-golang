package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	a int
	b int
	c int
}

type question1395 struct {
	params
	ans string
}

func Test_Problem1395(t *testing.T) {
	qs := []question1395{
		{
			params{a: 1, b: 1, c: 7},
			"ccaccbcc",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, longestDiverseString(p.a, p.b, p.c), ans)
	}
}
