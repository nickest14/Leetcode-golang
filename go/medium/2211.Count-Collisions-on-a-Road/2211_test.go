package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	directions string
}

type question2211 struct {
	params
	ans int
}

func Test_Problem2211(t *testing.T) {
	qs := []question2211{
		{
			params{directions: "RLRSLL"},
			5,
		},
		{
			params{directions: "LLRR"},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countCollisions(p.directions), ans)
	}
}
