package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	num string
}

type question2264 struct {
	params
	ans string
}

func Test_Problem2264(t *testing.T) {
	qs := []question2264{
		{
			params{"6777133339"},
			"777",
		},
		{
			params{"2300019"},
			"000",
		},
		{
			params{"42352338"},
			"",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, largestGoodInteger(p.num), ans)
	}
}
