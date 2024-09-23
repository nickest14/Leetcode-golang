package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s          string
	dictionary []string
}

type question2707 struct {
	params
	ans int
}

func Test_Problem2707(t *testing.T) {
	qs := []question2707{
		{
			params{s: "leetscode", dictionary: []string{"leet", "code", "leetcode"}},
			1,
		},
		{
			params{s: "sayhelloworld", dictionary: []string{"hello", "world"}},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minExtraChar(p.s, p.dictionary), ans)
	}
}
