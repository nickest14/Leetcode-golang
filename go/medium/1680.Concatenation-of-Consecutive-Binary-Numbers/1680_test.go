package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question1680 struct {
	params
	ans int
}

func Test_Problem1680(t *testing.T) {
	qs := []question1680{
		{
			params{1},
			1,
		},
		{
			params{3},
			27,
		},
		{
			params{12},
			505379714,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, concatenatedBinary(p.n), ans)
	}
}
