package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question3066 struct {
	params
	ans int
}

func Test_Problem3066(t *testing.T) {
	qs := []question3066{
		{
			params{nums: []int{2, 11, 10, 1, 3}, k: 10},
			2,
		},
		{
			params{nums: []int{1, 1, 2, 4, 9}, k: 20},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minOperations(p.nums, p.k), ans)
	}
}
