package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question2914 struct {
	params
	ans int
}

func Test_Problem2914(t *testing.T) {
	qs := []question2914{
		{
			params{"1001"},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minChanges(p.s), ans)
	}
}
