package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question75 struct {
	params
	ans []int
}

func Test_Problem75(t *testing.T) {

	qs := []question75{
		{
			params{nums: []int{2, 0, 2, 1, 1, 0}},
			[]int{0, 0, 1, 1, 2, 2},
		},
		{
			params{nums: []int{2, 0, 1}},
			[]int{0, 1, 2},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		sortColors(p.nums)
		fmt.Printf("[input]: %v        [output]: %v, [answer]: %v\n", p, p.nums, ans)
	}
}
