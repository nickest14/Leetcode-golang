package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2342 struct {
	params
	ans int
}

func Test_Problem2342(t *testing.T) {

	qs := []question2342{
		{
			params{[]int{18, 43, 36, 13, 7}},
			54,
		},
		{
			params{[]int{10, 12, 19, 14}},
			-1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v, [answer]: %v\n", p, maximumSum(p.nums), ans)
	}
}
