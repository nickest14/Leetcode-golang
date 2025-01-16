package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums1 []int
	nums2 []int
}

type question2425 struct {
	params
	ans int
}

func Test_Problem2425(t *testing.T) {
	qs := []question2425{
		{
			params{nums1: []int{2, 1, 3}, nums2: []int{10, 2, 5, 0}},
			13,
		},
		{
			params{nums1: []int{1, 2}, nums2: []int{3, 4}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, xorAllNums(p.nums1, p.nums2), ans)
	}
}
