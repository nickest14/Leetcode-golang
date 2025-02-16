package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question2698 struct {
	params
	ans int
}

func Test_Problem2698(t *testing.T) {
	qs := []question2698{
		{
			params{n: 10},
			182,
		},
		{
			params{n: 37},
			1478,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, punishmentNumber(p.n), ans)
	}
}
