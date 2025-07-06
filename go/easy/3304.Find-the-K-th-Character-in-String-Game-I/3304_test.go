package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	k int
}

type question3304 struct {
	params
	ans string
}

func Test_Problem3304(t *testing.T) {
	qs := []question3304{
		{
			params{5},
			"b",
		},
		{
			params{10},
			"c",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, kthCharacter(p.k), ans)
	}
}
