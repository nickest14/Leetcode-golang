package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3350 struct {
	params
	ans int
}

func Test_Problem3350(t *testing.T) {
	qs := []question3350{
		{
			params{nums: []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}},
			3,
		},
		{
			params{nums: []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxIncreasingSubarrays(p.nums), ans)
	}
}
