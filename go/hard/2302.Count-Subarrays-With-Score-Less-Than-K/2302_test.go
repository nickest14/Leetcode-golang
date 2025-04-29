package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int64
}

type question2302 struct {
	params
	ans int64
}

func Test_Problem2302(t *testing.T) {
	qs := []question2302{
		{
			params{nums: []int{2, 1, 4, 3, 5}, k: 10},
			6,
		},
		{
			params{nums: []int{1, 1, 1}, k: 5},
			5,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countSubarrays(p.nums, p.k), ans)
	}
}
