package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question2537 struct {
	params
	ans int64
}

func Test_Problem2537(t *testing.T) {
	qs := []question2537{
		{
			params{nums: []int{1, 1, 1, 1, 1}, k: 10},
			1,
		},
		{
			params{nums: []int{3, 1, 4, 3, 2, 2, 4}, k: 2},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countGood(p.nums, p.k), ans)
	}
}
