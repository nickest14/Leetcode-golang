package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question1963 struct {
	params
	ans int
}

func Test_Problem1963(t *testing.T) {
	qs := []question1963{
		{
			params{[]int{3, 1}},
			2,
		},
		{
			params{[]int{3, 2, 1, 5}},
			6,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countMaxOrSubsets(p.nums), ans)
	}
}
