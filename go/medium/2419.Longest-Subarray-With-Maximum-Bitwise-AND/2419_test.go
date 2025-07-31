package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2419 struct {
	params
	ans int
}

func Test_Problem2419(t *testing.T) {
	qs := []question2419{
		{
			params{nums: []int{1, 2, 3, 3, 2, 2}},
			2,
		},
		{
			params{nums: []int{1, 2, 3, 4}},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, longestSubarray(p.nums), ans)
	}
}
