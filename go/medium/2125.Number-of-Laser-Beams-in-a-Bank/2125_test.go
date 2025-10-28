package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	bank []string
}

type question2125 struct {
	params
	ans int
}

func Test_Problem2125(t *testing.T) {
	qs := []question2125{
		{
			params{[]string{"011001", "000000", "010100", "001000"}},
			8,
		},
		{
			params{[]string{"000", "111", "000"}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numberOfBeams(p.bank), ans)
	}
}
