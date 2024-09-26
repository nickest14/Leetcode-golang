package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	arr1 []int
	arr2 []int
}

type question3043 struct {
	params
	ans int
}

func Test_Problem3043(t *testing.T) {
	qs := []question3043{
		{
			params{arr1: []int{1, 10, 100}, arr2: []int{1000}},
			3,
		},
		{
			params{arr1: []int{1, 2, 3}, arr2: []int{4, 4, 4}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, longestCommonPrefix(p.arr1, p.arr2), ans)
	}
}
