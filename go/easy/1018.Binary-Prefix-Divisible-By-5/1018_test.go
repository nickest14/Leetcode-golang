package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question1018 struct {
	params
	ans []bool
}

func Test_Problem1018(t *testing.T) {
	qs := []question1018{
		{
			params{[]int{0, 1, 1}},
			[]bool{true, false, false},
		},
		{
			params{[]int{1, 1, 1}},
			[]bool{false, false, false},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, prefixesDivBy5(p.nums), ans)
	}
}
