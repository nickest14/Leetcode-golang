package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question1877 struct {
	params
	ans int
}

func Test_Problem1877(t *testing.T) {
	qs := []question1877{
		{
			params{[]int{3, 5, 2, 3}},
			7,
		},
		{
			params{[]int{3, 5, 4, 2, 4, 6}},
			8,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minPairSum(p.nums), ans)
	}
}
