package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question2779 struct {
	params
	ans int
}

func Test_Problem2779(t *testing.T) {
	qs := []question2779{
		{
			params{nums: []int{4, 6, 1, 2}, k: 2},
			3,
		},
		{
			params{nums: []int{1, 1, 1, 1}, k: 10},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximumBeauty(p.nums, p.k), ans)
	}
}
