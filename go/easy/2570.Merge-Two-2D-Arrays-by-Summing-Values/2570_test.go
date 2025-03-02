package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums1 [][]int
	nums2 [][]int
}

type question2570 struct {
	params
	ans [][]int
}

func Test_Problem2570(t *testing.T) {
	qs := []question2570{
		{
			params{nums1: [][]int{{1, 2}, {2, 3}, {4, 5}}, nums2: [][]int{{1, 4}, {3, 2}, {4, 1}}},
			[][]int{{1, 6}, {2, 3}, {3, 2}, {4, 6}},
		},
		{
			params{nums1: [][]int{{2, 4}, {3, 6}, {5, 5}}, nums2: [][]int{{1, 3}, {4, 3}}},
			[][]int{{1, 3}, {2, 4}, {3, 6}, {4, 3}, {5, 5}},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, mergeArrays(p.nums1, p.nums2), ans)
	}
}
