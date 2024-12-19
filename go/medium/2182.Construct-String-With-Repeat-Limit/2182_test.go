package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s           string
	repeatLimit int
}

type question2182 struct {
	params
	ans string
}

func Test_Problem2182(t *testing.T) {
	qs := []question2182{
		{
			params{s: "cczazcc", repeatLimit: 3},
			"zzcccac",
		},
		{
			params{s: "aababab", repeatLimit: 2},
			"bbabaa",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, repeatLimitedString(p.s, p.repeatLimit), ans)
	}
}
