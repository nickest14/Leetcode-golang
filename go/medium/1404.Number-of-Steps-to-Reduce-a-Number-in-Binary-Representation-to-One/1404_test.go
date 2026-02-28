package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question1404 struct {
	params
	ans int
}

func Test_Problem1404(t *testing.T) {
	qs := []question1404{
		{
			params{s: "1101"},
			6,
		},
		{
			params{s: "10"},
			1,
		},
		{
			params{s: "1"},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numSteps(p.s), ans)
	}
}
