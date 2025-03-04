package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums  []int
	pivot int
}

type question2161 struct {
	params
	ans []int
}

func Test_Problem2161(t *testing.T) {
	qs := []question2161{
		{
			params{nums: []int{9, 12, 5, 10, 14, 3, 10}, pivot: 10},
			[]int{9, 5, 3, 10, 10, 12, 14},
		},
		{
			params{nums: []int{-3, 4, 3, 2}, pivot: 2},
			[]int{-3, 2, 4, 3},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, pivotArray(p.nums, p.pivot), ans)
	}
}
