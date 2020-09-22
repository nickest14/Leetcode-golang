package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums   []int
	target int
}

type question1 struct {
	params
	ans []int
}

func Test_Problem1(t *testing.T) {
	qs := []question1{
		{
			params{[]int{1, 2, 3}, 5},
			[]int{1, 2},
		},
		{
			params{[]int{4, 2, 9, 1}, 3},
			[]int{1, 3},
		},
		{
			params{[]int{7, 1, 5, 2, 9}, 16},
			[]int{0, 4},
		},
	}

	for _, q := range qs {
		p := q.params
		ans := twoSum(p.nums, p.target)
		fmt.Printf("[input]: %v        [output]: %v\n", p, ans)
	}
}
