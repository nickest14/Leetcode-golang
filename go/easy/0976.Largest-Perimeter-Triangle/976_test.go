package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question976 struct {
	params
	ans int
}

func Test_Problem976(t *testing.T) {
	qs := []question976{
		{
			params{[]int{2, 1, 2}},
			5,
		},
		{
			params{[]int{1, 2, 1, 10}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, largestPerimeter(p.nums), ans)
	}
}
