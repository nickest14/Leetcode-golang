package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2799 struct {
	params
	ans int
}

func Test_Problem2799(t *testing.T) {
	qs := []question2799{
		{
			params{[]int{1, 3, 1, 2, 2}},
			4,
		},
		{
			params{[]int{5, 5, 5, 5}},
			10,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countCompleteSubarrays(p.nums), ans)
	}
}
