package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
	k int
}

type question2516 struct {
	params
	ans int
}

func Test_Problem2516(t *testing.T) {
	qs := []question2516{
		{
			params{s: "aabaaaacaabc", k: 2},
			8,
		},
		{
			params{s: "a", k: 1},
			-1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, takeCharacters(p.s, p.k), ans)
	}
}
