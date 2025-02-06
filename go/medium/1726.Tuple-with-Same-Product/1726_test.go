package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question1726 struct {
	params
	ans int
}

func Test_Problem1726(t *testing.T) {
	qs := []question1726{
		{
			params{[]int{2, 3, 4, 6}},
			8,
		},
		{
			params{[]int{1, 2, 4, 5, 10}},
			16,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, tupleSameProduct(p.nums), ans)
	}
}
