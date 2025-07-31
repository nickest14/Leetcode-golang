package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2411 struct {
	params
	ans []int
}

func Test_Problem2411(t *testing.T) {
	qs := []question2411{
		{
			params{nums: []int{1, 0, 2, 1, 3}},
			[]int{3, 3, 2, 2, 1},
		},
		{
			params{nums: []int{1, 2}},
			[]int{2, 1},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, smallestSubarrays(p.nums), ans)
	}
}
