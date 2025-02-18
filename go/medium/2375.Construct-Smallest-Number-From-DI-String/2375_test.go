package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	pattern string
}

type question2375 struct {
	params
	ans string
}

func Test_Problem2375(t *testing.T) {
	qs := []question2375{
		{
			params{"IIIDIDDD"},
			"123549876",
		},
		{
			params{"DDD"},
			"4321",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, smallestNumber(p.pattern), ans)
	}
}
