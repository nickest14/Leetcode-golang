package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums1 []int
	nums2 []int
}

type question1855 struct {
	params
	ans int
}

func Test_Problem1855(t *testing.T) {
	qs := []question1855{
		{
			params{nums1: []int{55, 30, 5, 4, 2}, nums2: []int{100, 20, 10, 10, 5}},
			2,
		},
		{
			params{nums1: []int{2, 2, 2}, nums2: []int{10, 10, 1}},
			1,
		},
		{
			params{nums1: []int{30, 29, 19, 5}, nums2: []int{25, 25, 25, 25, 25}},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxDistance(p.nums1, p.nums2), ans)
	}
}
