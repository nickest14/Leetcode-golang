package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2197 struct {
	params
	ans []int
}

func Test_Problem2197(t *testing.T) {
	qs := []question2197{
		{
			params{[]int{6, 4, 3, 2, 7, 6, 2}},
			[]int{12, 7, 6},
		},
		{
			params{[]int{2, 2, 1, 1, 3, 3, 3}},
			[]int{2, 1, 1, 3},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, replaceNonCoprimes(p.nums), ans)
	}
}
