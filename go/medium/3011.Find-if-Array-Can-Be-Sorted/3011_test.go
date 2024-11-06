package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3011 struct {
	params
	ans bool
}

func Test_Problem3011(t *testing.T) {
	qs := []question3011{
		{
			params{[]int{8, 4, 2, 30, 15}},
			true,
		},
		{
			params{[]int{3, 16, 8, 4, 2}},
			true,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, canSortArray(p.nums), ans)
	}
}
