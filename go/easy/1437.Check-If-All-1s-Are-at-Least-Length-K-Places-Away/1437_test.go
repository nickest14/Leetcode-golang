package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question1437 struct {
	params
	ans bool
}

func Test_Problem1437(t *testing.T) {
	qs := []question1437{
		{
			params{nums: []int{1, 0, 0, 0, 1, 0, 0}, k: 2},
			true,
		},
		{
			params{nums: []int{1, 0, 0, 1, 0, 1}, k: 2},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, kLengthApart(p.nums, p.k), ans)
	}
}
