package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	digits []int
}

type question590 struct {
	params
	ans []int
}

func Test_Problem590(t *testing.T) {

	qs := []question590{
		{
			params{[]int{1, 2, 3}},
			[]int{1, 2, 4},
		},
		{
			params{[]int{4, 3, 2, 1}},
			[]int{4, 3, 2, 2},
		},
		{
			params{[]int{9}},
			[]int{1, 0},
		},
	}

	for _, q := range qs {
		p := q.params
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, plusOne(p.digits), q.ans)
	}
}
