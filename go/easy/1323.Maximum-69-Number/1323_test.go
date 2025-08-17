package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	num int
}

type question1323 struct {
	params
	ans int
}

func Test_Problem1323(t *testing.T) {
	qs := []question1323{
		{
			params{9669},
			9969,
		},
		{
			params{9996},
			9999,
		},
		{
			params{9999},
			9999,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximum69Number(p.num), ans)
	}
}
