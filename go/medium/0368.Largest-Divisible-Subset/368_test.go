package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question368 struct {
	params
	ans []int
}

func Test_Problem368(t *testing.T) {

	qs := []question368{
		{
			params{nums: []int{1}},
			[]int{1},
		},
		{
			params{nums: []int{1, 2, 3}},
			[]int{1, 2},
		},
		{
			params{nums: []int{1, 2, 4, 8}},
			[]int{1, 2, 4, 8},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v, [answer]: %v\n", p, largestDivisibleSubset(p.nums), ans)
	}
}
