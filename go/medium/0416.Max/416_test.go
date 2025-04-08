package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question416 struct {
	params
	ans bool
}

func Test_Problem416(t *testing.T) {
	qs := []question416{
		{
			params{nums: []int{1, 5, 11, 5}},
			true,
		},
		{
			params{nums: []int{1, 2, 3, 5}},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, canPartition(p.nums), ans)
	}
}
