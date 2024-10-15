package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question2938 struct {
	params
	ans int
}

func Test_Problem2938(t *testing.T) {
	qs := []question2938{
		{
			params{"101"},
			1,
		},
		{
			params{"0011010"},
			5,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumSteps(p.s), ans)
	}
}
