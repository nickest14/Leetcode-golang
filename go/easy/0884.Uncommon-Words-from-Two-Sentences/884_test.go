package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s1 string
	s2 string
}

type question884 struct {
	params
	ans []string
}

func Test_Problem884(t *testing.T) {
	qs := []question884{
		{
			params{s1: "this apple is sweet", s2: "this apple is sour"},
			[]string{"sweet", "sour"},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, uncommonFromSentences(p.s1, p.s2), ans)
	}
}
