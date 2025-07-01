package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	word string
}

type question3330 struct {
	params
	ans int
}

func Test_Problem3330(t *testing.T) {
	qs := []question3330{
		{
			params{"abbcccc"},
			5,
		},
		{
			params{"abcd"},
			1,
		},
		{
			params{"aaaa"},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, possibleStringCount(p.word), ans)
	}
}
