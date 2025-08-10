package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question231 struct {
	params
	ans bool
}

func Test_Problem231(t *testing.T) {
	qs := []question231{
		{
			params{1},
			true,
		},
		{
			params{16},
			true,
		},
		{
			params{3},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, isPowerOfTwo(p.n), ans)
	}
}
