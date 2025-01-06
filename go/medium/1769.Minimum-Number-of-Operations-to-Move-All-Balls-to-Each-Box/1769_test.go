package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	boxes string
}

type question1769 struct {
	params
	ans []int
}

func Test_Problem1769(t *testing.T) {
	qs := []question1769{
		{
			params{"110"},
			[]int{1, 1, 3},
		},
		{
			params{"001011"},
			[]int{11, 8, 5, 4, 3, 4},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minOperations(p.boxes), ans)
	}
}
