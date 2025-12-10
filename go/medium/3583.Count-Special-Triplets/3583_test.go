package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3583 struct {
	params
	ans int
}

func Test_Problem(t *testing.T) {
	qs := []question3583{
		{
			params{nums: []int{6, 3, 6}},
			1,
		},
		{
			params{nums: []int{0, 1, 0, 0}},
			1,
		},
		{
			params{nums: []int{8, 4, 2, 8, 4}},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, specialTriplets(p.nums), ans)
	}
}
