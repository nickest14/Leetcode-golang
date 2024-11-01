package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question1957 struct {
	params
	ans string
}

func Test_Problem1957(t *testing.T) {
	qs := []question1957{
		{
			params{"leeetcode"},
			"leetcode",
		},
		{
			params{"ab"},
			"ab",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, makeFancyString(p.s), ans)
	}
}
