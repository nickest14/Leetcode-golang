package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question38 struct {
	params
	ans string
}

func Test_Problem38(t *testing.T) {

	qs := []question38{
		{
			params{4},
			"1211",
		},
		{
			params{1},
			"1",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v, [answer]: %v\n", p, countAndSay(p.n), ans)
	}
}
