package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question3512 struct {
	params
	ans int
}

func Test_Problem3512(t *testing.T) {
	qs := []question3512{
		{
			params{nums: []int{3, 9, 7}, k: 5},
			4,
		},
		{
			params{nums: []int{4, 1, 3}, k: 4},
			0,
		},
		{
			params{nums: []int{3, 2}, k: 6},
			5,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minOperations(p.nums, p.k), ans)
	}
}
