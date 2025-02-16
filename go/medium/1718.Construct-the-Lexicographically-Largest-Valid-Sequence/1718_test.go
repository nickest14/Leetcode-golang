package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question1718 struct {
	params
	ans []int
}

func Test_Problem1718(t *testing.T) {
	qs := []question1718{
		{
			params{3},
			[]int{3, 1, 2, 3, 2},
		},
		{
			params{5},
			[]int{5, 3, 1, 4, 3, 5, 2, 4},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, constructDistancedSequence(p.n), ans)
	}
}
