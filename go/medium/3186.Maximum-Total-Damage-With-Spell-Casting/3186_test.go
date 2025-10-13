package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	power []int
}

type question3186 struct {
	params
	ans int64
}

func Test_Problem3186(t *testing.T) {
	qs := []question3186{
		{
			params{[]int{1, 1, 3, 4}},
			6,
		},
		{
			params{[]int{7, 1, 6, 6}},
			13,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximumTotalDamage(p.power), ans)
	}
}
