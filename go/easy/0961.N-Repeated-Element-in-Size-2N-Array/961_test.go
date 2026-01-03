package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question961 struct {
	params
	ans int
}

func Test_Problem961(t *testing.T) {
	qs := []question961{
		{
			params{[]int{1, 2, 3, 3}},
			3,
		},
		{
			params{[]int{2, 1, 2, 5, 3, 2}},
			2,
		},
		{
			params{[]int{5, 1, 5, 2, 5, 3, 5, 4}},
			5,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, repeatedNTimes(p.nums), ans)
	}
}
