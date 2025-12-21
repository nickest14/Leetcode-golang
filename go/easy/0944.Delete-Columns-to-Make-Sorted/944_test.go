package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	strs []string
}

type question1018 struct {
	params
	ans int
}

func Test_Problem1018(t *testing.T) {
	qs := []question1018{
		{
			params{[]string{"cba", "daf", "ghi"}},
			1,
		},
		{
			params{[]string{"a", "b"}},
			0,
		},
		{
			params{[]string{"zyx", "wvu", "tsr"}},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minDeletionSize(p.strs), ans)
	}
}
