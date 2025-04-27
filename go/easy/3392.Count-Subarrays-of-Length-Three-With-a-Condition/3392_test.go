package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3392 struct {
	params
	ans int
}

func Test_Problem3392(t *testing.T) {
	qs := []question3392{
		{
			params{nums: []int{1, 2, 1, 4, 1}},
			1,
		},
		{
			params{nums: []int{1, 1, 1}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countSubarrays(p.nums), ans)
	}
}
