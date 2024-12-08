package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums          []int
	maxOperations int
}

type question1760 struct {
	params
	ans int
}

func Test_Problem(t *testing.T) {
	qs := []question1760{
		{
			params{nums: []int{9, 6}, maxOperations: 2},
			5,
		},
		{
			params{nums: []int{9}, maxOperations: 2},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumSize(p.nums, p.maxOperations), ans)
	}
}
