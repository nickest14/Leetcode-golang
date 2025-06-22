package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s    string
	k    int
	fill byte
}

type question2138 struct {
	params
	ans []string
}

func Test_Problem2138(t *testing.T) {
	qs := []question2138{
		{
			params{s: "abcdefghi", k: 3, fill: 'x'},
			[]string{"abc", "def", "ghi"},
		},
		{
			params{s: "abcdefghij", k: 3, fill: 'x'},
			[]string{"abc", "def", "ghi", "jxx"},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, divideString(p.s, p.k, p.fill), ans)
	}
}
