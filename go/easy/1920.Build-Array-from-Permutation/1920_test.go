package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question1920 struct {
	params
	ans []int
}

func Test_Problem1920(t *testing.T) {
	qs := []question1920{
		{
			params{nums: []int{0, 2, 1, 5, 3, 4}},
			[]int{0, 1, 2, 4, 5, 3},
		},
		{
			params{nums: []int{5, 0, 1, 2, 3, 4}},
			[]int{4, 5, 0, 1, 2, 3},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, buildArray(p.nums), ans)
	}
}
