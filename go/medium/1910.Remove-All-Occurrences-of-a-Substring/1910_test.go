package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s    string
	part string
}

type question1910 struct {
	params
	ans string
}

func Test_Problem1910(t *testing.T) {
	qs := []question1910{
		{
			params{s: "daabcbaabcbc", part: "abc"},
			"dab",
		},
		{
			params{s: "axxxxyyyyb", part: "xy"},
			"ab",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, removeOccurrences(p.s, p.part), ans)
	}
}
