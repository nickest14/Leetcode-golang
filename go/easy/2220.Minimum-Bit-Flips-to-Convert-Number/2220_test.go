package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	start int
	goal  int
}

type question2220 struct {
	params
	ans int
}

func Test_Problem2220(t *testing.T) {
	qs := []question2220{
		{
			params{10, 7},
			3,
		},
		{
			params{3, 4},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minBitFlips(p.start, p.goal), ans)
	}
}
