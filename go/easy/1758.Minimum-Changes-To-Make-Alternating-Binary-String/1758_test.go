package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question1758 struct {
	params
	ans int
}

func Test_Problem1758(t *testing.T) {
	qs := []question1758{
		{
			params{"0100"},
			1,
		},
		{
			params{"10"},
			0,
		},
		{
			params{"1111"},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minOperations(p.s), ans)
	}
}
