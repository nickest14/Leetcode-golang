package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question1963 struct {
	params
	ans int
}

func Test_Problem1963(t *testing.T) {
	qs := []question1963{
		{
			params{"][]["},
			1,
		},
		{
			params{"]]][[["},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minSwaps(p.s), ans)
	}
}
