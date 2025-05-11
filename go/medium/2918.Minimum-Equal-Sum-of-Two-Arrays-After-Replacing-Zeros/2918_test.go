package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums1 []int
	nums2 []int
}

type question2918 struct {
	params
	ans int64
}

func Test_Problem2918(t *testing.T) {
	qs := []question2918{
		{
			params{nums1: []int{3, 2, 0, 1, 0}, nums2: []int{6, 5, 0}},
			12,
		},
		{
			params{nums1: []int{2, 0, 2, 0}, nums2: []int{1, 4}},
			-1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minSum(p.nums1, p.nums2), ans)
	}
}
