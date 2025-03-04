package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question1780 struct {
	params
	ans bool
}

func Test_Problem1780(t *testing.T) {
	qs := []question1780{
		{
			params{12},
			true,
		},
		{
			params{46},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, checkPowersOfThree(p.n), ans)
	}
}
