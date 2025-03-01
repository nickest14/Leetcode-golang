package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2460 struct {
	params
	ans []int
}

func Test_Problem2460(t *testing.T) {
	qs := []question2460{
		{
			params{[]int{1, 2, 2, 1, 1, 0}},
			[]int{1, 4, 2, 0, 0, 0},
		},
		{
			params{[]int{0, 1}},
			[]int{1, 0},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, applyOperations(p.nums), ans)
	}
}
