package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question342 struct {
	params
	ans bool
}

func Test_Problem342(t *testing.T) {
	qs := []question342{
		{
			params{16},
			true,
		},
		{
			params{5},
			false,
		},
		{
			params{1},
			true,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, isPowerOfFour(p.n), ans)
	}
}
