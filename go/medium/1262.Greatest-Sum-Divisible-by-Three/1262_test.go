package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question1262 struct {
	params
	ans int
}

func Test_Problem1262(t *testing.T) {
	qs := []question1262{
		{
			params{[]int{3, 6, 5, 1, 8}},
			18,
		},
		{
			params{[]int{4}},
			0,
		},
		{
			params{[]int{1, 2, 3, 4, 4}},
			12,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxSumDivThree(p.nums), ans)
	}
}
