package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question869 struct {
	params
	ans bool
}

func Test_Problem869(t *testing.T) {
	qs := []question869{
		{
			params{1},
			true,
		},
		{
			params{10},
			false,
		},
		{
			params{128},
			true,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, reorderedPowerOf2(p.n), ans)
	}
}
