package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question912 struct {
	params
	ans int
}

func Test_Problem912(t *testing.T) {
	qs := []question912{
		{
			params{"()))"},
			2,
		},
		{
			params{"))()"},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minAddToMakeValid(p.s), ans)
	}
}
