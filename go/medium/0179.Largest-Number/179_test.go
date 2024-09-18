package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question179 struct {
	params
	ans string
}

func Test_Problem179(t *testing.T) {

	qs := []question179{
		{
			params{[]int{10, 2}},
			"210",
		},
		{
			params{[]int{3, 30, 34, 5, 9}},
			"9534330",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v, [answer]: %v\n", p, largestNumber(p.nums), ans)
	}
}
