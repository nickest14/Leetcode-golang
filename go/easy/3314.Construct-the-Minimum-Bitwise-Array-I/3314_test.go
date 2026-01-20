package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3314 struct {
	params
	ans []int
}

func Test_Problem3314(t *testing.T) {
	qs := []question3314{
		{
			params{nums: []int{2, 3, 5, 7}},
			[]int{-1, 1, 4, 3},
		},
		{
			params{nums: []int{11, 13, 31}},
			[]int{9, 12, 15},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minBitwiseArray(p.nums), ans)
	}
}
