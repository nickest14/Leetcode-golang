package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n string
}

type question1689 struct {
	params
	ans int
}

func Test_Problem1689(t *testing.T) {
	qs := []question1689{
		{
			params{"32"},
			3,
		},
		{
			params{"82734"},
			8,
		},
		{
			params{"27346209830709182346"},
			9,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minPartitions(p.n), ans)
	}
}
