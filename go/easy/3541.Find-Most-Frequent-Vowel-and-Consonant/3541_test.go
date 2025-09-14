package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question3541 struct {
	params
	ans int
}

func Test_Problem3541(t *testing.T) {
	qs := []question3541{
		{
			params{"successes"},
			6,
		},
		{
			params{"aeiaeia"},
			3,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxFreqSum(p.s), ans)
	}
}
