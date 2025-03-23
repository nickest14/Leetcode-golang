package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question2560 struct {
	params
	ans int
}

func Test_Problem2560(t *testing.T) {
	qs := []question2560{
		{
			params{nums: []int{2, 3, 5, 9}, k: 2},
			5,
		},
		{
			params{nums: []int{2, 7, 9, 3, 1}, k: 2},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minCapability(p.nums, p.k), ans)
	}
}
