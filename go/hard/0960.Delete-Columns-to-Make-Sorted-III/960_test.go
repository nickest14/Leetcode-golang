package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	strs []string
}

type question960 struct {
	params
	ans int
}

func Test_Problem960(t *testing.T) {
	qs := []question960{
		{
			params{strs: []string{"babca", "bbazb"}},
			3,
		},
		{
			params{strs: []string{"edcba"}},
			4,
		},
		{
			params{strs: []string{"ghi", "def", "abc"}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minDeletionSize(p.strs), ans)
	}
}
