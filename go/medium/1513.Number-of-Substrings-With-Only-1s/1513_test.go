package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question1513 struct {
	params
	ans int
}

func Test_Problem1513(t *testing.T) {
	qs := []question1513{
		{
			params{"0110111"},
			9,
		},
		{
			params{"101"},
			2,
		},
		{
			params{"111111"},
			21,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numSub(p.s), ans)
	}
}
