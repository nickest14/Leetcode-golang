package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums  []int
	n     int
	left  int
	right int
}

type question1508 struct {
	params
	ans int
}

func Test_Problem1508(t *testing.T) {
	qs := []question1508{
		{
			params{nums: []int{1, 2, 3, 4}, n: 4, left: 1, right: 5},
			13,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, rangeSum(p.nums, p.n, p.left, p.right), ans)
	}
}
