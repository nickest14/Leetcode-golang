package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question1593 struct {
	params
	ans int
}

func Test_Problem1593(t *testing.T) {
	qs := []question1593{
		{
			params{"ababccc"},
			5,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxUniqueSplit(p.s), ans)
	}
}
