package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2348 struct {
	params
	ans int
}

func Test_Problem2348(t *testing.T) {
	qs := []question2348{
		{
			params{nums: []int{1, 3, 0, 0, 2, 0, 0, 4}},
			6,
		},
		{
			params{nums: []int{0, 0, 0, 2, 0, 0}},
			9,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, zeroFilledSubarray(p.nums), ans)
	}
}
