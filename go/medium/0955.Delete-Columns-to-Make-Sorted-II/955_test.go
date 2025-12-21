package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	strs []string
}

type question955 struct {
	params
	ans int
}

func Test_Problem955(t *testing.T) {

	qs := []question955{
		{
			params{[]string{"ca", "bb", "ac"}},
			1,
		},
		{
			params{[]string{"xc", "yb", "za"}},
			0,
		},
		{
			params{[]string{"zyx", "wvu", "tsr"}},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v, [answer]: %v\n", p, minDeletionSize(p.strs), ans)
	}
}
