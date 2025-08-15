package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question326 struct {
	params
	ans bool
}

func Test_Problem326(t *testing.T) {
	qs := []question326{
		{
			params{27},
			true,
		},
		{
			params{0},
			false,
		},
		{
			params{-1},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, isPowerOfThree(p.n), ans)
	}
}
