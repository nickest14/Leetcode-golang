package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question1863 struct {
	params
	ans int
}

func Test_Problem1863(t *testing.T) {
	qs := []question1863{
		{
			params{[]int{1, 3}},
			6,
		},
		{
			params{[]int{5, 1, 6}},
			28,
		},
		{
			params{[]int{3, 4, 5, 6, 7, 8}},
			480,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, subsetXORSum(p.nums), ans)
	}
}
