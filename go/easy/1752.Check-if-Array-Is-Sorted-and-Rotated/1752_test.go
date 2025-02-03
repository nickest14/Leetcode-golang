package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question1752 struct {
	params
	ans bool
}

func Test_Problem1752(t *testing.T) {
	qs := []question1752{
		// {
		// 	params{[]int{3, 4, 5, 1, 2}},
		// 	true,
		// },
		// {
		// 	params{[]int{2, 1, 3, 4}},
		// 	false,
		// },
		// {
		// 	params{[]int{1, 2, 3}},
		// 	true,
		// },
		{
			params{[]int{3, 2, 1}},
			false,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, check(p.nums), ans)
	}
}
