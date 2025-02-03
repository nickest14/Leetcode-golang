package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3151 struct {
	params
	ans bool
}

func Test_Problem3151(t *testing.T) {
	qs := []question3151{
		{
			params{[]int{1}},
			true,
		},
		{
			params{[]int{2, 1, 4}},
			true,
		},
		{
			params{[]int{4, 3, 1, 6}},
			false,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, isArraySpecial(p.nums), ans)
	}
}
