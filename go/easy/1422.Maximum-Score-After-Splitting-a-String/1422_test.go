package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question1422 struct {
	params
	ans int
}

func Test_Problem1422(t *testing.T) {
	qs := []question1422{
		{
			params{s: "011101"},
			5,
		},
		{
			params{s: "1111"},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxScore(p.s), ans)
	}
}
