package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question2434 struct {
	params
	ans string
}

func Test_Problem2434(t *testing.T) {
	qs := []question2434{
		{
			params{"zza"},
			"aaz",
		},
		{
			params{"bac"},
			"abc",
		},
		{
			params{"bdda"},
			"addb",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, robotWithString(p.s), ans)
	}
}
