package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	mapping []int
	nums    []int
}

type question2191 struct {
	params
	ans []int
}

func Test_Problem2191(t *testing.T) {
	qs := []question2191{
		{
			params{mapping: []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, nums: []int{991, 338, 38}},
			[]int{338, 38, 991},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, sortJumbled(p.mapping, p.nums), ans)
	}
}
